require.config(requireOpt);
require(
	['ajax', 'util'],
	function (ajax, util){
    var main = new Vue({
      el : '#main-div',
      data: {
        rows: []
      },
      methods : {
      },

      mounted : function(){
    		ajax.NewClient("/api/app").getData("SElems").then((dat) => {
          let ojsonI = dat.fields.indexOf("o_json");
          this.rows = ajax.torows(dat, (k, i) =>{
            if(k === ojsonI){
              return JSON.parse(dat.columns[k][i]);
            }
            return dat.columns[k][i];
          });

          // console.log(dat);
  		    ajax.NewClient("/api/info").getData("Base");
    	  });
      }
    })

	}
);
