import request from '@/axios/request.js'

/**
    订单列表  getOrderList
    订单取消 cancelOrder
    订单评价 commentOrder
    计算订单金额 computedOrder
    订单确认 confirmOrder
    订单床架 createOrder
    订单删除 deleteOrder
    订单详情 detailOrder
    物流信息 expressOder
    获取订阅消息模板ID getSubscribeTemplate
    订单核销 verificOrder
    订单支付 payOrder
    订单产品信息 productOrder
    订单退款理由 reasonOrder
    订单退款审核 verifyOrder
    订单收货 takeOrder
 */

// 订单列表
export function getOrderList (params) {
    // {
    //     'limit': 10,
    //     'page': 1,
    //     'type': -1
    // }
    return request({
        url: '/order',
        method: 'get',
        params
    })
}

// 取消订单
export function cancelOrder (data) {
    // id非必须
    // {
    //     "id": ""
    // }
    return request({
        url: '/order/cancel',
        method: 'post',
        data
    })
}

// 计算订单金额
export function computedOrder (key, data) {
    // key与data非必须参数
    // {
    //     "addressId": "", 地址ID
    //     "bargainId": "", 砍价ID
    //     "combinationId": "", 拼团ID
    //     "couponId": "", 优惠券ID
    //     "payType": "", 支付方式
    //     "pinkId": "", 拼团ID
    //     "shipping_type": "", 配送方式 1=快递 ，2=门店自提
    //     "useIntegral": "" 使用积分 1-表示使用
    // }
    return request({
        url: `/order/computed/${key}`,
        method: 'post',
        data
    })
}

// 订单确认 结算详情
export function confirmOrder (data) {
    // {
    //     "cartId": "" 购物车id
    // }
    return request({
        url: '/order/confirm',
        method: 'post',
        data
    })
}

// 创建订单
export function createOrder (key, data) {
    // key必须 data非必须
    // {
    //     "addressId": "", 地址ID
    //     "bargainId": "", 砍价ID
    //     "combinationId": "", 拼团ID
    //     "couponId": "", 优惠券ID
    //     "from": "", 来源
    //     "isChannel": "", 支付渠道(0微信公众号1微信小程序)
    //     "mark": "", 备注
    //     "payType": "", 支付方式
    //     "phone": "", 门店电话
    //     "pinkId": "", 拼团id 0没有拼团
    //     "realName": "", 门店联系人
    //     "seckillId": "", 秒杀产品ID
    //     "shippingType": "", 配送方式 1=快递 ，2=门店自提
    //     "storeId": "", 门店ID
    //     "useIntegral": "" 使用积分 1-表示使用
    // }
    return request({
        url: `/order/create/${key}`,
        method: 'post',
        data
    })
}

// 删除订单
export function deleteOrder (data) {
    // {
    //     "uni": "" 订单ID
    // }
    return request({
        url: '/order/del',
        method: 'post',
        data
    })
}

// 订单详情
export function detailOrder (key) {
    return request({
        url: `/order/detail/${key}`,
        method: 'get'
    })
}

// 物流信息
export function expressOder (data) {
    // 必须参数
    // expressInfoDo: {
    //     "logisticCode": "", 物流单号
    //     "orderCode": "", 订单编号
    //     "shipperCode": "" 快递公司编码
    // }
    return request({
        url: '/order/express',
        method: 'post',
        data
    })
}

// 获取订阅消息模板ID
export function getSubscribeTemplate () {
    return request({
        url: '/order/getSubscribeTemplate',
        method: 'get'
    })
}

// 订单核销
export function verificOrder (data) {
    // {
    //     "isConfirm": 0, 订单核销状态：1确认0正常
    //     "verifyCode": "" 核销码
    // }
    return request({
        url: '/order/order_verific',
        method: 'post',
        data
    })
}

// 订单支付
export function payOrder (data) {
    // {
    //     "isConfirm": 0, 订单核销状态：1确认0正常
    //     "verifyCode": "" 核销码
    // }
    return request({
        url: '/order/pay',
        method: 'post',
        data
    })
}

// 订单产品信息
export function productOrder (data) {
    // {
    //     "unique": ""
    // }
    return request({
        url: '/order/product',
        method: 'post',
        data
    })
}

// 订单退款理由
export function reasonOrder () {
    return request({
        url: '/order/refund/reason',
        method: 'get'
    })
}

// 订单退款审核
export function verifyOrder (data) {
    // {
    //     "refund_reason_wap_explain": "", 退款备注
    //     "refund_reason_wap_img": "", 退款图片
    //     "text": "", 退款原因
    //     "uni": "" 订单唯一值
    // }
    return request({
        url: '/order/refund/verify',
        method: 'post',
        data
    })
}

// 订单收货
export function takeOrder (data) {
    // {
    //     "uni": "" 订单ID
    // }
    return request({
        url: '/order/take',
        method: 'post',
        data
    })
}

// 获取门店列表
export function getStoreList () {
    return request({
        url: '/store_list',
        method: 'get'
    })
}
