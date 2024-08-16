<template>
  <div class="card">
    <div class="title">马上登录即可</div>
    <div class="line_container">
      <svg-icon name="nice"
                class="icon"
                width="16"
                height="16" />
      <span>刷到更懂你的优质内容</span>
    </div>
    <div class="line_container">
      <svg-icon name="zhongcao"
                class="icon"
                width="16"
                height="16" />
      <span>搜索最新种草、拔草信息</span>
    </div>
    <div class="line_container">
      <svg-icon name="star"
                class="icon"
                width="16"
                height="16" />
      <span>查看收藏、点赞的笔记</span>
    </div>
    <div class="line_container">
      <svg-icon name="talk"
                class="icon"
                width="16"
                height="16" />
      <span>与他人更好地互动、交流</span>
    </div>
  </div>
</template>

<script setup>
import { } from "vue"
</script>

<style lang="scss" scoped>
.card{
	cursor: pointer;
    width: 100%;
    z-index: 100010;
    padding: 16px;
    height: 155px;
    border: 1px solid rgba(0, 0, 0, 0.08);
    border-radius: 16px;
    background: $layout-body-background;
    // opacity: 0;
    z-index: 0;
	.title{
		font-size: 14px;
    font-weight: 500;
    color: $gray-color;
    margin-bottom: 10px;
	}
	.line_container{
		display: flex;
		align-items: center;
		margin-bottom: 6px;
		// flex-start 从左到右
		justify-content: flex-start;

    display: flex;
    line-height: 120%;
		.icon{
			margin-right: 8px;
		}
		span{
			font-size: 14px;
    	color: rgba(51, 51, 51, 0.6);
		}
	}
}
</style>
