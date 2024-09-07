<template>
  <div class="parent-comment">
    <comment-item></comment-item>
    <replay-item></replay-item>
  </div>
</template>

<script setup>
import CommentItem from './comment-item.vue';
import ReplayItem from './replay-item.vue';
</script>

<style lang="scss" scoped>
.parent-comment {
	width: 100%;
	margin-bottom: 16px;
}
</style>
