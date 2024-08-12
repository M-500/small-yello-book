import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

// createApp(App).mount('#app')
const app = createApp(App)
app.use(ElementPlus)
app.mount('#app')
