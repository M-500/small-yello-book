<template>
  <div class="container">
    <div class="before_publish">
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
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { TabsPaneContext } from 'element-plus'
import YbUpload from '@/components/YbUpload/index.vue'
const activeName = ref('first')

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event)
}
</script>

<style lang="scss" scoped>

.container {
  .before_publish {
    height: 100%;
    .el-card{
      height: 100%;
        .el-tabs{
          
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
