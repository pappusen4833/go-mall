import request from '@/utils/request'

export function add(data) {
  return request({
    url: 'admin/promptPreset',
    method: 'post',
    data
  })
}

export function del(ids) {
  return request({
    url: 'admin/promptPreset/',
    method: 'delete',
    data: ids
  })
}

export function edit(data) {
  return request({
    url: 'admin/promptPreset',
    method: 'put',
    data
  })
}

export default { add, edit, del }
