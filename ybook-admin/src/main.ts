import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import service from './requests'
import ElMessage from 'element-plus'

// createApp(App).mount('#app')
const app = createApp(App)
app.use(ElementPlus)

app.config.globalProperties.$message = ElMessage;
app.config.globalProperties.$https = service;
app.mount('#app')
