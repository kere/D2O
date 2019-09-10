require.config(requireOpt);
require.config({
  paths: {
    notepad : MYENV + "/mylib/vue/notepad"
	}
})
require(
	['ajax', 'util', 'notepad', 'preparecookie'],
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
          this.baseinfo = {formfields: dat.formfields, areas: dat.areas};
          if(iid !== 'new'){
            ajax.NewClient("/api/app").send("LoadSElem", {iid: iid}, {busy: ths.$el}).then((formdata) => {
              ths.formdata = formdata;
            })
          }
    	  });
        
      }
    })

	}
);
