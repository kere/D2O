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
        baseinfo: {tags:[], fields:[], areas: []}
      },
      filter: {

      },
      methods : {
        _onSaved : function(obj){
      		ajax.NewClient("/api/app").send("SaveSElem", obj).then((dat) => {
      	    console.log(dat)
      	  })
        }
      },

      mounted : function(){
    		ajax.NewClient("/api/info").getData("TagsFormFieldsAreas").then((dat) => {
          this.baseinfo = {tags: ajax.torows(dat.tags), fields: ajax.torows(dat.fields), areas: dat.areas};
    	  })
      }
    })

	}
);
