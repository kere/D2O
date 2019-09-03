define('tags', ['util'], function(util){
  return {
    template:
    `<div class="gno-tagdatas">
      <div ref="tagbox">
        <el-tag v-for="(tag, index) in tagdatas" v-if="tag.selected"
            closable :key="tag"
            @close="_onTagClose(index)">
        {{tag.name}}</el-tag>
        <el-button class="button-new-tag" size="small"
          @click="showPane($event)"><i class="el-icon-plus"></i></el-button>
      </div>
      <div ref="tagpane" class="gno-tagdatas-select hide">
        <el-tag v-for="(tag, index) in tagdatas" :key="tag.name" :type="tag.selected ? '' : 'info'"
              @click="_onTagClick(index)">
          {{tag.name}}
        </el-tag>
        <el-tag v-show="isChanged" type="warning" @click="_onTagConfirm($event)">
          <i class="el-icon-check"></i>
        </el-tag>
      </div>
    </div>`,
    props : {
      tags: Array,
      tagdatas: Array
    },
    data(){
      return {
        isChanged : false
      }
    },
    watch:{
      tags(v){
        if(!v) return;
        // tags = [1,2,3];
        util.arrMix(v, this.tagdatas, (k, i) => {
          if(v[k] === this.tagdatas[i].iid){
            this.tagdatas[i].selected = true;
            this.tagdatas.splice(i, 1, this.tagdatas[i]);
            return true;
          }
          return false;
        })
      }
    },
    methods: {
      _onTagClose(i) {
        let obj = this.tagdatas[i];
        obj.selected = false;
        this.tagdatas.splice(i, 1, obj);
      },

      _onTagClick(i){
        let obj = this.tagdatas[i];
        obj.selected = !obj.selected;
        this.tagdatas.splice(i, 1, obj)
        this.isChanged = true;
      },

      getData(){
        let arr = [];
        for (var i = 0; i < this.tagdatas.length; i++) {
          if(!this.tagdatas[i].selected) continue;
          arr.push(this.tagdatas[i].iid);
        }
        return arr;
      },

      showPane(e) {
        util.$.show(this.$refs['tagpane']);
        util.$.hide(this.$refs['tagbox']);
      },

      _onTagConfirm(e){
        util.$.hide(this.$refs['tagpane']);
        util.$.show(this.$refs['tagbox']);
        this.isChanged = false;
      }

    }
  };
})
