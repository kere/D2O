package home

import (
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"regexp"
	"strconv"

	"github.com/kere/gno/db"
	"github.com/kere/gno/httpd"
	"github.com/kere/gno/libs/util"
	"github.com/microcosm-cc/bluemonday"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
	// blackfriday "gopkg.in/russross/blackfriday.v2"

	"github.com/kere/blackfriday"

	"onqee.visualstudio.com/D2O/app/model"
	"onqee.visualstudio.com/D2O/app/model/baseinfo"
	"onqee.visualstudio.com/D2O/app/model/selem"
	"onqee.visualstudio.com/D2O/app/page"
)

var policy = bluemonday.UGCPolicy()

func init() {
	policy.AllowAttrs("data-echo").OnElements("img")
}

// CellView page class
type CellView struct {
	httpd.P
}

// NewCellView func
func NewCellView() *CellView {
	d := &CellView{}
	d.Init("", "CellView", homeDir)
	page.Init(&d.PA, page.Option{HasHeader: true, HasFooter: false, NoRequireJS: true, NoPageLoad: true})

	d.PA.Bottom = append(d.PA.Bottom, page.EchojsRender)
	d.PA.Body = &CellViewRender{}

	return d
}

// CellViewData page data
type CellViewData struct {
	Title    template.HTML
	Content  template.HTML
	Subforms template.HTML
	Tags     template.HTML
}

// Page page
func (d *CellView) Page(ctx *fasthttp.RequestCtx) (interface{}, error) {
	v := ctx.UserValue(model.FieldIID).(string)
	iid, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return nil, err
	}

	return selem.PageViewData(iid)
}

// CellViewRender class
type CellViewRender struct{}

// RenderD func
func (t *CellViewRender) RenderD(w io.Writer, data interface{}) error {
	row := data.(db.MapRow)

	ojson := selem.OJSON{}
	json.Unmarshal(row.Bytes("o_json"), &ojson)
	contents := ojson.Contents
	if len(contents) == 0 {
		return errors.New("Content Not Found")
	}
	c1 := contents[0]

	w.Write(util.Str2Bytes(`<article id="articleMain" class="gno-cell-view container clearfix"><header id="header" class="header m-b-md"><h1 id="headerTitle"><a id="txtTitle" class="text-title">`))
	// title
	w.Write(util.Str2Bytes(template.HTMLEscapeString(c1.Title)))
	w.Write(util.Str2Bytes("</a></h1></header>\n")) // header end

	// text
	unsafe := blackfriday.Run(util.Str2Bytes(c1.Text))
	h := policy.SanitizeBytes(unsafe)
	w.Write(util.Str2Bytes(`<div id="content" class="content m-b-md">`))
	w.Write(h)
	w.Write(model.BDivEnd) // header end

	// subforms
	//   <div id="subforms" class="subforms m-b-md">{* .Subforms *}</div>
	w.Write(util.Str2Bytes(`<div id="subforms" class="subforms m-b-md">`))

	// area
	areas := ojson.Area
	areaRender(w, areas)

	// tags
	tags := row.Strings(model.FieldTags)
	tagsRender(w, tags)

	// datas
	renderArticleData(w, row, ojson.SubForms)

	w.Write(util.Str2Bytes("</article>\n")) // subforms end

	w.Write(util.Str2Bytes("<script>document.title='"))
	w.Write(util.Str2Bytes(c1.Title))
	w.Write(util.Str2Bytes("';_nick='"))
	w.Write(util.Str2Bytes(row.String(model.FieldNick)))
	w.Write(util.Str2Bytes(`';let ca = document.cookie.split(';'), str;
		for(let i=0;i < ca.length;i++) {
			str = ca[i].trim();
			if(str.substr(0,5)=='_nick' && str.split('=')[1]==_nick){
				let t = document.getElementById("txtTitle");
				t.className=""
				t.href = "/cell/edit/" + `))
	w.Write(util.Str2Bytes(row.String(model.FieldIID)))
	w.Write(util.Str2Bytes(`;
				break;
			}
		}
		let __imgs = document.querySelectorAll('#articleMain img');
		for (let i=0;i<__imgs.length;i++) {
			if(!__imgs[i].alt) continue;
			let span=document.createElement("span");
			span.className = 'img-desc';
			span.innerText = '-- ' + __imgs[i].alt + ' --';
			__imgs[i].parentNode.insertBefore(span, __imgs[i].nextSibling);
		}
		</script>`)) // subforms end

	return nil
}

func areaRender(w io.Writer, areas []baseinfo.Area) {
	w.Write(util.Str2Bytes(`<div id="area" class="area">`))
	l := len(areas)
	buf := bytebufferpool.Get()
	for i := 0; i < l; i++ {
		buf.WriteString(`<strong><a href="/area/`)
		buf.WriteString(areas[i].CN)
		buf.WriteString(`">`)
		buf.WriteString(areas[i].CN)
		buf.WriteString("</a></strong>")
	}
	w.Write(buf.Bytes())
	w.Write(model.BDivEnd)
}

