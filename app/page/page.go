package page

import (
	"fmt"
	"io/ioutil"

	"github.com/kere/gno/httpd"
)

var (
	rqs string
)

// Option page
type Option struct {
	HasElement bool
	HasVue     bool
	HasHeader  bool
	HasFooter  bool
	NoPageLoad bool
}

// Init page
func Init(pd *httpd.PageData, opt Option) {
	siteConf := httpd.Site.C.GetConf("site")

	viewport := httpd.NewStrRender(`<meta name="viewport" content="width=device-width, initial-scale=1.0,minimum-scale=1.0,maximum-scale=1.0, user-scalable=no">`)

	pd.Head = []httpd.IRender{viewport, httpd.NewJSSrc(requireOpt(), nil)}
	pd.CSS = make([]httpd.IRenderWith, 0, 3)
	pd.JS = make([]httpd.IRenderWith, 0, 6)
	if len(pd.Bottom) == 0 {
		pd.Bottom = make([]httpd.IRender, 0, 4)
	}

	if !opt.NoPageLoad {
		pd.Top = []httpd.IRender{httpd.NewStrRender(httpd.PageLoadOpen)}
	}
	// vue
	if opt.HasVue {
		pd.JS = append(pd.JS, httpd.NewJS(siteConf.DefaultString("vuejs", "vue.min.js")))
	}

	// element-ui
	if opt.HasElement {
		pd.CSS = append(pd.CSS, httpd.NewCSS(siteConf.Get("elementcss")))
		pd.JS = append(pd.JS, httpd.NewJS(siteConf.Get("elementjs")))
	}

	pd.CSS = append(pd.CSS, httpd.NewCSS("main.css"))
	pd.JS = append(pd.JS, httpd.RequireJS(pd, requireJS()))

	pd.JSPosition = httpd.JSPositionBottom

	if opt.HasHeader {
		pd.Top = append(pd.Top, httpd.NewTemplate("_header.htm"))
	}

	if !opt.NoPageLoad {
		pd.Bottom = append(pd.Bottom, httpd.NewStrRender(httpd.PageLoadClose))
	}
	if opt.HasFooter {
		pd.Bottom = append(pd.Bottom, httpd.NewTemplate("_bottom.htm"))
	}

	// pd.CacheOption.PageMode = httpd.CacheModePagePath
	if httpd.RunMode == httpd.ModeDev {
		pd.CacheOption.Store = httpd.CacheStoreNone
	} else {
		pd.CacheOption.PageMode = httpd.CacheModePage
		pd.CacheOption.Store = httpd.CacheStoreFile
		pd.CacheOption.HTTPHead = 1
	}
}

const (
	requireOptStrDev = `{
  waitSeconds :30,
  baseUrl : "/assets/js/",
  paths: {
    echarts : "echarts.min",
    util : "%s/mylib/util",
    zepto : "%s/mylib/zepto",
    accto : "%s/mylib/accto",
    tool : "%s/mylib/tool",
    ajax : "%s/mylib/ajax"
  }
}`
	requireOptStrPro = `{waitSeconds:15,baseUrl:"/assets/js/",paths:{echarts:"echarts.min"}}`
)

var requirejs []byte

func requireJS() []byte {
	if len(requirejs) > 0 {
		return requirejs
	}

	var err error
	requirejs, err = ioutil.ReadFile("./webroot/assets/js/require.js")
	if err != nil {
		panic(err)
	}
	return requirejs
}

func requireOpt() string {
	if rqs == "" {
		if httpd.RunMode == httpd.ModeDev {
			rqs = fmt.Sprintf(requireOptStrDev, httpd.ModeDev, httpd.ModeDev, httpd.ModeDev, httpd.ModeDev, httpd.ModeDev)
		} else {
			rqs = requireOptStrPro
		}
		rqs = "var requireOpt=" + rqs
	}

	return rqs
}
