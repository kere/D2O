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
	// home
	site.RegistGet("/", home.NewDefault())
	site.RegistGet("/cell", home.NewCell())

	site.RegistOpenAPI("/api/app", api.NewApp())
	site.RegistOpenAPI("/api/info", api.NewBaseInfo())

	site.RegistUpload("/upload/image", upload.NewImage())

	confServer(site.Server, site.C.GetConf("site"))

	site.Start()
}

func confServer(s *fasthttp.Server, a conf.Conf) {
	s.MaxConnsPerIP = a.DefaultInt("max_conns_per_ip", 1000)
	s.Concurrency = a.DefaultInt("concurrency", 2048*2)
	// MaxRequestBodySize 最大body限制
	s.MaxRequestBodySize = a.DefaultInt("max_request_body_size", 10*1024*1024)
}
