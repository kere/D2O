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
func Init(pa *httpd.PageAttr, opt Option) {
	siteConf := httpd.Site.C.GetConf("site")

	viewport := httpd.NewStrRender(`<meta name="viewport" content="width=device-width, initial-scale=1.0,minimum-scale=1.0,maximum-scale=1.0, user-scalable=no">`)

	pa.Head = make([]httpd.IRender, 3, 5)
	pa.CSS = make([]httpd.IRenderA, 0, 3)
	pa.JS = make([]httpd.IRenderA, 0, 5)
	pa.Top = make([]httpd.IRender, 0, 2)
	// if len(pa.Bottom) == 0 {
	pa.Bottom = make([]httpd.IRender, 0, 4)
	// }

	pa.Head[0] = viewport
	pa.Head[1] = httpd.NewJSSrc(requireOpt(), nil)
	pa.Head[2] = httpd.FaviconRender

	if !opt.NoPageLoad {
		pa.Top = append(pa.Top, httpd.NewStrRender(httpd.PageLoadOpen))
	}

	// vue
	if opt.HasVue {
		pa.JS = append(pa.JS, httpd.NewJS(siteConf.DefaultString("vuejs", "vue.min.js")))
	}

	// element-ui
	if opt.HasElement {
		pa.CSS = append(pa.CSS, httpd.NewCSS(siteConf.Get("elementcss")))
		pa.JS = append(pa.JS, httpd.NewJS(siteConf.Get("elementjs")))
	}

	pa.CSS = append(pa.CSS, httpd.NewCSS("main.css"))

	if !opt.NoRequireJS {
		pa.JS = append(pa.JS, httpd.RequireJSWithSrc(pa, requireJS()))
	}

	pa.JSPosition = httpd.JSPositionBottom

	if opt.HasHeader {
		pa.Top = append(pa.Top, httpd.NewTemplate("app/view/_header.htm"))
	}

	if !opt.NoPageLoad {
		pa.Bottom = append(pa.Bottom, httpd.NewStrRender(httpd.PageLoadClose))
	}
	if opt.HasFooter {
		pa.Bottom = append(pa.Bottom, httpd.NewTemplate("app/view/_footer.htm"))
	}

	// pa.CacheOption.PageMode = httpd.CacheModePagePath
	if httpd.RunMode == httpd.ModeDev {
		pa.CacheOption.Store = httpd.CacheStoreNone
	} else {
		pa.CacheOption.PageMode = httpd.CacheModePage
		pa.CacheOption.Store = httpd.CacheStoreFile
		pa.CacheOption.HTTPHead = 1
	}
}

const (
	requireOptStrDev = `{
  waitSeconds :30,
  baseUrl : "/assets/js/",
  paths: {
    echarts : "echarts.min",
		compressor: "dev/compressor",
    preparecookie : "dev/mylib/preparecookie",
    util : "dev/mylib/util",
    zepto : "dev/mylib/zepto",
    accto : "dev/mylib/accto",
    tool : "dev/mylib/tool",
    ajax : "dev/mylib/ajax"
  }
}`
	requireOptStrPro = `{waitSeconds:15,baseUrl:"/assets/js/",paths:{echarts:"echarts.min", compressor:"pro/compressor"}}`
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
