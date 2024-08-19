<template>
  <div class="container">
    <div class="left-logo">
      <logo />
    </div>
    <div class="right-info">
      <el-menu :default-active="activeIndex"
               class="el-menu-demo"
               mode="horizontal"
               :ellipsis="false"
               @select="handleSelect">
        <el-sub-menu index="2">
          <template #title>小红薯AbCs123</template>
          <el-menu-item index="2-1">退出登录</el-menu-item>
        </el-sub-menu>
      </el-menu>
    </div>
  </div>
</template>

<script setup>
import logo from '@/components/Logo/index.vue'
</script>

<style lang="scss" scoped>
.container{
	display: flex;
	background-color: white;
	justify-content: space-between;
	// box-shadow: rgba(0, 0, 0, 0.09) 0px 0px 20px 0px;
	padding: 0 20px;
	.left-logo{
		display: flex;
		align-items: center;
	}
	.right-info{
		display: flex;
		align-items: center;
	}
	.el-menu-demo{
		background-color: transparent;
		border: none;
		.el-menu-item{
			color: #fff;
		}
		.el-sub-menu{
			color: #fff;
		}
	}
}
</style>
