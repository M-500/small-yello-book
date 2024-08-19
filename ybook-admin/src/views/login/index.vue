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
              <div class="email-input">
                <el-input v-model="input1"
                          style="max-width: 600px"
                          placeholder="邮箱">
                  <template #prepend>Http://</template>
                </el-input>
              </div>
              <div class="code-input">

              </div>
            </div>
            <el-button>登录</el-button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import logo from '@/components/Logo/index.vue'
</script>

<style lang="scss" scoped>
.login-container{
  height: 100vh;
  width: 100%;
  .top-header{
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
        .title-box{
          margin-top: -60.78px;
          float: left;
          margin-left: 42px;
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
</style>
