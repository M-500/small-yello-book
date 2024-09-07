<template>
  <div class="container">
    <el-upload v-model:file-list="fileList"
               :action="uploadUrl"
               list-type="picture-card"
               :on-preview="handlePictureCardPreview"
               :on-success="handleSuccess"
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
import { ref, watch,reactive } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import { defineProps } from 'vue';
import type { UploadProps, UploadUserFile } from 'element-plus'
import { defineEmits } from 'vue';
import { computed } from 'vue';
import { el } from 'element-plus/es/locales.mjs';
// 获取环境变量中的 VITE_SERVER
const serverUrl = import.meta.env.VITE_SERVER;

const uploadUrl = computed(() => {
  return `${serverUrl}/api/v1/file/upload`;
});

const props = defineProps({
  coverImge: {
    type: String,
    default: ''
  }
})
const emit = defineEmits(['itemImgListChanged']);
const fileList = ref<UploadUserFile[]>([])
const dialogImageUrl = ref('')
const dialogVisible = ref(false)

const handleRemove: UploadProps['onRemove'] = (uploadFile, uploadFiles) => {
  console.log("删除咯",uploadFile, uploadFiles)
}

const handlePictureCardPreview: UploadProps['onPreview'] = (uploadFile) => {
  console.log("TMDE 哈哈哈哈",uploadFile)
  dialogImageUrl.value = uploadFile.url!
  dialogVisible.value = true
}

const handleSuccess: UploadProps['onSuccess'] = (response, file, fileList) => {
  console.log("上传成功",response, file, fileList)
  file.url = response.data
}

// 监听父组件传过来的值
watch(() =>props.coverImge,(newVal)=>{
  console.log("父组件传来的值：",newVal)
  fileList.value.push({
    name: 'cover.jpg',
    url: newVal
  })
  // console.log("此时的fileList",fileList)
})
// 监听fileList的变化，实时传递给父组件
watch(fileList,(newVal)=>{
  emit('itemImgListChanged',newVal)
})

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

::v-deep(.el-upload-list__item){
  img{
    background-size: cover;
    object-fit: cover;
  }
}
</style>