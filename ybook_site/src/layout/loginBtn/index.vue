<template>
  <div class="login_btn">
    <button class="btn">登录</button>
  </div>
</template>

<script setup name="loginBtn">
import { } from "vue"
</script>

<style lang="scss" scoped>
.login_btn{
	font-size: 16px;
	height: 48px;
	width: 100%;
	font-weight: 600;
	margin-bottom: 8px;
	.btn{
		width: 100%;
		line-height: 16px;
    padding: 0 24px;
		font-size: 16px;
    font-weight: 600;
		height: 40px;
		background-color: $primary-color-red;
    color: $text-color;
		border-radius: 100px;
	}
}
</style>
