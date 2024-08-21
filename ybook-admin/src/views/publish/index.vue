<template>
  <div class="container">
    <div class="before-publish"
         v-show="!step1">
      <el-card>
        <el-tabs v-model="activeName"
                 class="demo-tabs"
                 @tab-click="handleClick">
          <el-tab-pane label="上传视频"
                       name="first">
            <yb-upload mark="拖拽或者点击上传视频"
                       @updateData="handleUpdateData"
                       btnTitle="上传视频"></yb-upload>
          </el-tab-pane>
          <el-tab-pane label="上传图文"
                       name="second">
            <yb-upload mark="拖拽图片到此或点击上传"
                       @updateData="handleUpdateData"
                       btnTitle="上传图片"></yb-upload>
          </el-tab-pane>
        </el-tabs>
      </el-card>

    </div>

    <div class="main-publish"
         v-show="step1">
      <el-card>
        <div class="header">
          <span class="icon">
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
                <img-upload :coverImge="coverImgeUrl"></img-upload>
              </div>
              <div class="title-input">
                <el-input v-model="noteTitle"
                          style="width: 240px"
                          maxlength="20"
                          placeholder="填写标题，可能会有更多赞哦~"
                          show-word-limit
                          type="text" />
              </div>
              <div class="topic-container">
                <el-input v-model="noteContent"
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
                  <el-select v-model="value"
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
                  <el-radio-group v-model="isPublic">
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
                  <el-radio-group v-model="noCorn">
                    <el-radio :value="true"
                              size="large">
                      立即发布
                    </el-radio>
                    <el-radio :value="false"
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
import { onMounted, ref, watch } from 'vue'
import { TabsPaneContext } from 'element-plus'
import YbUpload from '@/components/YbUpload/index.vue'
import ImgUpload from '@/components/ImgUpload/index.vue'
const activeName = ref('first')

const coverImgeUrl = ref("")  // 来自yb-upload组件的属性

const noteTitle = ref('')
const noteContent = ref('')
const value = ref('')
const isPublic = ref(true)
const noCorn = ref(true)
const step1 = ref(true)
const options = [
  {
    value: 'Option1',
    label: '虚拟演绎，仅供娱乐',
  },
  {
    value: 'Option2',
    label: '笔记含AI合成内容',
  }
]
const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event)
}
function handleUpdateData(data: string) {
  step1.value = false
  coverImgeUrl.value = data
}

function handleResetUpload() {
  console.log('清空并重新上传')
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
