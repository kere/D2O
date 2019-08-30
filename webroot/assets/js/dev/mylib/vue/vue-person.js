define('person', ['util', 'ajax', 'Compressor'], function(util, ajax, Compressor){
  return {
    template:
    `<div class="gno-notepad p-a">
        <div class="gno-box clearfix" style="margin-bottom:8px;">
          <el-avatar class="pull-left" shape="square" :size="100" :src="avatar"></el-avatar>

          <el-form class="gno-small-form pull-left" ref="form" :model="form" label-width="80px" size="small">
            <el-form-item label="中文名称">
              <el-input v-model="form.name_ch"></el-input>
            </el-form-item>
            <el-form-item label="出生日期">
              <el-date-picker v-model="form.birthday" type="date" placeholder="出生日期"></el-date-picker>
            </el-form-item>
          </el-form>

        </div>

          <el-input v-model="form.memo" ref="text" type="textarea" :autosize="{ minRows: 6}" class="gno-textarea border w100 font-size-lg" placeholder="请输入内容 markdown"></el-input>

        <div class="gno-box">
          <el-divider>数据</el-divider>
          <div ref="formbox" class="formbox">
            <el-card v-for="(dat, index) in forms" class="box-card" shadow="never">
              <div slot="header" class="clearfix">
                <el-date-picker class="gno-date-picker" v-model="dat.date_on" type="date" placeholder="数据日期"> </el-date-picker>

                <el-button class="gno-btn-close" type="text" @click="_onCloseForm(index)"><i class="el-icon-close"></i></el-button>
              </div>
              <div v-for="(o, index) in dat.items" :key="index" class="text item">

                <el-input placeholder="输入数据" class="input-with-select" v-model="o.val" size="mini">
                  <el-select v-model="o.itype" slot="prepend" placeholder="请选择" allow-create :filterable="true">
                    <el-option label="餐厅名" :value="1"></el-option>
                    <el-option label="订单号" :value="2"></el-option>
                    <el-option label="用户电话" :value="3"></el-option>
                  </el-select>
                </el-input>
              </div>
            </el-card>
          </div>

          <el-button class="button-new-tag parent-p" size="small" @click="_clickAddForm($event)">
            <i class="el-icon-plus"></i>
          </el-button>
        </div>

        <el-divider>图片</el-divider>
        <el-upload ref="upload"
          list-type="picture-card"
          :action="upload"
          :on-success="_onImageSuccess"
          :http-request="_uploadHttpRequest"
          :on-remove="_onImageRemove">
          <i class="el-icon-plus"></i>
        </el-upload>

        <div class="gno-box">
          <el-divider>标签</el-divider>
          <div style="margin-top:-5px">
            <el-tag :key="tag" v-for="tag in tags" closable
              :disable-transitions="false" @close="_onTagClose(tag)">
            {{tag}}
            </el-tag>
          </div>
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
    data() {
      return {
        avatar: '',
        form: {birthday: '', name_ch: ''},
        forms : [],
        tags: ['标签一', '标签二', '标签三'],
        date_on : '',
        isTagPaneOpen : false
      };
    },
    methods: {
      _onTagClose(tag) {
        this.tags.splice(this.tags.indexOf(tag), 1);
      },

      _onImageRemove(file, fileList) {
        console.log(file, fileList);
      },

      _onImageSuccess(url, file, fileList){
        if(url.substring(0, 7)==="webroot"){
          url = url.substr(7);
        }
        file.response = url;
      },

      // _onImageBefore(file){
      //   var ts = (new Date()).getTime().toString(), ptoken = window['accpt'] || '';
      //   var str = ts+file.name +  file.size + file.lastModified + file.type+navigator.userAgent+ts+ptoken + window.location.hostname;
      //   var a = this.$refs['upload'].$refs['upload-inner'];
      //   a.headers = { Accto: accto(str), Accts: ts, AccPage: ptoken };
      //   a.data = { name : file.name, size : file.size, lastModified : file.lastModified, type : file.type };
      //   return true;
      // },

      // :before-upload="_onImageBefore"
      _uploadHttpRequest(e){
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

      _clickAddForm(){
        this.forms.push({date_on: '2019-01-02', items: [{itype: 2, val: 10}]});
      },

      _onCloseForm(index){
        this.forms.splice(index, 1);
      },

      _onClickSave(e){
        this.$emit('saved', {title:this.$refs.title.value, text:this.$refs.text.value, lang: 1});
      }
    }
  };

})
