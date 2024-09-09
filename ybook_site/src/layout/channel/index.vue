<template>
  <div class="channel_container">
    <router-link class="item active"
                 to="/home">
      <div class="icon">
        <img src="../../assets/icons/home.svg"
             alt="" />
      </div>
      <div class="title">发现</div>
    </router-link>
    <router-link class="item"
                 to="/publish">
      <div class="icon">
        <img src="../../assets/icons/publish.svg"
             alt="" />
      </div>
      <div class="title">发布</div>
    </router-link>
    <router-link class="item"
                 to="/notices">
      <div class="icon">
        <img src="../../assets/icons/msg.svg"
             alt="" />
        <div class="count">4</div>
      </div>
      <div class="title">通知</div>
    </router-link>
    <router-link class="item"
                 v-if="userStore.isLogin()"
                 to="/user/profile/1">
      <div class="icon">
        <img src="../../assets/imgs/avatar.jpeg"
             class="avatar"
             alt="" />
      </div>
      <div class="title">我</div>
    </router-link>
    <login_btn v-if="!userStore.isLogin()" />
  </div>
</template>

<script lang="ts" setup>
import useUserStore from '@/stores/modules/user';
import login_btn from "../loginBtn/index.vue";
import { RouterLink } from "vue-router";

const userStore = useUserStore();


</script>

<style lang="scss" scoped>
.channel_container{
	.active{
		background-color: $active-color;
	}
	.item{
		display: flex;
		padding-left: 16px;
		margin-bottom: 8px;
		align-items: center;
		height: 45px;
		cursor: pointer;
		border-radius: 50px;
		text-decoration: none; // 去掉a标签的下划线
		.icon {
			position: relative;
			img {
				width: 24px;
				height: 24px;
			}
			.avatar{
				border-radius: 50%;
			}
			.count{
				position: absolute;
				top: 0;
				right: 0;
				background-color: $primary-color-red;
				color: white;
				display: flex;
				align-items: center;
				justify-content: center;
				padding: 0 4px;
				min-width: 16px;
				height: 16px;
				border-radius: 999px;
				font-size: 12px;
				font-weight: 500;
				transform: translate(calc(100% - 4px), calc(-100% + 4px));
				z-index: 1;
			}
		}
		.title{
			font-size: 16px;
			font-weight: 600;
			margin-left: 12px;
			color: $gray-color;
		}
		&:hover{
			background-color: $active-color;
		}
	}
}
</style>
