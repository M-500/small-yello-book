<template>
  <div class="container">
    <el-card class="note-box">
      <div class="note-info">
        <yb-title title="笔记管理"></yb-title>

      </div>
      <div class="note-tabs">
        <el-tabs v-model="activeName"
                 class="demo-tabs"
                 @tab-click="handleClick">
          <el-tab-pane label="全部笔记"
                       name="first">User</el-tab-pane>
          <el-tab-pane label="已发布"
                       name="second">Config</el-tab-pane>
          <el-tab-pane label="审核中"
                       name="third">Role</el-tab-pane>
          <el-tab-pane label="未通过"
                       name="fourth">Task</el-tab-pane>
        </el-tabs>
      </div>
    </el-card>
  </div>

</template>

<script setup lang="ts">
import YbTitle from '@/components/YbTitle/index.vue'
import NoteManagerCard from '@/components/NoteManagerCard/index.vue'

import { ref } from 'vue'
import type { TabsPaneContext } from 'element-plus'

const activeName = ref('first')

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event)
}
</script>

<style lang="scss" scoped>
.container{
	height: 100%;
	width: 100%;
	.note-box{
		height: 100%;
		width: 100%;
		padding: 0px 8px 22px ;
	}
	.note-info{
		display: flex;
		width: 100%;
	}
	.note-tabs{

	}
}
::v-deep(.el-card){
	border-radius: 8px;
}
::v-deep(.el-card__body){
	padding: 0px;
}
::v-deep(.el-tabs__nav-scroll){
	margin: 0px 16px 0;
	position: relative;
	
}

// 替换掉伪元素样式
::v-deep(.el-tabs__nav-wrap):after{
	content: none; /* 去掉伪元素的内容 */
	display: none; /* 隐藏伪元素 */
}
::v-deep(.el-tabs__nav-wrap){
	// border-bottom: 1px solid rgb(247, 247, 247);
	border-bottom:none;
}
::v-deep(.el-tabs__header){
	border-bottom: 1px solid rgb(247, 247, 247);
}
::v-deep(.el-tabs__item.is-active){
	color: $primary-buttom-active-colder;
}
::v-deep(.el-tabs__item:hover){
	color: $primary-buttom-active-colder;
}
::v-deep(.el-tabs__active-bar ){
	background-color: $primary-buttom-active-colder;
	border-radius: 1.5px;
}
</style>
