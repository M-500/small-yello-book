<template>
  <div class="container">
    <!-- 左边头像 -->
    <div class="avatar_box">
      <!-- <img src="../../assets/imgs/avatar.jpeg"
           alt=""> -->
      <avatar imgURL="https://sns-avatar-qc.xhscdn.com/avatar/1000g2jo2obr4ip8k80005n65phn49t50rl79avo?imageView2/2/w/540/format/webp|imageMogr2/strip2" />
    </div>
    <!-- 右边主要信息区域 -->
    <div class="main">
      <div class="info">
        <!-- 作者名字 -->
        <div class="user_info">
          <router-link class="nickname"
                       to="">三冬四夏</router-link>
          <span class="user_tag">作者</span>
        </div>
        <!-- 副标题 -->
        <div class="interaction_hint">
          <span class="mark">回复了你的评论</span>
          <span class="timer">4天前</span>
        </div>
        <!-- 回复内容 -->
        <div class="interaction_content">
          为啥是老王
        </div>
        <!-- 评论内容 -->
        <div class="quote_info">
          累不死的牛是老王
        </div>
        <!-- 操作 -->
        <div class="action">
          <div class="action_repy">
            <svg-icon class="icon"
                      name="talk"
                      width="20"
                      high="20"
                      color="rgba(51,51,51,0.8)" />
            <div class="action_text">回复</div>
          </div>
          <div class="action_like">
            <svg-icon name="like"
                      width="20"
                      high="20"
                      color="rgba(51,51,51,0.8)" />
          </div>
        </div>
      </div>
      <div class="extra">
        <img src="../../assets/imgs/avatar.jpeg"
             class="extra_img"
             alt="">
      </div>
    </div>
  </div>
</template>

<script setup>
import { } from "vue"
import SvgIcon from "../global/SvgIcon.vue";
import avatar from "../avatar/index.vue";
</script>

<style lang="scss" scoped>

.container{
	display: flex;
	flex-direction: row;
	padding-top: 24px;
	max-width: 100vw;
	.avatar_box{
		// width: 100px;
		// height: 100px;
		// border-radius: 50%;
		overflow: hidden;
		margin-right: 24px;
	
	}
	.main{
		flex-grow: 1;
    flex-shrink: 1;
    display: flex;
    flex-direction: row;
    padding-bottom: 24px;
    border-bottom: 1px solid rgba(0,0,0,0.08);
		.info{
			flex-grow: 1; // 伸缩比例
    	flex-shrink: 1; // 伸缩比例
			.user_info{
				display: flex;
				flex-direction: row;
				align-items: center;
				font-size: 16px;
				font-weight: 600;
				margin-bottom: 4px;
				.nickname{
					color: #333;
					cursor: pointer; // 鼠标移动到上面会变成小手
					text-decoration: none;
    			background-color: transparent;
				}
				.user_tag{
					padding: 3px 4px;
					margin-left: 4px;
					background: rgba(0, 0, 0, 0.03);
					border-radius: 4px;
					font-size: 12px;
					color: rgba(51,51,51,0.6);
					font-weight: 400;
				}
			}
			.interaction_hint{
				font-size: 14px;
				color: rgba(51,51,51,0.6);
				margin-bottom: 8px;
				.timer{
					margin-left: 8px;
				}
			}
			.interaction_content{
				font-size: 14px;
				color: #333;
				margin-bottom: 12px;
				line-height: 140%;
				cursor: pointer;
			}
			.quote_info{
				font-size: 12px;
				display: flex;
				align-items: center;
				color: rgba(51,51,51,0.6);
				margin-bottom: 12px;
				cursor: pointer;
			}
			.quote_info:before{
				content: "";
				display: inline-block;
				border-radius: 8px;
				margin-right: 6px;
				width: 4px;
				height: 17px;
				background: rgba(0,0,0,0.08);	
			}
			.action{
				display: flex;
				color: rgba(51,51,51,0.8);
				.action_repy{
					cursor: pointer;
					width: 88px;
					height: 40px;
					border: 1px solid rgba(0,0,0,0.08);
					border-radius: 999px;
					display: flex;
					align-items: center;
					justify-content: center;
					color: rgba(51,51,51,0.8);
					.action_text{
						margin-left: 4px;
    				font-size: 16px;
					}
				}
				.action_like{
					cursor: pointer;
					width: 40px;
					height: 40px;
					margin-left: 12px;
					border: 1px solid rgba(0,0,0,0.08);
					border-radius: 999px;
					display: flex;
					align-items: center;
					justify-content: center;
					color: rgba(51,51,51,0.8);
				}
			}
		}
		
		.extra{
			min-width: 48px;
			flex-shrink: 0;
			margin-left: 24px;
			.extra_img{
				cursor: pointer;
				width: 48px;
				height: 48px;
				border: 1px solid rgba(0,0,0,0.02);
				border-radius: 6px;
				object-fit: cover;
			}
		}
	}
}
</style>
