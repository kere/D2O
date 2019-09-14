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
        isPageOK: false,
        formdata: null,
        baseinfo: {fields:[], areas: []}
      },
      methods : {
        _onSaved : function(iid){
          window.location = "/cell/view/"+iid;
        }
      },

      mounted : function(){
        let iid = util.getRouterParam(0);
        let ths = this;
    		ajax.NewClient("/api/info").getData("Base").then((dat) => {
          dat.tags = ajax.torows(dat.tags);
          this.baseinfo = dat;
          if(iid == 'new'){
            ths.isPageOK = true;
          }else{
            ajax.NewClient("/api/app").send("LoadSElem", {iid: iid}, {busy: ths.$el}).then((formdata) => {
              ths.isPageOK = true;
              ths.formdata = formdata;
            }, (e) => {
              util.tool.pageErr("错误", e.toString())
            })
          }
    	  });

      }
    })

	}
);
