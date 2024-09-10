<template>
  <div class="parent-comment">
    <comment-item :comment="commentItem"></comment-item>
    <replay-item v-for="(item,index) in commentItem.sub"
                 :key="index"
                 :replayItem="item"></replay-item>
  </div>
</template>

<script setup>
import CommentItem from './comment-item.vue';
import ReplayItem from './replay-item.vue';
import { defineProps } from 'vue';

const props = defineProps({
  commentItem: Object,
});
console.log("来了吗", props);
</script>

<style lang="scss" scoped>
.parent-comment {
	width: 100%;
	margin-bottom: 16px;
}
</style>
