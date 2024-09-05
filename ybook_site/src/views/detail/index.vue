<template>
  <div class="outer-link-container">
    <div class="note-container">
      <div class="left-card">
        <div class="note-cover">
          <img src="@/assets/imgs/avatar.jpeg"
               alt="">
        </div>
      </div>
      <div class="right-card">
        <div class="auth-container">
          <div class="author-wrapper">
            <div class="info">
              <router-link>
                <img class="avatar-item"
                     src="@/assets/imgs/avatar.jpeg"
                     alt="">
              </router-link>
              <router-link class="name">
                <span class="username">吴狗林</span>
              </router-link>
            </div>
            <div class="follow-btn">
              <button>关注</button>
            </div>
          </div>
        </div>
        <div class="note-scoller"></div>
        <div class="fotter"></div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { RouterLink } from 'vue-router';
</script>

<style lang="scss" scoped>
.outer-link-container {
	display: flex;
	justify-content: center;
	align-items: center;
	height: 100%;
	padding: 0 20px;
	.note-container {
		display: flex;
		width: 100%;
		height: 100%;
		border-radius: 16px;
		overflow: hidden;
		box-shadow: 0 0 10px rgba(0,0,0,0.1);
		.left-card {
			width: 65%;
			height: 100%;
			.note-cover {
				width: 100%;
				height: 100%;
				img {
					width: 100%;
					height: 100%;
					object-fit: cover;
				}
			}
		}
		.right-card {
			width: 35%;
			height: 100%;
			display: flex;
			flex-direction: column;
			.auth-container {
				display: flex;
				justify-content: space-between;
				align-items: center;
				.author-wrapper {
					display: flex;
					width: 100%;
					align-items: center;
					justify-content: space-between;
					width: 100%;
					padding: 24px;
					height: 100%;
					border-bottom: 1px solid transparent;
					.info {
						height: 100%;
						display: flex;
						justify-content: flex-start;
						.avatar-item {
							width: 40px;
							height: 40px;
							border-radius: 100%;
							object-fit: cover;
						}
						.name{
							margin-left: 12px;
							font-size: 16px;
							font-weight: 500;
							color: rgba(51,51,51,0.8);
							height: 40px;
							overflow: hidden;
							display: flex;
							align-items: center;
							justify-content: center;
							// 去掉a标签的下划线
							text-decoration: none;
							.username {
								overflow: hidden;  // 超出部分隐藏
								text-overflow: ellipsis;  // 文字超出部分显示省略号
								white-space: nowrap; // 不换行
								font-size: 16px;
								font-weight: 500;
								
							}
						}
						
					}
					.follow-btn {
						button {
							font-size: 16px;
							font-weight: 600;
							line-height: 16px;
							padding: 0 24px;
							height: 40px;
							background-color: #ff2e4d;
							color: #fff;
							border-radius: 100px;
						}
					}
				}
			}
		}
	}
}
</style>
