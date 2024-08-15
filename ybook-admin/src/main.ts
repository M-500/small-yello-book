import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import service from './requests'
import ElMessage from 'element-plus'
// 1. 引入pinia
import { createPinia } from 'pinia'

// createApp(App).mount('#app')
const app = createApp(App)
// 2. 创建pinia
const pinia = createPinia()

// 3. 安装pinia
app.use(pinia)
app.use(ElementPlus)

app.config.globalProperties.$message = ElMessage;
app.config.globalProperties.$https = service;
app.mount('#app')
