<template>
  <div class="container">
    <div class="user"></div>
    <div class="tab_list">
      <tab_item name="笔记"
                :active="true" />
      <tab_item name="收藏"
                :active="true" />
      <tab_item name="点赞"
                :active="false" />
    </div>
  </div>
</template>

<script setup>
import tab_item from "@/components/TabItem/index.vue";
</script>

<style lang="scss" scoped>
.container{
  .tab_list{
    display: flex;
  }
}
</style>
