require.config({
  paths: {
	}
})
define('notepad', ['ajax', 'accto'], function(ajax, accto){
  return {
    template:
    `<div class="gno-notepad p-a">
      <div v-show="!isEdit" class="w100 h100 as-table">
        <div class="border-b relative as-table-row font-size-elg">
          This is 这是一个title
        </div>
        <hr class="dashed m-y-sm"></hr>
        <div class="h100 as-table-row font-size-lg" style="display: table-row">
          This is 这是一个title
          This is 这是一个title
          This is 这是一个title
        </div>
        <div class="text-right as-table-row" style="display: table-row">
          <a @click="_onClickEdit">
            <i class="el-icon-edit-outline"></i>
          </a>
        </div>
      </div>
      <div v-show="isEdit" class="w100 h100 as-table">
        <div class="h100 as-table-row">
          <div class="as-table h100 w100">
            <div class="as-table-row">
              <el-date-picker v-model="date_on" type="date" placeholder="开始日期"> </el-date-picker>
              <el-date-picker class="pull-right" v-model="date_on" type="date" placeholder="结束日期"> </el-date-picker>
            </div>
            <div class="as-table-row">
              <input ref="title" type="text" class="note-title w100 p-x-sm p-y-sm border border-b-none font-size-lg">
            </div>
            <div class="h100 as-table-row">
              <textarea ref="text" class="border w100 h100 p-a-sm font-size-lg"></textarea>
            </div>
          </div>
        </div>
        <div class="p-t-sm">
          <el-select v-model="tags" multiple placeholder="请选择标签" class="w100">
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </div>
        <el-upload ref="upload" class="m-t-sm"
          list-type="picture-card"
          :action="upload"
          :before-upload="handleImageBefore"
          :on-success="handleImageSuccess"
          :on-remove="handleImageRemove">
          <i class="el-icon-plus"></i>
        </el-upload>
        <div class="text-right" style="display: table-row">
          <button @click="_onClickSave" type="button" class="gno-btn-min-w el-button el-button--primary el-button--mini">
            <i class="el-icon-check"></i>
          </button>
        </div>
      </div>
    </div>`,
    props : {
      upload : String,
      isEdit : Boolean
    },
    components:{
    },
    data() {
      return {
        options: [
          { value: '选项1', label: '黄金糕' }, { value: '选项2', label: '双皮奶' },
          { value: '选项3', label: '蚵仔煎' }, { value: '选项4', label: '龙须面' }, { value: '选项5', label: '北京烤鸭' }
        ],
        tags: '',
        date_on : '',
        dialogImageUrl: '',
        dialogVisible: false
      };
    },
    methods: {
      _onClickEdit(){
        this.isEdit = true;
      },

      handleImageRemove(file, fileList) {
        console.log(file, fileList);
      },
      handleImageSuccess(url, file, fileList){
        if(url.substring(0, 7)==="webroot"){
          url = url.substr(7);
        }
        file.response = url;
        // file.url = url;
      },
      handleImageBefore(file){
        var ts = (new Date()).getTime().toString(), ptoken = window['accpt'] || '';
        var str = ts+file.name +  file.size + file.lastModified + file.type+navigator.userAgent+ts+ptoken + window.location.hostname;
        var a = this.$refs['upload'];
        return new Promise(function(resolve, reject){
          a.headers = { Accto: accto(str), Accts: ts, AccPage: ptoken };
          a.data = { name : file.name, size : file.size, lastModified : file.lastModified, type : file.type };
          resolve();
        });
      },

      // handleRequest(dat){
      //   var up = ajax.NewUpload(this.upload, {onprogress: this.handleProgress});
      //   up.upload(dat.file).then((url)=>{
      //     if(url.substring(0, 7)==="webroot"){
      //       url = url.substr(7);
      //     }
      //     this.dialogImageUrl = url;
      //     this.dialogVisible = true;
      //   });
      // },
      _onClickSave(e){
        this.$emit('saved', {title:this.$refs.title.value, text:this.$refs.text.value, lang: 1});
      }
    }
  };

})
