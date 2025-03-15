import request from '@/utils/request'

export function getDepts(params) {
  return request({
    url: 'admin/gptkey',
    method: 'get',
    params
  })
}

export function add(data) {
  return request({
    url: 'admin/gptkey',
    method: 'post',
    data
  })
}

export function del(ids) {
  return request({
    url: 'admin/gptkey',
    method: 'delete',
    data: ids
  })
}

export function edit(data) {
  return request({
    url: 'admin/gptkey',
    method: 'put',
    data
  })
}

export default { add, edit, del }
