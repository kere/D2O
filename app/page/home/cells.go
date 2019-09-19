package home

import (
	"github.com/kere/gno/httpd"
	"onqee.visualstudio.com/D2O/app/page"
)

// Cells page class
type Cells struct {
	httpd.P
}

// NewCells func
func NewCells() *Cells {
	d := &Cells{}
	d.Init("内容列表", "Cells", homeDir)
	page.Init(&d.PA, httpd.PageOption{HasVue: true, HasHeader: true})

	return d
}

// // Page page
// func (d *Cells) Page(ctx *fasthttp.RequestCtx) error {
// 	return nil
// }

// // Auth page
// func (d *Cells) Auth(ctx *fasthttp.RequestCtx) error {
// 	return nil
// }
