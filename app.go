package main

import (
	"flag"

	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"

	"github.com/kere/gno/httpd"
	"github.com/kere/gno/libs/conf"
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

	site := httpd.Site
	c := site.C.GetConf("site")
	upload.StoreDir = "webroot" + c.DefaultString("upload_img_dir", "/upload")

	// page
	site.RegistGet("/", home.NewDefault())
	site.RegistGet("/login", home.NewLogin())

	site.RegistGet("/cell/new", home.NewCell(false))
	site.RegistGet("/cell/edit/:iid", home.NewCell(true))
	site.RegistGet("/cell/list", home.NewCells())

	// page cached
	site.RegistGet("/cell/view/:iid", home.GetCellView())

	// api
	site.RegistOpenAPI("/api/app", api.NewApp())
	site.RegistOpenAPI("/api/info", api.NewBaseInfo())

	// upload
	site.RegistUpload("/upload/image", upload.NewImage())

	confServer(site.Server, c)

	site.Start()
}

func confServer(s *fasthttp.Server, a conf.Conf) {
	s.MaxConnsPerIP = a.DefaultInt("max_conns_per_ip", 1000)
	// s.Concurrency = a.DefaultInt("concurrency", 2048*2)
	// MaxRequestBodySize 最大body限制
	s.MaxRequestBodySize = a.DefaultInt("max_request_body_size", 10*1024*1024)
}
