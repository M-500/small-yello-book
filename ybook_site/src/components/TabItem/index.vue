<template>
  <div class="tab_item"
       :class="active === true ? 'active': ''">
    <div class="container">
      <span>{{ name }}</span>
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
		span {
			color: #333;
			// 文字居中
			text-align: center;
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