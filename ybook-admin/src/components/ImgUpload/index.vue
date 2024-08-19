<template>
  <div class="container">
    <el-upload action="#"
               list-type="picture-card"
               :auto-upload="false">
      <el-icon>
        <Plus />
      </el-icon>

      <template #file="{ file }">
        <div>
          <img class="el-upload-list__item-thumbnail"
               :src="file.url"
               alt="" />
          <span class="el-upload-list__item-actions">
            <span class="el-upload-list__item-preview"
                  @click="handlePictureCardPreview(file)">
              <el-icon>
                <zoom-in />
              </el-icon>
            </span>
            <span v-if="!disabled"
                  class="el-upload-list__item-delete"
                  @click="handleDownload(file)">
              <el-icon>
                <Download />
              </el-icon>
            </span>
            <span v-if="!disabled"
                  class="el-upload-list__item-delete"
                  @click="handleRemove(file)">
              <el-icon>
                <Delete />
              </el-icon>
            </span>
          </span>
        </div>
      </template>
    </el-upload>

    <el-dialog v-model="dialogVisible">
      <img w-full
           :src="dialogImageUrl"
           alt="Preview Image" />
    </el-dialog>
  </div>

</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { Delete, Download, Plus, ZoomIn } from '@element-plus/icons-vue'

import type { UploadFile } from 'element-plus'

const dialogImageUrl = ref('')
const dialogVisible = ref(false)
const disabled = ref(false)

const handleRemove = (file: UploadFile) => {
  console.log(file)
}

const handlePictureCardPreview = (file: UploadFile) => {
  dialogImageUrl.value = file.url!
  dialogVisible.value = true
}

const handleDownload = (file: UploadFile) => {
  console.log(file)
}
</script>


<style lang="scss" scoped>
::v-deep(.el-upload--picture-card){
	width: 98px;
	height: 98px;
	flex: 0 0 98px;
	border-radius: 4px;
	border: 1px dashed #d9d9d9;
	background: #f7f7f7;
	display: flex;
	flex-direction: column;
	justify-content: center;
	padding-top: 24px;
	padding-bottom: 22px;
	margin-right: 10px;
	cursor: pointer;
}
::v-deep(.el-upload-list__item){
	width: 98px;
	height: 98px;
	flex: 0 0 98px;
}

.el-upload-list__item-thumbnail{
	width: 100%;
	height: 100%;
	object-fit: cover; /* 使图片铺满容器 */
}
</style>