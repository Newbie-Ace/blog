import axios from 'axios'

// 创建 axios 实例
const service = axios.create({
  // 配置请求的根路径
  baseURL: 'http://110.40.205.186:8000/api/v1',
  // 配置请求超时时间为 1s
  timeout: 1000
})

// 设置请求拦截器
service.interceptors.request.use((config: any) => {
  // config.headers.Authorization = window.sessionStorage.getItem('token')
  // 最后必须 return config
  return config
})

// 设置响应拦截器
service.interceptors.response.use((config: any) => {
  // 最后必须 return config
  return config
})

export default service
