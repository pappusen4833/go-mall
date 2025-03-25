import request from '@/axios/request.js'

import { requestConfig } from '@/api/index.js'

export const uploadURL = `${requestConfig.url}/upload`

// export function uploadData () {
//   return request({
//     url: '/api/api/upload',
//     method: ''
//   })
// }
// 获取城市列表
export function getCityData () {
  return request({
    url: '/city_list',
    method: 'get'
  })
}

// 获取地址列表
export function getAddressData () {
  return request({
    url: '/address',
    method: 'get'
  })
}

// 设置默认地址
export function setDefaultAddress (data) {
  return request({
    url: '/address/default/set',
    method: 'post',
    data
  })
}

// 编辑地址
export function editAddressData (data) {
  return request({
    url: '/address/edit',
    method: 'post',
    data
  })
}

// 删除地址
export function deleteAddress (data) {
  return request({
    url: '/address/del',
    method: 'post',
    data
  })
}

// 获取收藏、足迹
export function getCollectData (params) {
  return request({
    url: '/collect/user',
    method: 'get',
    params
  })
}

// 添加收藏、足迹
export function addCollection (data) {
  return request({
    url: '/collect/add',
    method: 'post',
    data
  })
}

// 删除收藏、足迹
export function deleteCollection (data) {
  return request({
    url: '/collect/del',
    method: 'post',
    data
  })
}

// 批量删除足迹
export function deleteAllCollection (data) {
  return request({
    url: '/collect/delFoot',
    method: 'delete',
    data
  })
}
