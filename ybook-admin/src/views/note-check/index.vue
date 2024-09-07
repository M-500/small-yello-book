<template>
  <el-card class="container">
    <div class="note-info">
      <yb-title title="笔记审核"></yb-title>
    </div>
    <div class="check-box">
      <el-table :data="tableData"
                stripe
                border
                height="100%"
                style="width: 100%">
        <el-table-column prop="date"
                         label="封面"
                         width="180">
          <template #default="{row}">
            <img :src="row.cover"
                 class="cover"
                 alt="" />
          </template>
        </el-table-column>
        <el-table-column prop="name"
                         label="作者"
                         width="180" />
        <el-table-column prop="noteTitle"
                         label="标题"
                         width="180" />
        <el-table-column prop="noteContent"
                         label="内容"
                         width="180" />
        <el-table-column prop="name"
                         label="状态"
                         width="180">
          <template #default="{row}">
            <el-tag v-if="row.status === 0"
                    type="primary">待审核</el-tag>
            <el-tag v-else-if="row.status === 1"
                    type="success">审核通过</el-tag>
            <el-tag v-else-if="row.status === 2"
                    type="danger">未通过</el-tag>
            <el-tag v-else-if="row.status === 3"
                    type="danger">已删除</el-tag>
            <el-tag v-else
                    type="info">其他</el-tag>
          </template>

        </el-table-column>
        <el-table-column prop="address"
                         label="发布日期" />
        <el-table-column fixed="right"
                         label="操作"
                         min-width="30">
          <template #default="{row}">
            <el-button type="primary"
                       v-if="row.status === 0"
                       size="small"
                       @click="passHandler(row.uuid)">
              通过
            </el-button>
            <el-button type="danger"
                       v-if="row.status === 0"
                       @click="refuseHandler"
                       size="small">拒绝</el-button>
            <el-button v-if="row.status === 1"
                       type="danger"
                       @click="refuseHandler"
                       size="small">下架</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- <div class="fotter">
      <el-pagination v-model:current-page="currentPage"
                     v-model:page-size="pageSize"
                     :size="size"
                     :disabled="disabled"
                     :background="background"
                     layout="prev, pager, next, jumper"
                     :total="1000"
                     @size-change="handleSizeChange"
                     @current-change="handleCurrentChange" />
    </div> -->
  </el-card>
</template>

<script lang="ts" setup>
import YbTitle from '@/components/YbTitle/index.vue'
import { getNoteList,passNote } from '@/api/note'
import type { queryNoteListForm } from '@/api/note/type'
import { ref } from 'vue'

const passHandler = (noteId:string) => {
  passNoteReq(noteId)
}
const refuseHandler = () => {
  console.log('click')
}
const tableData = ref([])

const passNoteReq = async (noteId:string) => {
  const res = await passNote(noteId)
  console.log(res)
}

// 获取笔记列表
const getNoteListData = async () => {
  const params: queryNoteListForm = {
    state: -1,
    page: 1,
    size: 10
  }
  const res = await getNoteList(params)
  tableData.value = res.data.list
}

getNoteListData()
</script>

<style lang="scss" scoped>
::v-deep(.el-card__body){
	padding: 0px;
  height: 100%;
  padding: 0 8px 22px;
}
::v-deep(.el-card){
	border-radius: 8px;
  height: 100% !important;
}
.container{
  height: 100%;
  width: 100%;

  .check-box{
    padding-top: 12px;
    height: calc(100% - 150px);
    width: 100%;
    overflow: auto;
    .cover{
      border-radius: 8px;
      height: 100px;
      // 自适应宽度
      width: 100%;
    }
  }
}
</style>
