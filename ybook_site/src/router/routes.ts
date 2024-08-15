// 对外暴露配置路由（常量路由 ） 存放所有路由的文件
import HomeView from '../views/HomeView.vue'


export const constantRoutes: Array<RouteRecordRaw> = [
	{
		path: '/',
		name: 'home',
		component: HomeView
	},
	{
		path: '/publish',
		name: 'publish',  // 命名路由，用于做权限的
		component: () => import('@/views/publish/index.vue')
	},
	{
		path: '/notices',
		name: 'notices',
		component: () => import('@/views/notices/index.vue')
	},
	{
		path: '/404',
		name: '404',
		component: () => import('@/views/404/index.vue')
	},
	{
		// 任意路由，都重定向到 404
		path: '/:pathMatch(.*)*',
		redirect: '/404',
		name: 'not-found'
	}
]