<template>
  <div class="content">
    <div class="left">
      <div class="media-area-new">
        <div class="cover-container">
          <div class="preview-box">
            <div class="cover">
              <img :src="coverImage"
                   alt="">
            </div>
            <div class="content">

              <div class="basic-info">
                <div class="title">优秀的封面会吸引更多人浏览笔记</div>
                <div class="info">
                  <div>视频大小：529.21 KB</div>
                  <div>视频时长：2.54s</div>
                </div>
              </div>
              <div class="operator">
                <span style="color: rgb(255, 36, 66);">修改封面</span>
                <span style="color: rgb(255, 36, 66);">替换视频</span>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="setting">
        正文内容
      </div>
      <div class="title-input">
        <el-input v-model="pubNotForm.noteTitle"
                  style="width: 240px"
                  maxlength="20"
                  placeholder="填写标题，可能会有更多赞哦~"
                  show-word-limit
                  type="text" />
      </div>
      <div class="topic-container">
        <el-input v-model="pubNotForm.noteContent"
                  style="width: 240px"
                  :rows="2"
                  tabindex="0"
                  maxlength="100"
                  show-word-limit
                  type="textarea"
                  placeholder="填写更全面的描述信息，让更多的人看到你吧！" />
      </div>
      <div class="setting">
        发布设置
      </div>
      <div class="formbox">
        <div class="flexbox">
          <div class="_title">自主申明</div>
          <el-select v-model="pubNotForm.statement"
                     placeholder="点击设置笔记声明"
                     size="default"
                     suffix-icon=""
                     style="width: 420px">
            <el-option v-for="item in options"
                       :key="item.value"
                       :label="item.label"
                       :value="item.value" />
          </el-select>
        </div>
        <div class="flexbox">
          <div class="_title">权限设置</div>
          <el-radio-group v-model="pubNotForm.private">
            <el-radio :value="true"
                      size="large">
              <span>公开</span>
              <span class="see">(所有人可见)</span>
            </el-radio>
            <el-radio :value="false"
                      size="large">
              <span>私有</span>
              <span class="see">(仅自己可见)</span>
            </el-radio>
          </el-radio-group>
        </div>
        <div class="flexbox">
          <div class="_title">发布时间</div>
          <el-radio-group v-model="pubNotForm.publishTime"
                          @change="handlePublishTimeChange">
            <el-radio :label="currentTimestamp"
                      size="large">
              立即发布
            </el-radio>
            <el-radio :label="0"
                      size="large">
              定时发布
            </el-radio>
          </el-radio-group>
        </div>
      </div>

      <div class="submit">
        <div class="submit-wrap">
          <el-button size="large"
                     color="#ff2442"
                     @click="handlePublish"
                     class="push">发布</el-button>
          <el-button size="large"
                     class="cancel">取消</el-button>
        </div>
      </div>
    </div>

    <div class="preview">
      <div class="phone-wrapper">
        <p style="height: 544px;">hah</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive,ref,watch,watchEffect, defineProps } from "vue"
import type { publishNoteForm,imageForm } from '@/api/note/type'
import { publishNote } from '@/api/note'
import {uploadFile} from '@/api/file'
const props = defineProps({
  videoUrl: {
    type: String,
    default: ''
  },
	videoFile:Object,
})
console.log('props',props.videoUrl)
const coverImage = ref('')
const mediaVal = ref({
	name:'',
	url:props.videoUrl
})
const pubNotForm:publishNoteForm = reactive({
  noteTitle: '',
  noteContent: '',
  private: true,
  statement: '',
	coverImageUrl:coverImage, 
  publishTime: 0,
  contentType: 2,
  imgList: [],
})
const options = [
  {
    value: 'VIRTUAL-ENTERTAINMENT',
    label: '虚拟演绎，仅供娱乐',
  },
  {
    value: 'AI-MERGER',
    label: '笔记含AI合成内容',
  }
]
watch(() => props.videoUrl,(newFile) => {
		console.log('我干你妈')
		console.log('newFile',newFile)
		if (newFile) {
			// generateVideoCover(newFile)
		}
  },
  { deep: true } // 使用 deep 选项，监听对象内部变化
)
// 监听到文件的变化
watchEffect(() => {
  const file = props.videoFile
	console.log("wo diiu ")
  if (file) {
    console.log('Video file changed:', file)  // 确认文件变化
    generateVideoCover(file).then((blob) => {
			if (blob) {
				const formData = new FormData()
				formData.append('file',blob)
				uploadFile(formData).then(res => {
					console.log('上传成功',res)
					coverImage.value = res.data
				}).catch(err => {
					console.log('上传失败',err)
				})
			}
		}).catch((error) => {
			console.error('Failed to generate video cover:', error);
		});
  }
})

const handlePublish = () => {
  // pubNotForm.imgList = imgList.value.map(item => ({ url: item.url, name: item.name }));
  publishNote(pubNotForm).then(res => {
    console.log(res)

    // TODO 发布成功后跳转到笔记管理页面
    // $router.push('/notes-manager')
  }).catch(err => {
    console.log(err)
  })
}
const currentTimestamp = Math.floor(Date.now() / 1000);
function handlePublishTimeChange (value:number) {
  if (value === currentTimestamp) {
    pubNotForm.publishTime = value;
  } else {
    pubNotForm.publishTime = value;
  }
}
// 随机选择视频中的一帧作为封面
// function generateVideoCover(file: File): Promise<string | null> {
//   return new Promise((resolve, reject) => {
//     // 创建 video 和 canvas 元素
//     const videoElement = document.createElement('video');
//     const canvas = document.createElement('canvas');
//     const context = canvas.getContext('2d');

//     if (!context) {
//       reject('Canvas context not supported');
//       return;
//     }

