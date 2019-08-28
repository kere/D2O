package page

import (
	"fmt"

	"github.com/kere/gno/httpd"
	"github.com/kere/gno/httpd/render"
	"github.com/kere/gno/libs/conf"
	"github.com/kere/gno/libs/util"
)

var (
	siteConf conf.Conf
	rqs      string
)

// Init page
func Init(pd *httpd.PageData, isElement bool) {
	if siteConf == nil {
		siteConf = httpd.GetConfig().GetConf("site")
	}

	if isElement {
		pd.Head = []render.IRender{
			render.NewCSS("main.css"),
			render.NewCSS(siteConf.Get("elementcss")),
			render.NewScript(requireOpt()),
		}
	} else {
		pd.Head = []render.IRender{
			render.NewCSS("main.css"),
			render.NewScript(requireOpt()),
		}
	}

	pd.Top = []render.IRender{render.NewTemplate("_header.htm")}

	// pd.CacheOption.PageMode = httpd.CacheModePagePath
	if httpd.RunMode == httpd.ModeDev {
		pd.CacheOption.Store = httpd.CacheStoreNone
	} else {
		pd.CacheOption.Store = httpd.CacheStoreMem
		pd.CacheOption.HTTPHead = 1
	}

	// requirejs
	data := make(map[string]string, 0)
	data["defer"] = ""
	data["async"] = "true"

	data["data-main"] = render.AssetsURL + util.PathToURL("/js/", httpd.RunMode+"/page", pd.Dir, pd.Name)
	data["src"] = "/assets/js/require.js"

	if isElement {
		pd.Bottom = []render.IRender{
			render.NewTemplate("_bottom.htm"),
			render.NewJS(siteConf.DefaultString("vuejs", "vue.min.js")),
			render.NewJS(siteConf.Get("elementjs")),
			render.Script("", data),
		}

	} else {
		pd.Bottom = []render.IRender{
			render.NewTemplate("_bottom.htm"),
			render.NewJS(siteConf.DefaultString("vuejs", "vue.min.js")),
			render.Script("", data),
		}
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
