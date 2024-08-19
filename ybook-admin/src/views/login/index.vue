<template>
  <div class="login-container">
    <div class="top-header">
      <logo />
    </div>
    <div class="content">
      <div class="con">
        <div class="video-box-container">
          <div class="mask"
               style="width: 697.76px; height: 418.656px;"></div>
          <div class="video-item-container"
               style="width: 178px;height: 178px;margin-top: 77px;margin-left: 0px;">
            <div class="circle balloon">
              <video autoplay
                     muted
                     loop
                     src="@/assets/videos/demo.mp4"></video>
            </div>
          </div>
          <div class="video-item-container"
               style="width: 222.163px; height: 222.163px; margin-top: 86.6436px; margin-left: 181.433px; border-radius: 222.163px;">
            <div class="circle balloon">
              <video autoplay
                     muted
                     loop
                     src="@/assets/videos/demo1.mp4"></video>
            </div>
          </div>
          <div class="video-item-container"
               style="width: 186.37px; height: 186.37px; margin-top: 195.627px; margin-left: 385.083px; border-radius: 186.37px;">
            <div class="circle balloon">
              <video autoplay
                     muted
                     loop
                     src="@/assets/videos/demo2.mp4"></video>
            </div>
          </div>
          <div class="video-item-container"
               style="width: 127.127px; height: 127.127px; margin-top: -9.87392px; margin-left: 335.096px; border-radius: 127.127px;">
            <div class="circle balloon">
              <video autoplay
                     muted
                     loop
                     src="@/assets/videos/demo3.mp4"></video>
            </div>
          </div>
          <div class="video-item-container"
               style="width: 120.956px; height: 120.956px; margin-top: 73.5607px; margin-left: 435.07px; border-radius: 120.956px;">
            <div class="circle balloon">
              <video autoplay
                     muted
                     loop
                     src="@/assets/videos/demo5.mp4"></video>
            </div>
          </div>
          <div class="video-item-container"
               style="width: 77.7571px; height: 77.7571px; margin-top: 157.366px; margin-left: 539.363px; border-radius: 77.7571px;">
            <div class="circle balloon">
              <video autoplay
                     muted
                     loop
                     src="@/assets/videos/demo6.mp4"></video>
            </div>
          </div>
          <div class="title-box">
            <span class="title-line">
              加入我们
              <br>
              解锁创作者专属功能
            </span>
            <span class="des">让创作发布、数据分析、商业变现更高效。</span>
          </div>
        </div>
        <div class="login-box-container">
          <div class="login-box">
            <div class="title">邮箱验证码登录</div>
            <div class="form-content">
              <div class="email-input-box">
                <el-input v-model="formData.email"
                          class="email-input"
                          placeholder="邮箱"
                          clearable />
              </div>
              <div class="code-input-box">
                <el-input v-model="formData.ver_code"
                          :disabled="canSend"
                          class="code-input"
                          max="6"
                          placeholder="验证码">
                </el-input>
                <div class="slot-right">
                  <div class="box">
                    <!-- <span class="send-btn">发送验证码</span> -->
                    <button class="send-btn"
                            :class="canSend ? 'active': ''"
                            @click="sendEmailBtn">{{ buttonText }}</button>
                  </div>
                </div>
              </div>
              <div class="help-btn">收不到验证码？</div>
            </div>
            <el-button @click="loginHandler">登录</el-button>
            <!-- :disabled="canSend" -->
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// 引入用户相关的pinia，用来存储token信息
import useUserStore from '@/stores/moudules/user';
import logo from '@/components/Logo/index.vue'
import { type emialForm } from '@/api/user/types';
import { reactive, ref, watch } from 'vue';
import { getEmailCaptchas } from '@/api/user';
import { ElMessage } from 'element-plus';
import { useRouter } from 'vue-router';
// const email = ref('1978992154@qq.com');
// const ver_code = ref('443254');
const formData = reactive({
  email: '1978992154@qq.com',
  ver_code: '123123'  
})
let $router = useRouter()
const canSend = ref(true);
const buttonText = ref('发送验证码');
let countdown = 60;
let timer = null;
let userStore = useUserStore()
// 监听Email输入框的值
watch(() => formData.email, (val) => {
  if (val) {
    canSend.value = true;
  } else {
    canSend.value = false;
  }
});

// 发送验证码点击事件
function sendEmailBtn () {
  let form = reactive({
    emial: formData.email,
    type_code: '1'
  })
  if (canSend.value) {
    getEmailCaptchas(form).then(res => {
      ElMessage.success('验证码发送成功');
    }).catch(err => {
      console.log(err);
    });
    canSend.value = false;
    buttonText.value = `${countdown}s后重发`;

    timer = setInterval(() => {
      countdown--;
      buttonText.value = `${countdown}s后重发`;

      if (countdown === 0) {
        clearInterval(timer);
        countdown = 60;
        buttonText.value = '发送验证码';
        canSend.value = true;
      }
    }, 1000);
  }
}

