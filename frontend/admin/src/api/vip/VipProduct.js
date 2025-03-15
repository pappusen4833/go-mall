import request from '@/utils/request'

export function add(data) {
  return request({
    url: 'admin/vipProduct',
    method: 'post',
    data
  })
}

export function del(ids) {
  return request({
    url: 'admin/vipProduct/',
    method: 'delete',
    data: ids
  })
}

export function edit(data) {
  return request({
    url: 'admin/vipProduct',
    method: 'put',
    data
  })
}

export default { add, edit, del }
