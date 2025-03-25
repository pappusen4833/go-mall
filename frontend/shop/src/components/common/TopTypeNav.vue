/**
    顶部路由组件
    参数：
*/
<template>
  <div class="TRouteNav">
    <div class="wrapper">
      <div class="TPNleft">
          <div>
            <router-link to="/">
              <el-image src="static/images/logo.png" style="width:58px;height:23px;"></el-image>
            </router-link>
          </div>
      </div>
      <div class="TPNright">
          <div class="content"
          v-for="item in this.arr.type"
          :key="item.value"
          @click="routLinkTo(item.route)">
              <div class="personalText"
              v-if="item.des != ''"
              :class="{active_path : activeRoute === item.route }">
                  {{ item.des }}
              </div>
          </div>
          <div class="searchItem">
            <input type="text" placeholder="请输入" v-model="searchKeyWord"><button @click="keyWordSearch">搜索</button>
          </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
    name: 'TopTypeNav',
    props: {
      obj: Object
    },
    data () {
      return {
        arr: {
          'type': [
            {
              'value': 0,
              'des': '首页',
              'route': '/indexpage/index'
            },
            {
              'value': 1,
              'des': '全部商品',
              'route': '/indexpage/productlist'
            },
            {
              'value': 2,
              'des': '秒杀专区',
              'route': '/indexpage/seckill'
            },
            {
              'value': 3,
              'des': '拼团专区',
              'route': '/indexpage/combination'
            },
          ]
        },
        activeRoute: '/indexpage/index',
        searchKeyWord: ''
      }
    },
    watch: {
      $route (to, from) {
        this.activeRoute = to.path
      },
      obj: {
        handler (nVal, oVal) {
          var data = nVal.componentContent.typeRoute
          this.arr.type = data.map((item, index) => {
            return {
              'value': index,
              'des': item.des,
              'route': item.linkTo
            }
          })
        }
      }
    },
    methods: {
      routLinkTo (val) {
        if (!val || val === this.$route.path) return
        if (val == '/indexpage/seckill' || val == '/indexpage/combination') {
          this.$message.warning('即将发布！！')
          return 
        }
        this.$router.push(val)
      },
      keyWordSearch () {
        this.$router.push({
          path: '/indexpage/productlist',
          query: {
            keyword: this.searchKeyWord
          }
        })
      }
    }
}
</script>

<style lang="scss">
.TRouteNav{
  height: 100px;
  line-height: 100px;
  background-color: #FFFFFF;
  .wrapper{
    max-width: 1440px;
    min-width: 1440px;
    margin: auto;
    div{
        display: inline-block;
    }
    .TPNleft{
        cursor: pointer;
        margin-left: 135px;
    }
    .TPNright{
        float: right;
        margin-right: 135px;
        cursor: pointer;
        color: #666666;
        font-size: 16px;
        font-family: PingFang SC;
        font-weight: 400;
        .content{
            margin-right: 30px;
            .personalText{
              height: 80px;
            }
            .active_path{
              border-bottom: 2px solid #F26E47;
            }
        }
        .searchItem{
          margin: 0;
          padding: 0;
          input{
            width: 220px;
            height: 34px;
            padding-left: 20px;
            background-color: #F6F6F8;
          }
          button{
            font-size: 16px;
            width: 62px;
            height: 34px;
            color: #FFFFFF;
            background-color: #333333;
            cursor: pointer;
          }
        }
    }
  }
}
</style>
