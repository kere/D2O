package home

import (
	"encoding/json"
	"io"
	"strconv"

	"github.com/kere/gno/httpd"
	"github.com/valyala/fasthttp"
	"onqee.visualstudio.com/D2O/app"
	"onqee.visualstudio.com/D2O/app/model/selem"
	"onqee.visualstudio.com/D2O/app/page"
)

// CellView page class
type CellView struct {
	httpd.P
}

type viewDataRender struct {
	Src []byte
}

var (
	bViewHTMLScript    = []byte("<script>var pagedata=")
	bViewHTML01        = []byte(";")
	bViewHTMLScriptEnd = []byte("</script>")
)

// Render func
func (t *viewDataRender) Render(w io.Writer) error {
	w.Write(bViewHTMLScript)
	w.Write(t.Src)
	w.Write(bViewHTML01)
	w.Write(bViewHTMLScriptEnd)
	return nil
}

// NewCellView func
func NewCellView() *CellView {
	d := &CellView{}
	d.D.Init("信息", "CellView", homeDir)

	d.D.Bottom = make([]httpd.IRender, 1, 4)
	d.D.Bottom[0] = &viewDataRender{}

	page.Init(&d.D, page.Option{HasHeader: true, HasFooter: true})

	return d
}

// Page page
func (d *CellView) Page(ctx *fasthttp.RequestCtx) error {
	v := ctx.UserValue(app.FieldIID).(string)
	iid, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return err
	}

	vo, err := selem.PageData(iid)
	if err != nil {
		return err
	}
	r := d.D.Bottom[0].(*viewDataRender)

	if vo.IID == 0 {
		r.Src = nil
		return nil
	}

	// set data
	src, err := json.Marshal(vo)
	if err != nil {
		return err
	}
	r.Src = src

	return nil
}

// // Auth page
// func (d *CellView) Auth(ctx *fasthttp.RequestCtx) error {
// 	return nil
// }
