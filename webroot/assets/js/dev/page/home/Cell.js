require.config(requireOpt);
require.config({
  paths: {
    notepad : MYENV + "/mylib/vue/notepad"
	}
})
require(
	['ajax', 'util', 'notepad'],
	function (ajax, util, notepad){
    var main = new Vue({
      el : '#main-div',
			components :{
				"notepad": notepad
			},
      data: {
        formdata: null,
        baseinfo: {tags:[], fields:[], areas: []}
      },
      filter: {

      },
      methods : {
        _onSaved : function(obj){
      	  // console.log(obj);
        }
      },

      mounted : function(){
        let iid = util.getRouterParam(0);
        let ths = this;
    		ajax.NewClient("/api/info").getData("Base").then((dat) => {
          this.baseinfo = {tags: ajax.torows(dat.tags), fields: ajax.torows(dat.fields), areas: dat.areas};
          if(iid !== 'new'){
            ajax.NewClient("/api/app").getData("LoadSElem", {iid: iid}, {busy: ths.$el}).then((formdata) => {
              ths.formdata = formdata;
            })
          }
    	  });
        window.closePageLoad();
      }
    })

	}
);
