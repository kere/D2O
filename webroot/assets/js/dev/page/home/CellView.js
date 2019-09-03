require.config(requireOpt);
require(
	['ajax', 'util', 'marked', 'purify'],
	function (ajax, util, marked, DOMPurify){
		var tempDiv = document.createElement("div");
		function HTMLEncode(html) {
			(tempDiv.textContent != null) ? (tempDiv.textContent = html) : (tempDiv.innerText = html);
			var output = tempDiv.innerHTML;
			return output;
		}

    let iid = util.getRouterParam(0);

		ajax.NewClient("/api/app").getData("SElemByIID", {iid: iid}).then((dat) => {
			let contents = dat.o_json.contents;
			let content = contents[0];
			let $item;
			document.getElementById('headerTitle').innerText = DOMPurify.sanitize(content.title);
			document.getElementById('content').innerHTML = DOMPurify.sanitize(marked(content.text, {headerIds: false}));

			// subforms -------
			let subforms = dat.o_json.subforms, html='', form, items, th1, th2;
			for (let i = 0; i < subforms.length; i++) {
				form = subforms[i];
				html += '<table class="border col-xs-12 col-sm-6"><thead>'

				th1 = false;th2 = false;
				if(form.date_on && form.date_on.length>0){
					th1 = HTMLEncode(util.date2str(form.date_on, 'date'));
				}
				if(form.title && form.title.length>0){
					th2 = HTMLEncode(form.title);
				}
				if(th1 && th2){
					html += '<th class="date_on col1">'+th1+'</th>';
					html += '<th class="sub-title col1">'+th2+'</th>';
				}else{
					html += '<tr><th class="col1">数据名称</th><th>数值</th></tr>';
				}
				html += '</tr></thead>'

				// items
				items = form.items;
				for (let k = 0; k < items.length; k++) {
					html += '<tr><th class="col1">' +  HTMLEncode(items[k].k) + '</th><td>' + HTMLEncode(items[k].v) + '</td></tr>'
				}
				html += '</table>'
			}
			document.getElementById('subforms').innerHTML = html;


			// images -------
			// let images = dat.o_json.images;
			// html = '';
			// for (let i = 0; i < images.length; i++) {
			// 	html += '<img src="'+images[i].url+'"/>'
			// }
			// document.getElementById('images').innerHTML = html;

			// tags -------
			ajax.NewClient("/api/info").getData("Base").then((baseinfo) => {
				let alltags = ajax.torows(baseinfo.tags);
				let tagEle = [], obj, span;
				let $tags = document.getElementById('tags');
				for (var i = 0; i < dat.tags.length; i++) {
					obj = util.find('iid', dat.tags[i], alltags);
					if(!obj) continue;
					span = document.createElement('span');
					span.className='el-tag el-tag--light';
					span.innerText = obj.name;
					$tags.appendChild(span);
				}

			});

	  });

		window.closePageLoad();
	}
);
