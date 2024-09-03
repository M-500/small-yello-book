import HomeView from '@/views/home/index.vue'
import Layout from '@/layout/index.vue'
import { createRouter, createWebHistory } from 'vue-router'
import { defineComponent, defineAsyncComponent } from 'vue'

// const IndeView = defineComponent(() => import('@/layout/index.vue'))
const PublishView = defineAsyncComponent(() => import('@/views/publish/index.vue'))
// const HomeView = defineComponent(() => import('@/views/home/index.vue'))
const NotesManagerView = defineAsyncComponent(() => import('@/views/notes/index.vue'))
const NotesDashboardView = defineAsyncComponent(() => import('@/views/notes/index.vue'))
const FansDashboardView = defineAsyncComponent(() => import('@/views/notes/index.vue'))

const router = createRouter({
  history: createWebHistory('/admin/'),
  routes: [
    {
      path: '/',
      name: 'home1',
      component: Layout,
      children: [
        {
          path: '/publish',
          name: 'publish',
          component: PublishView
        },
        {
          path: '/home',
          name: 'home',
          component: HomeView
        },
        {
          path: '/notes-manager',
          name: 'notes-manager',
          component: NotesManagerView
        },
        {
          path: '/notes-dashboard',
          name: 'notes-dashboard',
          component: NotesDashboardView
        },
        {
          path: '/fans-dashboard',
          name: 'fans-dashboard',
          component: FansDashboardView
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
