<template>
  <div class="tab_item"
       :class="active === true ? 'active': ''">
    <div class="container">
      <span>{{ name }}</span>
      <div v-show="tagCount > 0"
           class="count">{{ tagCount }}</div>
    </div>
  </div>

</template>

<script setup>


defineProps({
  active: {
    type: Boolean,
    default: false
  },
  name: {
    type: String,
    required: true
  },
  tagCount: {
    type: Number,
    default: 0
  }
})

</script>

<style lang="scss" scoped>
.tab_list{
	display: flex;
}
.tab_item {
	padding: 0px 16px;
	margin-right: 0px;
	font-size: 16px;
	display: flex;
	height: 40px;
	align-items: center;
	.container {
		width: 100%;
		height: 40px;
		padding: 16px;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-content: center;
		position: relative;
		span {
			color: #333;
			// 文字居中
			text-align: center;
		}
		.count{
			position: absolute;
			right: 14px;
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
	&:hover {
		font-weight: 500;
	}
}
.active {
	background-color: #f7f7f7;
	font-weight: 600;
	color: #333;
	border-radius: 50px;
}
</style>