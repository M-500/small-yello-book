<template>
  <div class="upload-content">
    <div class="upload-wrap">
      <el-upload class="upload-box"
                 drag
                 action="http://127.0.0.1:8122/api/v1/file/upload"
                 :on-success="handlePreview"
                 multiple>
        <el-icon class="el-icon--upload">
          <upload-filled />
        </el-icon>
        <div class="el-upload__text">
          <p class="desc">{{ mark }}</p>
          <el-button size="small"
                     type="primary">{{ btnTitle }}</el-button>
        </div>
      </el-upload>
    </div>
    <div class="tips-wrap">
      <div class="tip">
        <div class="box">
          <p class="tip-title">视频大小</p>
          <p class="tip-content">支持时长60分钟以内，</p>
          <p class="tip-content">最大20GB的视频文件</p>
        </div>
      </div>
      <div class="tip">
        <div class="box">
          <p class="tip-title">视频格式</p>
          <p class="tip-content">支持常用视频格式</p>
          <p class="tip-content">推荐使用mp4、mov</p>
        </div>
      </div>
      <div class="tip">
        <div class="box">
          <p class="tip-title">视频分辨率</p>
          <p class="tip-content">推荐上传720P(1280*720)及以上视频，</p>
          <p class="tip-content">超过1080P的视频用网页端上传画质更清晰</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { UploadProps } from 'element-plus';
import { ref }  from 'vue';
import { defineEmits } from 'vue';
import { defineProps } from 'vue';
const emit = defineEmits(['updateData']);

// 这里是子组件定义的属性
const coverImgUrl = ref('')

// defineExpose({coverImgUrl})
defineProps({
	// 接受父组件传递的参数
	mark: {
		type: String,
		default: ''
	},
	btnTitle: {
		type: String,
		required: true
	}
})

const handlePreview: UploadProps['onSuccess'] = (uploadFile) => {
	coverImgUrl.value = uploadFile.data.url
	emit('updateData', coverImgUrl.value)
}
</script>

<style lang="scss" scoped>
.updoad-content {
	display: flex;
	.upload-wrap {
		width: 100%;
		height: 300px;
		background-color: #f5f5f5;
		
	}
	
}
.tips-wrap {
		width: 100%;
		display: flex;
		flex-direction: row;
		justify-content: flex-start;
    padding-top: 50px;
    padding-bottom: 50px;
		// box-sizing: border-box;
		.tip{
			flex: 1;
			box-sizing: border-box;
			display: flex;
			flex-direction: column;
			justify-content: center;
			align-items: center;
			.box{
				.tip-title{
					text-align: left;
					font-size: 14px;
					line-height: 22px;
					color: #262626;
					font-family: PingFang SC;
					margin-bottom: 4px;
				}
				.tip-content{
					margin-top: 0;
					margin-bottom: 0;
					font-size: 12px;
					text-align: left;
					line-height: 20px;
					font-family: PingFang SC;
					color: #8c8c8c;
				}
				p{
					display: block;
					margin-block-start: 1em;
					margin-block-end: 1em;
					margin-inline-start: 0px;
					margin-inline-end: 0px;
					unicode-bidi: isolate;
				}
			}
		}
		// 如果不是最后一个元素，添加右边框
		.tip:not(:last-child){
			border-right: 1px solid #f4f4f4;
		}
	}

.upload-box{
		// flex: 1;
		// width: 100%;
		// position: relative;
		// border-radius: 4px;
		// display: flex;
		// flex-direction: column;
		// justify-content: center;
		// align-items: center;
		// border-width: 10px;
		// border-style: dashed;
		background: #f7f7f7;
	}

::v-deep(.el-button){
	background-color: #ff2442;
	border-color: #ff2442;
	width: 104px;
	height: 38px;
	span{
		font-size: 14px;
		font-weight: 500;
	}
}
::v-deep(.el-upload-dragger):hover{
	border: 1px dashed $primary-buttom-active-colder;
}

::v-deep(.el-tabs__header){
	maring-bottom: 28px;
}
.desc{
	font-size: 14px;
	color: #262626;
	font-family: PingFang SC;
	line-height: 22px;
	margin: 24px 0;
}
</style>
