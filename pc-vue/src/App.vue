<template>
  <div id="app">
    <!-- <v-header v-if="isRouterAlive"></v-header> -->
    <PersonalNav v-if="isRouterAlive"/>
    <TypeNav :obj="TypeRoute"/>
    <router-view></router-view>
    <VFoot :obj="FooterData"/>
  </div>
</template>

<script>
import VHeader from '@/components/base/v-header'
import VFoot from '@/components/base/v-foot'
import PersonalNav from '@/components/common/TopPersonalNav'
import TypeNav from '@/components/common/TopTypeNav'
import {
  getCanvasData
} from '@/api/indexpage.js'
export default {
  name: 'App',
  components: {VHeader, VFoot, PersonalNav, TypeNav},
  provide () {
    return {
      reload: this.reload
    }
  },
  data () {
    return {
      isRouterAlive: true,
      TypeRoute: {},
      FooterData: {}
    }
  },
  watch: {
    $route (to, from) {
      if (to.name !== from.name) {
        this.reload()
      }
    }
  },
  created () {
    this.getCanvas()
    // 在页面刷新时将vuex里的信息保存到localStorage里
    window.addEventListener('beforeunload', () => {
      localStorage.setItem('messageStore', JSON.stringify(this.$store.state))
    })
    // 在页面加载时读取localStorage里的状态信息
    localStorage.getItem('messageStore') && this.$store.replaceState(Object.assign(this.$store.state, JSON.parse(localStorage.getItem('messageStore'))))
  },
  methods: {
    async getCanvas () {
      var res = await getCanvasData({ terminal: '1' })
      if (!res.json) {
        return
      }
      var canvasData = JSON.parse(res.json)
      canvasData = canvasData.map(item => {
        if (item.type === 'NavType') {
          return item
        }
        if (item.type === 'Footer') {
          return item
        }
      }).filter(r => r)
      this.TypeRoute = this.giveTypeRes('NavType', canvasData)
      this.FooterData = this.giveTypeRes('Footer', canvasData)
    },
    giveTypeRes (type, data) {
      var result
      for (var key in data) {
        if (data[key].type === type) {
          result = data[key]
        }
      }
      return result
    },
    reload () {
      this.isRouterAlive = false
      this.$nextTick(function () {
        this.isRouterAlive = true
      })
    }
  }
}
</script>
<style lang="scss">
#app{
  min-width: 1440px;
  background-color: #F6F6F8;
}
.el-loading-mask{
  .el-loading-spinner{
    .path{
      stroke: #333333;
    }
  }
}
// loading
.el-loading-spinner{
  i,.el-loading-text{
    font-size: 18px;
    color: #333333;
  }
}
</style>
