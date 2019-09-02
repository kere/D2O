package page

import (
	"fmt"

	"github.com/kere/gno/httpd"
	"github.com/kere/gno/httpd/render"
	"github.com/kere/gno/libs/util"
)

var (
	rqs string
)

// Init page
func Init(pd *httpd.PageData, isElement bool) {
	siteConf := httpd.Site.C.GetConf("site")

	viewport := render.NewHead(`<meta name="viewport" content="width=device-width, initial-scale=1.0,minimum-scale=1.0,maximum-scale=1.0, user-scalable=no">`)

	// requirejs
	data := make(map[string]string, 0)
	data["defer"] = ""
	data["async"] = "true"

	data["data-main"] = util.PathToURL("/assets/js/", httpd.RunMode+"/page", pd.Dir, pd.Name)
	data["src"] = "/assets/js/require.js"

	pd.Head = []render.IRender{viewport, render.NewScript(requireOpt())}
	if isElement {
		pd.CSS = []render.IRenderWith{
			render.NewCSS(siteConf.Get("elementcss")),
			render.NewCSS("main.css"),
		}
		pd.JS = []render.IRenderWith{
			render.NewJS(siteConf.DefaultString("vuejs", "vue.min.js")),
			render.NewJS(siteConf.Get("elementjs")),
			render.Script("", data),
		}
	} else {
		pd.CSS = []render.IRenderWith{render.NewCSS("main.css")}
		pd.JS = []render.IRenderWith{
			render.NewJS(siteConf.DefaultString("vuejs", "vue.min.js")),
			render.Script("", data),
		}
	}

	pd.JSPosition = httpd.JSPositionBottom

	pd.Top = []render.IRender{render.NewTemplate("_header.htm")}

	// pd.CacheOption.PageMode = httpd.CacheModePagePath
	if httpd.RunMode == httpd.ModeDev {
		pd.CacheOption.Store = httpd.CacheStoreNone
	} else {
		pd.CacheOption.Store = httpd.CacheStoreMem
		pd.CacheOption.HTTPHead = 1
	}

	pd.Bottom = []render.IRender{
		render.NewTemplate("_bottom.htm"),
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
	requireOptStrPro = `{
  waitSeconds :15,
  baseUrl : "/assets/js/",
  paths: {
    "echarts" : "echarts.min"
  }
}`
)

func requireOpt() string {
	if rqs == "" {
		if httpd.RunMode == httpd.ModeDev {
			rqs = fmt.Sprintf(requireOptStrDev, httpd.ModeDev, httpd.ModeDev, httpd.ModeDev, httpd.ModeDev, httpd.ModeDev)
		} else {
			rqs = requireOptStrPro
		}
		rqs = "requireOpt=" + rqs
	}

	return rqs
}
