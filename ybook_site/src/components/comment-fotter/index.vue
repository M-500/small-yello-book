<template>
  <div class="comment-fotter-container">
    <div class="replay-content"
         v-if="replayContent">
      <div class="replay">回复：{{ replayContent.userNickName }}</div>
      <div class="content">{{ replayContent.content }}</div>
    </div>
    <div class="input-box">
      <!-- 左侧说点什么 -->
      <div class="content-edit">
        <p class="content-input"
           @click="inputHandler"
           ref="editableP"
           @input="handleInput"
           :contenteditable="inputActive"></p>
        <div class="not-active"
             @click="inputHandler"
             v-if="!inputActive">
          <div class="inner">
            <img src="@/assets/imgs/avatar.jpeg"
                 class="avatar-item"
                 alt="">
            <span>说点什么...</span>
          </div>
        </div>
      </div>
      <!-- 右侧点赞收藏 -->
      <div class="interact-container"
           v-show="!inputActive">
        <div class="box">
          <template v-if="!liked">
            <svg-icon name="like"
                      @click="handleLike(1)"
                      height="32"></svg-icon>
            <span class="count"
                  @click="handleLike(1)">点赞</span>
          </template>
          <template v-else>
            <svg-icon name="liked"
                      @click="handleLike(0)"
                      height="32"></svg-icon>
            <span class="count">1</span>
          </template>
        </div>
        <div class="box">
          <template v-if="!collectd">
            <svg-icon name="collect"
                      @click="handleCollect(1)"
                      height="36"></svg-icon>
            <span class="count"
                  @click="handleCollect(1)">收藏</span>
          </template>
          <template v-else>
            <svg-icon name="collected"
                      @click="handleCollect(0)"
                      height="36"></svg-icon>
            <span class="count">1</span>
          </template>
        </div>
        <div class="box"
             @click="inputHandler">
          <svg-icon name="comment"
                    height="32"></svg-icon>
          <span class="count">评论</span>
        </div>
      </div>
    </div>
    <div class="bottom-box"
         v-if="inputActive">
      <div class="btn-box">
        <div class="left-area">
          <div class="box">
            <SvgIcon name="pic"
                     width="32"></SvgIcon>
          </div>
        </div>
        <div class="right-area">
          <button class="send"
                  @click="addComment"
                  :class="inputing ? 'active': ''">发送</button>
          <button class="cancel"
                  @click="cancelSub">取消</button>
        </div>
      </div>

    </div>

  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { likeRequest } from '@/api/interactive';
import type { IntrLikeForm } from '@/api/interactive/types';
import type { CommentForm } from '@/api/comment/types';
import { addCommentRequest } from '@/api/comment';
import { defineProps } from 'vue';
import { watch } from 'vue';
const inputActive = ref(false)
const inputing = ref(false) // 是否有输入文字
const liked = ref(false) // 是否点赞
const collectd = ref(false) // 是否收藏
const commentContent = ref('') // 评论内容
const editableP = ref(null);  // 获取p标签的引用
const replayContent = ref(null)
const parentId = ref(0)
const props = defineProps({
	detail: {
		type: Object,
		default: () => {}
	},
	triggerAction: {
		type:Object
	}
})

// 监听commentContent的变化
watch(commentContent, (newVal) => {
	if (newVal) {
		inputing.value = true
	} else {
		inputing.value = false
	}
})
// 监听triggerAction的变化
watch(() => props.triggerAction, (newVal) => {
  if (newVal) {
    console.log("来活啦！",newVal)
		inputActive.value = true
		replayContent.value = newVal
		parentId.value = newVal.id
  }
});


const handleInput = () => {
		// 获取p标签内的内容
		if (editableP.value) {
			commentContent.value = editableP.value.innerText;
			console.log(commentContent.value)
		}
	};




const likeForm = ref({
	resource_id: "",
	action: 'like',
	owner_id: "",
	resource_type: "note"
})

const requestLike = async () => {
	try {
		likeForm.value.resource_id = props.detail.uuid
		likeForm.value.owner_id = props.detail.authorId
		likeRequest(likeForm.value).then((res) => {
			liked.value = !liked.value
		}).catch((err) => {
			console.log(err)
		})
	} catch (error) {
		console.log(error)
	}
}

// 点击评论按钮
const inputHandler = () => {
  inputActive.value = true
}

