<template>
  <div class="container">
    <div class="content">
      <el-card>
        <div class="fans-data-container">

          <yb-title title="粉丝数据概览"></yb-title>

        </div>
      </el-card>
      <div class="fan-bottom">
        <el-card>
          <div class="fan-portrait-container">

            <yb-title title="粉丝画像"></yb-title>

          </div>
        </el-card>
      </div>

    </div>
    <div class="fans-content">
      <el-card>
        <div class="fans-active">
          <yb-title title="我的活跃粉丝"></yb-title>
        </div>
      </el-card>

    </div>
  </div>
</template>

<script setup>
import YbTitle from '@/components/YbTitle/index.vue'
</script>

<style lang="scss" scoped>
.container{
  height: 100%;
  width: 100%;
  display: flex;
  .content{
    width: calc(100% - 372px);
    // padding: 0px 24px 32px;
    .fans-data-container{
      padding: 0px 24px 32px;
    }
    .fan-bottom{
      margin-top: 16px;
      .fan-portrait-container{
        padding: 0px 24px 32px;
      }
    }
    
  }
  .fans-content{
    margin-left: 16px;
    width: 372px;
    .fans-active{
      padding: 0 24px;
    }
  }
}

::v-deep(.el-card){
  border-radius: 8px;
  height: 100%;
}
::v-deep(.el-card__body){
  padding: 0;
  height: 100%;
}
</style>
