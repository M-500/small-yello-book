import { loginRequest, getUserInfo } from '@/api/user/index'
import { type loginForm } from '@/api/user/types'
import { type loginResponseData } from '@/api/user/types'
import { defineStore } from 'pinia'

const useUserStore = defineStore({
  id: 'User',
  state: () => ({
    token: localStorage.getItem('x-token') || '',
    userInfo: JSON.parse(localStorage.getItem('USER_INFO') || '{}')
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
    async queryUserInfo() {
      const result: any = await getUserInfo()
      if (result.code === 0) {
        this.setUserInfo(result.data)
        return 'ok'
      } else {
        return Promise.reject(new Error(result.data.data.msg))
      }
    },
    userLogout() {
      this.clearToken()
      this.userInfo = {}
    },
    setToken(token: string) {
      this.token = token
      localStorage.setItem('x-token', token)
    },
    clearToken() {
      localStorage.clear()
      this.token = ''
    },
    // 判断用户是否登录
    isLogin() {
      return this.token !== ''
    },
    setUserInfo(userInfo: any) {
      this.userInfo = userInfo
      localStorage.setItem('USER_INFO', JSON.stringify(userInfo))
    }
  }
})

export default useUserStore
