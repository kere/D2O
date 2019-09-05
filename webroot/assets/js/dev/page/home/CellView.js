require.config(requireOpt);
require(
	['util', 'marked', 'purify'],
	function (util, marked, DOMPurify){
		var tempDiv = document.createElement("span");
		function HTMLEncode(html) {
			(tempDiv.textContent != null) ? (tempDiv.textContent = html) : (tempDiv.innerText = html);
			var output = tempDiv.innerHTML;
			return output;
		}

    // let iid = util.getRouterParam(0);
		let dat = pagedata;


		let contents = dat.o_json.contents;
		let content = contents[0];
		document.getElementById('headerTitle').innerText = DOMPurify.sanitize(content.title);
		document.getElementById('content').innerHTML = DOMPurify.sanitize(marked(content.text, {headerIds: false}));


		// subforms -------
		let subforms = dat.o_json.subforms, html='', form, items, th1, th2;
		let alllinks  = [];
		for (let i = 0; i < subforms.length; i++) {
			// 取出“参考连接”,然后统一至于页面
			form = subforms[i];
			items = form.items;
			for (let k = 0; k < items.length; k++) {
				if(items[k].k != '参考连接') continue;
				alllinks.push(items[k]);
				items[k] = false;
			}
		}

		let formsHTML;
		html = '';

		for (let i = 0; i < subforms.length; i++) {
			form = subforms[i];
			formsHTML = '<table class="border col-xs-12 col-sm-6"><thead>'

			th1 = false;th2 = false;
			if(form.date_on && form.date_on.length>0){
				th1 = HTMLEncode(util.date2str(form.date_on, 'date'));
			}
			if(form.title && form.title.length>0){
				th2 = HTMLEncode(form.title);
			}
			if(th1 && th2){
				formsHTML += '<th class="date_on col1">'+th1+'</th>';
				formsHTML += '<th class="sub-title col1">'+th2+'</th>';
			}else{
				formsHTML += '<tr><th class="col1">数据名称</th><th>数值</th></tr>';
			}
			formsHTML += '</tr></thead>'

			// items
			let isOK = false;
			items = form.items;
			for (let k = 0; k < items.length; k++) {
				if(!items[k]) continue
				formsHTML += '<tr><th class="col1">' +  HTMLEncode(items[k].k) + '</th><td>' + HTMLEncode(items[k].v) + '</td></tr>';
				isOK = true;
			}
			formsHTML += '</table>';

			if(isOK){
				html += formsHTML;
			}
		}

		html += '<h3>参考连接：</h3><ul class="gno-ref-links">';
		for (var i = 0; i < alllinks.length; i++) {
			html += DOMPurify.sanitize('<li><a href="'+alllinks[i].v+'">' + alllinks[i].v + '</a></li>');
		}
		html += '</ul>';
		document.getElementById('subforms').innerHTML = html;


		// images -------
		// let images = dat.o_json.images;
		// html = '';
		// for (let i = 0; i < images.length; i++) {
		// 	html += '<img src="'+images[i].url+'"/>'
		// }
		// document.getElementById('images').innerHTML = html;

		// tags -------
		html = '';
		var tags = dat.tags;
		for (var i = 0; i < tags.length; i++) {
			html += '<span class="el-tag el-tag--light"><a href="">' + HTMLEncode(tags[i]) + '</a></span>';
		}
		document.getElementById('tags').innerHTML = html;

		window.closePageLoad();
	}
);
