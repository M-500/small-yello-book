<template>

  <svg :style="{width,height}">
    <use :xlink:href="`#${prefix}-${name}`"
         :fill="color" />
  </svg>
</template>

<script setup>

// 接受父组件传递的参数
defineProps({
  // xlink:href 的前缀
  prefix: {
    type: String,
    default: 'icon'
  },
  name: {
    type: String,
    required: true
  },
  color: {
    type: String,
    default: ''
  },
  wight: {
    type: String,
    default: '16px'
  },
  height: {
    type: String,
    default: '16px'
  }
})
</script>

<style lang="scss" scoped>

</style>
