<template>
    <div class="sider_bar">
        
        <div class="item">
            <div class="box active-box">
                <div class="icon">
                    <img src="../../assets/icon/home.svg" alt="">
                </div>
                <div class="text">发现</div>
            </div>
            <div class="box">
                <div class="icon">
                    <img src="../../assets/icon/publish.svg" alt="">
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
        <Login />
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import Login from '../Login.vue';
const dialogVisible: Ref<boolean> = ref(false);
function loginBtnClick(){
    console.log('click');
    dialogVisible.value = true;
}
</script>
<style scoped>
.sider_bar{
    /* 高度撑满屏幕 */
    height: calc(100vh - 72px);
    background-color: #ffffff;
    padding-left: 16px;
}
.item{
    padding-top: 16px;
    /* height: 150; */
}


.active-box{
    background: rgba(51,51,51,0.1);
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

.logo img{
    height: 24px;
}
</style>