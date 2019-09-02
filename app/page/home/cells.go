package home

import (
	"github.com/kere/gno/httpd"
	"github.com/valyala/fasthttp"
	"onqee.visualstudio.com/D2O/app/page"
)

// Cells page class
type Cells struct {
	httpd.P
}

// NewCells func
func NewCells() *Cells {
	d := &Cells{}
	d.D.Title = []byte("Cells Page")
	d.D.Name = "Cells"
	d.D.Dir = homeDir

	page.Init(&d.D, page.Option{HasVue: true})

	return d
}

// Page page
func (d *Cells) Page(ctx *fasthttp.RequestCtx) error {
	return nil
}

// // Auth page
// func (d *Cells) Auth(ctx *fasthttp.RequestCtx) error {
// 	return nil
// }
