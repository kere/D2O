package home

import (
	"github.com/kere/gno/httpd"
	"github.com/valyala/fasthttp"
	"onqee.visualstudio.com/D2O/app/page"
)

// CellView page class
type CellView struct {
	httpd.P
}

// NewCellView func
func NewCellView() *CellView {
	d := &CellView{}
	d.D.Init("信息", "CellView", homeDir)
	page.Init(&d.D, page.Option{HasHeader: true, HasFooter: true})

	return d
}

// Page page
func (d *CellView) Page(ctx *fasthttp.RequestCtx) error {
	return nil
}

// // Auth page
// func (d *CellView) Auth(ctx *fasthttp.RequestCtx) error {
// 	return nil
// }
