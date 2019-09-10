define('preparecookie', ["ajax","util","accto"], function(ajax, util, accto){
	let cookies = [];
  let date = new Date();
  date.setTime(date.getTime()+30*86400000);
  let expires = "; expires="+date.toUTCString();

	let uiid = window.localStorage.getItem("_uiid");
	if(!uiid){
		let ts = ajax.serverTime.utctime().toString();
		uiid = accto(ts + navigator.userAgent+ts+(window.accpt?window.accpt:"") + window.location.hostname);
		window.localStorage.setItem("_uiid", uiid);
		cookies.push("_uiid="+uiid+expires+"; path=/");
	}
  let nick = window.localStorage.getItem("_nick");
	if(nick && !util.hasCookie('_nick')){
  	let token = window.localStorage.getItem("_token");
		cookies.push("_nick="+nick+expires+"; path=/")
	}
	
  let token = window.localStorage.getItem("_token");
	if(token && !util.hasCookie('_token')){
  	let token = window.localStorage.getItem("_token");
		cookies.push("_token="+token+expires+"; path=/")
	}else if(!token && util.hasCookie('_token')){
		token = util.getCookie('_token');
		window.localStorage.setItem("_token", token);
	}
	for (var i = 0; i < cookies.length; i++) {
		document.cookie = cookies[i];
	}

	return {};
});
