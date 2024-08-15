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
    <div class="content">
      <div class="session">
        <Card />
      </div>
      <div class="session">
        <Card />
      </div>
      <div class="session">
        <Card />
      </div>
      <div class="session">
        <Card />
      </div>
      <div class="session">
        <Card />
      </div>
      <div class="session">
        <Card />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import Card from './Card.vue'
</script>

<style scoped>
.scroll-container {
  display: flex;
  justify-content: space-start;
  align-items: center;
  user-select: none;
  -webkit-user-select: none;
}
.box {
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
.box-active {
  background-color: rgba(0, 0, 0, 0.03);
  color: #333;
  font-weight: bold;
  border-radius: 50px;
}
.content {
  padding-top: 10px;
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  /* 每行显示最多4个 */
  justify-content: space-between;
  align-items: center;
}
.session {
  margin-right: 25px;
  margin-bottom: 16px;
}
</style>
