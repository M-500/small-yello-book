<template>
    <div class="sider_bar">
        
        <div class="item">
            <div class="box active-box">
                <div class="icon">
                    <img src="../../assets/icon/msg.svg" alt="">
                </div>
                <div class="text">发现</div>
            </div>
            <div class="box">
                <div class="icon">
                    <img src="../../assets/icon/msg.svg" alt="">
                </div>
                <div class="text">发布</div>
            </div>
            <div class="box">
                <div class="icon">
                    <img src="../../assets/icon/msg.svg" alt="">
                </div>
                <div class="text">通知</div>
            </div>
           
            <div class="box-login" @click="loginBtnClick">
                登 陆
            </div>
           
        </div>

        <el-dialog
    v-model="dialogVisible"
    width="500"
  >
    <div class="login_container">
        <div class="left">
            <div class="mark">登陆后推荐更懂你的笔记</div>
            <div class="logo">
                <img src="../../assets/logo.svg" alt="慕课网">
            </div>
            <div class="wechar_code">
                微信二维码
            </div>
        </div>

        <div class="right">
            <div class="login_title">
                账号密码登录
            </div>
            <div class="login_form">
                <el-input class="form_item" placeholder="请输入手机号"></el-input>
                <el-input class="form_item" placeholder="请输入密码"></el-input>
                <el-button type="primary" class="login_btn">登陆</el-button>
            </div>
            <div class="protocl">
                <el-checkbox></el-checkbox>
                <a href="#">我已阅读并同意《用户协议》</a>
            </div>

        </div>
    </div>
  </el-dialog>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const dialogVisible: Ref<boolean> = ref(false);
function loginBtnClick(){
    console.log('click');
    dialogVisible.value = true;
}
</script>
<style scoped>
.sider_bar{
    /* 高度撑满屏幕 */
    height: 100vh;
    /* width: calc(-18px + 25vw); */
    background-color: #ffffff;
}
.item{
    padding-top: 16px;
    /* height: 150; */
}


.active-box{
    background-color: rgba(190, 187, 187, 0.6);
}

.box{
    display: flex;
    padding-left: 16px;
    /* justify-content: space-start; */
    margin-bottom: 8px;
    align-items: center;
    height: 45px;
    cursor: pointer;
    border-radius: 50px;
}

.box-login{
    display: flex;
    padding-left: 16px;
    justify-content: center;
    height: 45px;
    /* 上下居中 */
    align-items: center;
    color: #ffffff;
    border-radius: 50px;
    background-color: #ff2e4d;
}

.icon img{
    width: 24px;
    height: 24px;
    margin-right: 10px;
}

.login{
    position: absolute;
    bottom: 0;
    width: 100%;
    padding: 16px 0;
    text-align: center;
    background-color: #f5f5f5;
    border-top: 1px solid #ebebeb;
}

.login_container{
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 16px;
    height: 100%;
}
.logo img{
    height: 24px;
}
.left{
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    border-right: 1px solid #ebebeb;
}

/deep/.el-dialog{
    display: flex;
    position: relative;
    justify-content: center;
    width: 800px;
    height: 480px;
    border-radius: 16px;
}

.login_btn{
    background-color: #ff2e4d;
    border: #ebebeb;
    border-radius: 50px;
    width: 100%;
    height: 48px;
    margin-top: 16px;
}
/deep/ .login_btn:hover{
    background-color: #ff2e4d;
}

.form_item{
    margin-top: 16px;
    height: 48px;
    font-size: 16px;
    caret-color: #ff2e4d;
    color: #333;
}

/deep/ .el-input__wrapper{
    border-radius: 50px;
}

/deep/  .is_focus {
    border: 1px solid #ff2e4d;
}
</style>