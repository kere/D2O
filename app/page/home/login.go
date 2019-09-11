package home

import (
	"github.com/kere/gno/httpd"
	"onqee.visualstudio.com/D2O/app/page"
)

// Login page class
type Login struct {
	httpd.P
}

// NewLogin func
func NewLogin() *Login {
	d := &Login{}
	d.Init("用户登录", "Login", homeDir)
	page.Init(&d.PA, page.Option{HasHeader: true})

	return d
}
