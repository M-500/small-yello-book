<template>
  <div class="layout-container">
    <el-container class="main-container">
      <el-header class="header">
        <yb-header />
      </el-header>
      <el-container class="container">
        <el-aside class="yb-asider">
          <asider></asider>
        </el-aside>
        <el-main class="yb-main">
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import asider from '@/layout/asider/index.vue'
import YbHeader from '@/layout/YbHeader/index.vue'
import { RouterView } from 'vue-router'
</script>

<style lang="scss" scoped>
.layout-container{
	height: 100vh;
		.main-container{
			height: 100vh;
			.el-container{
				height: calc(100vh - $primary-header-height);
				overflow:auto;
				.yb-asider{
					width: $primary-asider-width;
					background-color: $primary-header-bg-color;
				}
			}
			.el-header{
				background-color: $primary-header-bg-color;
				box-shadow: rgba(0, 0, 0, 0.09) 0px 0px 20px 0px;
				unicode-bidi: isolate;
				height: $primary-header-height;
			}
			.yb-main{
				// min-width: 1200px;
				width: 100%;
				overflow: auto;
				box-sizing: content-box;
				padding: 20px 20px 32px;
				font-family: "PingFang SC";
				color: rgb(38, 38, 38);
				background-color: rgb(238, 241, 248);
			}
		}
}
</style>
