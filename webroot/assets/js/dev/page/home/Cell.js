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
          var obj = {

          }

      		ajax.NewClient("/api/app").send("SaveSElem", obj).then(function(result){
      	    console.log(result)
      	  })
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
