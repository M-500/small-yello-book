<template>
  <div class="container">
    <div class="box"
         @click="showLogout">
      <img src="@/assets/imgs/curry.jpeg"
           alt="">
      <span class="user-name">巧舌如簧的哑巴</span>
      <el-icon class="icon">
        <ArrowDown />
      </el-icon>

      <div class="menu-list"
           v-show="showLogoutVisable"
           @click="open">
        <div class="popover_text">
          <el-icon class="logout-icon">
            <SwitchButton />
          </el-icon>
          <span class="inline">退出登录</span>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import useUserStore from '@/stores/moudules/user';
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue';
import { useRouter } from 'vue-router';
let userStore = useUserStore()
let $router = useRouter()
const showLogoutVisable = ref(false)

function showLogout(){
	showLogoutVisable.value = !showLogoutVisable.value
}
const open = () => {
	ElMessageBox.confirm(
    '确定要退出登录吗？',
    {
			customClass: 'message-logout',
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
      // ElMessage({
      //   type: 'info',
      //   message: '取消',
      // })
    })
}
</script>

<style lang="scss" scoped>
.container{
	.box{
		display: flex;
		justify-content: center;
		img{
			width: 28px;
			height: 28px;
			border-radius: 100%;
			margin-right: 4px;
		}
		.user-name{
			display: flex;
			justify-content: center;
			align-items: center;
		}
		.icon{
			height: 100%;
			margin-left: 7px;
			display: flex;
			justify-content: center;
			align-items: center;
		}
		.menu-list{
			position: absolute;
			width: 164px;
			right: 40px;
			top: 50px;
			padding: 8px 0;
			background: #fff;
			box-shadow: 0 2px 8px #e4e6f2;
			border-radius: 4px;
			display: flex;
			justify-content: center;
			align-items: center;
			cursor: pointer;
			.popover_text{
				width: 100%;
				height: 40px;
				// padding-left: 18px;
				font-family: PingFang SC;
				font-style: normal;
				font-weight: 400;
				font-size: 14px;
				line-height: 14px;
				color: #2d2d2d;
				display: flex;
				justify-content: flex-start;
				align-items: center;
				.logout-icon{
					display: flex;
					justify-content: center;
					align-items: center;
					margin-right: 14px;
					margin-left: 14px;
				}
				span{
					font-weight: 900;
				}
			}
			.popover_text:hover{
				background: #f5f5f5;
				span{
					color: $primary-buttom-active-colder;
				}
			}
		}
	}
}


</style>

<style >
/* el-message-box的样式修改只能这么玩 */
.message-logout{
	border-radius: 10px !important;
}
::v-deep(.el-button--primary){
	background-color: rgb(58, 100, 255) !important;
}
</style>