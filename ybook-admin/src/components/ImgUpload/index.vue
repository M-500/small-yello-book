<template>
  <div class="container">
    <el-upload v-model:file-list="fileList"
               action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
               list-type="picture-card"
               :on-preview="handlePictureCardPreview"
               :on-remove="handleRemove">
      <el-icon>
        <Plus />
      </el-icon>
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
import { Plus } from '@element-plus/icons-vue'

import type { UploadProps, UploadUserFile } from 'element-plus'

const fileList = ref<UploadUserFile[]>([
  {
    name: 'food.jpeg',
    url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100',
  },
  {
    name: 'plant-1.png',
    url: '/images/plant-1.png',
  },
  {
    name: 'food.jpeg',
    url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100',
  },
  {
    name: 'plant-2.png',
    url: '/images/plant-2.png',
  },
  {
    name: 'food.jpeg',
    url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100',
  },
  {
    name: 'figure-1.png',
    url: '/images/figure-1.png',
  },
  {
    name: 'food.jpeg',
    url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100',
  },
  {
    name: 'figure-2.png',
    url: '/images/figure-2.png',
  },
])

const dialogImageUrl = ref('')
const dialogVisible = ref(false)

const handleRemove: UploadProps['onRemove'] = (uploadFile, uploadFiles) => {
  console.log(uploadFile, uploadFiles)
}

const handlePictureCardPreview: UploadProps['onPreview'] = (uploadFile) => {
  dialogImageUrl.value = uploadFile.url!
  dialogVisible.value = true
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