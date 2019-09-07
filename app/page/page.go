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
	HasElement  bool
	HasVue      bool
	HasHeader   bool
	HasFooter   bool
	NoPageLoad  bool
	NoRequireJS bool
}

// Init page
func Init(pd *httpd.PageData, opt Option) {
	siteConf := httpd.Site.C.GetConf("site")

	viewport := httpd.NewStrRender(`<meta name="viewport" content="width=device-width, initial-scale=1.0,minimum-scale=1.0,maximum-scale=1.0, user-scalable=no">`)

	pd.Head = make([]httpd.IRender, 0, 4)
	pd.CSS = make([]httpd.IRenderWith, 0, 3)
	pd.JS = make([]httpd.IRenderWith, 0, 6)
	if len(pd.Bottom) == 0 {
		pd.Bottom = make([]httpd.IRender, 0, 4)
	}

	pd.Head = append(pd.Head, viewport, httpd.NewJSSrc(requireOpt(), nil))

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

	if !opt.NoRequireJS {
		pd.JS = append(pd.JS, httpd.RequireJS(pd, requireJS()))
	}

	pd.JSPosition = httpd.JSPositionBottom

	if opt.HasHeader {
		pd.Top = append(pd.Top, httpd.NewTemplate("_header.htm"))
	}

	if !opt.NoPageLoad {
		pd.Bottom = append(pd.Bottom, httpd.NewStrRender(httpd.PageLoadClose))
	}
	if opt.HasFooter {
		pd.Bottom = append(pd.Bottom, httpd.NewTemplate("_footer.htm"))
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

// EchojsRender 延迟加载
var EchojsRender = httpd.NewStrRender(`<script type="text/javascript">
  !function(t,e){"function"==typeof define&&define.amd?define(function(){return e(t)}):"object"==typeof exports?module.exports=e:t.echo=e(t)}(this,function(t){"use strict";var e,n,o,r,c,a={},u=function(){},d=function(t){return null===t.offsetParent},l=function(t,e){if(d(t))return!1;var n=t.getBoundingClientRect();return n.right>=e.l&&n.bottom>=e.t&&n.left<=e.r&&n.top<=e.b},i=function(){(r||!n)&&(clearTimeout(n),n=setTimeout(function(){a.render(),n=null},o))};return a.init=function(n){n=n||{};var d=n.offset||0,l=n.offsetVertical||d,f=n.offsetHorizontal||d,s=function(t,e){return parseInt(t||e,10)};e={t:s(n.offsetTop,l),b:s(n.offsetBottom,l),l:s(n.offsetLeft,f),r:s(n.offsetRight,f)},o=s(n.throttle,250),r=n.debounce!==!1,c=!!n.unload,u=n.callback||u,a.render(),document.addEventListener?(t.addEventListener("scroll",i,!1),t.addEventListener("load",i,!1)):(t.attachEvent("onscroll",i),t.attachEvent("onload",i))},a.render=function(n){for(var o,r,d=(n||document).querySelectorAll("[data-echo], [data-echo-background]"),i=d.length,f={l:0-e.l,t:0-e.t,b:(t.innerHeight||document.documentElement.clientHeight)+e.b,r:(t.innerWidth||document.documentElement.clientWidth)+e.r},s=0;i>s;s++)r=d[s],l(r,f)?(c&&r.setAttribute("data-echo-placeholder",r.src),null!==r.getAttribute("data-echo-background")?r.style.backgroundImage="url("+r.getAttribute("data-echo-background")+")":r.src!==(o=r.getAttribute("data-echo"))&&(r.src=o),c||(r.removeAttribute("data-echo"),r.removeAttribute("data-echo-background")),u(r,"load")):c&&(o=r.getAttribute("data-echo-placeholder"))&&(null!==r.getAttribute("data-echo-background")?r.style.backgroundImage="url("+o+")":r.src=o,r.removeAttribute("data-echo-placeholder"),u(r,"unload"));i||a.detach()},a.detach=function(){document.removeEventListener?t.removeEventListener("scroll",i):t.detachEvent("onscroll",i),clearTimeout(n)},a});

	  echo.init({ offset: 100, throttle: 250, unload: false,
	    callback: function (element, op) {
	      // console.log(element, 'has been', op + 'ed')
	    }
	  });

</script>
`)

// FooterViewEndRender render
var FooterViewEndRender = httpd.NewStrRender(` <footer id="page-footer">
  ------- <small>end</small> -------
</footer>`)
