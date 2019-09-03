require.config(requireOpt);
require(
	['ajax', 'util', 'tool'],
	function (ajax, util, tool){
    var main = new Vue({
      el : '#main-div',
      data: {
      },
      filter: {

      },
      methods : {
        _onClick : function(){

        }
      },

      mounted : function(){
				
				window.closePageLoad();
      }
    })

	}
);

function chartOption(){
}
