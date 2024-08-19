import request from '@/utils/requests'
import { type emialForm } from './types'

// 枚举所有API接口
enum UserApi {
  REGISTER = '/register',
  LOGIN = '/login',
  LOGOUT = '/logout',
  GET_USER_INFO = '/user_info',
  GET_USER_LIST = '/user_list',
  GET_USER_DETAIL = '/user_detail',
  GET_USER_COUNT = '/user_count',
  GET_USER_CAPTCHAS = '/user_captcha',
  GET_EMAIL_CAPTCHAS = '/email_captcha'
}

// 通过接口调用后端API
export const getEmailCaptchas = (data: emialForm) => {
  return request.get(UserApi.GET_EMAIL_CAPTCHAS, {
    params: {
      data
    }
  })
}
