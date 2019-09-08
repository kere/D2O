define('contents', [], function(){
  return {
    template:
    `<div class="gno-textarea">
      <div class="content-langs">
        <span v-for="(item, index) in formdata" :class="{'current-pane': index==currentPaneI}" @click="_onLangClick(index)">{{item.lang}}</span>
      </div>
      <div v-for="(item, index) in formdata" v-show="currentPaneI==index">
        <el-input v-model="item.title" placeholder="输入题目" class="gno-title"></el-input>
        <el-input v-model="item.text" ref="text" type="textarea" :autosize="{ minRows: 6}"
            class="border w100 font-size-lg" placeholder="markdown"></el-input>
      </div>
    </div>`,
    props : {
      formdata: Array
    },
    data(){
      return {
        currentPaneI : 0
      }
    },
    watch:{
      formdata(v){
        let hasEN = false;
        for (var i = 0; i < v.length; i++) {
          if(v[i].lang=='en') hasEN = true;
        }
        if(!hasEN){
          v.push({title:'', text:'', lang:'en'})
        }
      }
    },
    methods: {
      _onLangClick(index){
        this.currentPaneI =index;
      },
      getData() {
        return this.formdata;
      }
    }
  };
})
