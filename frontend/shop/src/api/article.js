import request from '@/axios/request'

export function articleList () {
    return request({
        url: '/api/article/list',
        method: 'get'
    })
}

export function articleDetial (id) {
    return request({
        url: `/api/article/details/${id}`,
        method: 'get'
    })
}
