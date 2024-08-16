<template>
  <div class="channel_container">
    <div class="item active_box">
      <div class="icon">
        <img src="../../assets/icons/home.svg"
             alt="" />
      </div>
      <div class="title">发现</div>
    </div>
    <div class="item">
      <div class="icon">
        <img src="../../assets/icons/publish.svg"
             alt="" />
      </div>
      <div class="title">发布</div>
    </div>
    <div class="item">
      <div class="icon">
        <img src="../../assets/icons/msg.svg"
             alt="" />
      </div>
      <div class="title">通知</div>
    </div>
    <login_btn />
  </div>
</template>

<script setup>

import login_btn from "../loginBtn/index.vue";
</script>

<style lang="scss" scoped>
.channel_container{
	.active_box{
		background-color: $active-color;
	}
	.item{
		display: flex;
		padding-left: 16px;
		margin-bottom: 8px;
		align-items: center;
		height: 45px;
		cursor: pointer;
		border-radius: 50px;
		.icon img {
			width: 24px;
			height: 24px;
		}
		.title{
			font-size: 16px;
			font-weight: 600;
			margin-left: 12px;
			color: $gray-color;
		}
	}
}
</style>
