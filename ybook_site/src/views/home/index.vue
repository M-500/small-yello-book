<script setup lang="ts">
import tab_item from "@/components/TabItem/index.vue";
import NoteItem from "@/components/note-item/index.vue";
import { getNoteListRequest } from "@/api/note";
import type { NoteFeedQuery } from "@/api/note/types";
import { onMounted } from "vue";
import { ref } from "vue";

const noteList = ref([]);
const selectTab = ref('美妆');
const tabs = ref([
  { name: '美妆', label:'meizhuang'},
  { name: '健身', label:'jianshen'},
  { name: '游泳', label:'youyong'},
  { name: '育儿', label:'yuer'},
]);


const getFeedNoteList = async () => {
  let params: NoteFeedQuery = {
    tagId: 0,
    page: 1,
    size: 20,
  };
  getNoteListRequest(params).then((res) => {
    noteList.value = res.data.list
  });
};

onMounted(() => {
  getFeedNoteList()
});

</script>

<template>
  <div class="container">
    <div class="channel_container">
      <tab_item v-for="(tab,index) in tabs"
                :key="index"
                :active="selectTab === tab.name"
                :name="tab.name"
                @click="selectTab = tab.name" />
    </div>

    <div class="note_container">
      <note-item v-for="(item,index) in noteList"
                 :key="index"
                 :item="item" />
    </div>

  </div>
</template>

<style scoped lang="scss">  
.container {
  .channel_container{
			overflow: hidden;
			display: flex;
			user-select: none;
			-webkit-user-select: none;
			align-items: center;
			font-size: 16px;
			color: var(--color-secondary-label);
			height: 40px;
			white-space: nowrap;
			margin-bottom: 16px;
		}
		.note_container{
			display: flex; // flex布局
      flex-wrap: wrap; // 自动换行
      justify-content: flex-start; // 两端对齐
      gap: 20px;  // 间距

		}
}
@media (max-width: 768px){
  .container{
    .note_container{
      display: flex; // flex布局
      flex-wrap: wrap; // 自动换行
      justify-content: flex-start; // 两端对齐
      gap: 12px;  // 间距
    }
  }
}
</style>