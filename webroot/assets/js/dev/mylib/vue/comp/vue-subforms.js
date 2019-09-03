define('subforms', ['util'], function(util){
  return {
    template:
    `<div class="gno-sub-form">
      <div v-for="(dat, index) in formdata" :key="index" class="gno-form-card">
        <div class="gno-card-header clearfix">
          <el-date-picker class="gno-date-picker" v-model="dat.date_on" type="date"
            formart="yyyy-MM-dd" value-format="yyyy-MM-dd"
            placeholder="日期" size="mini"></el-date-picker>
          <el-input v-model="dat.title" size="mini" class="gno-card-title"></el-input>
          <el-button class="gno-btn-close" type="text" @click="_onCloseForm(index)">
            <i class="el-icon-close"></i>
          </el-button>
        </div>

        <div v-for="(item, i) in dat.items" :key="i">
          <el-autocomplete size="mini" v-model="item.k" class="el-col el-col-10"
              @select="_onTypeSelect" value-key="name"
              :fetch-suggestions="querySearch" placeholder="数据名称"></el-autocomplete>

          <div class="el-col el-col-14">
            <el-input class="input-with-select" size="mini" v-model="item.v"
              @keyup.native="_onInputValueChanged($event, index, i)">
            </el-input>
          </div>
        </div>

      </div>

      <el-button class="button-new-tag m-t-sm parent-p" size="small" @click="_clickAddForm($event)">
        <i class="el-icon-plus"></i>
      </el-button>
    </div>`,
    props : {
      formdata: Array,
      fields : Array
    },
    watch:{
      formdata(dat){
        if(!dat || dat.length==0) return;
        for (var i = 0; i < dat.length; i++) {
          dat[i].date_on = util.date2str(dat[i].date_on, 'date');
        }
      }
    },
    methods: {
      getData(){
        var arr = [];
        for (var i = 0; i < this.formdata.length; i++) {
          let dat = this.formdata[i];
          let obj = {date_on: dat.date_on, title: dat.title, items:[]};
          for (var k = 0; k < dat.items.length; k++) {
            let item = dat.items[k];
            if(!item.v || !item.k) continue;
            obj.items.push(item);
          }
          arr.push(obj);
        }
        return arr;
      },

      _onCloseForm(index){
        this.formdata.splice(index, 1);
      },

      _clickAddForm(){
        this.formdata.push({items:[{}]});
      },

      _onTypeSelect(item){
        console.log(item);
      },
      createFilter(queryString) {
        return (dat) => {
          return (dat.name.toLowerCase().indexOf(queryString.toLowerCase()) === 0);
        };
      },
      querySearch(queryString, cb){
        var results = queryString ? this.fields.filter(this.createFilter(queryString)) : this.fields;
        cb(results);
      },

      _onInputValueChanged(e, index, i) {
        let v = e.target ? e.target.value : e;
        let items = this.formdata[index].items;
        // 添加节点
        if(v){
          // 最后一个item不为空
          if(items.length > 0 && items[items.length - 1].v){
            this.formdata[index].items.push({});
          }
        }
        if(!v && items.length > 1){ // 最后item不为空
          this.formdata[index].items.splice(i, 1);
        }
      }
    }
  };
})
