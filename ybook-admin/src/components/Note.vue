<template>
    <div class="maincontainer">
        <div class="scroll-container">
            <div class="box box-active">推荐</div>
            <div class="box">摄影</div>
            <div class="box">教育</div>
            <div class="box">穿搭</div>
            <div class="box">美食</div>
            <div class="box">育儿</div>
            <div class="box">彩妆</div>
            <div class="box">健身</div>
        </div>
    </div>
</template>

<script setup lang="ts">
</script>


<style scoped>
.scroll-container{
    display: flex;
    justify-content: space-between;
    align-items: center;
    user-select: none;
    -webkit-user-select: none;
}
.box{
    height: 40px;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 0 16px;
    cursor: pointer;
    -webkit-user-select: none;
    user-select: none;
    color: #afafaf;
}
.box-active{
    background-color: rgba(0,0,0,0.03);
    color: #333;
    font-weight: bold;
    border-radius: 50px;
}
</style>