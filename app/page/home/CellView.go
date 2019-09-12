package home

import (
	"encoding/json"
	"io"
	"strconv"

	"github.com/kere/gno/db"
	"github.com/kere/gno/httpd"
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

// NewCellView func
func NewCellView() *CellView {
	d := &CellView{}

	d.PA.Body = &CellViewRender{}

	d.Init("", "CellView", homeDir)
	page.Init(&d.PA, page.Option{HasHeader: true, HasFooter: false, NoRequireJS: true, NoPageLoad: true})
	d.PA.Bottom = append(d.PA.Bottom, page.EchojsRender)

	return d
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
