define('contents', [], function(){
  return {
    template:
    `<div class="gno-textarea">
      <div v-for="item in formdata">
        <el-input v-model="item.title" placeholder="输入题目" class="gno-title"></el-input>
        <el-input v-model="item.text" ref="text" type="textarea" :autosize="{ minRows: 6}"
            class="border w100 font-size-lg" placeholder="markdown"></el-input>
      </div>
    </div>`,
    props : {
      formdata: Array
    },
    watch:{
      // formdata(v){
      //   console.log(v);
      // }
    },
    methods: {
      getData() {
        return this.formdata;
      }
    }
  };
})