// 取消评论
const cancelSub = () => {
  inputActive.value = false
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

const addComment = () => {
	publishComment(props.detail.uuid,parentId.value, '')
	inputActive.value = false
	commentContent.value = ''
	inputing.value = false
	editableP.value.innerText = ''
}

const publishComment = async( resource_id:string,parent_id:number,media_url:string) => {
	try {
		const commentForm: CommentForm = {
			content: commentContent.value,
			resource_id: resource_id,
			parent_id: parent_id,
			media_url: media_url,
			is_pinned: false
		}
		let res = await addCommentRequest(commentForm)
		console.log('评论回复内容为',res)
	} catch (error) {
		console.log(error)
	}
}
// 点击 收藏 按钮
const handleCollect = (data:number) => {
  if (data === 1) {
    collectd.value = true
  } else {
    collectd.value = false
  }
}

</script>

<style lang="scss" scoped>
.comment-fotter-container{
	position: relative;
	height: unset; 
	.replay-content{
		margin-bottom: 12px;
    padding: 0 16px;
    line-height: 16px;
		.replay{
			color: rgba(51,51,51,0.6);
			font-size: 14px;
		}
		.content{
			color:rgba(51,51,51,0.8);
			font-size: 14px;
			width: 100%;
			margin-top: 4px;
			overflow: hidden;
			text-overflow: ellipsis;
			white-space: nowrap;
		}
	}
	.input-box{
			display: flex;  // flex 布局
			align-items: center;  // 垂直居中
			margin-bottom: 16px;  // 底部间距
			transition: all .2s;  // 过渡动画
			min-height: 40px;  // 确保有显式高度
			.content-edit{
				flex: 1;  // 自适应宽度
				position: relative;  // 相对定位
				min-width: 70px;  // 最小宽度
				min-height: 40px;
				.content-input{
					cursor: text;  // 文本输入光标
					caret-color: $primary-color-red;
					margin: 0;
					width: 100%;
					min-height: 40px;
					max-height: 100px;
					background-color: rgba(0,0,0,0.03);
					// background-color: red;
					border: none;
					border-radius: 4px;
					padding: 9px 16px;
					border-radius: 20px;
					outline: none;
					overflow-y: auto;
					text-rendering: optimizeLegibility;
					word-break: break-all;
					white-space: break-spaces;
					font-size: 16px;
					line-height: 22px;
					color: #333;
				}
				.not-active{
					position: absolute;
					height: 100%;
					display: flex;
					align-items: center;
					justify-content: center;
					left: 0;
					top: 0;
					padding-left: 8px;
					.inner{
						width: 100%;
						height: 100%;
						display: flex;
						align-items: center;
						justify-content: center;
						justify-content: flex-start;
						color: rgba(51,51,51,0.6);
						font-size: 14px;
						.avatar-item{
							margin-right: 6px;
							width: 24px;
							height: 24px;
							border-radius: 12px;
							overflow: hidden;
							color: rgba(51,51,51,0.3);
							border-style: none;
						}
					}
				}
			}
			.interact-container{
				display: flex;
				justify-content: flex-end;
				.box{
					height: 40px;
					display: flex;
					align-items: center;
					justify-content: center;
					cursor: pointer;
					border-radius: 100%;
					padding: 0 8px;
					// background: rgba(0,0,0,0.08);
					.count{
						margin-left: 4px;
						margin-right: 0;
						font-weight: 500;
						font-size: 14px;
					}
				}
			}
			
			
		}
	.bottom-box{
		display: flex;
		width: 100%;
		margin-bottom: 16px; 
		justify-content: space-between;
		.btn-box{
			display: flex;
			width: 100%;
			justify-content: space-between;
			.left-area{
				flex: 1;
				display: flex;
				justify-content: flex-start;
				.box{
					width: 40px;
					height: 40px;
					display: flex;
					align-items: center;
					justify-content: center;
					cursor: pointer;
					border-radius: 100%;
					// background: rgba(0,0,0,0.08);
					svg{
						width: 24px;
						height: 24px;
					}
				}
			}
			.right-area{
				display: flex;
				button{
					margin-left: 8px;
					width: 64px;
					height: 40px;
					display: flex;
					align-items: center;
					justify-content: center;
					cursor: pointer;
					flex-shrink: 0;
					border-radius: 44px;
					font-size: 16px;
					font-weight: 600;
				}
				.cancel{
					border:1px solid rgba(0,0,0,0.08);
					color: rgba(51,51,51,0.8);
				}
				.send{
					color: #ffffff;
					font-weight: 600;

					background:#ff9aa2;	
				}
				.active{
    			background:$primary-color-red;	
				}
			}
		}
	}
	
}
</style>

