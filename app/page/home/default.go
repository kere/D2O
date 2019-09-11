package home

import (
	"github.com/kere/gno/httpd"
	"onqee.visualstudio.com/D2O/app/page"
)

const (
	homeDir = "home"
)

// Default page class
type Default struct {
	httpd.P
}

// NewDefault func
func NewDefault() *Default {
	d := &Default{}
	d.Init("Default Page", "Default", homeDir)
	page.Init(&d.PA, page.Option{HasVue: true, HasHeader: true, HasFooter: true})

	return d
}

// // Page page
// func (d *Default) Page(ctx *fasthttp.RequestCtx) error {
// 	// time.Sleep(3 * time.Second)
// 	return nil
// }

// // Auth page
// func (d *Default) Auth(ctx *fasthttp.RequestCtx) error {
// 	return nil
// }
