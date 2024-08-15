import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";
import { ElMessage } from "element-plus";

export interface ResponseData {
  code: number;
  data?: any;
  message: string;
}

// console.log('import.meta.env: ', import.meta.env);

// 创建 axios 实例
let service: AxiosInstance | any;

const baseURL =
  import.meta.env.VITE_API_BASE_URL || "http://localhost:5173/api/v1";

if (import.meta.env.MODE === "development") {
  service = axios.create({
    baseURL, // api 的 base_url
    timeout: 50000, // 请求超时时间
    withCredentials: true, // 如果需要发送cookie等凭证信息
    headers: {
      "Content-Type": "application/json",
    },
  });
} else {
  // 生产环境下
  service = axios.create({
    baseURL: "/api/v1",
    timeout: 50000,
    withCredentials: true, // 如果需要发送cookie等凭证信息
    headers: {
      "Content-Type": "application/json",
    },
  });
}

// request 拦截器 axios 的一些配置
service.interceptors.request.use(
  (config: AxiosRequestConfig) => {
    return config;
  },
  (error: any) => {
    // Do something with request error
    console.error("error:", error); // for debug
    Promise.reject(error);
  }
);

// respone 拦截器 axios 的一些配置
service.interceptors.response.use(
  (res: AxiosResponse) => {
    // Some example codes here:
    // code == 0: success
    if (res.status === 200) {
      const data: ResponseData = res.data;
      if (data.code === 0) {
        return data.data;
      } else {
        ElMessage({
          message: data.message,
          type: "error",
        });
      }
    } else {
      ElMessage({
        message: "网络错误!",
        type: "error",
      });
      return Promise.reject(new Error(res.data.message || "Error"));
    }
  },
  (error: any) => Promise.reject(error)
);
//因为别的地方要用，所以就把实例暴露出去，导出
export default service;
