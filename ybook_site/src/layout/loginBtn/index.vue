<template>
  <div class="login-container">
    <div class="login_btn">
      <button class="btn"
              @click="handleClick">登录</button>
    </div>
    <el-dialog class="login-dialog"
               v-model="dialogTableVisible">
      <div class="content">
        <div class="left">

        </div>
        <div class="right">
          <div class="title">邮箱验证码登录</div>
          <div class="input-container">
            <el-form :model="form"
                     class="form"
                     label-width="auto">
              <el-form-item class="emial">

                <el-input placeholder="请输入邮箱"
                          v-model="form.email" />
                <SvgIcon name="email"
                         class="icon"
                         width="25"
                         height="25"></SvgIcon>
              </el-form-item>
              <el-form-item class="code">
                <el-input placeholder="请输入验证码"
                          v-model="form.ver_code"></el-input>
                <span class="verfi-code">获取验证码</span>
              </el-form-item>
              <div class="err-msg">

              </div>
              <el-form-item class="btn">
                <el-button type="primary"
                           @click="onSubmit">登陆</el-button>
              </el-form-item>
            </el-form>
          </div>
          <div class="agree"></div>
          <div class="user-tips"> 新用户可直接登录 </div>
        </div>
      </div>

    </el-dialog>
  </div>

</template>

<script setup lang="ts" name="loginBtn">
import { useRouter } from 'vue-router';
import { ref } from 'vue';
const dialogTableVisible = ref(true)
const form = ref({
	email: '',
	ver_code: ''
})
let $router = useRouter()
function handleClick () {
  dialogTableVisible.value = true
  // $router.push('/login')
}
</script>

<style lang="scss" scoped>
.login-container{
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
	.login-dialog{
		width: 800px;
		height: 100%;
		// border-radius: 50px;
		.content{
			width: 100%;
			display: flex;
			justify-content: space-between;
			
			.left{
				width: 400px;
				height: 100%;
				border-radius: 15px 0 0 15px;
				
			}
			.right{
				border-left: 1px solid rgba(0,0,0,0.08);
				padding-top: 28px;
				display: flex;
				align-items: center;
    		flex-direction: column;
				width: 400px;
				height: 100%;
				border-radius: 0 15px 15px 0;
				.title{
					font-size: 18px;
					color: #333;
					font-weight: 600;
					line-height: 120%;
					height: 24px;
				}
				.input-container{
					margin-top: 32px;
					padding: 16px 0;
					width: 304px;
					display: flex;
					flex-direction: column;
					.form{
						::v-deep(.el-input__wrapper){
							border-radius: 999px;
							padding: 0 22px;
							background: rgba(0,0,0,0.03);
						}
						.emial{
							margin-bottom: 16px;
							::v-deep(.el-input__wrapper){
								padding-left: 60px;
							}
							position: relative;
							.icon{
								position: absolute;
								left: 20px;
								top: 10px;
								align-content: center;
								font-size: 16px;
								color: #ff2442;
							}
						}
						.code{
							position: relative;
							margin-bottom: 16px;
							.verfi-code{
								position: absolute;
								height: 48px;
								right: 20px;
								top: 0;
								align-content: center;
								font-size: 16px;
								color: #ff2442;
							}
						}
						.err-msg{
							margin-top: 9.5px;
							height: 28px;
							line-height: 140%;
							color: #ff2442;
							font-size: 14px;
							font-weight: 400;
							display: flex;
							align-items: center;
							justify-content: center;
						}
						.btn{
							margin-top: 16px;
    					height: 48px;
							::v-deep(.el-button){
								width: 100%;
								border-radius: 999px;
								height: 48px;
								background: #ff2442;
							}
						}
					}
				}
				.agree{

				}
				.user-tips{
					height: 28px;
					margin-top: auto;
					margin-bottom: 32px;
					font-size: 14px;
					font-weight: 400;
					line-height: 120%;
					display: flex;
					align-items: center;
					justify-content: center;
					color:rgba(51,51,51,0.6);
				}
			}
		}
		
	}
}

::v-deep(.el-dialog){
	border-radius: 15px;
	width: 800px;
}
::v-deep(.el-input){
	height: 48px;
	border-radius: 999px;
}


</style>
