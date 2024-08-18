<template>
  <div class="layout-container">
    <el-container>
      <el-header class="header">这里是头哦</el-header>
      <el-container class="container">
        <el-aside width="200px">这里是左侧</el-aside>
        <el-main>
          <p style="height: 1000000px;">这里是主要的main</p>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
</script>

<style lang="scss" scoped>
.layout-container{
	height: 100vh;
	.el-container{
		height: calc(100vh - 64px);
		overflow:auto;
	}
	.el-header{
		background-color: $primary-header-bg-color;
	}
}
</style>
