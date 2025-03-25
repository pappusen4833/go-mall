import router from '@/router/index.js'
import store from '@/store'
import { getToken } from '@/axios/cookies.js'

router.beforeEach((to, from, next) => {
  const token = getToken()
  const whiteList = [
    '/',
    '/indexpage',
    '/indexpage/index',
    '/indexpage/productlist', // 商品列表
    '/indexpage/seckill', // 秒杀中心
    '/indexpage/discount', // 折扣中心
    '/indexpage/combination', // 拼团中心
    '/indexpage/getcoupon', // 优惠卷中心
    '/indexpage/noticedetail', // 图文详情
    '/productDetail' // 商品详情
  ]
  if (token) {
    next()
  } else {
    if (whiteList.indexOf(to.path) !== -1) {
      next()
    } else {
      store.commit('setLoginVisible', true)
      next('/')
    }
  }
})
