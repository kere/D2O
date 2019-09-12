define('util', [], function(){
	function indexOfSortedI(val, arr, b, e, desc){
		if(b == e){
			return arr[b] == val ? b : -1;
		}else if(b>e){
			return -1
		}
		let l = e-b+1, i = b+Math.floor(l/2), v = arr[i];

		if(v == val){
			return i;
		}else if(v > val){
			if(desc){
				return indexOfSortedI(val, arr, i+1, e)
			}
			// small zone
			return indexOfSortedI(val, arr, b, i-1)
		}else{
			if(desc){
				return indexOfSortedI(val, arr, b, i-1)
			}
			return indexOfSortedI(val, arr, i+1, e)
		}
	}

	// b begin; e end index
	function getSortedI(field, val, arr, b, e, desc) {
		if(b == e){
			return arr[b][field] == val ? b : -1;
		}else if(b>e){
			return -1
		}
		let l = e-b+1, i = b+Math.floor(l/2), v;
		v = arr[i][field]

		if(v == val){
			return i;
		}else if(v > val){
			if(desc){
				return getSortedI(field, val, arr, i+1, e)
			}
			// small zone
			return getSortedI(field, val, arr, b, i-1)
		}else{
			if(desc){
				return getSortedI(field, val, arr, b, i-1)
			}
			return getSortedI(field, val, arr, i+1, e)
		}
	}

	var util = {
		DATE_DAY : 86400000,
		DATE_HOUR : 3600000,
	  env : () => {
	    let o = {}
	    let agent = navigator.userAgent.toLowerCase();
	    if(/android/.test(agent)){
	      o.os = 'android';
	    }else if(/iphone|ipod|ipad|ios/.test(agent)) {
	      o.os = 'ios';
	    }else if(/windows/.test(agent)){
	      o.os = 'windows';
	    }
	    o.iswxwork = /wxwork/.test(agent);
	    o.inwx = /micromessenger/.test(agent);

	    return o;
	  },

		language : function() {
			let lang = this.getCookie('lang');
			if(!lang){
				lang = navigator.language || navigator.userLanguage;
			}
			return lang.replace('_', '-').toLowerCase();
		},

		getRouterParam : (index) => {
	    let arr = window.location.pathname.split('/'),
	      l = arr.length,
	      i = l-1-index;
	    if (i<0) return null;

	    return arr[i];
		},

	  num2str : (v, deci) => {
			if(!v) return '';
			if(typeof(v)== 'string') v= parseFloat(v);

			let isPad = false, arr;
			if(deci < 0){// 左补齐
				let s = v.toString();
				arr = v.toString().split('.')
				if(arr.length == 1) return s;
				deci = -deci;
				isPad = true;
			}else if(deci==0){
	      return v.toFixed(deci);
	    }

	    let deciV = Math.pow(10, deci);
	    let s = (Math.round(v * deciV)/deciV).toFixed(deci);
			if(!isPad) return s;

			let l = s.length;
			arr = new Array(l);
			for (let i = 0; i < l; i++) {
				arr[i] = s[i];
			}

			for (let i = l-1; i > -1; i--) {
				if(arr[i]!= '0') break;
				arr.pop();
			}

			return arr.join('');
	  },

		// n 长度，年月日小时分钟
		timeAgoStr : function(b, e, n) {
			if(!n) n = 2;
	    let arr = this.timeAgo(b, e), str='';
	    let l = arr.length, isskip = true, i, k=0;
	    for (i = l-1; i > -1; i--) {
	      if(isskip && arr[i].value==0) continue;
	      isskip = false;
	      str += arr[i].value + arr[i].label + ' ';
				k++
				if(k == n) break;
	    }
	    return str;
	  },

		timeAgo : (b, e) => {
	    let diff;
	    if(typeof(b)=='number' && typeof(e) == 'undefined'){
	      diff = b;
	    }else{
	      diff = Math.abs(e.getTime() - b.getTime());
	    }
	    let v, n, arr=[];

	    for(let i=0;i<6;i++) {
	      switch (i) {
	        case 0: // second
	          v = diff % 60000;
	          arr.push({value:Math.floor(v/1000), 'label':'秒', 'ext': 's'});
	          diff= Math.floor(diff/60000);
	          break;
	        case 1: // 分钟
	          v = diff % 60;
	          arr.push({value:v, 'label':'分', 'ext': 'm'});
	          diff= Math.floor(diff/60);
	          break;
	        case 2: //小时
	          v = diff % 24;
	          arr.push({value:v, 'label':'小时', 'ext': 'h'});
	          diff= Math.floor(diff/24);
	          break;
	        case 3: //天
	          v = diff % 30;
	          arr.push({value:v, 'label':'天', 'ext': 'd'});
	          diff= Math.floor(diff/30);
	          break;
	        case 4: //月
	          v = diff % 12;
	          arr.push({value:v, 'label':'月', 'ext': 'n'});
	          diff= Math.floor(diff/12);
	          break;
	        case 5:
	          arr.push({value:v, 'label':'年', 'ext': 'y'});
	          break;
	      }
	      if(Math.floor(diff) < 0) break;
	    }

	    return arr;
		},

		str2date : (str) => {
	    if(!str) return;
	    if(typeof str != 'string') return str;

			str = str.replace(/[A-Za-z日]/g, ' ').substr(0,19);
	    str = str.replace(/[年月]/g, '-');

	    let d = new Date(Date.parse(str));
	    if(!d || isNaN(d.getFullYear())){
	      str = str.replace(/[-]/g, '/');
	      return new Date(Date.parse(str));
	    }
	    return d;
		},

		date2str : function(time, ctype) {
			if(!time){
				return '';
			}
	    ctype = ctype ? ctype : 'date';

			switch(typeof(time)){
				case 'number':
					time = new Date(time);
					break;
				case 'string':
					time = this.str2date(time);
					break;
			}
			switch (ctype) {
				case 'date':
					return time.getFullYear()+'-'+this.lpad(time.getMonth()+1, '0', 2)+'-'+this.lpad(time.getDate(), '0', 2)
	      case 'date2':
	        return this.lpad(time.getMonth()+1, '0', 2)+'-'+this.lpad(time.getDate(), '0', 2)
	      case 'dateCH':
	        return time.getFullYear()+'年'+this.lpad(time.getMonth()+1, '0', 2)+'月'+this.lpad(time.getDate(), '0', 2) + '日'
	      case 'date2CH':
	        return this.lpad(time.getMonth()+1, '0', 2)+'月'+this.lpad(time.getDate(), '0', 2) + '日'
				case 'datetime':
					return time.getFullYear()+'-'+this.lpad(time.getMonth()+1, '0', 2)+'-'+this.lpad(time.getDate(), '0', 2)+' '+this.lpad(time.getHours(), '0', 2)+':'+this.lpad(time.getMinutes(), '0', 2)
				case 'time':
					return this.lpad(time.getHours(), '0', 2)+':'+this.lpad(time.getMinutes(), '0', 2)
			}
			return 'unknow'
		},

		lpad : (str, padString, l) => {
			if(typeof(str)!='string')
				str = str.toString();
			while (str.toString().length < l)
				str = padString + str;
			return str;
		},

		//pads right
		rpad : (str, padString, l) => {
			if(typeof(str)!='string')
				str = str.toString();
			while (str.toString().length < l)
				str = str + padString;
			return str;
		},

		tisvo : (s) => {
			let str = window.atob(s);
			let l = str.length,dat = "";
			for (let i = 0; i < l; i++)
				dat += String.fromCharCode(str.charCodeAt(i) ^ ((i%7 << 4 ) + (i%15)));
			return window.atob(dat);
		},

		find : (field, value, arr) => {
			if(!arr)
				return null;

			let i,len = arr.length
			for(i=0;i<len;i++)
				if(arr[i] && arr[i][field] == value){
					return arr[i];
				}
			return null;
		},

		indexOfSortedI : (val, arr) => {
			if(!arr) return -1;
			let isdesc = false;
			if(arr.length> 1)
				isdesc = arr[0] > arr[1];
			return indexOfSortedI(val, arr, 0, arr.length-1, isdesc);
		},

		findSortedI : (field, val, arr) => {
			if(!arr) return -1;
			let isdesc = false;
			if(arr.length> 1)
				isdesc = arr[0][field] > arr[1][field];
			return getSortedI(field, val, arr, 0, arr.length-1, isdesc);
		},

		findSorted : (field, val, arr) => {
			if(!arr) return null;
			let isdesc = false;
			if(arr.length> 1)
				isdesc = arr[0][field] > arr[1][field];
			let i = getSortedI(field, val, arr, 0, arr.length-1, isdesc);
			if(i < 0) return null;
			return arr[i];
		},

		findIndex : (field, value, data) => {
			if(!data)
				return -1;

			let i,len = data.length
			for(i=0;i<len;i++)
				if(data[i] && data[i][field] == value){
					return i;
				}
			return -1;
		},

		arrMix: (arr, all, callback) => {
			let finded = [];
			for (let i = 0; i < all.length; i++) {

				for (let k = 0; k < arr.length; k++) {
					if(finded.indexOf(k) !== -1) continue;
					if(callback(k, i)){
						finded.push(k)
					}
				}
			}
		},

		inArray : (val, arr) => {
			for(let i in arr){
				if(arr[i]==val)
					return true;
			}
			return false;
		},

	  setCookie :  (name,value,days) => {
			let expires = ""
      if (days) {
        let date = new Date();
        date.setTime(date.getTime()+days*86400000);
        expires = "; expires="+date.toGMTString();
      }
      document.cookie = encodeURIComponent(name)+"="+encodeURIComponent(value)+expires+"; path=/";
	  },
		hasCookie : (sKey)=>{
			return (new RegExp("(?:^|;\\s*)" + encodeURIComponent(sKey).replace(/[-.+*]/g, "\\$&") + "\\s*\\=")).test(document.cookie);
		},
		getCookie : (name) => {
			// return decodeURIComponent(document.cookie.replace(new RegExp("(?:(?:^|.*;)\\s*" + encodeURIComponent(sKey).replace(/[-.+*]/g, "\\$&") + "\\s*\\=\\s*([^;]*).*$)|^.*$"), "$1")) || null;
	    let nameEQ = name + "=";
	    let ca = document.cookie.split(';');
	    for(let i=0;i < ca.length;i++) {
	        let c = ca[i];
	        while (c.charAt(0)==' ') c = c.substring(1,c.length);
	        if (c.indexOf(nameEQ) == 0){
	          let s = c.substring(nameEQ.length,c.length);
	          if(s[0] == '"'){
	            return s.substring(1, s.length-1);
	          }
	          return c.substring(nameEQ.length,c.length);
	        }
	    }
	    return '';
	  },

	  deleteCookie : (name) => {
	      util.SetCookie(name,"",-1);
	  },

	  copy : (o) => {
	    let dat = {};
			if(typeof(o) == 'object' && o.hasOwnProperty('length')){
				dat = new Array(o.length);
			}
	    for (let k in o) {
	      if (!o.hasOwnProperty(k) || k.substr(0,2)=="__") continue;
				if(o[k]== null){
					dat[k] = null
					continue;
				}
	      if(typeof(o[k]) == 'object'){
	        dat[k] = util.copy(o[k])
	      }else{
	        dat[k] = o[k];
	      }
	    }
	    return dat;
	  },

		dasit : (str) => {
			let b64 = window.btoa(str), l = b64.length, dat="";
			for (let i = 0; i < l; i++)
				dat += String.fromCharCode(b64.charCodeAt(i) ^ ((i%7 << 4 ) + (i%15)));
			return window.btoa(dat);
		},

		$ : {
			get : function(sel) {
				if(typeof(sel)==='string'){
					return document.querySelector(sel);
				}
				return sel;
			},
			getAll : function(sel) {
				if(typeof(sel)==='string'){
					return document.querySelectorAll(sel);
				}
				if(!sel.length)
					return [sel];

				return sel;
			},

			each : function(el, f) {
				el = this.getAll(el);
				if(el.length > 0){
					for (let i = 0; i < el.length; i++) {
						f(el[i], i);
					}
					return;
				}
				return f(el, -1);
			},

			addClass : function(sel, clas) {
				this.each(sel, (el) =>{
					if(this.hasClass(el, clas)) return;
					if(el.className.length==0){
						el.className = clas;
					}else{
						el.className += ' ' +clas;
					}
				});
			},

			hasClass : function(sel, clas) {
				let isfound = false;
				this.each(sel, (el) =>{
					if(!el.className){
						return false;
					}
					let arr = el.className.split(' ');
					for (let i = 0; i < arr.length; i++) {
						if(arr[i]==clas) {
							isfound = true;
							return;
						};
					}
				})
				return isfound;
			},

			removeClass : function(sel, clas)  {
				this.each(sel, (el) =>{
					if(!el.className){
						return;
					}
					let arr = el.className.split(' ');
					for (let i = 0; i < arr.length; i++) {
						if(arr[i]==clas) {
							arr.splice(i, 1);
						};
					}
					el.className = arr.join(' ');
				});
			},

			show : function(el){
				if(this.hasClass(el, 'hide')){
					this.removeClass(el, 'hide');

					if(this.hasClass(el, 'fade')){
						var ths = this;
						setTimeout(function(){
							ths.addClass(el, 'in');
						}, 50)
					}
				}

				if(el.style.display == 'none'){
					el.style.display = '';
				}
			},

			hide : function(el){
				if(this.hasClass(el, 'fade')){
					this.removeClass(el, 'in');
					var ths = this;
					setTimeout(function(){
						ths.addClass(el, 'hide');
					}, 200)
				}else{
					this.addClass(el, 'hide');
				}
			}
		}
	}

	util.tool = {
		pageErr(title, msg){
			util.$.each(document.body.children, (e) =>{
				if(e.tagName == "SCRIPT") return;
				e.hidden = true
			})

			let el = document.createElement("DIV");
			el.id="pageError"
			el.className = 'page-not-found-div'
			el.innerHTML = '<h1>'+title+'</h1><h5>'+msg+'</h5>'
			document.body.appendChild(el);

			let style = document.createElement("style");
			style.type = "text/css";
			style.innerHTML = `html,body{height: 100%;margin: 0;padding: 0}
.page-not-found-div{line-height:4rem;text-align: center;margin: 0 auto;position: absolute;top: 36%;width: 100%}
.page-not-found-div h1{font-size: 82px}.page-not-found-div h5{font-weight: 200;font-size: 26px;}
body{font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif; color: #343434; }`
			document.getElementsByTagName("HEAD").item(0).appendChild(style);
		},
		showLoading : function(n){
			this.loading(0, n);
		},
		showSuccess : function(n){
			this.loading(1, n);
		},
		loading : function(itype, n){
			let el = util.$.get('#toast');
			if(!el){
				el = document.createElement("DIV");
				el.id = 'toast';
				el.className = 'weui-toast hide';
				el.innerHTML = `<div class="weui-mask"></div><div class="weui-toast">
  <i class="toast-success weui-icon-success-no-circle weui-icon_toast hide"></i>
  <p class="toast-success weui-toast__content hide">已完成</p>
  <i class="toast-loading weui-loading weui-icon_toast hide"></i>
  <p class="toast-loading weui-toast__content hide">数据加载中</p>
</div>`
				document.body.appendChild(el);
			}

			util.$.show(el);
			if(itype){
				util.$.each(el.querySelector('.toast-loading'), e => {
					util.$.hide(e)
				})
				util.$.each(el.querySelector('.toast-success'), e => {
					util.$.show(e)
				})
			}else{
				util.$.each(el.querySelector('.toast-success'), e => {
					util.$.hide(e)
				})
				util.$.each(el.querySelector('.toast-loading'), e => {
					util.$.show(e)
				})
			}

			n = n || 1;
			setTimeout(function(){
				if(itype==1){
					util.$.each(el.querySelector('.toast-success'), e => {
						util.$.hide(e)
					})
				}else{
					util.$.each(el.querySelector('.toast-loading'), e => {
						util.$.hide(e)
					})
				}
				util.$.hide(el);
			}, n * 1000);
		},

		hideToast : ()=>{
			let el = util.$.get('#toast');
			util.$.each(el.querySelector('.toast-success'), e => {
				util.$.hide(e)
			})
			util.$.each(el.querySelector('.toast-loading'), e => {
				util.$.hide(e)
			})
			util.$.hide(el);
		},

		showBusy : function(el, n){
			el = util.$.get(el);
			let t = el.querySelector('div.el-loading-mask');
			if(!t){
				t = document.createElement("DIV");
				t.className = 'el-loading-mask';
				t.innerHTML = `<div class="el-loading-spinner"><svg viewBox="25 25 50 50" class="circular"><circle cx="50" cy="50" r="20" fill="none" class="path"></circle></svg></div>`
				el.appendChild(t);
			}

			util.$.show(t);
			n = n || 3;
			setTimeout(function(){
				util.$.hide(t);
			}, n * 1000)
		},

		hideBusy : function(el){
			el = util.$.get(el);
			let t = el.querySelector('div.el-loading-mask');
			if(!t) return;
			util.$.hide(t);
		},

	  viewImage : function(url) {
			let el = util.$.get('#viewImage');
	    if(!el){
				el = document.createElement("DIV");
				el.id = 'viewImage';
				el.className = 'fade hide';
				el.innerHTML = '<div style="z-index:1000;position:fixed;width:100%;height:100%;top:0;left:0;text-align:center;top:0;left:0;"><div style="width:100%;height:100%;background: #000;opacity: 0.6;" class="view-image fade in"></div><img src="" style="z-index:1001;max-width:98%;max-height:98%;transform: translate(-50%, -50%);top:50%;position:absolute;"></div>';
				document.body.appendChild(el);

	      el.addEventListener('click', function(e){
					util.$.hide(el);
	      })
	    }

			let img = el.querySelector('img');
			img.src = url;
			util.$.show(el);
	  },

	  taggle : function(e){
	    let el = e.currentTarget;
	    if(e.currentTarget){
				if(util.$.hasClass(el, 'parent-pp')){
					el = el.parentElement.parentElement;
				}else if(util.$.hasClass(el, 'parent-p')){
					el = el.parentElement;
				}
	    }else{
	      el = e;
	    }
			let box = el.nextElementSibling;
			if(util.$.hasClass(el, 'open')){
				util.$.hide(box);
				util.$.removeClass(el, 'open');
				return false;
			}
			util.$.show(box);
			util.$.addClass(el, 'open');
			return true;
	  }

	}

	if(window.closePageLoad)
  	window.closePageLoad();

	return util;
});
