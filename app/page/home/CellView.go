package home

import (
	"encoding/json"
	"errors"
	"html/template"
	"io"
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
	d.D.Init("信息", "CellView", homeDir)

	page.Init(&d.D, page.Option{HasHeader: true, HasFooter: false, NoRequireJS: true, NoPageLoad: true})
	d.D.Bottom = append(d.D.Bottom, page.FooterViewEndRender, page.EchojsRender)
	d.D.Body = append(d.D.Body, &CellViewRender{})

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

	row, err := db.NewQuery(selem.Table).Where("iid=?", iid).QueryOne()
	if err != nil {
		return nil, err
	}

	if row.IsEmpty() {
		return nil, errors.New("Page Not Found")
	}

	return row, nil
}

// <div class="gno-cell-view container">
//   <div id="header" class="header m-b-md">
//     <h1 id="headerTitle"></h1>
//   </div>
//   <div id="content" class="content m-b-md">{* .Content *}</div>
//   <div id="subforms" class="subforms m-b-md">{* .Subforms *}</div>
//   <div id="tags" class="tags m-b-md"></div>
// </div>

// CellViewRender class
type CellViewRender struct{}

// RenderWithData func
func (t *CellViewRender) RenderWithData(w io.Writer, data interface{}) error {
	row := data.(db.MapRow)
	ojson := selem.OJSON{}
	json.Unmarshal(row.Bytes("o_json"), &ojson)
	contents := ojson.Contents
	if len(contents) == 0 {
		return errors.New("Content Not Found")
	}
	c1 := contents[0]

	w.Write(util.Str2Bytes(`<article id="articleMain" class="gno-cell-view container clearfix"><header id="header" class="header m-b-md"><h1 id="headerTitle">`))

	// title
	w.Write(util.Str2Bytes(template.HTMLEscapeString(c1.Title)))
	w.Write(util.Str2Bytes("</h1></header>\n")) // header end

	// text
	unsafe := blackfriday.Run(util.Str2Bytes(c1.Text))
	h := policy.SanitizeBytes(unsafe)
	w.Write(util.Str2Bytes(`<div id="content" class="content m-b-md">`))
	w.Write(h)
	w.Write(util.Str2Bytes("</div>\n")) // header end

	// subforms
	//   <div id="subforms" class="subforms m-b-md">{* .Subforms *}</div>
	w.Write(util.Str2Bytes(`<div id="subforms" class="subforms m-b-md">`))
	renderArticleData(w, row, ojson.SubForms)

	w.Write(util.Str2Bytes("</article>\n")) // subforms end

	w.Write(util.Str2Bytes(`<script>
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

func tagsRender(w io.Writer, tags []string) {
	w.Write(util.Str2Bytes(`<div id="tags" class="tags m-b-md">`))
	l := len(tags)
	for i := 0; i < l; i++ {
		w.Write(util.Str2Bytes(`<strong class="el-tag el-tag--light"><a href="`))
		w.Write(util.Str2Bytes(`">`))
		w.Write(util.Str2Bytes(template.HTMLEscapeString(tags[i])))
		w.Write(util.Str2Bytes("</a></strong>"))
	}
	w.Write(util.Str2Bytes(`</div>`))
}

func renderArticleData(w io.Writer, row db.MapRow, subforms []selem.SubForm) {
	tags := row.Strings(model.FieldTags)
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
			buf.WriteString(template.HTMLEscapeString(items[k].Value))
			buf.WriteString(`</span><li>`)
			isOK = true
		}
		buf.WriteString(`</ul>`)

		if isOK {
			w.Write(buf.Bytes())
		}

		bytebufferpool.Put(buf)
	}

	tagsRender(w, tags)

	l = len(alllinks)
	if l > 0 {
		w.Write(util.Str2Bytes(`<h3>参考连接：</h3><ul class="gno-ref-links m-b-md">`))
		for i := 0; i < l; i++ {
			w.Write(util.Str2Bytes(`<li><a href="`))
			w.Write(util.Str2Bytes(alllinks[i].Value))
			w.Write(util.Str2Bytes("\">"))
			w.Write(util.Str2Bytes(alllinks[i].Value))
			w.Write(util.Str2Bytes("</a></li>"))
		}
		w.Write(util.Str2Bytes("</ul>"))
	}

	renderArticleFooter(w, row)
}

func renderArticleFooter(w io.Writer, row db.MapRow) {
	dateon := row.String(model.FieldDateON)
	author := "someone"
	w.Write(util.Str2Bytes(`<footer class="article-footer"><p class="date_on">`))
	if len(dateon) > 10 {
		w.Write(util.Str2Bytes(dateon)[:10])
	}
	w.Write(util.Str2Bytes(`</p><p class="author">`))
	w.Write(util.Str2Bytes(author))
	w.Write(util.Str2Bytes(`</p></footer>`))
}

// // Auth page
// func (d *CellView) Auth(ctx *fasthttp.RequestCtx) error {
// 	return nil
// }

// type viewDataRender struct {
// 	Src []byte
// }
//
// var (
// 	bViewHTMLScript    = []byte("<script>var pagedata=")
// 	bViewHTML01        = []byte(";")
// 	bViewHTMLScriptEnd = []byte("</script>")
// )
//
// // Render func
// func (t *viewDataRender) Render(w io.Writer) error {
// 	w.Write(bViewHTMLScript)
// 	if len(t.Src) == 0 {
// 		w.Write(util.Str2Bytes("null"))
// 	} else {
// 		w.Write(t.Src)
// 	}
// 	w.Write(bViewHTML01)
// 	w.Write(bViewHTMLScriptEnd)
// 	return nil
// }
