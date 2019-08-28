package main

import (
	"flag"

	_ "github.com/lib/pq"

	"github.com/kere/gno/httpd"
	"onqee.visualstudio.com/D2O/app/api"
	"onqee.visualstudio.com/D2O/app/page/home"
	"onqee.visualstudio.com/D2O/app/upload"
)

var (
	cf string
)

func main() {
	flag.Parse()
	flag.StringVar(&cf, "conf", "app/app.conf", "conf file name")
	httpd.Init(cf)

	// home
	httpd.Site.RegistGet("/", home.NewDefault())
	httpd.Site.RegistGet("/cell", home.NewCell())

	httpd.Site.RegistOpenAPI("/api/app", api.NewApp())
	httpd.Site.RegistUpload("/upload/image", upload.NewImage())

	httpd.Site.Start()
}
