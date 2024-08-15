import { createRouter, createWebHistory } from 'vue-router'
import { constantRoutes } from './routes'

// 创建路由器实例
const router = createRouter({
  // 使用 HTML5 历史模式
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: constantRoutes,
  // 滚动行为
  // scrollBehavior(to, from, savedPosition) {
  //   if (savedPosition) {
  //     return savedPosition
  //   } else {
  //     return { top: 0 }
  //   }
  // }
  scrollBehavior() {
    return { top: 0,left:0 }
  }
})

export default router
