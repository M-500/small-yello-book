// 用户相关的小仓库
import { login } from '@/api/user'
import { type loginForm } from '@/api/user/types'
import { defineStore } from 'pinia'

// 创建用户小仓库
const useUserStore = defineStore('User', {
  // state: () => {
  // 	return {
  // 		token: '',
  // 		userInfo: {}
  // 	}
  // }

  state: () => ({
    token: localStorage.getItem('TOKEN'),
    userInfo: {}
  }),
  // 计算属性
  getters: {
    getToken(): string {
      return this.token
    },
    getUserInfo(): any {
      return this.userInfo
    }
  },
  // 处理异步|逻辑的地方
  actions: {
    async userLogin(form: loginForm) {
      const result: any = await login(form)
      // 要么成功 200
      // 要么失败 != 200
      if (result.code === 0) {
        // 由于pinia|vuex存储数据其实就是用的js对象 所以还需要持久化存储
        this.token = result.data.token
        localStorage.setItem('TOKEN', result.data.token)
        return 'ok'
      } else {
        return Promise.reject(result.data.msg)
      }
    },
    setToken(token: string) {
      this.token = token
    },
    setUserInfo(userInfo: any) {
      this.userInfo = userInfo
    }
  }
})

// 对外暴露小仓库
export default useUserStore
// Compare this snippet from src/utils/requests.ts:
