package page

import (
	"github.com/kere/gno/httpd"
)

var (
	rqs string
)

// Init page
func Init(pa *httpd.PageAttr, opt httpd.PageOption) {
	httpd.PageInit(pa, opt)
	pa.Head = append(pa.Head, httpd.NewJSSrc(requireOpt(), nil))

	pa.CacheOption.PageMode = httpd.CacheModePage
	pa.CacheOption.Store = httpd.CacheStoreFile
	pa.CacheOption.HTTPHead = 1
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