//     // 为 video 元素设置文件 URL
//     videoElement.src = URL.createObjectURL(file);
//     videoElement.load();

//     // 监听视频元数据加载完毕的事件
//     videoElement.onloadedmetadata = () => {
//       const randomTime = Math.random() * videoElement.duration; // 随机时间点
//       videoElement.currentTime = randomTime;

//       // 监听视频 seek 事件
//       videoElement.onseeked = () => {
//         // 设置 canvas 尺寸为视频的宽高
//         canvas.width = videoElement.videoWidth;
//         canvas.height = videoElement.videoHeight;

//         // 将当前帧绘制到 canvas 上
//         context.drawImage(videoElement, 0, 0, canvas.width, canvas.height);

//         // 将 canvas 内容转换为 base64 格式的图片
				
//         coverImage.value = canvas.toDataURL('image/jpeg');
//         // resolve(coverImage.value);
// 				// console.log("mabi ",coverImage)
// 				// const videoCoverElement = document.getElementById('coverImg');
// 				// if (!videoCoverElement) {
// 				// 	console.log('没有找到元素')
// 				// 	return;
// 				// }
// 				// videoCoverElement.style.backgroundImage = `url(${coverImage.value})`; // 将封面图设置为背景图
//         // // 释放 URL 对象，防止内存泄露
//         URL.revokeObjectURL(videoElement.src);
//       };
//     };

//     // 处理视频加载错误
//     videoElement.onerror = (error) => {
//       reject('Failed to load video file');
//     };
//   });
// }

function generateVideoCover(file: File): Promise<Blob | null> {
  return new Promise((resolve, reject) => {
    const videoElement = document.createElement('video');
    const canvas = document.createElement('canvas');
    const context = canvas.getContext('2d');

    if (!context) {
      reject('Canvas context not supported');
      return;
    }

    videoElement.src = URL.createObjectURL(file);
    videoElement.load();

    videoElement.onloadedmetadata = () => {
      const randomTime = Math.random() * videoElement.duration; // 随机选择视频时间点
      videoElement.currentTime = randomTime;

      videoElement.onseeked = () => {
        canvas.width = videoElement.videoWidth;
        canvas.height = videoElement.videoHeight;

        context.drawImage(videoElement, 0, 0, canvas.width, canvas.height);

        // 将 canvas 内容转换为 blob 对象，便于上传
        canvas.toBlob((blob) => {
          if (blob) {
            resolve(blob); // 返回 Blob 对象
          } else {
            reject('Failed to generate video cover');
          }
        }, 'image/jpeg');

        URL.revokeObjectURL(videoElement.src); // 释放资源
      };
    };

    videoElement.onerror = () => {
      reject('Failed to load video file');
    };
  });
}

// 长传封面图到后台
</script>

<style lang="scss" scoped>

.content{
	position: relative;
	width: 100%;
	height: 100%;
	padding: 0 300px 0 13px;
	.left{
		width: 100%;
		.media-area-new{
			width: 100%;
			.cover-container{
				width: 100%;
				.preview-box{
					display: flex;
					justify-content: flex-start;
					.cover{
						// background: #fafafa;
						width: 75px;
						height: 100px;
						border-radius: 6px;
						overflow: hidden;
						.coverImage{
							height: 100%;
						}
					}
					.content{
						display: flex;
						justify-content: flex-start;
						margin-left: 10px;
						background: #fafafa;
						border-radius: 6px;
						flex: 1;
						padding: 24px 30px 20px;
						display: flex;
						.basic-info{
							flex: 1;
							.title{
								color: #262626;
								font-family: PingFang SC;
								font-size: 14px;
								font-style: normal;
								font-weight: 400;
								line-height: 24px;
							}
							.info{
								display: flex;
								justify-content: flex-start;
								margin-top: 6px;
    						width: 100%;
								div{
									color: rgba(0, 0, 0, .45);
									font-family: PingFang SC;
									font-size: 12px;
									font-style: normal;
									font-weight: 400;
									line-height: 22px;
								}
							}
						}
						.operator{
							display: flex;
							flex-direction: column;
							width: 160px;
							align-items: center;
							span{
								text-align: center;
								font-family: PingFang SC;
								font-size: 14px;
								font-style: normal;
								font-weight: 400;
								line-height: 28px;
								height: 28px;
								padding: 0 12px;
								white-space: nowrap;
							}
						}
					}
				}
			}
		}
		.title-input{
			width: 100%;
			margin-bottom: 12px;
		}
		.topic-container{
			width: 100%;
			margin-bottom: 12px;
		}
		.setting{
				margin-top: 24px;
				margin-bottom: 12px;
				color: #262626;
				font-family: PingFang SC;
				font-size: 18px;
				font-style: normal;
				font-weight: 500;
				line-height: 26px;
		}
		.formbox{
			.flexbox{
				display: flex;
				height: 48px;
				align-items: center;
				._title{
					align-items: center;
					margin: 0;
					font-size: 14px;
					color: #262626;
					padding: 8px 24px 8px 0;
				}
			}
		}
		.submit{
			margin-top: 20px;
			display: flex;
			align-items: center;
			.submit-wrap{
				.push{
					width: 110px;
					margin-right: 20px;
				}
				.cancel{
					width: 110px;
				}
			}
		}
	}
	
	.preview{
		width: 252px;
    padding: 10px 36px 69px 4px;
    box-sizing: content-box;
		position: absolute;
		right: 0;
		top: 0;
		.phone-wrapper{
			width: 100%;
			height: 100%;
			position: relative;
			background-image: url('@/assets/imgs/phone_bg.png');
			background-size: 100% 100%;
			overflow: hidden;
		}
	}
}</style>
