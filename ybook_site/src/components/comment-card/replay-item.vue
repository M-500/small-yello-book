<template>
  <div class="replay-container">
    <div class="list">
      <CommentSubItem></CommentSubItem>
      <CommentSubItem></CommentSubItem>
      <CommentSubItem></CommentSubItem>
    </div>
  </div>
</template>

<script setup>
import CommentSubItem from './comment-sub-item.vue';
</script>

<style lang="scss" scoped>
.replay-container {
  width: 100%;
  .list{
    margin-left: 52px;
    // position: relative;
  }
}
</style>
