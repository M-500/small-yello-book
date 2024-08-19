<template>
  <div class="container">
    <el-card>
      <div class="note-info">
        <yb-title title="笔记管理"></yb-title>
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
		padding: 0px 8px 24px ;
		display: flex;
	}
}
::v-deep(.el-card){
	border-radius: 8px;
}
::v-deep(.el-card__body){
	padding: 0px;
}
</style>
