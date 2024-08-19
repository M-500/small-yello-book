<template>
  <div class="container">
    <div class="before-publish"
         v-show="false">
      <el-card>
        <el-tabs v-model="activeName"
                 class="demo-tabs"
                 @tab-click="handleClick">
          <el-tab-pane label="上传视频"
                       name="first">
            <yb-upload mark="拖拽或者点击上传视频"
                       btnTitle="上传视频"></yb-upload>
          </el-tab-pane>
          <el-tab-pane label="上传图文"
                       name="second">
            <yb-upload mark="拖拽图片到此或点击上传"
                       btnTitle="上传图片"></yb-upload>
          </el-tab-pane>
        </el-tabs>
      </el-card>

    </div>

    <div class="main-publish">
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
              </div>
              <div class="img-upload-area">
                <img-upload></img-upload>
              </div>
            </div>
          </div>
        </div>
      </el-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { TabsPaneContext } from 'element-plus'
import YbUpload from '@/components/YbUpload/index.vue'
import ImgUpload from '@/components/ImgUpload/index.vue'
const activeName = ref('first')

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event)
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
              
            }
            .img-upload-area{
              
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

</style>
