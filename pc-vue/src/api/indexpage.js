import request from '@/axios/request.js'

// 获取首页内容
export function getIndexPageList () {
    return request({
        url: '/index',
        method: 'get'
    })
}

// 热门搜索关键字获取
export function getKeyWord () {
    return request({
        url: '/search/keyword',
        method: 'get'
    })
}

// 获取商品列表
export function getProducts (params) {
    // {
    //     'isIntegral': 0, // 是否积分兑换商品 0否  1是
    //     'keyword': '', // 关键
    //     'news': '', // 是否新品  0否  1是
    //     'priceOrder': '', // 价格排序  desc正序  asc倒序
    //     'salesOrder': '', // 销量排序
    //     'sid': '', // 分类ID
    //     'type': '', // 类别 不传
    //     'limit': 10,
    //     'page': 1
    // }
    return request({
        url: '/products',
        method: 'get',
        params
    })
}

// 普通商品详情
export function getNormalProductDetail (id) {
    return request({
        url: `/product/detail/${id}`,
        method: 'get'
    })
}

// 商品详情海报
export function getProductPoster (id, params) {
    return request({
        url: `/product/poster/${id}`,
        method: 'get',
        params
    })
}

// 获取分类id
export function getCategory () {
    return request({
        url: '/category',
        method: 'get'
    })
}

// 热门推荐
export function getHotProduct () {
    return request({
        url: '/product/hot',
        method: 'get'
    })
}

export function getCanvasData (params) {
    return request({
        url: '/getCanvas',
        method: 'get',
        params
    })
}
