define('areas', ['util'], function(util){
  return {
    template:
    `<template>
      <el-select v-model="area" filterable placeholder="请选择" multiple @change="_onChange($event)">
        <el-option v-for="a in areas" :key="a.id"
          :label="a.cn + ' ' + a.en" :value="a.id">
          {{ a.cn }}
        </el-option>
      </el-select>
    </template>`,

    props : {
      areas: Array,
      area: Array
    },
    data: function(){
      return {
        areas : []
      }
    },
    model: {
      props: 'area',
      event: 'cc'
    },
    methods: {
      fullData(){
        let arr = new Array(this.area.length);
        for (var i = 0; i < this.area.length; i++) {
          let k = util.findSortedI('id', this.area[i], this.areas)
          if(k < 0) continue
          arr[i] = this.areas[k];
        }
        return arr;
      },
      _onChange(v) {
        this.$emit('cc', v);
      }
    }
  };
})
