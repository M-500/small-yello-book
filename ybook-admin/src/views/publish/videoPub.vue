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
                  <div class="info-bar">视频大小：529.21 KB</div>
                  <div class="info-bar">视频时长：2.54s</div>
                </div>
              </div>
              <div class="operator">
                <span style="color: rgb(255, 36, 66);"
                      @click="dialogVisible = true">修改封面</span>
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
      <!-- 视频预览 -->
      <video_phone :videoUrl="mediaVal.url"
                   :coverImgUrl="coverImage" />
    </div>

    <el-dialog v-model="dialogVisible"
               title="设置封面"
               width="500">
      <div class="cover-choice-container">
        <video_edit :videoUrl="mediaVal.url"
                    :coverImgUrl="coverImage"
                    @update:coverImgUrl="coverImage = $event" />
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">Cancel</el-button>
          <el-button type="primary"
                     @click="dialogVisible = false">
            Confirm
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { reactive,ref,watch,watchEffect, defineProps } from "vue"
import type { publishNoteForm,imageForm } from '@/api/note/type'
import video_phone from "./compents/video_phone.vue"
import video_edit from "./compents/video_edit.vue"
import { publishNote } from '@/api/note'
import {uploadFile} from '@/api/file'
const props = defineProps({
  videoUrl: {
    type: String,
    default: ''
  },
	videoFile:Object,
})
const dialogVisible = ref(false)
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
// 监听到文件的变化  生成封面图上传到后台
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

// 生成视频封面图
function generateVideoCover(file) {
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
						img{
							height: 100%;
              width: 100%;
              // 填充模式
              object-fit: cover;
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
								.info-bar{
									color: rgba(0, 0, 0, .45);
									font-family: PingFang SC;
									font-size: 12px;
									font-style: normal;
									font-weight: 400;
									line-height: 22px;
                  // 如果不是第一个元素，就添加左边距
                  &:not(:first-child){
                    margin-left: 24px;
                  }
								}
                
							}
						}
						.operator{
							display: flex;
							// flex-direction: column;
							width: 160px;
							align-items: center;
							span{
								text-align: center;
                cursor: pointer;
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
			position: relative;
			// background-image: url('@/assets/imgs/phone_bg.png');
			background-size: 100% 100%;
      height: 544px;
			overflow: hidden;
      z-index: 1000;
      // 在最上层
      z-index: 1;
      .cover-base{
        z-index: 100;
        width: 100%;
        // height: 100%;
        padding: 1.5px;
        border-radius: 34px;
        background: #000;
        z-index: 2;
        // 在父元素的下层
        img{
          z-index: 111;
          width: 100%;
          height: 100%;
          object-fit: cover;
        }
      }
		}
	}
}</style>
