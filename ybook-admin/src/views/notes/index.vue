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
          <el-tab-pane :label="tab.label"
                       v-for="(tab,index) in tabs"
                       :key="index"
                       :class="{active: selectTab === tab.name}"
                       @click="selectTab = tab.name">
            <div class="note-container-box">
              <note-manager-card v-for="(item,index) in noteList"
                                 :key="index"
                                 :item="item"></note-manager-card>
            </div>
          </el-tab-pane>
        </el-tabs>
      </div>
    </el-card>
  </div>

</template>

<script lang="ts" setup>
import YbTitle from '@/components/YbTitle/index.vue'
import NoteManagerCard from '@/components/NoteManagerCard/index.vue'
import { getNoteList } from '@/api/note'
import type { queryNoteListForm } from '@/api/note/type'

import { onMounted, ref } from 'vue'
import type { TabsPaneContext } from 'element-plus'


const activeName = ref('all')
const noteList = ref([])
const selectTab = ref('all');
const tabs = ref([
  { name: 'all', label: '全部笔记',state: -1 },
  { name: 'published', label: '已发布',state: 0 },
  { name: 'checking', label: '审核中',state: 1 },
  { name: 'refuse', label: '未通过',state: 2 },
]); 

const handleClick = (tab:TabsPaneContext, event: Event) => {
  console.log("我尼玛",tab, event)
  // getNoteListData()
}

// 获取笔记列表
const getNoteListData = async (state:number) => {
  const params: queryNoteListForm = {
    state: state,
    page: 1,
    size: 10
  }
  const res = await getNoteList(params)
  noteList.value = res.data.list
}

onMounted(() => {
  getNoteListData(-1)
})
</script>

<style lang="scss" scoped>
::v-deep(.el-card__body){
  padding: 0px;
  height: 100%;
}
::v-deep(el-tabs__content){
  height: 100%;
  overflow: auto;
}
::v-deep(.el-tab-pane){
    height: 100%;
    overflow: auto;
  }
.container{
	height: 100%;
	width: 100%;
  .el-card{
    height: 100%;
    width: 100%;
    padding: 0px 8px 22px ;
    .note-info{
      display: flex;
      width: 100%;
    }
    .note-tabs{
      height: calc(100% - 76px);
      .el-tabs{
        height: 100%;
        .note-container-box{
          display: flex;
          flex-direction: column;
          // height: 100%;
          overflow: auto;
        }
      }
    }
  }
  
	// .note-box{
	// 	height: 100%;
	// 	width: 100%;
	// 	padding: 0px 8px 22px ;
	// }
	// .note-info{
	// 	display: flex;
	// 	width: 100%;
	// }
	// .note-container-box{
	// 	display: flex;
	// 	flex-direction: column;
	// 	height: 100%;
	// 	overflow: auto;
	// }
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
