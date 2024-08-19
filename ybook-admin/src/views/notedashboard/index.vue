<template>
  <div class="container">
    <el-card>
      <div class="note-info">
        <yb-title title="单篇笔记数据"></yb-title>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import YbTitle from '@/components/YbTitle/index.vue'
</script>

<style lang="scss" scoped>
.container{
  height: 100%;
  
  .note-info{
    padding: 0px 24px ;
    display: flex;
  }
}

::v-deep(.el-card){
  border-radius: 8px;
  height: 100%;
}
::v-deep(.el-card__body){
  padding: 0;
  height: 100%;
}
</style>
