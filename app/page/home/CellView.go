package home

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/kere/gno/db"
	"github.com/kere/gno/httpd"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
	// blackfriday "gopkg.in/russross/blackfriday.v2"

	"onqee.visualstudio.com/D2O/app/model"
	"onqee.visualstudio.com/D2O/app/model/selem"
	"onqee.visualstudio.com/D2O/app/page"
	"onqee.visualstudio.com/D2O/app/qtpl"
)

// CellView page class
type CellView struct {
	httpd.P
}

var cpage *CellView

// GetCellView var
func GetCellView() *CellView {
	if cpage != nil {
		return cpage
	}
	cpage := &CellView{}

	cpage.PA.Body = &CellViewRender{}

	cpage.Init("", "CellView", homeDir)
	page.Init(&cpage.PA, httpd.PageOption{HasHeader: true, HasFooter: false, NoRequireJS: true, NoPageLoad: true})
	cpage.PA.Bottom = append(cpage.PA.Bottom, page.EchojsRender, httpd.RequireJSWithSrc(&cpage.PA, httpd.ReadRequireJS()))

	// set cache
	cpage.PA.CacheOption.PageMode = httpd.CacheModePagePath

	return cpage
}

// ClearCache page
func (d *CellView) ClearCache(iid int64) {
	// clear page cache
	buf := bytebufferpool.Get()
	buf.WriteString("/cell/view/")
	buf.WriteString(fmt.Sprint(iid))
	httpd.ClearCache(buf.Bytes(), d)
	bytebufferpool.Put(buf)
}

// Page page
func (d *CellView) Page(ctx *fasthttp.RequestCtx) (interface{}, error) {
	v := ctx.UserValue(model.FieldIID).(string)
	iid, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return nil, err
	}

	row, err := selem.PageViewData(iid)
	if err != nil {
		return row, err
	}
	if row.IsEmpty() {
		return row, model.ErrDataNotFound
	}
	return row, nil
}

// CellViewRender class
type CellViewRender struct{}

// RenderD func
func (t *CellViewRender) RenderD(w io.Writer, data interface{}) error {
	row := data.(db.MapRow)
	ojson := selem.OJSON{}
	json.Unmarshal(row.Bytes("o_json"), &ojson)
	if len(ojson.Contents) == 0 {
		return model.ErrDataNotFound
	}

	qtpl.WriteCellView(w, row, &ojson)
	return nil
}
