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
func NewCell(isEdit bool) *Cell {
	d := &Cell{}
	if isEdit {
		d.Init("编辑信息", "Cell", homeDir)
	} else {
		d.Init("新建信息", "Cell", homeDir)
	}
	page.Init(&d.PA, page.Option{HasVue: true, HasElement: true})

	return d
}

// Auth page
func (d *Cell) Auth(ctx *fasthttp.RequestCtx) error {
	return page.Auth(ctx)
}
