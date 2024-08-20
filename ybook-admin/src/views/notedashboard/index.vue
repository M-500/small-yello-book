<template>
  <div class="container">
    <el-card>
      <div class="note-info">
        <yb-title title="单篇笔记数据"></yb-title>
      </div>
      <div class="inline-btn">
        <el-button type="text">查看更多</el-button>
      </div>
      <div class="note-list-line">
        <note-info-card></note-info-card>
        <note-info-card></note-info-card>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import YbTitle from '@/components/YbTitle/index.vue'
import NoteInfoCard from '@/components/NoteInfoCard/index.vue'
</script>

<style lang="scss" scoped>
.container{
  height: 100%;
  
  .note-info{
    padding: 0px 24px ;
    display: flex;
  }
  .inline-btn{
    display: flex;
    padding: 0px 24px ;
  }
  .note-list-line{
    display: flex;
    flex-wrap: wrap;
    grid-template-columns: auto;
    gap: 20px 16px;
    margin-top: 24px;
    padding: 0px 24px 24px;
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
