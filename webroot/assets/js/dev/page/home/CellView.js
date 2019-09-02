require.config(requireOpt);
require(
	['ajax', 'util'],
	function (ajax, util){
    var main = new Vue({
      el : '#main-div',
      data: {
        title: '',
        text: '',
				dat: {}
      },
      methods : {
      },

      mounted : function(){
        let iid = util.getRouterParam(0);

    		ajax.NewClient("/api/app").getData("SElemByIID", {iid: iid}).then((dat) => {
					let contents = dat.o_json.contents;
					let content = contents[0];
					this.title = content.title;
					this.text = content.text;
					this.dat = dat;

    	  });
      }
    })

	}
);
