// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
// import '@/config/rem' // 配置flex
import 'normalize.css' // 重置样式表
import VuePhotoZoomPro from 'vue-photo-zoom-pro'
// import '@/styles/element-ui-style.css' // 重置样式表
import maps from 'qqmap'
import Moment from 'moment'

import '@/router/permission.js'

Vue.use(VuePhotoZoomPro)
Vue.use(maps)
Vue.use(ElementUI)

Vue.filter('comverTime', function (data, format) {
  return Moment(data).format(format)
})
Vue.prototype.$moment = Moment;//赋值使用
Vue.prototype.$message = ElementUI.Message

Vue.prototype.openLoading = function () {
  const loading = this.$loading({ // 声明一个loading对象
    lock: true, // 是否锁屏
    text: 'loading...',
    spinner: 'el-icon-loading',
    background: 'rgba(255, 255, 255, 1)'
  })
  // setTimeout(function () { // 设定定时器，超时5S后自动关闭遮罩层，避免请求失败时，遮罩层一直存在的问题
  //   loading.close() // 关闭遮罩层
  // }, 5000)
  return loading
}

// 阻止启动生产消息
Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
