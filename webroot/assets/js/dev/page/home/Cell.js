require.config(requireOpt);
require.config({
  paths: {
    notepad : MYENV + "/mylib/vue/vue-notepad"
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
      },
      filter: {

      },
      methods : {
        _onSaved : function(dat){
          console.log(dat);
        }
      },

      mounted : function(){

      }
    })


		// ajax.NewClient("/api/app").send("PageData", {name:'tom', age: 22}).then(function(result){
	  //   console.log(result)
	  // })
	}
);
