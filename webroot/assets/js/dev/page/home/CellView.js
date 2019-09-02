require.config(requireOpt);
require(
	['ajax', 'util', 'marked', 'purify'],
	function (ajax, util, marked, DOMPurify){
    let iid = util.getRouterParam(0);

		ajax.NewClient("/api/app").getData("SElemByIID", {iid: iid}).then((dat) => {
			let contents = dat.o_json.contents;
			let content = contents[0];

			document.getElementById('title').innerText = content.title;
			document.getElementById('content').innerHTML = DOMPurify.sanitize(marked(content.text, {headerIds: false}));
	  });

	}
);
