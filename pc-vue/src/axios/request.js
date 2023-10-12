/* eslint-disable */
import Vue from 'vue'
import axios from 'axios'
import store from '@/store'
import { requestConfig } from '@/api/index.js'
import { getToken, removeToken } from '@/axios/cookies.js'

const service = axios.create({
    baseURL: requestConfig.url,
    timeout: requestConfig.timeout
})

service.interceptors.request.use(
    /**
     * 请求拦截器，在请求之前做配置
    */
    config => {
        config.headers['Content-Type'] = 'application/json'
        // if (store.getters.token) {
        //     config.headers['Authorization'] = 'Bearer ' + getToken()
        // }
        if (getToken()) {
            config.headers['Authorization'] = 'Bearer ' + getToken()
        }
        return config
    },
     error => {
         console.error(error)
         return Promise.reject(error)
     }
)

service.interceptors.response.use(
    /**
     * 响应拦截器，在请求之前做配置
    */
    response => {
        const res = response.data
        if (response.data.code === 504) {} 

        if (response.data.status === 401) {
            Vue.prototype.$message.warning('请先登录')
            // router.go(-1)
            removeToken()
            store.commit('setLoginVisible', true)
            store.commit('reset')
        }
        return res
    },
    error => {
        console.log('err' + error.response.data.msg)
        // 失败响应
        // if (err && err.response) {
        // switch (err.response.status) {
        //   case 400:
        //     err.message = '请求无效，请检查参数是否正确!'
        //     break
  
        //   case 401:
        //     err.message = '未经授权，访问被拒!'
        //     break
  
        //   case 403:
        //     err.message = '拒绝访问!'
        //     break
  
        //   case 404:
        //     err.message = `地址不存在!`
        //     break
  
        //   case 408:
        //     err.message = '请求超时!'
        //     break
  
        //   case 500:
        //     err.message = '系统错误!'
        //     break
  
        //   case 501:
        //     err.message = '该方法未实现!'
        //     break
  
        //   case 502:
        //     err.message = '网管出错!'
        //     break
  
        //   case 503:
        //     err.message = '服务不可用!'
        //     break
  
        //   case 504:
        //     err.message = '网关请求超时'
        //     break
  
        //   case 505:
        //     err.message = 'HTTP版本不受支持'
        //     break
  
        //   default:
        // }
        // if (err.response.data.error) {
        //   err.message = err.response.data.error
        // }
        // Vue.prototype.$message.closeAll()
        // Vue.prototype.$message.error(err.message)
        Vue.prototype.$message.warning(error.response.msg)
        console.log(error.response.status)
        if (error.response.status === 401) {
            Vue.prototype.$message.warning('请先登录')
            // router.go(-1)
            removeToken()
            store.commit('setLoginVisible', true)
            store.commit('reset')
        }
        // 401无token
        if (error.response.status === 401 || error.response.data.status === 401) {
            store.dispatch('user/resetToken').then(() => {
                location.reload()
            })
        }
        Vue.prototype.$message.error(error.response.data.msg)
        return Promise.reject(error)
    }
)

export default service
