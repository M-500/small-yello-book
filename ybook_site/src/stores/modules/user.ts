import { loginRequest } from '@/api/user/index'
import { type loginForm } from '@/api/user/types'
import { type loginResponseData } from '@/api/user/types'
import { defineStore } from 'pinia'

const useUserStore = defineStore({
  id: 'User',
  state: () => ({
    token: localStorage.getItem('x-token') || '',
    userInfo: {}
  }),
  // 计算属性
  getters: {
    getToken(): string | null {
      return this.token
    },
    getUserInfo(): object | null {
      return this.userInfo
    }
  },
  actions: {
    async login(data: loginForm) {
      const res: any = await loginRequest(data)
      if (res.code === 0) {
        this.setToken(res.data.token)
        return 'ok'
      } else {
        return res.msg
      }
    },
    userLogout() {
      this.clearToke()
    },
    setToken(token: string) {
      this.token = token
      localStorage.setItem('x-token', token)
    },
    clearToke() {
      localStorage.clear()
      this.token = ''
    },
    // 判断用户是否登录
    isLogin() {
      return this.token !== ''
    },
    setUserInfo(userInfo: any) {
      this.userInfo = userInfo
    }
  }
})

export default useUserStore
