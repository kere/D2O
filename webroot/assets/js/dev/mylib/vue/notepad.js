require.config({
  paths: {
    areas : MYENV + "/mylib/vue/comp/vue-areas",
    tags : MYENV + "/mylib/vue/comp/vue-tags",
    contents : MYENV + "/mylib/vue/comp/vue-contents",
    subforms : MYENV + "/mylib/vue/comp/vue-subforms"
	}
})
define('notepad',
['util', 'ajax', 'Compressor', 'tags', 'subforms', 'areas', 'contents'],
function(util, ajax, Compressor, tags, subforms, areas, contents){
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

          <el-form class="pull-left" ref="form" label-width="60px" size="mini">
            <el-form-item label="日期">
              <el-date-picker v-model="date_on" type="date" placeholder="输入日期"
                formart="yyyy-MM-dd" value-format="yyyy-MM-dd"></el-date-picker>
            </el-form-item>
            <el-form-item label="地点">
              <areas :area="area" v-model="area" :areas="baseinfo.areas"></areas>
            </el-form-item>
          </el-form>

        </div>

        <contents ref="contents" :formdata="ojson.contents"></contents>

        <el-upload ref="upload" class="gno-upload m-t"
          list-type="picture-card"
          :file-list="imageList"
          :action="upload"
          :on-preview="_onImagePreview"
          :http-request="_uploadImg"
          :on-success="_onImgSuccess"
          :on-remove="_onImageRemove">
          <i class="el-icon-plus"></i>
        </el-upload>

        <div class="gno-subforms">
          <el-divider>数据</el-divider>
          <subforms ref="subforms" :formdata="ojson.subforms" :fields="baseinfo.fields"></subforms>
        </div>

        <div class="gno-tags">
          <el-divider>标签</el-divider>
          <tags ref="tags" :tags="tags" v-model="tags"></tags>
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
      contents: contents,
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
        ojson: {contents: [{title: '', text: '', lang: 'zh'}], avatar: null, subforms:[], images:[]}
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
        setTimeout(() => {
          this.rebuildImagesLinkBtn();
        }, 200);
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
      _onImagePreview(file){
        let a = this.$refs.contents.$el.querySelector('textarea');
        // ![avatar](http://baidu.com/pic/doge.png)
        let url = file.url.substr(0,5) == 'blob:'? file.response : file.url;
        a.value = a.value.substring(0, a.selectionStart) + '\n!['+file.name+']('+url+")\n"+ a.value.substr(a.selectionStart);
        this.ojson.contents[this.$refs.contents.currentPaneI].text = a.value;
      },

      _onImageLink(e){
        console.log(e);
      },

      _onImageRemove(file, fileList) {
        console.log(file, fileList);
        let i = util.findIndex("name", file.name, this.imageList);
        if(i < 0) return;
        this.imageList.splice(i, 1);

      },

      _onImgSuccess(arr, file){
        // console.log(url, file);
        file.name = arr[0];
        file.response = arr[1];
        setTimeout(() => {
          this.rebuildImagesLinkBtn();
        }, 200);
      },

      // 把preview icon变成link
      rebuildImagesLinkBtn(){
        util.$.each(this.$refs.upload.$el.querySelectorAll('i.el-icon-zoom-in'), (t)=>{
          t.className = 'el-icon-link';
        });
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
              e.onSuccess(arr, blob);
            })
          },
          error(err) {
            e.onError(err)
          },
        });
      },

      _onClickSave(e){
        let ojs = util.copy(this.ojson);
        let tags = this.$refs.tags.confirm();
        ojs.subforms = this.$refs["subforms"].getData();
        ojs.contents = this.$refs["contents"].getData();
        // this.$emit('saved', {});
        let obj = {
            iid: this.iid,
            o_json: ojs,
            tags: tags,
            itype: this.itype,
            area: this.area,
            date_on: this.date_on
          };
        if(obj.date_on=="") obj.date_on= '100-01-01';

    		ajax.NewClient("/api/app").send("SaveSElem", obj, {loading:true}).then((dat) => {
          this.$emit("onSaved", obj);
          this.isSuccess = true;
          this.isError = false;
          let cls = this;
          setTimeout(() => {
            cls.isSuccess = false;
          }, 1000);

    	  }).catch((err) =>{
          this.isError = true;
          this.errMessage = err;
        })
      }
    }
  };

})
