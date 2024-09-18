// 用户相关的小仓库
import { login, getUserInfo } from '@/api/user'
import { type loginForm } from '@/api/user/types'
import { useId } from 'element-plus'
import { defineStore } from 'pinia'

// 创建用户小仓库
const useUserStore = defineStore('User', {

  state: () => ({
    token: localStorage.getItem('TOKEN'),
    userInfo: JSON.parse(localStorage.getItem('USER_INFO') || '{}')
  }),
  // 计算属性
  getters: {
    getToken(): string | null {
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
        this.setToken(result.data.token)
        return 'ok'
      } else {
        return Promise.reject(new Error(result.data.data.msg))
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
      localStorage.setItem('TOKEN', token)
    },
    setUserInfo(userInfo: any) {
      this.userInfo = userInfo
      localStorage.setItem('USER_INFO', JSON.stringify(userInfo))
    },
    clearToken() {
      localStorage.clear()
      this.token = ''
    }
  }
})

// 对外暴露小仓库
export default useUserStore
// Compare this snippet from src/utils/requests.ts:
