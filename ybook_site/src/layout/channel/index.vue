<template>
  <div class="channel_container">
    <router-link class="item"
                 v-for="(tab,index) in tabs"
                 :key="index"
                 :class="{active: selectTab === tab.name}"
                 @click="selectTab = tab.name"
                 :to="tab.to">
      <div class="icon">
        <img :src="tab.icon"
             alt="" />
        <div class="count"
             v-if="tab.count > 0 ">{{ tab.count }}</div>
      </div>
      <div class="title">{{ tab.label }}</div>
    </router-link>
    <router-link class="item"
                 v-if="userStore.isLogin()"
                 :to="`/user/profile/${userInfo.globalNumber}`">
      <div class="icon">
        <img :src="userInfo.avatar ? userInfo.avatar :'../../assets/imgs/avatar.jpeg'"
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
import { ref } from 'vue';
import homeIcon from '@/assets/icons/home.svg';
import publishIcon from '@/assets/icons/publish.svg';
import msgIcon from '@/assets/icons/msg.svg';

const userStore = useUserStore();
const userInfo  = ref(userStore.getUserInfo);
const tabs = ref([
	{ name: 'home', label:'发现',icon:homeIcon,to:'/home',count:0},
	{ name: 'publish', label:'发布',icon:publishIcon,to:'/publish',count:0},
	{ name: 'notices', label:'通知',icon:msgIcon,to:'/notices',count:4},
]);
const selectTab = ref('home');

// function getImageUrl(icon) {
// 	return require(`${icon}`);
// }




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
