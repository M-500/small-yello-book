// 统一管理项目用户相关接口

import request from '@/utils/request'
import { type loginForm, type loginResponseData, type emailForm } from './types' // 妈的 非要加type

enum UserApi {
  // 登录
  LoginAPI = '/api/v1/na/login',
  // 获取用户信息
  UserInfoAPI = '/info',
  // 退出登录
  LogoutAPI = '/logout',
  // 获取验证码
  GetCaptchaAPI = '/api/v1/na/email/send',
  // 获取用户自己信息
  GET_USER_INFO = '/api/v1/users/info',
  // 查询其他用户的基本信息
  GET_USER_INFO_BY_UUID = '/api/v1/users/:uuid'
}

// 对外暴露请求函数

export const loginRequest = (data: loginForm) => {
  // any是请求body的数据类型，loginResponseData是响应的数据类型
  return request.post<any, loginResponseData>(UserApi.LoginAPI, data)
}

// 获取邮箱验证码
export const getCaptchaRequest = (data: emailForm) => {
  return request.post(UserApi.GetCaptchaAPI, data)
}

// 获取用户信息
export const getUserInfo = () => {
  return request.get(UserApi.GET_USER_INFO)
}

export const getUserInfoByUUID = (uuid: string) => {
  return request.get(UserApi.GET_USER_INFO_BY_UUID.replace(':uuid', uuid))
}
