<template>
  <el-container>
    <el-header>
      <Header />
    </el-header>
    <el-container>
      <el-aside>
        <Aside />
      </el-aside>
      <el-main>
        <Note />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import Header from './layout/Header.vue'
import Aside from './layout/Aside.vue'
import Note from './Note.vue'
</script>

<style scoped></style>
