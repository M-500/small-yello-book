import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus' // 引入ElementPlus
import 'element-plus/dist/index.css' // 引入ElementPlus的样式
import zhCn from 'element-plus/es/locale/lang/zh-cn' // 引入ElementPlus的中文语言包
import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(ElementPlus, {
  locale: zhCn // 使用ElementPlus的中文语言包
}) // 使用ElementPlus

import 'virtual:svg-icons-register' // 引入svg图标
import globalComponent from '@/components' // 引入全局组件

app.use(globalComponent) // 使用全局组件
app.mount('#app')
