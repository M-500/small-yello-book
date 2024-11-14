<template>
  <div class="video-container">
    <img src="@/assets/imgs/video-bg.png"
         class="phone-bg"
         alt="">
    <div class="video-area">
      <video :src="videoUrl"
             class="video-box"
             controls
             preload="auto"
             oncontextmenu="return false"
             controlslist="nofullscreen noplaybackrate nodownload"
             style="width: 100%; height: 100%;">
      </video>
      <img src="@/assets/imgs/play.png"
           class="play"
           v-show="!isPlay"
           @click="playVideo"
           alt="">
    </div>
  </div>
</template>

<script setup>
import { defineProps, ref } from "vue"
const isPlay = ref(false)
const props = defineProps({
  videoUrl: String,
  coverImgUrl: String
})

const playVideo = () => {
  isPlay.value = true
  const video = document.querySelector('video')
  video.play()
  video.addEventListener('ended', () => {
    isPlay.value = false
  })
}
</script>

<style lang="scss" scoped>
video::-webkit-media-controls-fullscreen-button {
	display: none;
}
//播放按钮
video::-webkit-media-controls-play-button {
	display: none;
}
//进度条
video::-webkit-media-controls-timeline {
	display: none;
}
//所有控件
video::-webkit-media-controls-enclosure{ 
        display: none;
    }

.video-container {
	position: relative;
	width: 100%;
	height: 544px;
	overflow: hidden;
	z-index: 1000;
	width: 100%;
	padding: 1.5px;
	border-radius: 34px;
	background: #000;
	.phone-bg {
		z-index: 111;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    width: 100%;
    height: 100%;
    position: absolute;
    overflow: hidden;
    left: 0;
    top: 0;
	}
	.video-area {
		// position: absolute;
		margin-top: 0;
    height: 480px;
    overflow: hidden;
		width: 100%;
		border-radius: 32px;
    border-bottom-left-radius: 0;
    border-bottom-right-radius: 0;
		// display: flex;
		// justify-content: center;
		// align-items: center;
		.video-box{
			width: 100%;
			height: 100%;
			position: relative;
			border-bottom-left-radius: 0;
			border-bottom-right-radius: 0;
			object-fit: contain;
			z-index: 110;
			object-position: center center;
			cursor: pointer;
			object-fit: cover;
		}
		.play {
			width: 40px;
			height: 44px;
			position: absolute;
			top: 50%;
			left: 50%;
			transform: translateX(-50%) translateY(-39px);
			z-index: 130;
			opacity: .7;
			cursor: pointer;
		}
	}
}
</style>