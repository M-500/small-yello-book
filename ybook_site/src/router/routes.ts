// 对外暴露配置路由（常量路由 ） 存放所有路由的文件

import component from 'element-plus/es/components/tree-select/src/tree-select-option.mjs'
import { pa } from 'element-plus/es/locales.mjs'
import path from 'path'
import type { RouteRecordRaw } from 'vue-router'

export const constantRoutes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'index',
    component: () => import('@/layout/index.vue'),
    children: [
      {
        path: '/home',
        name: 'home',
        component: () => import('@/views/home/index.vue')
      },
      {
        path: '/publish',
        name: 'publish', // 命名路由，用于做权限的
        component: () => import('@/views/publish/index.vue')
      },
      {
        path: '/notices',
        name: 'notices',
        component: () => import('@/views/notices/index.vue')
      },
      {
        path: '/user/profile/:id',
        name: 'user-profile',
        component: () => import('@/views/user/index.vue')
      },
      {
        path: '/note/detail/:uuid',
        name: 'note-detail',
        component: () => import('@/views/detail/index.vue')
      },
      {
        path: '/404',
        name: '404',
        component: () => import('@/views/404/index.vue')
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
