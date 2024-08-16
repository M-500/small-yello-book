<template>
  <router-link class="logo_container"
               to="/home">
    <img src="../../assets/icons/logo.svg"
         alt="logo" />
  </router-link>
</template>

<script setup>

</script>

<style lang="scss" scoped>
.logo_container {
	width: 100%;
	height: $layout-header-height;
	background-color: $layout-sider-background;
	display: flex;
	justify-content: flex-start;
	padding-left: 32px;
	align-items: center;
	img{
		width: 50px;
		height: 50px;
	}
}
</style>
