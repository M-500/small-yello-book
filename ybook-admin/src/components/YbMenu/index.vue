<template>
  <div class="">
    <el-menu default-active="2"
             class="el-menu-vertical-demo"
             @open="handleOpen"
             router="true"
             @close="handleClose">

      <el-menu-item index="/home">
        <el-icon>
          <HomeFilled />
        </el-icon>
        <span>首页</span>
      </el-menu-item>
      <el-menu-item index="/notes-manager">
        <el-icon>
          <Tickets />
        </el-icon>
        <span>笔记管理</span>
      </el-menu-item>
      <el-menu-item index="/notes-check">
        <el-icon>
          <Stamp />
        </el-icon>
        <span>笔记审核</span>
      </el-menu-item>
      <el-sub-menu index="1">
        <template #title>
          <el-icon>
            <Histogram />
          </el-icon>
          <span>数据看板</span>
        </template>
        <el-menu-item index="/notes-dashboard">笔记数据</el-menu-item>
        <el-menu-item index="/fans-dashboard">粉丝数据</el-menu-item>
        <!-- <el-menu-item-group>
          <el-menu-item index="1-1">笔记数据</el-menu-item>
          <el-menu-item index="1-2">粉丝数据</el-menu-item>
        </el-menu-item-group> -->
      </el-sub-menu>

    </el-menu>
  </div>
</template>

<script setup lang="ts">

const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
</script>

<style lang="scss" scoped>

::v-deep(.el-menu-item.is-active){
  color: $primary-buttom-active-colder;
  background-color: rgb(242, 247, 255);
  border-radius: 8px;
}
::v-deep(.el-sub-menu.is-active){
  color: $primary-buttom-active-colder;
  background-color: rgb(242, 247, 255);
  border-radius: 8px;
}

::v-deep(.el-menu-item){
  margin: 0 10px 0;
  height: 44px;
}

::v-deep(.el-menu-item):hover{
  color: $primary-buttom-active-colder;
  // 背景色透明
  background-color:transparent;
}
::v-deep(.el-sub-menu){
  margin: 0 10px 0;
  height: 44px;
}
::v-deep(.el-sub-menu__title):hover{
  color: $primary-buttom-active-colder;
  // 背景色透明
  background-color:transparent;
}

::v-deep(.el-sub-menu__title){
  height: 44px;
  line-height: 44px;
}

::v-deep(.el-menu){
  // 去掉右边框
  border-right: none;
}
</style>
