<template>
  <div class="footer">
    <a href=""
       class="title">
      <span>{{ props.item.noteTitle }}</span>
    </a>
    <div class="author_wrapper">
      <a href=""
         class="author">
        <img :src="props.item.author.avatar"
             class="author-avatar"
             alt="">
        <span>{{ props.item.author.nickName   }}</span>
      </a>
      <span class="like_wrapper"
            @click="handleLike(1)">
        <span class="like-lottie"></span>
        <SvgIcon name="heart"
                 width="16"
                 height="16"
                 class="like" />
        <span class="count"
              v-if="props.item.interactiveInfo.likeCnt > 0">{{ props.item.interactiveInfo.likeCnt }}</span>
      </span>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { defineProps } from 'vue';
import { ref } from 'vue';
import { likeRequest } from '@/api/interactive';
import { type IntrLikeForm } from '@/api/interactive/types';
const liked = ref(false) // 是否点赞
const props:any = defineProps({
  item: {
    type: Object,
    required: true,
    default: () => ({} as any),
  },
});
const likeForm = ref({
	resource_id: "",
	action: 'like',
	owner_id: "",
	resource_type: "note"
})
const requestLike = async () => {
	try {
		likeForm.value.resource_id = props.item.uuid
		likeForm.value.owner_id = props.item.authorId
		likeRequest(likeForm.value).then((res) => {
			liked.value = !liked.value
		}).catch((error) => {
			console.log(error)
		})
	} catch (error) {
		console.log(error)
	}
}

// 点击 喜欢 按钮
const handleLike = (data:number) => {
  if (data === 1) {
    liked.value = true
		requestLike()
  } else {
    liked.value = false
  }
}
</script>

<style lang="scss" scoped>
.footer{
	padding: 12px;
	width: 100%;
	.title{
		margin-bottom: 8px;
    word-break: break-all;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 2;
    overflow: hidden;
    font-weight: 500;
    font-size: 14px;
    line-height: 140%;
		color: #333;
		// 去掉a标签下划线样式
		text-decoration: none;
	}
	.author_wrapper{
		display: flex;
    align-items: center;
    justify-content: space-between;
    height: 20px;
    color: rgba(51,51,51,0.8);
    font-size: 12px;
    transition: color 1s;
		// width: 100%;

		.author{
			width: 100%;
			display: flex;
			align-items: center;
			color: inherit;
			overflow: hidden;
			text-overflow: ellipsis;
			// white-space: nowrap;
			margin-right: 12px;
			// 去掉a标签下划线样式
			text-decoration: none;
    	background-color: transparent;
			.author-avatar{
				margin-right: 6px;
				width: 20px;
				height: 20px;
				border-radius: 20px;
				border: 1px solid rgba(0,0,0,0.08);
				// flex-shrink: 0;
			}
		}
		.like_wrapper{
			position: relative;
			cursor: pointer;
			display: flex;
			align-items: center;
			.like-lottie{
				width: 16px;
    		height: 16px;
				position: absolute;
				left: 0;
				top: 0;
				transform: scale(1.7);
			}
			.count{
				margin-left: 2px;
				font-size: 12px;
				text-wrap: nowrap;
			}
		}
	}
	
}
</style>
