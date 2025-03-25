<template>
  <div class="header-con">
    <div class="head-box" :class="flag?'grey':''">
      <div class="content banxin">
        <div class="annunciate-box">
          <el-image class="tzgg-icon" src="../../../static/images/tzgg.png" />
          <label>通知公告</label>
        </div>
        <div class="group-icon">
          <router-link to="/coupon">
            <label class="cur-poi">优惠券</label>
          </router-link>
          <label> / </label>
          <router-link to="/myOrder">
            <label class="cur-poi">我的订单</label>
          </router-link>
          <router-link to="/shoppingCart">
            <el-badge :value="cartNum" class="cartNum">
              <el-image class="gwc-icon" src="../../../static/images/gwc.png" />
            </el-badge>
          </router-link>
            <el-image class="kf-icon" src="../../../static/images/kf.png" />
          <router-link to="/userInfo">
            <el-image class="wd-icon" src="../../../static/images/wd.png" />
          </router-link>
        </div>
      </div>
    </div>
    <div class="title-box" :class="flag?'white':''">
      <div class="title-con banxin">
        <div class="logo-box">
          <el-image class="logo-icon" src="../../../static/images/logo.png" />
        </div>
        <div class="classify-box fs16">
          <label>首页</label>
          <label>男士</label>
          <label>女士</label>
          <label>新品上新</label>
          <label>卫衣</label>
          <label>衬衫/T恤</label>
          <div class="serch">
            <input class="serch-input" placeholder="请输入" />
            <div class="serch-text">搜索</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {mapGetters, mapMutations} from 'vuex'
import {cartCount} from '@/api/cart.js'
  export default {
    name: 'v-header',
    data () {
      return {
        flag: this.$route.name === 'UserInfo' || this.$route.name === 'Address' || this.$route.name === 'Footprint' || this.$route.name === 'Coupon' || this.$route.name === 'MyOrder' || this.$route.name === 'OrderDetail' || this.$route.name === 'MyEvaluate' || this.$route.name === 'Evaluate' || this.$route.name === 'ASRecord' || this.$route.name === 'ASRecordDetail' || this.$route.name === 'Favorites'
      }
    },
    computed: {
        ...mapGetters([
          'cartNum' // 购物车商品数量
        ])
    },
    mounted () {
      this.getCartNum()
    },
    methods: {
      ...mapMutations({
          setCartNum: 'SET_CARTNUM' // 修改购物车商品数量
        }),
      // 获取购物车商品数量
      async getCartNum () {
        var res = await cartCount()
        if (res.status === 200) {
          this.setCartNum(res.data.count)
        } else {
          this.$message.warning(res.msg)
        }
      }
    }
  }
</script>

<style scoped lang="scss">
.header-con{
  .head-box{
    display: flex;
    justify-content: center;
    background-color:#FFFFFF;
    height: 40px;
    .content{
      display: flex;
      justify-content: space-between;
      align-items: center;
    }
  }
}

.annunciate-box{
  display: flex;
  align-items: center;
  .tzgg-icon{
    width: 13px;
    height: 14px;
  }
  label{
    color: #666666;
    font-size: 14px;
    margin-left: 8px;
  }
}
.group-icon{
  label{
    color: #666666;
    font-size: 14px;
  }
  .gwc-icon{
    width: 16px;
    height: 14px;
    margin-left: 32px;
  }
  .kf-icon{
    width: 14px;
    height: 14px;
    margin-left: 30px;
  }
  .wd-icon{
    width: 14px;
    height: 14px;
    margin-left: 30px;
  }
}
.title-box{
  background-color: #F6F6F8;
  height: 100px;
  display: flex;
  justify-content: center;
  .title-con{
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
}
.classify-box{
  display: flex;
  flex-direction: row;
  align-items: center;
  color: #666666;
  input::placeholder { /* 大部分现代浏览器 */
    color:#CCCCCC;
    font-size: 14px;
    padding-left: 18px;
  }
  .logo-icon{
    width: 60px;
    height: 22px;
  }
  label{
    margin-left: 54px;
  }
  .serch{
    display: flex;
    flex-direction: row;
    margin-left: 60px;
    .serch-input{
      width: 220px;
      height: 34px;
    }
    .serch-text{
      width: 62px;
      height: 34px;
      background-color: #333333;
      line-height: 34px;
      text-align: center;
      color: #ffffff;
    }
  }
}
.grey{
  background-color: #F6F6F6 !important;
}
.white{
  background-color: #FFFFFF !important;
  .serch-input{
    background-color: #F6F6F6;
  }
}
</style>
