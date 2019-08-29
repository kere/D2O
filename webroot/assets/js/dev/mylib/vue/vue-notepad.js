require.config({
  paths: {
	}
})
define('notepad', ['util', 'ajax', 'accto', 'Compressor'], function(util, ajax, accto, Compressor){
  return {
    template:
    `<div class="gno-notepad p-a">
      <div class="w100 h100 as-table">
        <div class="gno-box">
          <div>
            <el-date-picker v-model="date_on" type="date" placeholder="开始日期"> </el-date-picker>
            <el-date-picker class="pull-right" v-model="date_on" type="date" placeholder="结束日期"> </el-date-picker>
          </div>

          <el-input v-model="data[0].title" ref="title" type="text" class="gno-note-title w100 border border-b-none font-size-lg" placeholder="标题"></el-input>

          <el-input v-model="data[0].text" ref="text" type="textarea" :autosize="{ minRows: 16}" class="gno-textarea border w100 font-size-lg" placeholder="请输入内容 markdown"></el-input>

        </div>

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
            <el-button class="button-new-tag parent-p" size="small" @click="_clickAddTag($event)">
              <i v-if="isTagPaneOpen" class="el-icon-arrow-up"></i>
              <i v-else class="el-icon-arrow-down"></i>
            </el-button>
          </div>
          <div class="hide fade m-t-sm">
            <el-checkbox-group v-model="tags">
              <el-checkbox label="复选框 A"></el-checkbox>
              <el-checkbox label="复选框 B"></el-checkbox>
              <el-checkbox label="复选框 C"></el-checkbox>
              <el-checkbox label="复选框 D"></el-checkbox>
              <el-checkbox label="复选框 E"></el-checkbox>
              <el-checkbox label="复选框 F"></el-checkbox>
              <el-checkbox label="复选框 G"></el-checkbox>
            </el-checkbox-group>
          </div>
        </div>
        <div class="gno-box">
          <el-divider>数据</el-divider>
          <div ref="formbox" class="formbox">
            <el-card v-for="(dat, index) in forms" class="box-card" shadow="never">
              <div slot="header" class="clearfix">
                <el-date-picker v-model="dat.date_on" type="date" placeholder="数据日期"> </el-date-picker>

                <el-button class="gno-btn-close" type="text" @click="_onCloseForm(index)"><i class="el-icon-close"></i></el-button>
              </div>
              <div v-for="(o, index) in dat.items" :key="index" class="text item">

                <el-input placeholder="输入数据" class="input-with-select" v-model="o.val">
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

        <div class="gno-box text-right">
          <hr class="line m-b-md"></hr>
          <button @click="_onClickSave" type="button" class="gno-btn-min-w el-button el-button--primary el-button--mini">
            <i class="el-icon-check"></i>
          </button>
        </div>

      </div>

    </div>`,
    props : {
      upload : String
    },
    components:{
    },
    data() {
      return {
        data: [{title: '', text: ''}],
        forms : [],
        tags: ['标签一', '标签二', '标签三'],
        date_on : '',
        isTagPaneOpen : false
      };
    },
    methods: {
      _clickAddTag(e){
        this.isTagPaneOpen = util.tool.taggle(e);
      },

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
          quality: 0.7,
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
