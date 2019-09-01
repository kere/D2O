require.config({
  paths: {
    areas : MYENV + "/mylib/vue/comp/vue-areas",
    tags : MYENV + "/mylib/vue/comp/vue-tags",
    subforms : MYENV + "/mylib/vue/comp/vue-subforms"
	}
})
define('notepad', ['util', 'ajax', 'Compressor', 'tags', 'subforms', 'areas'], function(util, ajax, Compressor, tags, subforms, areas){
  return {
    template:
    `<div class="gno-notepad p-a">
        <div class="gno-dtypes m-b">
          <el-radio-group v-model="itype">
            <el-radio-button label="1">事件</el-radio-button>
            <el-radio-button label="2">人物</el-radio-button>
            <el-radio-button label="3">物件</el-radio-button>
          </el-radio-group>
        </div>

        <div class="gno-top-form clearfix">
          <el-upload class="avatar-uploader pull-left" action="upload" :class="{'gno-avatar-success': form.avatar}"
            :show-file-list="false"
            :http-request="_uploadA"
            :before-upload="beforUpload">
            <img v-if="form.avatar" :src="form.avatar[1]" class="gno-avatar">
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>

          <el-form class="pull-left" ref="form" :model="form" label-width="60px" size="mini">
            <el-form-item label="名称">
              <el-input v-model="form.title" placeholder="输入题目"></el-input>
            </el-form-item>
            <el-form-item label="日期">
              <el-date-picker v-model="date_on" type="date" placeholder="输入日期"
                formart="yyyy-MM-dd" value-format="yyyy-MM-dd"></el-date-picker>
            </el-form-item>
            <el-form-item label="地点">
              <areas v-model="area" :areas="baseinfo.areas"></areas>
            </el-form-item>
          </el-form>

        </div>

        <div class="gno-textarea">
          <el-input v-model="form.text" ref="text" type="textarea" :autosize="{ minRows: 6}" class="border w100 font-size-lg" placeholder="markdown"></el-input>
        </div>

        <div class="gno-subforms">
          <el-divider>数据</el-divider>
          <subforms ref="subforms" :formdata="form.subforms" :fields="baseinfo.fields"></subforms>
        </div>

        <el-divider>图片</el-divider>
        <el-upload ref="upload" class="gno-upload"
          list-type="picture-card"
          :action="upload"
          :http-request="_uploadImg"
          :on-remove="_onImageRemove">
          <i class="el-icon-plus"></i>
        </el-upload>

        <div class="gno-tags">
          <el-divider>标签</el-divider>
          <tags ref="tags" :tags="tags" :alltags="baseinfo.tags"></tags>
        </div>


        <div class="gno-box text-right">
          <hr class="line m-b-md"></hr>
          <button @click="_onClickSave" type="button" class="gno-btn-min-w el-button el-button--primary el-button--mini">
            <i class="el-icon-check"></i>
          </button>
        </div>

    </div>`,
    props : {
      upload : String,
      baseinfo : Object
    },
    components:{
      subforms: subforms,
      areas: areas,
      tags: tags
    },
    data() {
      return {
        iid : 0,
        itype: 1,
        area: [],
        tags: [],
        date_on: '',
        form: {title: '', text: '', avatar: [], subforms:[], images:[]}
      };
    },
    methods: {
      _onImageRemove(file, fileList) {
        console.log(file, fileList);
      },

      // _onImgOK(s, file, fileList){
      //   let arr = s.split(",")
      //   file.response = arr[1];
      // },
      // _onAOK(s, file, fileList){
      //   let arr = s.split(",")
      //   file.response = arr[1];
      // },

      beforUpload(file) {
        if (!(file.type === 'image/jpeg' || file.type === 'image/png')) {
          this.$message.error('上传图片只能是 JPG,PNG 格式!');
          return false;
        }

        if (file.size / 1024 / 1024 > 2) {
          this.$message.error('上传头像图片大小不能超过 2MB!');
          return false;
        }
        return true;
      },

      _uploadA(e){
        this.doUpload(e, 0);
      },
      _uploadImg(e){
        this.doUpload(e, 1);
      },
      doUpload(e, itype){
        let cls =this;
        new Compressor(e.file, {
          quality: 0.5,
          maxWidth: 1920,
          maxHeight: 1920,
          success(blob) {
            ajax.NewUpload(cls.upload).upload(blob, {"onProgress": e.onProgress}).then(str => {
              let arr = str.split(',');
              if(itype){
                cls.form.images.push(arr);
              }else{
                cls.form.avatar = arr;
              }
              e.onSuccess(arr[1], blob);
            })
          },
          error(err) {
            e.onError(err)
          },
        });
      },

      _onClickSave(e){
        let dat = util.copy(this.form);
        let tags = this.$refs["tags"].getData();
        dat.subforms = this.$refs["subforms"].getData();
        // this.$emit('saved', {});
        this.$emit("onsave", {
          iid: this.iid,
          o_json: dat,
          tags: tags,
          itype: this.itype,
          area: this.area,
          date_on: this.date_on});
      }
    }
  };

})
