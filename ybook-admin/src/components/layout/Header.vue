<template>
    <div class="main_header">
        <div class="logo">
            <img src="../../assets/logo.svg" alt="慕课网">
        </div>
        <div class="search">
            <el-input class="search_input" placeholder="登陆探索更多内容" :suffix-icon="Search"></el-input>
        </div>
        <div class="center">
            <el-button type="text">创作中心</el-button>
            <el-button type="text">业务合作</el-button>
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
    /* box-shadow: 0 2px 4px 0 rgba(0,0,0,.1); */
}

.logo img{
    height: 24px;
    margin-left: 16px;
}
/* 深度修改 覆盖elemet的样式 */
/* .search /deep/  .el-input--suffix{
    padding: 0 84px 0 16px;
    color: #333;
    caret-color: #ff2442;
    border-radius: 50px;
}
.search .el-input__suffix .el-input__suffix-inner{
    color: #333;
    font-size: 14px;
} */
</style>