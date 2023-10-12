import request from '@/axios/request.js'

// 提交评论
export function commentOrder (data) {
  // data非必须
  // {
  //     "comment": "", 商品评论内容
  //     "pics": "", 商品评论图片地址
  //     "productScore": "", 商品评分
  //     "serviceScore": "", 服务评分
  //     "unique": "" 订单唯一值
  // }
  return request({
      url: '/order/comment',
      method: 'post',
      data
  })
}

// 提交评论
export function commentOrders (data,id) {
  // data非必须
  // {
  //     "comment": "", 商品评论内容
  //     "pics": "", 商品评论图片地址
  //     "productScore": "", 商品评分
  //     "serviceScore": "", 服务评分
  //     "unique": "" 订单唯一值
  // }
  return request({
      url: `/order/comments/${id}`,
      method: 'post',
      data
  })
}

// 获取商品评论数据
export function getReplyData (id) {
  return request({
    url: `/reply/config/${id}`,
    method: 'get'
  })
}

// 获取商品评论列表
export function getReplyList (id, params) {
  return request({
    url: `/reply/list/${id}`,
    method: 'get',
    params
  })
}
