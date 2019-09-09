require.config(requireOpt);
require(
	['ajax', 'util', 'accto', 'preparecookie'],
	function (ajax, util, accto){

		let uiid = window.localStorage.getItem("_uiid");
		let client = ajax.NewClient("/api/info");
		console.log(util.getCookie('_token'));
		client.send("IsLogin").then((dat) =>{
			console.log(dat);
		})

		document.getElementById('btnLogin').addEventListener('click', ()=>{
			let ts = ajax.serverTime.utctime().toString();
			let nick = document.getElementById('txtNick').value, pwd = document.getElementById('txtPwd').value;
			// ts + md5(pwd) + ts + pageToken + uiid + ts
			let aa = accto(ts + accto(pwd) + ts + window.accpt + uiid + ts);
			let obj = {
				nick: nick,
				src : util.dasit(aa),
				sign:util.dasit(accto(aa+ts))
			}
			client.send("DoUserLogin", obj, {ts: ts, busy: "#loginDiv" }).then((dat) =>{
				console.log(dat);
				if(dat.value){
					window.localStorage.setItem("_nick", nick);
				}
			})

		})

    window.closePageLoad();
	}
);
