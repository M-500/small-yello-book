<template>
  <div class="layout_container">
    <div class="layout_slide">
      <Logo />
      <div class="sider_main">
        <div class="top">
          <Channel />
          <card v-if="!userStore.getToken" />
        </div>
        <div class="buttom">
          <more />
        </div>
      </div>
    </div>

    <div class="layout_header">
      <search />
    </div>

    <div class="layout_main">
      <RouterView />
    </div>
  </div>
</template>

<script lang="ts" setup>
import Logo from "./logo/index.vue";
import more from "./more/index.vue";
import Channel from "./channel/index.vue";
import card from "./card/index.vue";
import search from "../components/search/index.vue";
import { RouterView } from 'vue-router'
import useUserStore from '@/stores/modules/user';
const userStore = useUserStore();
</script>

<style scoped lang="scss" >
.layout_container {
	width: 100%;
	height: 100vh;
	background-color: $layout-body-background;
	.layout_slide{
		width: $layout-sider-width;
		height: 100vh; // 这里用100vh是为了让左侧栏高度和屏幕一样高 不用100%是因为100%是相对于父元素的高度
		background-color: $layout-sider-background;// 使用左侧栏的背景色
		float: left;
		.sider_main{
			height: calc(100% - $layout-header-height);
			padding-top: 16px;
			margin-left: 16px;
			display: flex;
			flex-direction: column;
			justify-content: space-between;
			.top{
				display: flex;
				flex-direction: column;
				justify-content: flex-start;
			}

		}
	}
	.layout_header{
		width: calc(100% - $layout-sider-width);
		height: $layout-header-height;
		background-color: $layout-body-background;
		// 固定定位
		position: fixed;
		top: 0;
		left: $layout-sider-width; // 这里是为了让顶部栏和左侧栏对齐
		
		display: flex;
		align-items: center;

		z-index: 1000; /* 确保在 container 之上 */
	}
	.layout_main{
		width: calc(100% - $layout-sider-width);
		height: calc(100% - $layout-header-height);
		background-color: $layout-body-background;
		// 绝对定位
		position: absolute;
		top: $layout-header-height;
		left: $layout-sider-width;
		padding: 16px 32px 0; // 这里是为了让内容不要贴边 24 32
		overflow: auto; // 这里是为了防止撑开整个页面，出现滚动条
		
	}
}

@media (max-width: 768px){
	.layout_container{
		.layout_slide{
			display: none;
		}
		.layout_header{
			width: 100%;
			left: 0;
		}
		.layout_main{
			width: 100%;
			left: 0;
			padding: 16px 12px 0; // 这里是为了让内容不要贴边 24 32
			overflow: auto; // 这里是为了防止撑开整个页面，出现滚动条
		}
	}
}
</style>
 