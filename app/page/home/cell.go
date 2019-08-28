package home

import (
	"github.com/kere/gno/httpd"
	"github.com/valyala/fasthttp"
	"onqee.visualstudio.com/D2O/app/page"
)

// Cell page class
type Cell struct {
	httpd.P
}

// NewCell func
func NewCell() *Cell {
	d := &Cell{}
	d.D.Title = []byte("Cell Page")
	d.D.Name = "Cell"
	d.D.Dir = homeDir

	page.Init(&d.D, true)

	return d
}

// Page page
func (d *Cell) Page(ctx *fasthttp.RequestCtx) error {
	// time.Sleep(3 * time.Second)
	return nil
}

// // Auth page
// func (d *Cell) Auth(ctx *fasthttp.RequestCtx) error {
// 	return nil
// }
