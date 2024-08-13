<template>

  <el-dialog v-model="dialogVisible">
    <div class="login_container">
      <div class="left">
        <div class="header">
          <div class="login_reason">登陆后推荐更懂你的笔记</div>
          <div class="logo">
            <img src="../assets/logo.svg" alt="">
          </div>
        </div>

        <div class="wechar_code">
          <img src="../assets/background.png" alt="">
        </div> 
      </div>

      <div class="right">
        <div class="login_title">
          账号密码登录
        </div>
        <div class="login_form" >
          <el-input class="form_item" v-model="LoginForm.phone" placeholder="请输入手机号"></el-input>
          <el-input class="form_item" v-model="LoginForm.password"  placeholder="请输入密码"></el-input>
          <el-button type="primary" class="login_btn">登陆</el-button>
        </div>
        <div class="protocl">
          <el-checkbox class="check_agree"></el-checkbox>
          <span>我已阅读并同意
            <a href="#">《用户协议》</a></span>
        </div>

      </div>
    </div>
  </el-dialog>

</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
const dialogVisible: Ref<boolean> = ref(true);

const LoginForm = reactive({
  phone: '18574945291',
  password: 'wulinlin'
});

function handleClose(done) {
  dialogVisible.value = false;
  done();
}
</script>


<style scoped>

.login_container{
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  height: 100%;
}

.left{
  width: 400px;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  /* align-items: center; */
  border-right: 1px solid #ebebeb;
}
.left .header{
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.left .header .logo img{
  width: 100px;
}
.left .wechar_code{
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  padding: 20px;
}
.wechar_code img{
  width: 100%;
  height: 100%;
}
.login_reason{
  height: 48px;
  padding: 0 20px;
  background:  #d0e3ff;
  color: #3d8af5;
  border-radius: 999px;
  font-size: 16px;
  font-weight: 600;
  line-height: 120%;
  margin-bottom: 20px;
  max-width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}
.right{
  margin: 0;
  width: 400px;
  display: flex;
  align-items: center;
  flex-direction: column;
  padding-top: 48px;
}

.login_form{
  margin-top: 32px;
  padding: 16px 20px;
  width: 304px;
  display: flex;
  flex-direction: column;
}

.el-dialog{
  border-radius: 16px;
}
::v-deep(.el-dialog) {
  /* display: flex;
  position: relative;
  justify-content: center; */
  width: 800px;
  height: 480px;
  border-radius: 16px;
  /* padding: 0; */
}
::v-deep(.el-input__inner) {
  padding: 12px;
}
.login_btn{
  background-color: #ff2e4d;
  border: #ebebeb;
  border-radius: 50px;
  width: 100%;
  height: 48px;
  margin-top: 16px;
}
.login_btn:hover{
  background-color: #ff2e4d;
  border: #ebebeb;
  border-radius: 50px;
  width: 100%;
  height: 48px;
  margin-top: 16px;
}
::v-deep(.el-button--primary){
  background-color: #ff2e4d;
}

.form_item{
  margin-top: 16px;
  height: 48px;
  font-size: 16px;
  caret-color: #ff2e4d;
  color: #333;
}

::v-deep(.el-input__wrapper){
  border-radius: 50px;
}

::v-deep(.is_focus) {
  border: 1px solid #ff2e4d;
}
.login_title{
  font-size: 18px;
  color: #333;
  font-weight: 600;
  height: 24px;
  line-height: 120%;
}

.protocl{
  /* width: 100%; */
  padding: 0px 48px;
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
}
.protocl span{
  margin-left: 4px;
}

</style>
