<template>
  <div class="sub-container">
    <div class="comment-inner">
      <div class="avatar">
        <router-link to="/home"
                     class="item">
          <img src="@/assets/imgs/avatar.jpeg"
               class="img-item"
               alt="">
        </router-link>
      </div>
      <div class="right">
        <div class="author-wrap">
          <div class="author">
            <router-link to="/home">
              王木木咯
              <span>作者</span>
            </router-link>
          </div>
          <div class="comment-menu">
            ...
          </div>
        </div>
        <div class="content">{{ item.content }}</div>
        <div class="info">
          <div class="date">
            <span class="date-info">09-01</span>
            <span class="location">广东</span>
          </div>
          <div class="interaction">
            <div class="like">

            </div>
            <div class="replay"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineProps } from 'vue';
const props = defineProps({
  item: Object
});
</script>

<style lang="scss" scoped>
.sub-container {
	position: relative;
	display: flex;
	padding: 8px;
	.comment-inner{
		position: relative;
    display: flex;
    z-index: 1;
    width: 100%;
    flex-shrink: 0;
		.avatar{
			flex: 0 0 auto; // 不伸缩，不收缩
			.img-item{
				width: 24px;
				height: 24px;
				display: flex;
				align-items: center;
				justify-content: center;
				cursor: pointer;
				border-radius: 100%;
				border: 1px solid rgba(0,0,0,0.08);
				object-fit: cover;
			}
		}
		.right{
			border-bottom: none;
			padding-bottom: 0;
			margin-left: 12px;
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
				font-size: 14px;
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
					.like{
						box-sizing: border-box;
						-webkit-user-select: auto;
						user-select: auto;
						scrollbar-width: none;
						-webkit-tap-highlight-color: transparent;
					}
				}
			}
		}
	}
}
</style>
