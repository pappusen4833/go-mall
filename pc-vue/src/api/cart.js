/**
 * 购物车接口封装
 */
import request from '@/axios/request.js'

// 添加到购物车
export function cartAdd (data) {
    // {
    //     "bargainId": 0, 产品砍价ID
    //     "cartNum": 0, 购物车数量
    //     "combinationId": 0, 产品拼团ID
    //     "new": 0, 是否新购买
    //     "productId": 0, 产品ID
    //     "secKillId": 0, 产品秒杀ID
    //     "uniqueId": "" 唯一的ID
    // }
    return request({
        url: '/cart/add',
        method: 'post',
        data
    })
}

// 获取数量
export function cartCount () {
    return request({
        url: '/cart/count',
        method: 'get'
    })
}

// 删除购物车产品
export function cartDel (data) {
    // {
    //     "ids": [] 购物车ID，多个用,分隔开
    // }
    return request({
        url: '/cart/del',
        method: 'post',
        data
    })
}

// 获取购物车列表
export function cartList () {
    return request({
        url: '/carts',
        method: 'get'
    })
}

// 更新购物车
export function cartUpdate (data) {
    // {
    //     "id": 0, 购物车ID
    //     "number": 0  购物车数量
    // }
    return request({
        url: '/cart/num',
        method: 'post',
        data
    })
}
