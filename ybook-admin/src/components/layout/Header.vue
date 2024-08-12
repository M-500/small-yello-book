<template>
    <div class="main_header">
        <div class="logo">
            <img src="../../assets/logo.svg" alt="慕课网">
        </div>
        <div class="search">
            <el-input class="search_input" placeholder="登陆探索更多内容" :suffix-icon="Search"></el-input>
        </div>
        <div class="center">
            <button class="rds_btn" >创作中心</button>
            <button class="rds_btn">业务合作</button>
        </div>
    </div>
  
</template>

<script setup lang="ts">
import { Search } from '@element-plus/icons-vue'
</script>

<style scoped>

.main_header{
    height: 72px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px 0;
    background-color: #fff;

    line-height: 72px;
}

.logo img{
    height: 24px;
    margin-left: 16px;
}


.center{
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-right: 16px;
}


 /deep/ .el-input__wrapper{
    width: 300px;
    height: 40px;
    border-radius: 50px;
    background-color: #f5f5f5;
    padding: 0 84px 0 16px;
    color: #333;
    caret-color: #ff2442;
 }
 /deep/ .el-input__wrapper .is_focus {
    border: 1px solid #ff2e4d;
}
.rds_btn{
    height: 40px !important;
    color: rgba(51,51,51,0.6);
    font-size: 16px;
    line-height: 120%;
    padding: 0 16px;
    background: transparent;
    border: transparent;
}
</style>