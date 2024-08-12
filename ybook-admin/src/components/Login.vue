<template>
    
<el-dialog
    v-model="dialogVisible"
    title="Tips"
    width="500"
    :before-close="handleClose"
  >
    <div class="login-container">
        登陆组件啊傻逼
    </div>
  </el-dialog>
</template>



<script setup>
import { ref } from 'vue';
const dialogVisible = ref(false)

function handleClose(done) {
    dialogVisible.value = false;
    done();
}
</script>
