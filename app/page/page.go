package page

import (
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

	pd.Head = make([]httpd.IRender, 0, 5)
	pd.CSS = make([]httpd.IRenderWith, 0, 3)
	pd.JS = make([]httpd.IRenderWith, 0, 5)
	pd.Top = make([]httpd.IRender, 0, 2)
	// if len(pd.Bottom) == 0 {
	pd.Bottom = make([]httpd.IRender, 0, 4)
	// }

	pd.Head = append(pd.Head, viewport, httpd.NewJSSrc(requireOpt(), nil))

	if !opt.NoPageLoad {
		pd.Top = append(pd.Top, httpd.NewStrRender(httpd.PageLoadOpen))
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
		pd.Top = append(pd.Top, httpd.NewTemplate("app/view/_header.htm"))
	}

	if !opt.NoPageLoad {
		pd.Bottom = append(pd.Bottom, httpd.NewStrRender(httpd.PageLoadClose))
	}
	if opt.HasFooter {
		pd.Bottom = append(pd.Bottom, httpd.NewTemplate("app/view/_footer.htm"))
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
    preparecookie : "dev/mylib/preparecookie",
    util : "dev/mylib/util",
    zepto : "dev/mylib/zepto",
    accto : "dev/mylib/accto",
    tool : "dev/mylib/tool",
    ajax : "dev/mylib/ajax"
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
			rqs = requireOptStrDev
		} else {
			rqs = requireOptStrPro
		}
		rqs = "var requireOpt=" + rqs
	}

	return rqs
}

// EchojsRender 延迟加载
var EchojsRender = httpd.NewStrRender(`<script type="text/javascript">
(function(c,m){"function"===typeof define&&define.amd?define(function(){return m(c)}):"object"===typeof exports?module.exports=m:c.echo=m(c)})(this,function(c){var m,q,r,t,h={},n=function(){},e,u,v,p,f=function(){if(v||!e)clearTimeout(e),e=setTimeout(function(){h.render();e=null},u)};h.init=function(b){b=b||{};var g=b.offset||0,d=b.offsetVertical||g;g=b.offsetHorizontal||g;m=parseInt(b.offsetTop||d,10);q=parseInt(b.offsetBottom||d,10);r=parseInt(b.offsetLeft||g,10);t=parseInt(b.offsetRight||g,10);
u=parseInt(b.throttle||250,10);v=!1!==b.debounce;p=!!b.unload;n=b.callback||n;h.render();document.addEventListener?(c.addEventListener("scroll",f,!1),c.addEventListener("load",f,!1)):(c.attachEvent("onscroll",f),c.attachEvent("onload",f))};h.render=function(b){b=(b||document).querySelectorAll("[data-echo], [data-echo-background]");for(var g=b.length,d,a,f={l:0-r,t:0-m,b:(c.innerHeight||document.documentElement.clientHeight)+q,r:(c.innerWidth||document.documentElement.clientWidth)+t},e=0;e<g;e++){a=
b[e];var k=a;var l=f;null===k.offsetParent?l=!1:(k=k.getBoundingClientRect(),l=k.right>=l.l&&k.bottom>=l.t&&k.left<=l.r&&k.top<=l.b);l?(p&&a.setAttribute("data-echo-placeholder",a.src),null!==a.getAttribute("data-echo-background")?a.style.backgroundImage="url("+a.getAttribute("data-echo-background")+")":a.src!==(d=a.getAttribute("data-echo"))&&(/(?:javascript|jav\s+ascript|&#\d+|&#x)/i.test(d)?("img-desc"==a.nextSibling.className&&a.parentElement.removeChild(a.nextSibling),a.parentElement.removeChild(a)):
a.src=d),p||(a.removeAttribute("data-echo"),a.removeAttribute("data-echo-background")),n(a,"load")):p&&(d=a.getAttribute("data-echo-placeholder"))&&(null!==a.getAttribute("data-echo-background")?a.style.backgroundImage="url("+d+")":/(?:javascript|jav\s+ascript|&#\d+|&#x)/i.test(d)?("img-desc"==a.nextSibling.className&&a.parentElement.removeChild(a.nextSibling),a.parentElement.removeChild(a)):a.src=d,a.removeAttribute("data-echo-placeholder"),n(a,"unload"))}g||h.detach()};h.detach=function(){document.removeEventListener?
c.removeEventListener("scroll",f):c.detachEvent("onscroll",f);clearTimeout(e)};return h});
echo.init({ offset: 100, throttle: 250, unload: false,
 callback:function(element,op){}
});
</script>`)

// FooterViewEndRender render
var FooterViewEndRender = httpd.NewStrRender(`<footer id="page-footer"> ----- <small>end</small> ----- </footer>`)
