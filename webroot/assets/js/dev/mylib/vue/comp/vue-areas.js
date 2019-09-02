define('areas', [], function(){
  return {
    template:
    `<template>
      <el-select v-model="area" filterable placeholder="请选择" multiple @change="_onChange($event)">
        <el-option v-for="a in areas" :key="a.id"
          :label="a.name + ' ' + a.en" :value="a.id">
          {{ a.name }}
        </el-option>
      </el-select>
    </template>`,
    props : {
      areas: Array,
      area: Array
    },
    model: {
      props: 'area',
      event: 'cc'
    },
    methods: {
      _onChange(v) {
        this.$emit('cc', v)
        this.area = v;
      }
    }
  };
})
