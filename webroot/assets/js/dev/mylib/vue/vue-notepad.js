define('notepad', ['ajax', 'accto', 'compressor'], function(ajax, accto, Compressor){
  return {
    template:
    `<div class="gno-notepad p-a">
      <div class="w100 h100 relative" style="display:table">
        <div class="relative" style="display:table-row;">
          <input ref="title" type="text" class="note-title w100 p-x-sm p-y border border-b-none font-size-lg">
        </div>
        <div class="h100 relative" style="display: table-row">
          <textarea ref="text" class="border w100 h100 p-a-sm font-size-lg"></textarea>
        </div>
        <el-upload ref="upload" class="m-t-sm"
          list-type="picture-card"
          :action="upload"
          :before-upload="handleBefore"
          :on-success="handleSuccess"
          :on-remove="handleRemove">
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
      upload : String
    },
    data() {
      return {
        dialogImageUrl: '',
        dialogVisible: false
      };
    },
    methods: {
      handleRemove(file, fileList) {
        console.log(file, fileList);
      },
      handleSuccess(url, file, fileList){
        if(url.substring(0, 7)==="webroot"){
          url = url.substr(7);
        }
        file.response = url;
        // file.url = url;
      },
      handleBefore(file){
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
        this.$emit('saved', {title:this.$refs.title.value, text:this.$refs.text.value});
      }
    }
  };

})
