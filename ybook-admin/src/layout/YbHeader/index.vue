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
          <el-menu-item index=""
                        @click="open">退出登录</el-menu-item>
        </el-sub-menu>
      </el-menu>
    </div>
  </div>
</template>

<script setup lang="ts">
import logo from '@/components/Logo/index.vue'
import useUserStore from '@/stores/moudules/user';
import { ElMessage, ElMessageBox } from 'element-plus'
import { useRouter } from 'vue-router';
let userStore = useUserStore()
let $router = useRouter()
const open = () => {
	ElMessageBox.confirm(
    '确定要退出登录吗？',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  )
    .then(() => {
			userStore.userLogout()
			$router.push('/login')
      ElMessage({
        type: 'success',
        message: '退出登录成功',
      })
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '',
      })
    })
}
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
::v-deep(.el-sub-menu__title.el-tooltip__trigger.el-tooltip__trigger):hover{ 
	color: $primary-buttom-active-colder;
}

::v-deep(.el-message-box){
	border-radius: 10px;
}
</style>
