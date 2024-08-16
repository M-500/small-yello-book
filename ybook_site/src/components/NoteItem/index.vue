<template>
  <div class="note-item">
    <div class="container">
      <a href=""></a>
      <a href=""
         class="cover">
        <img src="../../assets/imgs/avatar.jpeg"
             alt="">
      </a>
      <NoteFooter />
    </div>
  </div>
</template>

<script setup>
import NoteFooter from "@/components/NoteFooter/index.vue"
</script>

<style lang="scss" scoped>
.note-item{
	// position: absolute;
	// left: 0;
	// top: 0;
	width: 221.33px; // 这个需要后期需要根据屏幕来适配
	.container{
		width: 100%;
		height: 100%;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		.cover{
			position: relative;
			width: 221.33px; // 这个需要后期需要根据屏幕来适配
			display: flex;
			border-radius: 16px;
			overflow: hidden;
			outline: 1px solid rgba(0,0,0,0.08);
			outline-offset: -1px;
			transition: background .2s;
			transform: translateZ(0);
		}
	}
}
</style>
