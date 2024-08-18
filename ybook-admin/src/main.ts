import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus' // 引入ElementPlus
import 'element-plus/dist/index.css' // 引入ElementPlus的样式
import zhCn from 'element-plus/es/locale/lang/zh-cn' // 引入ElementPlus的中文语言包
import '@/styles/index.scss' // 引入全局SASS样式

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(ElementPlus, {
  locale: zhCn
}) // 使用ElementPlus
app.mount('#app')
