require.config({
  paths: {
    tags : MYENV + "/mylib/vue/comp/vue-tags",
    subform : MYENV + "/mylib/vue/comp/vue-subform"
	}
})
define('notepad', ['util', 'ajax', 'Compressor', 'tags', 'subform'], function(util, ajax, Compressor, tags, subform){
  return {
    template:
    `<div class="gno-notepad p-a">
        <div class="gno-dtypes m-b">
          <el-radio-group v-model="form.itype">
            <el-radio-button label="1">事件</el-radio-button>
            <el-radio-button label="2">人物</el-radio-button>
            <el-radio-button label="3">物件</el-radio-button>
          </el-radio-group>
        </div>
        <div class="gno-top-form clearfix">
          <el-upload class="avatar-uploader pull-left" action="upload" :class="{'gno-avatar-success': form.avatar}"
            :show-file-list="false"
            :on-success="_onAOK"
            :http-request="_uploadA"
            :before-upload="beforUpload">
            <img v-if="form.avatar" :src="form.avatar" class="gno-avatar">
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>

          <el-form class="pull-left" ref="form" :model="form" label-width="60px" size="mini">
            <el-form-item label="名称">
              <el-input v-model="form.name_ch" placeholder="输入题目"></el-input>
            </el-form-item>
            <el-form-item label="日期">
              <el-date-picker v-model="form.birthday" type="date" placeholder="输入日期"></el-date-picker>
            </el-form-item>
            <el-form-item label="地点">
              <el-select v-model="form.area_id" filterable placeholder="请选择">
                <el-option v-for="item in areas" :key="item.value"
                  :label="item.label" :value="item.value">
                </el-option>
              </el-select>
            </el-form-item>
          </el-form>

        </div>

        <div class="gno-textarea">
          <el-input v-model="form.memo" ref="text" type="textarea" :autosize="{ minRows: 6}" class="border w100 font-size-lg" placeholder="markdown"></el-input>
        </div>

        <div class="gno-subforms">
          <el-divider>数据</el-divider>
          <subform :formdata="form.subform"></subform>
        </div>

        <el-divider>图片</el-divider>
        <el-upload ref="upload" class="gno-upload"
          list-type="picture-card"
          :action="upload"
          :on-success="_onImgOK"
          :http-request="_uploadImg"
          :on-remove="_onImageRemove">
          <i class="el-icon-plus"></i>
        </el-upload>

        <div class="gno-tags">
          <el-divider>标签</el-divider>
          <tags :tags="form.tags"></tags>
        </div>


        <div class="gno-box text-right">
          <hr class="line m-b-md"></hr>
          <button @click="_onClickSave" type="button" class="gno-btn-min-w el-button el-button--primary el-button--mini">
            <i class="el-icon-check"></i>
          </button>
        </div>

    </div>`,
    props : {
      upload : String
    },
    components:{
      subform: subform,
      tags: tags
    },
    data() {
      return {
        areas: [{label: "a", value: 1}],
        form: {birthday: '', name_ch: '', itype: 1, avatar: '', tags:[], subform:[]}
      };
    },
    methods: {
      _onImageRemove(file, fileList) {
        console.log(file, fileList);
      },

      _onImgOK(url, file, fileList){
        this.upOK(str, file, 1)
      },
      _onAOK(str, file){
        this.upOK(str, file)
      },
      upOK(str, file, itype){
        let arr = str.split(",")
        file.response = arr[1];
      },

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
        this.upload(e, 0);
      },
      _uploadImg(e){
        this.upload(e, 1);
      },
      upload(e, itype){
        var url = this.upload;
        new Compressor(e.file, {
          quality: 0.5,
          maxWidth: 1024*2,
          maxHeight: 1024*2,
          success(blob) {
            ajax.NewUpload(url).upload(blob, {"onProgress": e.onProgress}).then(url => {
              e.onSuccess(url, blob);

            })
          },
          error(err) {
            e.onError(err)
          },
        });
      },

      _onClickSave(e){
        // this.$emit('saved', {});
        console.log(this.form);
      }
    }
  };

})
