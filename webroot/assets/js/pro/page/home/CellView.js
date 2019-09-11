require.config(requireOpt);
require(
	['util', 'marked', 'purify'],
	function (util, marked, DOMPurify){

		if(!pagedata){
			document.title = "没有找到相应的数据";
			// document.getElementById('headerTitle').innerText = document.title;
			document.getElementById('headerTitle').innerText = "记录不存在，可能已经被删除了。";
			return;
		}

		var tempDiv = document.createElement("span");
		function HTMLEncode(html) {
			(tempDiv.textContent != null) ? (tempDiv.textContent = html) : (tempDiv.innerText = html);
			var output = tempDiv.innerHTML;
			return output;
		}

    // let iid = util.getRouterParam(0);
		let dat = pagedata;
		dat.o_json = JSON.parse(dat.o_json);

		let contents = dat.o_json.contents;
		let content = contents[0];
		document.getElementById('headerTitle').innerText = DOMPurify.sanitize(content.title);
		document.getElementById('content').innerHTML = DOMPurify.sanitize(marked(content.text, {headerIds: false}));
		document.title = content.title;


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
			formsHTML = '<ul class="data-list m-b">'

			th1 = false;th2 = false;
			if(form.date_on && form.date_on.length>0){
				th1 = HTMLEncode(util.date2str(form.date_on, 'date'));
			}
			if(form.title && form.title.length>0){
				th2 = HTMLEncode(form.title);
			}
			if(th1 || th2){
				formsHTML += '<li class="sub-title">';
				if(th1)
					formsHTML += '<strong class="date_on m-r">' + th1 + '</strong>';
				if(th2)
					formsHTML += '<strong class="title">' + th2 + '</strong>';
				formsHTML += '</li>'
			}

			// items
			let isOK = false;
			items = form.items;
			for (let k = 0; k < items.length; k++) {
				if(!items[k]) continue
				formsHTML += '<li><strong class="m-r-sm">' +  HTMLEncode(items[k].k) + ':</strong><span>' + HTMLEncode(items[k].v) + '</span></li>';
				isOK = true;
			}
			formsHTML += '</ul>';

			if(isOK){
				html += formsHTML;
			}
		}

		if(alllinks.length>0){
			html += '<h3>参考连接：</h3><ul class="gno-ref-links">';
			for (var i = 0; i < alllinks.length; i++) {
				html += DOMPurify.sanitize('<li><a href="'+alllinks[i].v+'">' + alllinks[i].v + '</a></li>');
			}
			html += '</ul>';
		}
		document.getElementById('subforms').innerHTML = html;

		// tags -------
		if(dat.tags){
			html = '';
			var tags = dat.tags;
			for (var i = 0; i < tags.length; i++) {
				html += '<span class="el-tag el-tag--light"><a href="">' + HTMLEncode(tags[i]) + '</a></span>';
			}
			document.getElementById('tags').innerHTML = html;
		}

	}
);
