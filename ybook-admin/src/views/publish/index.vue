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
                      @updateData="handleUpdateData"
                      btnTitle="上传视频"></YbUpload>
          </el-tab-pane>
          <el-tab-pane label="上传图文"
                       name="second">
            <YbUpload mark="拖拽图片到此或点击上传"
                      @updateData="handleUpdateData"
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
          <span>
            发布图文
          </span>
        </div>

        <div class="content">
          <div class="media-area-new">
            <div class="img-list">
              <div class="top">
                <div class="title-area">
                  <div class="title">图片编辑</div>
                  <div class="status">(1/18)</div>
                </div>
                <span class="reset-upload"
                      @click="handleResetUpload">清空并重新上传</span>
              </div>
              <div class="img-upload-area">
                <ImgUpload :coverImge="coverImgeUrl"
                           @itemImgListChanged="handleImgListChanged"></ImgUpload>
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
          </div>
        </div>
      </el-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref, watch } from 'vue'
import type { publishNoteForm,image } from '@/api/note/type'
import { publishNote } from '@/api/note'
import type {  TabsPaneContext } from 'element-plus'
import YbUpload from '@/components/YbUpload/index.vue'
import ImgUpload from '@/components/ImgUpload/index.vue'
import type { UploadUserFile } from 'element-plus'
import { useRouter } from 'vue-router'

let $router = useRouter()
const activeName = ref('first')

const coverImgeUrl = ref("")  // 来自yb-upload组件的属性
const imgList = ref<UploadUserFile[]>([])  // 来自img-upload组件的属性 
// const noteTitle = ref('')
const value = ref('')
const step1 = ref(true)

const currentTimestamp = Math.floor(Date.now() / 1000);

function handlePublishTimeChange(value:any) {
  if (value === currentTimestamp) {
    pubNotForm.publishTime = value;
  } else {
    pubNotForm.publishTime = value;
  }
}

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

const pubNotForm:publishNoteForm = reactive({
  noteTitle: '库里起飞',
  noteContent: '就是库里咯',
  private: true,
  statement: '',
  publishTime: 0,
  contentType: 1,
  imgList: imgList.value as image[]
})

const handlePublish = () => {
  // console.log("hahahah ",imgList)
  pubNotForm.imgList = imgList.value.map(item => ({ url: item.url, name: item.name }));
  publishNote(pubNotForm).then(res => {
    console.log(res)

    // TODO 发布成功后跳转到笔记管理页面
    $router.push('/notes-manager')
  }).catch(err => {
    console.log(err)
  })
}



const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event)
}
function handleUpdateData(data: string) {
  step1.value = false
  coverImgeUrl.value = data
  // 追加写入到imgList
  imgList.value.push({
    url: data,
    name: 'coverImge',
  })
  console.log('imgList', imgList.value)
}

function handleResetUpload() {
  console.log('清空并重新上传')
  imgList.value = []
}

function handleImgListChanged(data: UploadUserFile[]) {
  console.log("子组件传回来的",data)
  imgList.value = data
  // console.log('我干你娘', pubNotForm)
}

function goback (){
  imgList.value = []
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
