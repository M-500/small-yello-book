import { createApp } from 'vue'
import ElementPlus from 'element-plus' // 引入ElementPlus
import 'element-plus/dist/index.css' // 引入ElementPlus的样式
import zhCn from 'element-plus/es/locale/lang/zh-cn' // 引入ElementPlus的中文语言包
import '@/styles/index.scss' // 引入全局SASS样式
import * as ElementPlusIconsVue from '@element-plus/icons-vue' // 引入ElementPlus的ICON
import 'virtual:svg-icons-register' // 引入SVG ICON
import globalComponent from '@/components' // 引入全局组件
// 将 Message 组件挂载到全局
// import { ElMessage } from 'element-plus'
import App from './App.vue'
import router from './router'
import pinia from './stores/index'

const app = createApp(App)
// 注册所有ElementUI的ICON
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

app.use(router)
// 安装仓库
app.use(pinia)

app.use(ElementPlus, {
  locale: zhCn
}) // 使用ElementPlus

app.use(globalComponent) // 使用全局组件

//如果没有全局引用element，还需写下面一句
//app.config.globalProperties.$message = ElMessage;
// app.provide('$message', ElMessage)
app.mount('#app')