func tagsRender(w io.Writer, tags []string) {
	w.Write(util.Str2Bytes(`<div id="tags" class="tags m-b-md">`))
	l := len(tags)
	buf := bytebufferpool.Get()
	for i := 0; i < l; i++ {
		buf.WriteString(`<strong><a href="/tag/`)
		buf.WriteString(tags[i])
		buf.WriteString(`">`)
		buf.WriteString(tags[i])
		buf.WriteString("</a></strong>")
	}
	w.Write(buf.Bytes())
	w.Write(model.BDivEnd)
	bytebufferpool.Put(buf)
}

// var linkReg = regexp.MustCompile(`\[(.*)\]\((http[s]?:.+\))`)
var linkReg = regexp.MustCompile(`^(?:(\S+)\s*\|\s*)?(http[s]?:\/\/\S+)`)

func renderArticleData(w io.Writer, row db.MapRow, subforms []selem.SubForm) {
	// subforms = dat.o_json.subforms, html='', form, items, th1, th2;
	alllinks := make([]selem.FormItem, 0, 10)
	l := len(subforms)
	for i := 0; i < l; i++ {
		// 取出“参考连接”,然后统一至于页面
		form := subforms[i]
		items := form.Items
		n := len(items)
		for k := 0; k < n; k++ {
			if items[k].Key != "参考连接" {
				continue
			}
			alllinks = append(alllinks, items[k])
			items[k] = selem.FormItem{}
		}
	}

	buf := bytebufferpool.Get()
	for i := 0; i < l; i++ {
		form := subforms[i]
		buf.WriteString(`<ul class="data-list m-b">`)

		var th1 string
		var th2 string
		if form.DateON != "" {
			th1 = form.DateON
		}

		if form.Title != "" {
			th2 = template.HTMLEscapeString(form.Title)
		}

		if th1 != "" || th2 != "" {
			buf.WriteString(`<li class="sub-title">`)
			if th1 != "" {
				buf.WriteString(`<strong class="date_on m-r">`)
				buf.WriteString(th1)
				buf.WriteString(`</strong>`)
			}
			if th2 != "" {
				buf.WriteString(`<strong class="title">`)
				buf.WriteString(th2)
				buf.WriteString(`</strong>`)
			}
			buf.WriteString("</li>")
		}

		// items
		isOK := false
		items := form.Items
		count := len(items)
		for k := 0; k < count; k++ {
			if items[k].Value == "" {
				continue
			}

			buf.WriteString(`<li><strong class="m-r-sm">`)
			buf.WriteString(template.HTMLEscapeString(items[k].Key))
			buf.WriteString(`:</strong><span>`)

			match := linkReg.FindAllSubmatch(util.Str2Bytes(items[k].Value), -1)
			if len(match) > 0 && len(match[0]) == 3 {
				buf2 := bytebufferpool.Get()
				buf2.WriteString(`<a href="`)
				buf2.Write(match[0][2])
				buf2.WriteString("\">")
				if len(match[0][1]) == 0 {
					template.HTMLEscape(buf2, match[0][2])
				} else {
					template.HTMLEscape(buf2, match[0][1])
				}
				buf2.WriteString("</a>")
				buf.Write(policy.SanitizeBytes(buf2.Bytes()))
				bytebufferpool.Put(buf2)
			} else {
				buf.WriteString(template.HTMLEscapeString(items[k].Value))
			}

			buf.WriteString(`</span><li>`)
			isOK = true
		}
		buf.WriteString(`</ul>`)

		if isOK {
			w.Write(buf.Bytes())
		}
	}
	bytebufferpool.Put(buf)

	l = len(alllinks)
	if l > 0 {
		w.Write(util.Str2Bytes(`<strong>参考连接：</strong><ul class="gno-ref-links m-b-md">`))
		for i := 0; i < l; i++ {
			w.Write(util.Str2Bytes(`<li>`))
			match := linkReg.FindAllSubmatch(util.Str2Bytes(alllinks[i].Value), -1)
			if len(match) > 0 && len(match[0]) == 3 {
				buf2 := bytebufferpool.Get()
				buf2.WriteString(`<a href="`)
				buf2.Write(match[0][2])
				buf2.WriteString("\">")
				if len(match[0][1]) == 0 {
					template.HTMLEscape(buf2, match[0][2])
				} else {
					template.HTMLEscape(buf2, match[0][1])
				}
				buf2.WriteString("</a>")
				w.Write(policy.SanitizeBytes(buf2.Bytes()))
				bytebufferpool.Put(buf2)
			} else {
				w.Write(util.Str2Bytes(template.HTMLEscapeString(alllinks[i].Value)))
			}

			w.Write(util.Str2Bytes("</li>"))

		}
		w.Write(util.Str2Bytes("</ul>"))
	}

	renderArticleFooter(w, row)
}

func renderArticleFooter(w io.Writer, row db.MapRow) {
	dateon := row.String(model.FieldDateON)
	author := row.String(model.FieldNick)
	w.Write(util.Str2Bytes(`<footer class="article-footer"><p class="date_on">`))
	if len(dateon) > 10 {
		w.Write(util.Str2Bytes(dateon)[:10])
	}
	w.Write(util.Str2Bytes(`</p><p class="author">`))
	w.Write(util.Str2Bytes(author))
	w.Write(util.Str2Bytes(`</p></footer>`))
}
