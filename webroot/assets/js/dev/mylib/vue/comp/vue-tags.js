define('tags', ['util'], function(util){
  return {
    template:
    `<div class="gno-alltags">
      <div ref="tagbox">
        <el-tag v-for="(tag, index) in alltags" v-if="tag.selected"
            closable :key="tag"
            @close="_onTagClose(index)">
        {{tag.name}}</el-tag>
        <el-button class="button-new-tag" size="small"
          @click="showPane($event)"><i class="el-icon-plus"></i></el-button>
      </div>
      <div ref="tagpane" class="gno-alltags-select hide">
        <el-tag v-for="(tag, index) in alltags" :key="tag.name" :type="tag.selected ? '' : 'info'"
              @click="_onTagClick(index)">
          {{tag.name}}
        </el-tag>
        <el-tag v-show="isChanged" type="warning" @click="_onTagConfirm($event)">
          <i class="el-icon-check"></i>
        </el-tag>
      </div>
    </div>`,
    props : {
      tags: Array
    },
    data: function(){
      return {
        isChanged : false,
        alltags: [{
          value: 1,
          name: '黄金糕'
        }, {
          value: 2,
          name: '双皮奶'
        }, {
          value: 3,
          name: '蚵仔煎'
        }]
      }
    },
    methods: {
      _onTagClose(i) {
        let obj = this.alltags[i];
        obj.selected = false;
        this.alltags.splice(i, 1, obj);
      },

      _onTagClick(i){
        let obj = this.alltags[i];
        obj.selected = !obj.selected;
        this.alltags.splice(i, 1, obj)
        this.isChanged = true;
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
