<template>
  <div class="container">
    <div class="before-publish"
         v-show="step1">
      <el-card>
        <el-tabs v-model="activeName"
                 class="demo-tabs"
                 @tab-click="handleClick">
          <el-tab-pane label="上传视频"
                       name="first">
            <YbUpload mark="拖拽或者点击上传视频"
                      @updateData="handlerUpdateVideo"
                      @updateFile="handlerUpdateFile"
                      :tipList="videoTips"
                      accept="video/*"
                      btnTitle="上传视频"></YbUpload>
          </el-tab-pane>
          <el-tab-pane label="上传图文"
                       name="second">
            <YbUpload mark="拖拽图片到此或点击上传"
                      @updateData="handleUpdateData"
                      :tipList="imageTips"
                      accept="image/*"
                      btnTitle="上传图片"></YbUpload>
          </el-tab-pane>
        </el-tabs>
      </el-card>

    </div>

    <div class="main-publish"
         v-show="!step1">
      <el-card>
        <div class="header">
          <span class="icon"
                @click="goback">
            <el-icon>
              <ArrowLeft />
            </el-icon>
          </span>
          <span v-if="noteType === 'image'">
            发布图文
          </span>
          <span v-else>
            发布视频
          </span>
        </div>
        <ImagePub v-if="noteType === 'image'"
                  :coverImgeUrl="coverImgeUrl" />
        <VideoPub v-else
                  :videoFile="videoFile"
                  :videoUrl="videoUrl" />
      </el-card>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref, watch } from 'vue'
import { TabsPaneContext } from 'element-plus'
import YbUpload from '@/components/YbUpload/index.vue'
import { UploadUserFile } from 'element-plus'
import { useRouter } from 'vue-router'
import ImagePub from './imagePub.vue'
import VideoPub from './videoPub.vue'
import { da } from 'element-plus/es/locales.mjs'

let $router = useRouter()
const noteType = ref('image')
const activeName = ref('first')
const step1 = ref(true)
const videoTips = ref([
  {
    title: '视频大小',
    content: ['支持时长60分钟以内，', '最大20GB的视频文件']
  },
  {
    title: '视频格式',
    content: ['支持常用视频格式', '推荐使用mp4、mov']
  },
  {
    title: '视频分辨率',
    content: ['推荐上传720P(1280*720)及以上视频，', '超过1080P的视频用网页端上传画质更清晰']
  }
]);
const imageTips = ref([
  {
    title: '图片大小',
    content: ['支持JPG、PNG、GIF格式', '最大20MB的图片文件']
  },
  {
    title: '图片格式',
    content: ['支持上传的图片格式，', '推荐使用png、jpg、jpeg、webp']
  },
  {
    title: '图片分辨率',
    content: ['推荐上传宽高比为3:4、分辨率不低于720*960的照片，', '超过1280P的图片用网页端上传画质更清晰']
  }
]);

const coverImgeUrl = ref('')
const videoUrl = ref('')
const videoFile = ref(null)

const handleClick = (tab, event) => {
  console.log(tab, event)
}
function handleUpdateData (data) {
  step1.value = false
  noteType.value = 'image'
  coverImgeUrl.value = data
}

function handlerUpdateVideo (data) {
  step1.value = false
  noteType.value = 'video'
  videoUrl.value = data
  console.log('videoUrl', videoUrl.value)
}

const handlerUpdateFile = (filters) => {
  // 将视频文件传递给子组件
  videoFile.value = file
  console.log("妈的", videoFile.value)
}
function goback () {
  coverImgeUrl.value = ""
  step1.value = true
}
</script>

<style lang="scss" scoped>

.container {
  .before-publish {
    height: 100%;
    .el-card{
      height: 100%;
        .el-tabs{
          
        }
      }
  }
  .main-publish{
    height: 100%;
    .el-card{
      height: 100%;
      .header{
        line-height: 28px;
        font-size: 20px;
        font-weight: 900;
        margin-bottom: 24px;
        display: flex;
        align-items: center; /* 垂直居中 */
        // justify-content: center; /* 水平居中 */
        .icon{
          display: flex;
          align-content: center;
          margin-right: 4px; /* 图标和文字之间的间距 */
          width: 20px;
          height: 20px;
        }
      }
      .content{
        position: relative;
        width: 100%;
        padding: 0 300px 0 13px;
        .media-area-new{
          width: 100%;
          .img-list{
            .top{
              display: flex;
              justify-content: space-between;
              align-items: center;
              margin-bottom: 12px;
              .title-area{
                display: flex;
                align-items: center;
                .title{
                  font-weight: 900;
                  font-size: 18px;
                  line-height: 26px;
                  color: #262626;
                  margin-right: 6px;
                }
                .status{
                  color: rgba(51, 51, 51, .6);
                  font-family: PingFang SC;
                  font-size: 14px;
                  font-style: normal;
                  font-weight: 400;
                  line-height: 18px;
                }
              }
              .reset-upload{
                color: #ff2442;
                font-size: 14px;
                font-family: PingFang SC;
                font-weight: 400;
                cursor: pointer; // 鼠标移上去变成小手
              }
            }
            .img-upload-area{
              margin-bottom: 12px;
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
        }
      }
    }
  }
}

.el-card{
	height: 100%;
	.header{
		line-height: 28px;
		font-size: 20px;
		font-weight: 900;
		margin-bottom: 24px;
		display: flex;
		align-items: center; /* 垂直居中 */
		// justify-content: center; /* 水平居中 */
		.icon{
			display: flex;
			align-content: center;
			margin-right: 4px; /* 图标和文字之间的间距 */
			width: 20px;
			height: 20px;
		}
	}
	
}
.el-tabs__header{
  margin-bottom: 28px;
  font-size: 20px;
  font-weight: 900;
  .el-tabs__nav{
  padding: 12px 24px;
  position: relative;
  cursor: pointer;
  
  }
}
::v-deep(.el-tabs__active-bar) {
  background-color: #ff2442 !important;
}
// 深度覆盖element-plus的样式
::v-deep(.el-tabs__item.is-top.is-active) {
  color: #ff2442 !important;
}

::v-deep(.el-tabs__item:hover){
  color: #ff2442 !important;
}

::v-deep(.el-input){
  width: 100% !important;
}
::v-deep(.el-textarea){
  width: 100% !important;
}
::v-deep(.el-input__wrapper):hover{
  border: 1px solid $primary-buttom-active-colder !important;
}

// ::v-deep(.el-select__wrapper){
//   width: 420px;
// }

::v-deep(.el-radio .is-checked){
  .el-radio__input{
    color: $primary-buttom-active-colder !important;
  }
  
}
</style>
