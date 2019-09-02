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
	d.D.Title = []byte("CellView Page")
	d.D.Name = "CellView"
	d.D.Dir = homeDir

	page.Init(&d.D, false)

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
