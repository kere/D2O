require.config(requireOpt);
require.config({
  paths: {
    notepad : MYENV + "/mylib/vue/vue-notepad",
    person : MYENV + "/mylib/vue/vue-person"
	}
})
require(
	['ajax', 'util', 'notepad', 'person'],
	function (ajax, util, notepad, person){
    var main = new Vue({
      el : '#main-div',
			components :{
				"notepad": notepad,
				"person": person
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
