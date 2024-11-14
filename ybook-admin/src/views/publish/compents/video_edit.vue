<template>
  <div class="video-screenshot">
    <!-- 视频播放器 -->
    <video ref="videoRef"
           :src="videoUrl"
           controls
           @loadedmetadata="onLoadedMetadata"
           @play="onPlay"
           @pause="onPause"
           @timeupdate="onTimeUpdate"></video>

    <!-- 截图按钮 -->
    <button @click="captureFrame">截取封面图</button>

    <!-- 显示截图的图片 -->
    <div v-if="screenshot"
         class="screenshot-preview">
      <h3>封面图预览:</h3>
      <img :src="screenshot"
           alt="封面图" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';

const props = defineProps({
  videoUrl: String,
});

const videoRef = ref(null);
const canvasRef = ref(null);
const screenshot = ref('');
const videoDuration = ref(0);

// 视频加载完成时获取视频总时长
const onLoadedMetadata = () => {
  videoDuration.value = videoRef.value.duration;
};

// 播放时的回调（可选）
const onPlay = () => {
  console.log('视频播放中...');
};

// 暂停时的回调（可选）
const onPause = () => {
  console.log('视频已暂停');
};

// 捕获当前帧作为封面图
const captureFrame = () => {
  const video = videoRef.value;
  if (video) {
    // 创建 canvas 元素
    const canvas = document.createElement('canvas');
    canvas.width = video.videoWidth;
    canvas.height = video.videoHeight;
    const ctx = canvas.getContext('2d');

    // 绘制视频的当前帧到 canvas 上
    ctx.drawImage(video, 0, 0, canvas.width, canvas.height);

    // 将 canvas 转换为图片 URL
    screenshot.value = canvas.toDataURL('image/png');
  }
};
</script>

<style  lang="scss" scoped>
.video-screenshot {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
}

video {
  max-width: 640px;
  width: 100%;
}

button {
  padding: 10px 20px;
  background-color: #3498db;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.screenshot-preview {
  margin-top: 20px;
}

.screenshot-preview img {
  max-width: 100%;
  border: 1px solid #ddd;
}
</style>