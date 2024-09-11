<template>
  <div class="outer-link-container">
    <div class="note-container">
      <div class="left-card">
        <div class="note-cover">
          <template v-if="noteDetail.contentType === 1">
            <el-carousel indicator-position="outside"
                         :autoplay="false">
              <el-carousel-item v-for="item in noteDetail.imgList"
                                :key="item">
                <img class="note-img"
                     :src="item.localPath"
                     alt="">
              </el-carousel-item>
            </el-carousel>
          </template>
          <template v-else>
            <!-- 展示视频了 -->
            <video width="100%"
                   style="background-color: #000"
                   height="100%"
                   controls>
              <source src="http://127.0.0.1:8122/static/1725721792_%E2%80%9C%E6%88%91%E4%B8%8D%E9%9C%80%E8%A6%81%E5%90%91%E4%BD%A0%E8%A7%A3%E9%87%8A!%E2%80%9D.mp4"
                      type="video/mp4">
            </video>
          </template>
        </div>
      </div>
      <div class="right-card">
        <div class="auth-container">
          <div class="author-wrapper">
            <div class="info">
              <router-link to="/">
                <img class="avatar-item"
                     src="@/assets/imgs/avatar.jpeg"
                     alt="">
              </router-link>
              <router-link class="name">
                <!-- <span class="username">{{ noteDetail.value.author.nickName }}</span> -->
                <span class="username">王木木</span>
              </router-link>
            </div>
            <div class="follow-btn">
              <button @click="handlerFollowBtn(1)"
                      v-if="!isFollowed"
                      class="active">关注</button>
              <button v-else
                      @click="handlerFollowBtn(0)">已关注</button>
            </div>
          </div>
        </div>
        <div class="note-scoller">
          <div class="note-content">
            <div class="title">
              {{noteDetail.noteTitle}}
            </div>
            <div class="desc">
              <span class="detail-desc">{{ noteDetail.noteContent }}</span>
            </div>
            <div class="bottom-container">
              <span class="date">08-08 浙江</span>
            </div>
          </div>
          <div class="divider"></div>
          <div class="comment-el">
            <div class="comment-container">
              <div class="total">共360条评论</div>
              <div class="list-container">
                <comment-card v-for="(item,index) in commentList"
                              :key="index"
                              :commentItem="item"></comment-card>
              </div>
            </div>
          </div>
        </div>
        <div class="fotter">
          <div class="bar-container">
            <comment-fotter :detail="noteDetail"></comment-fotter>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { RouterLink } from 'vue-router';
import { getNoteDetailRequest } from '@/api/note';
import { getCommentListRequest } from '@/api/comment';
import { useRoute } from 'vue-router';
import { onMounted, ref } from 'vue';

import CommentCard from '@/components/comment-card/index.vue';
import CommentFotter from '@/components/comment-fotter/index.vue';
// 获取路由参数
const route = useRoute();
const noteId = route.params.uuid; // 获取路由参数上的uuid
const noteDetail = ref({});
const commentList = ref([]);
const isFollowed = ref(false); // 是否关注


// 获取数据
const getNoteDetail = async () => {
  const res = await getNoteDetailRequest(noteId);
  noteDetail.value = res.data;
};
// 获取评论数据
const getCommentList = async () => {
  const res = await getCommentListRequest(noteId);
  commentList.value = res.data;
};


// 点击关注按钮
const handlerFollowBtn = (data) => {
  if (data === 1) {
    isFollowed.value = true;
  } else {
    isFollowed.value = false;
  }
};


onMounted(() => {
  getNoteDetail();
  getCommentList();
});
</script>

<style lang="scss" scoped>
::v-deep(.el-carousel__container){
	// height: 100%;
}
::v-deep(.el-carousel__container){
	// height: 100%;
}
.outer-link-container {
	display: flex;
	justify-content: center;
	align-items: center;
	height: 100%;
	padding: 0 100px 20px;
	.note-container {
		display: flex;
		width: 100%;
		height: 100%;
		border-radius: 16px;
		overflow: hidden;
		border:1px solid rgba(0,0,0,0.08);
		box-shadow: 0 0 10px rgba(0,0,0,0.1);
		.left-card {
			width: 65%;
			height: 100%;
			border-right: 1px solid rgba(0,0,0,0.1);;
			.note-cover {
				width: 100%;
				height: 100%;
				.outside{
					height: 100%;
					.note-img {
						max-width: 100%;
    				max-height: 100%;
						width: 100%;
						height: 100%;
						object-fit: contain;
    				overflow: clip;
					}
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
							// background-color: $primary-color-red;
							border-radius: 100px;
							background: transparent;
							border: 1px solid rgba(0,0,0,0.08);
							color: rgba(51,51,51,0.8);
						}
						.active{
							background-color: $primary-color-red;
							color: #fff;
						}
					}
				}
			}
			.note-scoller{
				transition: scroll .4s; // 滚动效果
				overflow-y: scroll; // 纵向滚动
				flex-grow: 1;  // 自适应高度
				.note-content{
					padding: 0 20px 20px;
					.title{
						margin-bottom: 8px;
						font-weight: 600;
						font-size: 18px;
						line-height: 140%;
					}
					.desc{
						margin: 0;
						font-weight: 400;
						font-size: 16px;
						line-height: 150%;
						color: #333;
						white-space: pre-wrap;
						overflow-wrap: break-word;
						.detail-desc{
							margin-top: 12px;
							span{
								font-size: 16px;
								line-height: 150%;
								color: #333;
							}
						}
					}
					.bottom-container{
						display: flex;
						justify-content: space-between;
						align-items: center;
						margin-top: 12px;
						.date{
							font-size: 14px;
							line-height: 120%;
						}
					}
					
				}
				.divider{
						margin: 0 20px;
						list-style: none;
						height: 1px;
						border: solid rgba(0,0,0,0.08);
						border-width: 1px 0 0;
				}
				.comment-el{
					position: relative;
					// height: 100%;
					.comment-container{
						padding: 16px 8px;
						.total{
							font-size: 14px;
							color: rgba(51,51,51,0.6);
							margin-left: 8px;
							margin-bottom: 12px;
						}
					}
				}
			}
			.fotter{
				padding: 16px 16px env(safe-area-inset-bottom); // 底部安全区域
				border-top: 1px solid rgba(0,0,0,0.08);  // 顶部边框
				flex-basis: unset; // 重置flex-basis 用来撑开高度
				height: unset;  // 重置高度 用来撑开高度
				.bar-container{
					width: 100%;
    			position: relative;
					// height: 100%;
					// width: 100%;
				}
			}
		}
	}
}
::v-deep(.el-carousel__container){
	height: 100%;
	img{
		width: 100%;
		height: 100%;
		object-fit: cover;
	}
}
::v-deep(.el-carousel){
	height: 100%;
}
::v-deep(.el-carousel__container img){
	// height: 100%;
	object-fit: contain;
}
</style>
