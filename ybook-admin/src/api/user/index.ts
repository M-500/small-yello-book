import request from '@/utils/requests'
import { type emialForm, type loginForm, type loginResponse } from './types'

// 枚举所有API接口
enum UserApi {
  // REGISTER = '/register',

  // LOGOUT = '/logout',
  // GET_USER_INFO = '/user_info',
  // GET_USER_LIST = '/user_list',
  // GET_USER_DETAIL = '/user_detail',
  // GET_USER_COUNT = '/user_count',
  // GET_USER_CAPTCHAS = '/user_captcha',
  LOGIN = '/api/v1/na/login',
  GET_EMAIL_CAPTCHAS = '/api/v1/na/email/send'
}

// 通过接口调用后端API
export const getEmailCaptchas = (data: emialForm) => {
  return request.post(UserApi.GET_EMAIL_CAPTCHAS, data)
}

export const login = (data: loginForm) => {
  return request.post<any, loginResponse>(UserApi.LOGIN, data)
}
