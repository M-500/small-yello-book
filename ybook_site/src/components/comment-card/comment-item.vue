<template>
  <div class="comment-item">
    <div class="comment-inner-container">
      <div class="avatar">
        <router-link to="/home"
                     class="item">
          <img :src="props.comment.userAvatar"
               alt="">
        </router-link>
      </div>
      <div class="right">
        <div class="author-wrap">
          <div class="author">
            <router-link to="/home">
              {{ props.comment.userNickName }}
              <span v-show="props.comment.isAuthor">作者</span>
            </router-link>
          </div>
          <div class="comment-menu">
            ...
          </div>
        </div>
        <div class="content">{{ props.comment.content }}</div>
        <div class="picture">
          <div class="img-box"
               v-if="props.comment.mediaUrl">
            <img :src="props.comment.mediaUrl"
                 alt="">
          </div>
        </div>
        <!-- <div class="labels">
          <span class="top">置顶评论</span>
        </div> -->
        <div class="info">
          <div class="date">
            <span class="date-info">09-01</span>
            <span class="location">广东</span>
          </div>
          <div class="interaction">
            <div class="like">
              <template v-if="!liked">
                <svg-icon name="like"
                          @click="handleLike(1)"
                          height="16"></svg-icon>
                <span class="count">1</span>
              </template>
              <template v-else>
                <svg-icon name="liked"
                          @click="handleLike(0)"
                          height="16"></svg-icon>
                <span class="count">1</span>
              </template>
            </div>
            <div class="replay"
                 @click="replayHandler">
              <svg-icon name="comment"
                        height="16"></svg-icon>
              <span class="count">1</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { RouterLink } from "vue-router";
import { defineProps } from 'vue';
import { ref } from 'vue';
import { defineEmits } from 'vue';
const liked = ref(false) // 是否点赞
const props = defineProps({
  comment: Object
});

// 点击 喜欢 按钮
const handleLike = (data) => {
  if (data === 1) {
    liked.value = true
  } else {
    liked.value = false
  }
}
// 定义子组件发出的事件
const emit = defineEmits(['updateValue']);

const replayHandler = () => {
  emit('updateValue', props.comment);
}
</script>

<style lang="scss" scoped>
.comment-item{
	position: relative;
	display: flex;
	padding: 8px;
	.comment-inner-container{
		position: relative;
    display: flex;
    z-index: 1;
    width: 100%;
    flex-shrink: 0;
		.avatar{
			flex: 0 0 auto; // 
			.item{
				img{
					width: 40px;
					height: 40px;
					display: flex;
					align-items: center;
					justify-content: center;
					cursor: pointer;
					border-radius: 100%;
					border: 1px solid rgba(0,0,0,0.08);
					object-fit: cover;
				}
			}
		}
		.right{
			margin-left: 12px;
			display: flex;
			flex-direction: column;
			font-size: 14px;
			flex-grow: 1;
			.author-wrap{
				display: flex;
				justify-content: space-between;
				align-items: center;
				.author{
					display: flex;
    			align-items: center;
					font-size: 14px;
					// 去掉a标签的下划线
					a{
						text-decoration: none;
						color: #333;
						span{
							margin-left: 8px;
							font-size: 12px;
							color: rgba(51,51,51,0.6);
							background-color: rgba(0,0,0,0.03);
							border-radius: 10.5px;
							margin-left: 4px;
							padding: 3px 6px;
							font-size: 10px;
						}
					}
				}
			}
			.content{
				margin-top: 4px;
    		line-height: 140%;
				color: #333;
			}
			.labels{
				margin-top: 8px;
				.top{
					padding: 0 8px;
					background: rgba(255,36,66,0.06);
					border-radius: 999px;
					font-weight: 500;
					font-size: 12px;
					line-height: 24px;
					height: 24px;
					color: $primary-color-red;
					display: inline-block;
				}
			}
			.picture{
				.img-box{
					margin-top: 8px;
					width: 120px;
					height: 160px;
					overflow: hidden;
					background:rgba(0,0,0,0.03);
					cursor: zoom-in;
					position: relative;
					img{
						object-fit: cover;
						width: 100%;
						height: 100%;
						border-radius: 8px;
						border-style: none;
					}
				}
				
			}
			.info{
				display: flex;
				flex-direction: column;
				justify-content: spance-between;
				font-size: 12px;
				line-height: 16px;
				color: rgba(51,51,51,0.6);
				.date{
					margin: 8px 0;
					.location{
						margin-left: 4px;
					}
				}
				.interaction{
					display: flex;
					margin-left: -2px;
					color:rgba(51,51,51,0.6);
					.like{
						display: flex;
						align-items: center;
						box-sizing: border-box;
						-webkit-user-select: auto;
						user-select: auto;
						scrollbar-width: none;
						-webkit-tap-highlight-color: transparent;
						span{
							margin-left: 4px;
							font-size: 12px;
							font-weight: 500;
						}
					}
					.replay{
						margin-left: 16px;
						display: flex;
						align-items: center;
						box-sizing: border-box;
						-webkit-user-select: auto;
						user-select: auto;
						scrollbar-width: none;
						-webkit-tap-highlight-color: transparent;
						span{
							margin-left: 4px;
							font-size: 12px;
							font-weight: 500;
						}
					}
				}
			}
		}
	}
}

</style>
