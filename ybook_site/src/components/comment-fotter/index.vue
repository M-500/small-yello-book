<template>
  <div class="comment-fotter-container">

    <div class="input-box">
      <!-- 左侧说点什么 -->
      <div class="content-edit">
        <p class="content-input"
           @click="inputHandler"
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
        点赞
      </div>
    </div>
    <div class="bottom-box"
         v-if="inputActive">
      <div class="btn-box">
        <div class="left-area">
          <button>表情</button>
          <button>图片</button>
          <button>视频</button>
        </div>
        <div class="right-area">
          <button class="send">发送</button>
          <button class="cancel"
                  @click="cancelSub">取消</button>
        </div>
      </div>

    </div>

  </div>
</template>

<script setup>
import { tr } from 'element-plus/es/locales.mjs';
import { ref } from 'vue';
const inputActive = ref(false)

const inputHandler = () => {
  inputActive.value = true
}

const cancelSub = () => {
  inputActive.value = false
}
</script>

<style lang="scss" scoped>
.comment-fotter-container{
	position: relative;
	height: unset; 
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
					caret-color: #ff2442;
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
			
		}
	.bottom-box{
		display: flex;
		width: 100%;
		justify-content: space-between;
		.btn-box{
			display: flex;
			width: 100%;
			justify-content: space-between;
			.left-area{
				flex: 1;
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
    			background:#ff2442;	
				}
			}
		}
	}
	
}
</style>

