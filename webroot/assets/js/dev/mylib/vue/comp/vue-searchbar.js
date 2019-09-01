define('searchbar', [], function(){
  return {
    template:
    `<div class="search-div">
    <div class="weui-search-bar search-bar" :class="{\'weui-search-bar_focusing\':words && words!=\'\'}">
      <form class="weui-search-bar__form">
        <div class="weui-search-bar__box">
          <i class="weui-icon-search"></i>
          <input :value="words" class="weui-search-bar__input" placeholder="搜索" type="search" v-on:blur="onblur($event)" v-on:input="oninput($event)">
          <a v-on:click="onclear($event)" href="javascript:" class="weui-icon-clear searchClear"></a>
        </div>
        <label v-on:click="onclick($event)" class="weui-search-bar__label searchText" style="transform-origin: 0px 0px 0px; opacity: 1; transform: scale(1, 1);">
            <i class="weui-icon-search"></i>
            <span>{{label}}</span>
        </label>
      </form>
      <a v-on:click="oncancel($event)" href="javascript:" class="weui-search-bar__cancel-btn searchCancel">取消</a>
    </div>
    <div class="weui-cells weui-cells_access search_show_div" style="display:none">
    </div></div>`
    ,
    props : ['title', 'words'],
    computed: {
      label : function(){
        if(!this.title || this.title=='') return '搜索';
        return this.title;
      }
    },
    methods: {
      onclick: function(e){
        var $el = $(this.$el);
        $el.children('div.search-bar').addClass('weui-search-bar_focusing');
        $el.find('div.search-bar input').focus();
      },
      onblur: function(e){
        var $this = $(e.currentTarget);
        if(e.currentTarget.value.length==0) this.cancelSearch();
      },
      oninput: function(e){
        if(!this.$resultDiv){
          this.$resultDiv = $(this.$el).children('div.search_show_div');
        }
        this.$emit('dosearch', e.currentTarget.value)
        if(e.currentTarget.value.length) {
          // this.$resultDiv.show();
        } else {
          this.$resultDiv.hide();
        }
      },
      hideSearchResult : function (){
        var $el = $(this.$el);
        $el.find('div.search_show_div').hide();
        $el.find('input').val('');
      },
      cancelSearch : function (){
        this.hideSearchResult();
        var $el = $(this.$el);
        $el.find('div.search-bar').removeClass('weui-search-bar_focusing');
        $el.find('input').show();
      },
      onclear: function(){
        this.hideSearchResult();
        var $el = $(this.$el);
        $el.find('input').focus();
        this.$emit('doclear');
      },
      oncancel: function(){
        this.cancelSearch();
        var $el = $(this.$el);
        $el.find('input').blur();
      },
      onresultClick: function(){
        // this.$emit('onresult', 'aaa', 111)
      }
    }
  };

})
