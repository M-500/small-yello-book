<template>
  <div class="container">
    <div class="notifaction_page">
      <div class="ybook_tab_wrap">
        <div class="tab_list">
          <tab_item :name="tab.label"
                    v-for="(tab,index) in tabs"
                    :key="index"
                    :class="{active: selectTab === tab.name}"
                    @click="selectTab = tab.name" />
        </div>
      </div>

      <div class="tab_page">
        <div class="comment_list">
          <CommentCard />
          <CommentCard />
          <CommentCard />
          <CommentCard />
          <CommentCard />
          <CommentCard />
          <div class="end_content"> - THE END - </div>
        </div>
      </div>
    </div>
  </div>

</template>

<script setup>
import tab_item from "@/components/TabItem/index.vue";
import CommentCard from "@/components/CommentCard/index.vue";
import { ref } from "vue";

const selectTab = ref('comment');
const tabs = ref([
  { name: 'comment', label: '评论和@' },
  { name: 'like', label: '赞和关注' },
  { name: 'follow', label: '新增关注' },
]); 
</script>

<style lang="scss" scoped>
.container{
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
  // align-items: center;
  .notifaction_page{
    width: 755px;
    margin: 0 auto;
    // 这里布局有问题
    .ybook_tab_wrap{
      
      .tab_list{
        display: flex;
        justify-content: flex-start;
        height: 72px;
        padding-top: 16px;
      }
    }
    .tab_page{
      // width: calc(-33.6px + 40vw);
      height: calc(100% - 72px);
      overflow: auto;
      padding: 0 24px 32px 24px;
      display: flex;
      flex-direction: column;
      .comment_list{
        display: flex;
        flex-direction: column;
        // gap: 16px;
        .end_content{
          display: flex;
          justify-content: center;
          margin-top: 24px;
          margin-bottom: 24px;
          font-size: 12px;
          line-height: 16px;
          color: rgba(51,51,51,0.6);
        }
      }
    }
  }
}

</style>
