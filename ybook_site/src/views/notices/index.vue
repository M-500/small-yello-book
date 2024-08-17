<template>
  <div class="notifaction_page">
    <div class="ybook_tab_wrap">
      <div class="tab_list">
        <tab_item name="评论和@"
                  :active="true" />
        <tab_item name="赞和关注"
                  :active="false" />
        <tab_item name="新增关注"
                  :active="false" />
      </div>
    </div>

    <div class="tab_page">
      <p style="height: 100000px;">hahah</p>
    </div>
  </div>
</template>

<script setup>
import tab_item from "@/components/TabItem/index.vue";
</script>

<style lang="scss" scoped>
.notifaction_page{
  height: 100%;

  .ybook_tab_wrap{
    width: 100%;
    .tab_list{
      display: flex;
      justify-content: flex-start;
      height: 72px;
      // padding-top: 16px;
    }
  }
  .tab_page{
    height: calc(100% - 72px);
    overflow: auto;
  }
}
</style>
