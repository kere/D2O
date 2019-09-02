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
            <el-radio-button :label="1">事件</el-radio-button>
            <el-radio-button :label="2">人物</el-radio-button>
            <el-radio-button :label="3">物件</el-radio-button>
          </el-radio-group>
        </div>

        <div class="gno-top-form clearfix">
          <el-upload class="avatar-uploader pull-left" action="upload" :class="{'gno-avatar-success': isAvatar}"
            :show-file-list="false"
            :http-request="_uploadA"
            :before-upload="beforUpload">
            <img v-if="isAvatar" :src="ojson.avatar.url" class="gno-avatar">
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>

          <el-form class="pull-left" ref="form" :model="ojson" label-width="60px" size="mini">
            <el-form-item label="名称">
              <el-input v-model="ojson.title" placeholder="输入题目"></el-input>
            </el-form-item>
            <el-form-item label="日期">
              <el-date-picker v-model="date_on" type="date" placeholder="输入日期"
                formart="yyyy-MM-dd" value-format="yyyy-MM-dd"></el-date-picker>
            </el-form-item>
            <el-form-item label="地点">
              <areas :area="area" v-model="area" :areas="baseinfo.areas"></areas>
            </el-form-item>
          </el-form>

        </div>

        <div class="gno-textarea">
          <el-input v-model="ojson.text" ref="text" type="textarea" :autosize="{ minRows: 6}" class="border w100 font-size-lg" placeholder="markdown"></el-input>
        </div>

        <div class="gno-subforms">
          <el-divider>数据</el-divider>
          <subforms ref="subforms" :formdata="ojson.subforms" :fields="baseinfo.fields"></subforms>
        </div>

        <el-divider>图片</el-divider>
        <el-upload ref="upload" class="gno-upload"
          list-type="picture-card"
          :file-list="imageList"
          :action="upload"
          :http-request="_uploadImg"
          :on-remove="_onImageRemove">
          <i class="el-icon-plus"></i>
        </el-upload>

        <div class="gno-tags">
          <el-divider>标签</el-divider>
          <tags ref="tags" :tags="tags" :tagdatas="baseinfo.tags"></tags>
        </div>


        <div class="gno-box text-right">
          <hr class="line m-b-md"></hr>
          <el-alert v-show="isSuccess" title="成功保存数据" type="success" center show-icon :closable="false"></el-alert>
          <el-alert v-show="isError" title="数据保存失败" type="error" :description="errMessage" show-icon></el-alert>
          <hr v-show="isSuccess || isError" class="line m-b-md"></hr>

          <button @click="_onClickSave" type="button" class="gno-btn-min-w el-button el-button--primary el-button--mini">
            <i class="el-icon-check"></i>
          </button>
        </div>

    </div>`,
    components:{
      subforms: subforms,
      areas: areas,
      tags: tags
    },
    props : {
      upload : String,
      formdata : Object,
      baseinfo : Object
    },
    data() {
      return {
        isSuccess : false,
        isError : false,
        errMessage: '',
        iid : 0,
        itype: 1,
        area: [],
        tags: [],
        imageList: [],
        date_on: '',
        ojson: {title: '', text: '', avatar: null, subforms:[], images:[]}
      };
    },
    watch:{
      formdata(dat){
        if(!dat) return;
        this.iid = dat.iid;
        this.itype = parseInt(dat.itype);
        this.area = dat.area;
        this.tags = dat.tags;
        this.date_on = util.date2str(dat.date_on, 'date');
        this.ojson = dat.o_json;

        this.imageList = this.ojson.images;
      }
    },
    computed:{
      isNew( ){
        return !this.formdata;
      },
      isAvatar(){
        return this.ojson.avatar;
      }
    },
    methods: {
      _onImageRemove(file, fileList) {
        console.log(file, fileList);
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
                cls.ojson.images.push({name:arr[0], url:arr[1]});
              }else{
                cls.ojson.avatar = {name:arr[0], url:arr[1]};
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
        let dat = util.copy(this.ojson);
        let tags = this.$refs["tags"].getData();
        dat.subforms = this.$refs["subforms"].getData();
        // this.$emit('saved', {});
        let obj = {
            iid: this.iid,
            o_json: dat,
            tags: tags,
            itype: this.itype,
            area: this.area,
            date_on: this.date_on
          };

    		ajax.NewClient("/api/app").send("SaveSElem", obj, {loading:true}).then((dat) => {
          this.$emit("onSaved", obj);
          this.isSuccess = true;
          this.isError = false;
          let cls = this;
          setTimeout(() => {
            cls.isSuccess = false;
          }, 1500);

    	  }).catch((err) =>{
          this.isError = true;
          this.errMessage = err;
        })
      }
    }
  };

})
