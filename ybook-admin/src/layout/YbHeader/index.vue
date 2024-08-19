<template>
  <div class="container">
    <div class="left-logo">
      <logo />
    </div>
    <div class="right-info">
      <yb-user-header></yb-user-header>
    </div>
  </div>
</template>

<script setup lang="ts">
import logo from '@/components/Logo/index.vue'
import YbUserHeader from '@/components/YbUserHeader/index.vue'

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
	
}
::v-deep(.el-sub-menu__title.el-tooltip__trigger.el-tooltip__trigger):hover{ 
	color: $primary-buttom-active-colder;
}
::v-deep(.el-button--primary){
	background-color: $primary-buttom-active-colder !important;
}
</style>

