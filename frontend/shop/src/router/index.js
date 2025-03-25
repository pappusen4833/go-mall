import Vue from 'vue'
import Router from 'vue-router'
import layout from '@/views/layout'

import IndexPage from '@/views/PageIndex/Index.vue' // 首页母视图
import indexAll from '@/views/PageIndex/indexAll/Index.vue'

import userCenter from '@/views/userCenter/index' // 个人中心母视图
import userInfo from '@/views/userCenter/userInfo/index'

import Cart from '@/views/cart/Index.vue' // 购物车母视图
// import shoppingCart from '@/views/cart/shoppingCart/index.vue'

// import { component } from 'vue/types/umd'

// const error = (resolve) => {
//   import('@/components/error/error').then((module) => {
//     resolve(module)
//   })
// }
Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/',
      name: 'layout',
      component: layout,
      children: [
        {
          path: '/',
          redirect: '/indexpage/index'
        },
        // 首页
        {
          path: '/indexpage',
          name: 'Index',
          component: IndexPage,
          redirect: '/indexpage/index',
          children: [
            {
              path: '/indexpage/index',
              name: 'IndexAll',
              component: indexAll
            },
            // 首页下分类查询页面
            {
              path: '/indexpage/productlist',
              name: 'ProductList',
              component: () => import('@/views/PageIndex/productList/Index.vue')
            },
            // 公告详情
            // {
            //   path: '/indexpage/noticedetail',
            //   name: 'NoticeDetail',
            //   component: () => import('@/views/PageIndex/noticeDetail/index')
            // }
          ]
        },
        // 商品详情
        {
          path: '/productDetail',
          name: 'ProductDetail',
          component: () => import('@/views/productDetail/index')
        },
        // 购物车
        {
          path: '/cart',
          name: 'cart',
          component: Cart,
          redirect: '/cart/shoppingCart',
          children: [
            // 购物车
            {
              path: '/cart/shoppingCart',
              name: 'ShoppingCart',
              component: () => import('@/views/cart/shoppingCart/index.vue')
            },
            // 结算
            {
              path: '/cart/settlement',
              name: 'Settlement',
              component: () => import('@/views/cart/settlement/index.vue')
            },
            // 支付
            {
              path: '/cart/payment',
              name: 'Payment',
              component: () => import('@/views/cart/payment/index')
            }
          ]
        },
        // 个人中心
        {
          path: '/userCenter',
          name: 'UserCenter',
          component: userCenter,
          redirect: '/userInfo',
          children: [
            // 基本资料
            {
              path: '/userInfo',
              name: 'UserInfo',
              component: userInfo
            },
            // 地址管理
            {
              path: '/address',
              name: 'Address',
              component: () => import('@/views/userCenter/address/index')
            },
            // 我的收藏
            {
              path: '/favorites',
              name: 'Favorites',
              component: () => import('@/views/userCenter/favorites/index')
            },
            // 我的订单
            {
              path: '/myOrder',
              name: 'MyOrder',
              component: () => import('@/views/userCenter/myOrder/index')
            },
            // 订单详情
            {
              // 考虑一下组件复用
              path: '/orderDetail',
              name: 'OrderDetail',
              component: () => import('@/views/userCenter/Detail/Index'),
              beforeEnter: (to, from, next) => {
                if (to.fullPath === '/personal/detail' || to.query.id.length <= 0) {
                  next({path: '/'})
                }
                next()
              }
            },
            // 评价
            {
              path: '/myEvaluate',
              name: 'MyEvaluate',
              component: () => import('@/views/userCenter/myEvaluate/index')
            },
            // 提交评价
            {
              path: '/evaluate',
              name: 'Evaluate',
              component: () => import('@/views/userCenter/myEvaluate/evaluate/index')
            },
           
          ]
        }
      ]
    }
  ]
})

// router.beforeEach((to, from, next) => {
//   console.log('beforeEach')
//   // next('/indexpage/index')
//   next()
//   // cookie没有值就跳到登陆页
//   // let area = sessionStorage.getItem('area')
//   // if (area === null && to.name !== 'login' && to.name !== 'overview') {
//   //   // Vue.prototype.$message.error('用户未登陆')
//   //   next('/login')
//   // } else {
//   // next()
//   // }
// })

const originalPush = Router.prototype.push
Router.prototype.push = function push (location, onResolve, onReject) {
  if (onResolve || onReject) {
    return originalPush.call(this, location, onResolve, onReject)
  }
  return originalPush.call(this, location).catch(err => err)
}

export default router
