<template>
  <div class="header-container">
    <svg-icon @click="handleClick"
              name="logo"
              width="72"
              height="72" />
    <div class="title">创作服务平台</div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
let $router = useRouter()
function handleClick(){
	$router.push('/home')
}
</script>

<style lang="scss" scoped>
.header-container {
	display: flex;
	align-items: center;
	height: 72px;
	.title {
		font-size: 20px;
		font-weight: bold;
		color: #333;
		margin-left: 20px;
	}
}
</style>
