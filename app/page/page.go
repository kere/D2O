package page

import (
	"fmt"
	"io/ioutil"

	"github.com/kere/gno/httpd"
	"github.com/kere/gno/httpd/render"
	"github.com/kere/gno/libs/util"
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
}

// Init page
func Init(pd *httpd.PageData, opt Option) {
	siteConf := httpd.Site.C.GetConf("site")

	viewport := render.NewHead(`<meta name="viewport" content="width=device-width, initial-scale=1.0,minimum-scale=1.0,maximum-scale=1.0, user-scalable=no">`)

	// requirejs
	data := make(map[string]string, 0)
	data["defer"] = ""
	data["async"] = "true"
	data["data-main"] = util.PathToURL("/assets/js/", httpd.RunMode+"/page", pd.Dir, pd.Name)
	// data["src"] = "/assets/js/require.js"

	pd.Head = []render.IRender{viewport, render.NewScript(requireOpt())}
	pd.CSS = make([]render.IRenderWith, 0, 3)
	pd.JS = make([]render.IRenderWith, 0, 6)

	if opt.HasVue {
		pd.JS = append(pd.JS, render.NewJS(siteConf.DefaultString("vuejs", "vue.min.js")))
	}

	if opt.HasElement {
		pd.CSS = append(pd.CSS, render.NewCSS(siteConf.Get("elementcss")))
		pd.JS = append(pd.JS, render.NewJS(siteConf.Get("elementjs")))
	}

	pd.CSS = append(pd.CSS, render.NewCSS("main.css"))
	pd.JS = append(pd.JS, render.ScriptB(requireJS(), data))

	pd.JSPosition = httpd.JSPositionBottom

	if opt.HasHeader {
		pd.Top = []render.IRender{render.NewTemplate("_header.htm")}
	}

	// pd.CacheOption.PageMode = httpd.CacheModePagePath
	if httpd.RunMode == httpd.ModeDev {
		pd.CacheOption.Store = httpd.CacheStoreNone
	} else {
		pd.CacheOption.PageMode = httpd.CacheModePage
		pd.CacheOption.Store = httpd.CacheStoreFile
		pd.CacheOption.HTTPHead = 1
	}

	if opt.HasFooter {
		pd.Bottom = []render.IRender{
			render.NewTemplate("_bottom.htm"),
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
    echarts : "echarts.min"
  }
}`
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
		rqs = "requireOpt=" + rqs
	}

	return rqs
}
