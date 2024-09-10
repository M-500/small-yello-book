<template>
  <div class="replay-container">
    <div class="list">
      <CommentSubItem :item="replayItem"></CommentSubItem>
    </div>
  </div>
</template>

<script setup>
import CommentSubItem from './comment-sub-item.vue';
import { defineProps } from 'vue';

const props = defineProps({
  replayItem: Object,
});
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
