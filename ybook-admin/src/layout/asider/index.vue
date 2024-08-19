<template>
  <div class="container">
    <div class="publish-video">
      <router-link class="yb-btn"
                   to="/publish">发布笔记</router-link>
    </div>
    <div class="menu-box">
      <YbMenu />
    </div>
  </div>
</template>

<script setup>
import YbMenu from '@/components/YbMenu/index.vue'
</script>

<style lang="scss" scoped>
.container{
	.publish-video{
		width: $primary-asider-width;
    height: 72px;
    padding: 16px;
    border-top: 1px solid #f3f3f3;
    border-bottom: 1px solid #f3f3f3;
    // border-right: 1px solid #f3f3f3;
    background: #fff;
		.yb-btn{
			display: block;
			width: 176px;
			height: 40px;
			line-height: 40px;
			font-size: 16px;
			font-weight: 500;
			border-radius: 4px;
			text-align: center;
			background: #ff2442;
			color: #fff;
			cursor: pointer;
			text-decoration: none;
		}
	}
	.menu-box{
		padding-top: 20px;
	}
}
</style>
