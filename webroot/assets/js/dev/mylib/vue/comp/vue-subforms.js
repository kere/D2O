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
        <el-input v-for="(item, i) in dat.items" :key="i"
            placeholder="输入数据" class="input-with-select" size="mini"
            v-model="item.v" :type="item.type"
            @keyup.native="_onInputValueChanged($event, index, i)">
          <el-select v-model="item.k" slot="prepend" placeholder="请选择" :filterable="true"
              @change="_onSelectChanged($event, index, i)">
            <el-option v-for="field in fields" :label="field.name" :value="field.name"></el-option>
          </el-select>
        </el-input>
      </div>

      <el-button class="button-new-tag parent-p" size="small" @click="_clickAddForm($event)">
        <i class="el-icon-plus"></i>
      </el-button>
    </div>`,
    props : {
      formdata: Array,
      fields : Array
    },
    // data: function(){
    //   return {
    //     dat: {}
    //   }
    // },
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

      _onSelectChanged(e, index, i){

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
