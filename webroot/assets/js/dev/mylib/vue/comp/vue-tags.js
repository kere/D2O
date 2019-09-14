define('tags', ['util'], function(util){
  return {
    template:
    `<div class="gno-tagdatas">
      <div ref="tagbox">
        <el-tag v-for="(tag, index) in tags" closable :key="index"
              @close="_onTagClose(index)">{{tag}}</el-tag>
        <el-button class="button-new-tag" size="small"
          @click="showPane($event)"><i v-if="tagstr" class="el-icon-edit"></i><i v-else class="el-icon-plus"></i></el-button>
      </div>
      <div ref="tagpane" class="gno-tagdatas-select hide">
        <el-input ref="taginput" placeholder="请输入标签，用空格分隔" v-model="tagstr" size="mini"
            v-on:keyup.native.enter="confirm">
          <el-button @click="confirm()" slot="append" icon="el-icon-check"></el-button>
        </el-input>
      </div>
    </div>`,
    props : {
      tags: Array
    },
    model: {
      props: 'tags',
      event: 'cc'
    },
    watch:{
      tags(v){
        if(!v || v.length== 0) return;
        this.tagstr = v.join(' , ')
      }
    },
    data(){
      return {
        oval: '',// 打开input时的value，用于判断十分有改动
        tagstr: ''
      }
    },
    methods: {
      getData(){
        if(this.oval == this.tagstr) return this.tags;

        let v = this.tagstr, a=' ';
        if(!v || v.length==0) return;
        let s = v.trim().replace(/[,，。\.]/g, a);
        s = s.replace(/\s{2,}/g, a);
        this.tags = s.split(a)
        this.$emit('cc', this.tags);

        return this.tags;
      },

      _onTagClose(i) {
        this.tags.splice(i, 1);
      },

      showPane(e) {
        util.$.show(this.$refs['tagpane']);
        util.$.hide(this.$refs['tagbox']);
        this.oval = this.tagstr;
        this.$refs.taginput.focus();
      },

      confirm(e){
        util.$.hide(this.$refs['tagpane']);
        util.$.show(this.$refs['tagbox']);
        // 有改动
        return this.getData();
      }

    }
  };
})