const loginHandler = async()=>{
  try {
    await userStore.userLogin(formData)
    $router.push('/home')
  }catch (err){
    ElMessage.success(err)
  }
}

// 登录点击事件
// function loginHandler () {
  
//   // login({
//   //   email: email.value,
//   //   ver_code: ver_code.value
//   // }).then(res => {
//   //   ElMessage.success('登录成功');

//   //   let jwt = res.token
//   //   console.log(res, jwt)
//   //   userStore.setToken(jwt) // 设置JWT信息
//   //   // 获取JWT信息
//   // }).catch(err => {
//   //   console.log("")
//   // });
// }
</script>

<style lang="scss" scoped>
.login-container{
  height: 100vh;
  width: 100%;
  .top-header{
    padding: 0 20px;
    height: $primary-header-height;
    background-color: #fff;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  .content{
    height: calc(100vh - #{$primary-header-height});
    width: 100%;
    background-image: url('@/assets/imgs/background.png'); // 确保路径正确
    background-size: cover; // 使背景图片覆盖整个容器
    background-position: center; // 使背景图片居中
    position: absolute;
    display: flex;
    justify-content: space-between;
    align-items: center;
    flex-direction: column;
    .con{
      width: 100%;
      max-width: 1546px;
      height: calc(100% - 84px);
      padding-top: 20px;
      display: flex;
      justify-content: center;
      align-items: center;
      .video-box-container{
        width: 663.04px;
        height: 397.824px;
        top: 0;
        .mask{
          position: absolute;
          z-index: 10;
        }
        .video-item-container{
          position: absolute;
          display: flex;
          justify-content: center;
          align-items: center;
          .circle{
            display: flex;
            justify-content: center;
            align-items: center;
            overflow: hidden;
            border-radius: 100%;
            width: 85%;
            height: 85%;
          }
        }
        .title-box{
          margin-top: -60.78px;
          float: left;
          margin-left: 42px;
          position: relative;
          top: 0;
          .title-line{
            font-size: 2.5vw;
            line-height: 150%;
            margin-bottom: .7vw;
            display: block;
            width: 100%;
            font-size: 36px;
            font-weight: 600;
            line-height: 150%;
            margin: 0 0 10px;
            color: #363f4d;
          }
        }
      }
    }
  }
}

.login-box-container{
  width: 320px;
  height: 362px;
  margin-left: 44px;
  box-shadow: 0 0 40px rgba(0, 0, 0, .1);
  border-radius: 16px;
  .login-box{
    padding: 40px 32px;
    width: 320px;
    height: 363px;
    border-radius: 16px;
    background-color: white;
    position: relative;
    .title{
      font-size: 24px;
      font-weight: 500;
      color: #333;
      margin-bottom: 20px;
    }
    .form-content{
      height: 180px;
      .email-input-box{
        margin-bottom: 20px;
        .email-input{
          height: 48px;
          border-color: rgb(228, 234, 242);
          border-radius: 8px;
          font-size: 14px;
        }
      }
      .code-input-box{
        position: relative;
        // margin-bottom: 20px;
        .code-input{
          height: 48px;
          border-color: rgb(228, 234, 242);
          border-radius: 8px;
          font-size: 14px;
        }
        .slot-right{
          height: 48px;
          position: absolute;
          right: 10px;
          top: 0px;
          bottom: 1px;
          overflow: hidden;
          display: flex;
          -webkit-box-align: center;
          align-items: center;
          border-top-right-radius: 3px;
          border-bottom-right-radius: 3px;
          .box{
            border-left: 1px solid rgb(228, 234, 242);
            display: flex;
            align-content: center;
            .send-btn{
              color: rgb(136, 136, 136);
              cursor: not-allowed;
              padding: 2px 0px;
              margin-left: 14px;
              font-size: 14px;
              width: 70px;
              text-align: center;
              cursor: pointer;
            }
            .active{
              color: rgb(22, 119, 255);
              cursor: pointer;
            }
          }
        }
      }
      .help-btn{
        margin-top: 5px;
        color: rgb(22, 119, 255);
        font-size: 12px;
        cursor: pointer;
      }
    }
    
  }
}
::v-deep(.el-button){
  width: 100%;
  min-width: 88px;
  width: 256px;
  height: 48px;
  font-size: 16px;
  font-weight: 500;
  margin-bottom: 12px;
  background-color: blue;
  border-radius: 8px;
  span{
    color: white;
  }
}
::v-deep(.el-input__wrapper){
  border-radius: 8px;
}

::v-deep(.email-input){
  width: 100%;
}

@keyframes balloon-move {
  0%, 100% {
    transform: translate(0, 0);
  }
  25% {
    transform: translate(2px, -2px);
  }
  50% {
    transform: translate(-2px, 2px);
  }
  75% {
    transform: translate(2px, 2px);
  }
}

.balloon {
  animation: balloon-move 3s ease-in-out infinite;
}
</style>
