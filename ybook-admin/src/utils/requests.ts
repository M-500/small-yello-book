// 对axios进行二次封装：统一处理请求异常，统一处理请求loading，统一处理请求结果
import router from '@/router'
import axios from 'axios'
import { ElMessage } from 'element-plus'
import useUserStore from '@/stores/moudules/user'
// axios.defaults.baseURL = import.meta.env.VITE_APP_BASE_API
// 创建axios实例 使用Create方法
const request = axios.create({
  baseURL: import.meta.env.VITE_SERVER, // api的base_url(基础路径配置)
  timeout: 5000 // 请求超时时间
})

// request 请求拦截器
request.interceptors.request.use((config) => {
  // 请求前的处理
  const token = localStorage.getItem('TOKEN') || sessionStorage.getItem('TOKEN')
  if (token) {
    // 如果存在 token，则在请求头中携带
    config.headers['Authorization'] = `Bearer ${token}`
  }
  return config
})

// request 响应拦截器
request.interceptors.response.use(
  (response) => {
    // 请求成功后的处理 这个成功只是网络成功，不是业务成功，业务成功需要根据业务状态码来判断
    // 业务成功需要和后端约定好状态码，一般是200，或者0，或者其他的
    // console.log('哈哈？', response.data)
    // if (response.data.code === 0) {
    //   return response.data.data // 简化数据，只返回data
    // }
    return response.data // 简化数据，只返回data
  },
  (error) => {
    // 请求失败后的处理：一般处理http的网络错误
    let message: string = ''
    const status = error.response.status // HTTP状态码
    const userStore = useUserStore()
    switch (status) {
      case 400:
        message = '请求错误'
        break
      case 401:
        message = '登录失效，请登录'
        // 退出登录
        userStore.userLogout()
        router.push('/login')
        break
      case 403:
        message = '权限不够，拒绝访问'
        break
      case 404:
        message = `请求地址出错: ${error.response.config.url}`
        break
      case 500:
        if (error.response.data.msg) {
          message = error.response.data.msg
          break
        }
        message = '服务器内部错误:'
        break
      default:
        message = '网络异常'
        break
    }
    ElMessage({
      message,
      type: 'error',
      duration: 5 * 1000
    })
    // 返回一个失败的Promise
    return Promise.reject(error)
  }
)

// 对外暴露
export default request
