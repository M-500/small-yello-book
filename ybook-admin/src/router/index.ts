import component from 'element-plus/es/components/tree-select/src/tree-select-option.mjs'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home1',
      component: () => import('@/layout/index.vue'),
      children: [
        {
          path: '/publish',
          name: 'publish',
          component: () => import('@/views/publish/index.vue')
        },
        {
          path: '/home',
          name: 'home',
          component: () => import('@/views/home/index.vue')
        },
        {
          path: '/notes-manager',
          name: 'notes-manager',
          component: () => import('@/views/notes/index.vue')
        },
        {
          path: '/notes-dashboard',
          name: 'notes-dashboard',
          component: () => import('@/views/notedashboard/index.vue')
        },
        {
          path: '/fans-dashboard',
          name: 'fans-dashboard',
          component: () => import('@/views/fans/index.vue')
        }
      ]
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/login/index.vue')
    }
  ]
})

export default router
