// 对外暴露配置路由（常量路由 ） 存放所有路由的文件

import component from 'element-plus/es/components/tree-select/src/tree-select-option.mjs'
import 'path'
import type { RouteRecordRaw } from 'vue-router'

export const constantRoutes: Array<RouteRecordRaw> = [
  {
    path: '/',
    redirect: '/home',
    component: () => import('@/layout/index.vue'),
    children: [
      {
        path: '/home',
        name: 'home',
        component: () => import('@/views/home.vue')
      },
      {
        path: '/publish',
        name: 'publish', // 命名路由，用于做权限的
        component: () => import('@/views/publish.vue')
      },
      {
        path: '/notices',
        name: 'notices',
        component: () => import('@/views/notices.vue')
      },
      {
        path: '/user/profile/:uuid',
        name: 'user-profile',
        component: () => import('@/views/user.vue')
      },
      {
        path: '/note/detail/:uuid',
        name: 'note-detail',
        component: () => import('@/views/detail.vue')
      },
      {
        path: '/404',
        name: '404',
        component: () => import('@/views/404.vue')
      }
    ]
  },
  {
    // 任意路由，都重定向到 404
    path: '/:pathMatch(.*)*',
    redirect: '/404',
    name: 'not-found'
  }
]
