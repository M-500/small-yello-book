// 统一管理项目用户相关接口

import request from "@/utils/request";
import { type loginForm, type loginResponseData } from "./types"; // 妈的 非要加type

enum UserApi {
	// 登录
	LoginAPI = "/login",
	// 获取用户信息
	UserInfoAPI = "/info",
	// 退出登录
	LogoutAPI = "/logout",   
	// 获取验证码
	GetCaptchaAPI = "/captcha"
}

// 对外暴露请求函数

export const loginRequest = (data:loginForm) => {
	// any是请求body的数据类型，loginResponseData是响应的数据类型
	return request.post<any,loginResponseData>(UserApi.LoginAPI, data);
}

// 获取邮箱验证码
export const getCaptchaRequest = (email: string) => {
	return request.get(UserApi.GetCaptchaAPI, {
		params: {
			email
		}
	})
}