<template>
  <div class="container">
    <!-- 404页面 -->
    <div class="not-found">
      <!-- <img src="@/assets/images/404.png" /> -->
      <p class="title">你访问的页面不在了</p>
      <p>3秒后自动返回首页</p>
      <button class="goback_btn">返回首页</button>
    </div>
  </div>
</template>

<script setup>

</script>

<style lang="scss" scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  .not-found {
    text-align: center;
    img {
      width: 100px;
      height: 100px;
    }
    p {
      font-size: 16px;
      color: #333;
      margin-top: 20px;
    }
    .title{
      font-size: 24px;
      color: #333;
      margin-top: 20px;
      font-weight: 600;
    }
    .goback_btn {
      margin-top: 20px;
      padding: 8px 16px;
      border: none;
      background-color: $primary-color-red;
      color: #fff;
      border-radius: 40px;
      cursor: pointer;
    }
  }
}
</style>
