// 用来定义请求和响应的数据类型

export interface emailForm {
  email: string
  type_code: number
}
// 登陆接口需要的请求数据类型TS
export interface loginForm {
  email: string
  ver_code: string
}

// 登陆接口需要的请求数据类型TS
export interface loginResponseData {
  code: number
  msg: string
  data: {
    token: string
  }
}
