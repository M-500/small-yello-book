<template>
  <div class="title-container">
    <div class="title-box">
      <span class="icon"></span>
      <span class="title-name">{{ title }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps({
	// 接受父组件传递的参数
	title: {
		type: String,
		default: ''
	},
	paddingTop: {
		type: String,
		default: '24px'
	}
})
</script>

<style lang="scss" scoped>
.title-container{
	padding: 24px 0;
	.title-box{
		display: flex;
		align-items: center;
		float: left;
		.icon{
			display: inline-block;
			width: 5px;
			height: 18px;
			margin-right: 8px;
			border-radius: 16px;
			background: rgb(58, 100, 255);
		}
		.title-name{
			font-size: 20px;
			font-weight: 500;
			color: rgb(45, 45, 45);
			padding-right: 9px;
		}
	}
}
</style>
