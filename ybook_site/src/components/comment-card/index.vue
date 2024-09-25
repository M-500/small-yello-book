<template>
  <div class="parent-comment">
    <comment-item :comment="commentData"
                  @updateValue="handlerData"></comment-item>
    <replay-item v-for="(item,index) in commentData.sub"
                 :key="index"
                 :replayItem="item"></replay-item>
  </div>
</template>

<script setup>
import CommentItem from './comment-item.vue';
import ReplayItem from './replay-item.vue';
import { defineEmits } from 'vue';
import { defineProps } from 'vue';

const props = defineProps({
  commentData: Object,
});
const emit = defineEmits(['replayHdl']);

// 继续传递给父组件
const handlerData = (data) => {
  emit('replayHdl', data);
}


</script>

<style lang="scss" scoped>
.parent-comment {
	width: 100%;
	margin-bottom: 16px;
}
</style>
