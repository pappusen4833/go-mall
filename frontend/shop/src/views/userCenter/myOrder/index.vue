/**
    我的订单组件
    rotue: /personal/oerder
    涉及字段：
    type: 默认为0未支付 1待发货 2待收货 3待评价 4已完成 5退款中 6已退款 7退款
*/
<template>
  <div class="personalMyOrder">
      <nav>我的订单</nav>
      <ul class="MyOrderSearch">
        <li class="orderSearchType" v-for="(item,index) in searchList" :key="index" @click="searchType(item.value,index)">
          <div :class="{search_active: isActive == index}"  >{{ item.label }}</div>
          <div class="orderSearchTypeTool" v-if="index != (searchList.length - 1)"></div>
        </li>
        <div class="searchBtn">
          <input
          ref="searchOrder"
          @keyup.enter="searchBtn"
          type="text"
          placeholder="请输入订单号"
          ><button @click="searchBtn">搜索</button>
        </div>
      </ul>
      <!-- 订单列表 -->
      <section v-loading="loading" v-if="this.itemList">
        <ListItem
        :list="itemList"
        @del="isDel"
        @cancel="isCancel"
        @invit="isInvi"
        @take="isTake"/>
      </section>
      <section class="notfound" v-else >
        <NFInfo />
      </section>
      <el-pagination
        class="elePag"
        layout="prev, pager, next"
        :total="this.totalPage"
        :current-page="this.orderQuery.page"
        @current-change="currentchange">
      </el-pagination>
      <!-- 删除订单弹框 -->
      <el-dialog
      class="delDia"
      :visible.sync="delDialogVisible"
      width="600px"
      center>
          <div class="diaContent">
              <div class="diaTitle">确认删除{{ this.orderId }}订单吗？</div>
          </div>
          <div class="diaBtn">
              <button @click="delDialogVisible = false">取 消</button>
              <el-button :loading="loadingBtn" class="yes" @click="dleSure(oid)">确 定</el-button>
          </div>
      </el-dialog>
      <!-- 取消订单弹框 -->
      <el-dialog
      class="delDia"
      :visible.sync="cancelDialogVisible"
      width="600px"
      center>
          <div class="diaContent">
              <div class="diaTitle">确认取消订单吗？</div>
          </div>
          <div class="diaBtn">
              <button @click="cancelDialogVisible = false">取 消</button>
              <el-button :loading="loadingBtn" class="yes" @click="canOrder(orderId)">确 定</el-button>
          </div>
      </el-dialog>
      <!-- 我的订单-邀请好友 -->
      <el-dialog
      class="delDia"
      :visible.sync="inviteDialogVisible"
      width="670px"
      center>
            <div class="diaContent noBottom">
                <div class="diaTitle">邀请好友</div>
                <div class="diaText">扫描下方二维码或复制链接分享给好友吧</div>
                <img class="pink" :src="this.qrcode" alt="">
                <!-- <VueQrcode
                class="qrcode"
                :width=270
                :margin=1
                :value="this.qrcode" /> -->
            </div>
            <div class="diaBtn">
                <!-- <span class="shareTxt">{{ this.qrcode }}</span> -->
                <button class="cpqr" @click="useClipboard(qrcode,$event)">复制链接</button>
            </div>
        </el-dialog>
        <!-- 确认收货 -->
        <el-dialog
        class="delDia"
        :visible.sync="receDialogVisible"
        width="600px"
        center>
            <div class="diaContent">
                <div class="diaTitle">确认收到货了吗</div>
                <div class="diaText">为保障您的售后权益，请核对商品无误后确认收货</div>
            </div>
            <div class="diaBtn">
                <button @click="receDialogVisible = false">取 消</button>
                <el-button :loading="loadingBtn" class="yes" @click="takOrder(orderId)">确 定</el-button>
            </div>
        </el-dialog>
  </div>
</template>

<script>
import ListItem from './components/orderItem'
import NFInfo from '@/components/NotFound/NotFoundInfo.vue'
import {
  getOrderList,
  detailOrder,
  deleteOrder,
  cancelOrder,
  takeOrder
} from '@/api/order.js'
import VueQrcode from 'vue-qrcode'
import Clipboard from 'clipboard'

export default {
  components: {ListItem, NFInfo, VueQrcode},
  data () {
    return {
      itemList: [],
      isActive: 0,
      loading: false,
      searchList: [
        {
          'label': '全部订单',
          'value': -1
        },
        {
          'label': '待付款',
          'value': 0
        },
        {
          'label': '待发货',
          'value': 1
        },
        {
          'label': '待收货',
          'value': 2
        },
        {
          'label': '待评价',
          'value': 3
        },
        {
          'label': '已完成',
          'value': 4
        }
      ],
      orderQuery: {
        'limit': 10,
        'page': 1,
        'type': -1
      },
      totalPage: 10,
      // 弹框相关
      delDialogVisible: false,
      cancelDialogVisible: false,
      inviteDialogVisible: false,
      receDialogVisible: false,
      oid: null, // 需要操作的订单id
      orderId: null, // 需要操作的订单号
      qrcode: '',
      loadingBtn: false
    }
  },
  created () {
    this.orderList(this.orderQuery)
  },
  methods: {
    // 订单查询
    async orderList (val) {
      this.loading = true
      var res = await getOrderList(val)
      if (res.status === 200) {
        this.totalPage = res.totalPage * 10
        this.getItemData(res)
        this.loading = false
      } else {
        this.$message.error(res.msg || '请求错误！')
      }
    },
    // 根据类型查询 、 修改导航选中状态
    searchType (value, index) {
      this.orderQuery.page = 1
      this.orderQuery.type = value
      this.orderList(this.orderQuery)
      this.isActive = index
    },
    // 分页查询
    currentchange (page) {
      this.orderQuery.page = page
      this.orderList(this.orderQuery)
    },
    // 查询按钮部分——根据订单号或商品查询
    searchBtn () {
      var orderid = `${this.$refs.searchOrder.value}`
      if (orderid.length === 0) {
        this.$message.warning('请输入要搜索的商品或订单号')
        return
      }
      detailOrder(this.$refs.searchOrder.value).then(res => {
        console.log(res)
        if (res.status === 200) {
          this.getObjData(res)
          this.orderQuery = {
            'limit': 10,
            'page': 1,
            'type': -1
          }
          this.isActive = 0
        }
      })
    },
    // 解析多个订单数据
    getItemData (res) {
      if (!res) { return }
      /**
       * 涉及字段
       * 订单唯一id res.data[0].unique       String
       * 订单状态 res.data[0].status 订单状态（-1 : 申请退款 -2 : 退货成功 0：待发货；1：待收货；2：已收货；3：待评价；-1：已退款）  Number
       * 订单日期 res.data[0].createTime    String
       * 订单id res.data[0].id              Number
       * 订单号 res.data[0].orderId         String
       * 支付状态 res.data[0].paid          Number
       * 支付时间 res.data[0].payTime       String
       * 实付金额 res.data[0].payPrice      Number
       * 是否拼团 res.data[0].pinkId        Nubmer 拼团id 0没有拼团
       * 商品信息：res.data[0].cartInfo
       * 商品数量 res.data[0].cartInfo.cartNum        Number
       * 拼团id res.data[0].cartInfo.combinationId    Number
       */
      // 未支付：10 拼团：11 拼团失败: 12
      // -1 : 申请退款
      // -2 : 退货成功
      // 0：待发货
      // 1：待收货
      // 2：已收货
      // 3：待评价
      // 4: 已评价
      // -1：已退款
      console.log(res.data)
      this.itemList = res.data && res.data.map(item => {
        if (item.paid === 0) { // 已支付为1
          // 未支付
          item.status = 10
        }
        if (item.paid === 1 && item.pinkId !== 0) {
          // 拼团订单
          item.status = 11
        }
        item.cartInfo.forEach(i => {
          if (i.isReply !== 0) {
            item.status = 4
          }
        })
        return {
          'status': item.status,
          'unique': item.unique,
          'id': item.id,
          'createTime': item.createTime,
          'orderId': item.orderId,
          'paid': item.paid,
          'payTime': item.payTime,
          'payPrice': item.payPrice,
          'pinkId': item.pinkId,
          'combinationId': item.combinationId,
          'cartInfo': item.cartInfo.map(i => {
            return {
              'id': i.id,
              'productInfoId': i.productInfo.id,
              'attrid': i.productInfo.attrInfo.id,
              'isReply': i.isReply,
              'combinationId': i.combinationId,
              'productId': i.productId,
              'attrProductId': i.productInfo.attrInfo.productId,
              'cartNum': i.cartNum,
              'truePrice': i.truePrice,
              'storeName': i.productInfo.storeName,
              'img': i.productInfo.attrInfo.image,
              'sku': i.productInfo.attrInfo.sku
            }
          })
        }
      })
    },
    // 解析单个订单数据
    getObjData (res) {
      if (!res) { return }
      if (res.data.paid === 0) { // 已支付为1
        // 未支付
        res.data.status = 10
      }
      if (res.data.paid === 1 && res.data.pinkId !== 0) {
        // 拼团订单
        res.data.status = 11
      }
      this.itemList = [{
          'status': res.data.status,
          'unique': res.data.unique,
          'id': res.data.id,
          'createTime': res.data.createTime,
          'orderId': res.data.orderId,
          'paid': res.data.paid,
          'payTime': res.data.payTime,
          'payPrice': res.data.payPrice,
          'pinkId': res.data.pinkId,
          'combinationId': res.data.combinationId,
          'cartInfo': res.data.cartInfo.map(i => {
            return {
              'cartid': i.id,
              'productInfoId': i.productInfo.id,
              'attrid': i.productInfo.attrInfo.id,
              'combinationId': i.combinationId,
              'productId': i.productId,
              'attrProductId': i.productInfo.attrInfo.productId,
              'cartNum': i.cartNum,
              'truePrice': i.truePrice,
              'storeName': i.productInfo.storeName,
              'img': i.productInfo.attrInfo.image,
              'sku': i.productInfo.attrInfo.sku
            }
          })
      }]
    },
    // 删除订单相关
    isDel (val) { this.delDialogVisible = this.isShow(val) },
    dleSure (id) {
      this.loadingBtn = true
      // id为订单unique码
      deleteOrder({
          'uni': `${id}`
      }).then(res => {
        this.$message.warning(res.msg)
        // this.delDialogVisible = false
        this.loadingBtn = false
      }).catch(err => {
        console.log(err)
        this.$message.error('订单尚未完成，无法删除！')
        this.delDialogVisible = false
        this.loadingBtn = false
      })
    },
    // 取消订单相关
    isCancel (val) { this.cancelDialogVisible = this.isShow(val) },
    canOrder (id) {
      this.loadingBtn = true
      cancelOrder(JSON.stringify({'id': `${id}`})).then(res => {
        if (res.status === 200) {
          this.$message.success('取消订单成功!')
          this.orderList(this.orderQuery)
          this.cancelDialogVisible = false
        } else {
          this.$message.error(res.msg)
        }
      })
      this.loadingBtn = false
    },
    // 签收相关
    isTake (val) { this.receDialogVisible = this.isShow(val) },
    takOrder (orderId) {
      this.loadingBtn = true
      takeOrder({'uni': orderId}).then(res => {
        if (res.status === 200 && res.success) {
          this.$message.success('收货成功')
          this.receDialogVisible = false
          this.orderList(this.orderQuery)
        } else {
          this.$message.error(res.msg)
        }
        this.loadingBtn = false
      })
    },
    isShow (val) {
      this.oid = val.id
      this.orderId = val.orderid
      return val.show
    },
    // 邀请相关
    isInvi (val) {
      getPoster({
          'from': 'h5',
          'id': `${val.id}`
        }).then(res => {
          if (res.status === 200) {
            this.qrcode = res.data.url
          } else {
            this.$message.error(res.msg)
          }
        })
      this.inviteDialogVisible = true
    },
    // 一键复制
    useClipboard (text, event) {
        var clipboard = new Clipboard(event.target, {text: () => text})
        clipboard.on('success', e => {
            this.$message.success('复制成功')
            clipboard.destroy()
        })
        clipboard.on('error', e => {
            this.$message.error('该浏览器不支持自动复制')
            clipboard.destroy()
        })
    }
  }
}
</script>

<style lang="scss">
.personalMyOrder{
  padding: 30px 30px 30px 30px;
  background-color: #FFFFFF;
  nav{
    font-size: 30px;
    cursor: default;
  }
  .MyOrderSearch{
    height: 66px;
    line-height: 65px;
    margin: 20px 0;
    display: flex;
    .orderSearchType{
      font-size: 18px;
      cursor: pointer;
      color: #666666;
      display: flex;
    }
    .orderSearchTypeTool{
      width: 1px;
      height: 18px;
      margin: 24px 20px 0 19px;
      background-color: #CDCDCD;
    }
    .search_active{
      color: #333333;
      border-bottom: 2px solid #F26E47;
    }
    .searchBtn{
      margin-left: 80px;
      padding: 0;
      input{
        width: 200px;
        height: 40px;
        line-height: 40px;
        padding-left: 20px;
        font-size: 14px;
        background-color: #F6F6F8;
      }
      button{
        width: 62px;
        height: 40px;
        color: #FFFFFF;
        background-color: #333333;
        cursor: pointer;
      }
    }
  }
  .elePag{
    margin-bottom: 20px;
    display: flex;
    justify-content: center;
    button{
      width: 50px;
      height: 50px;
      background-color: #FFFFFF !important;
      border: 1px solid #333333;
      i{
        color: #333333;
      }
    }
    li{
      width: 50px;
      height: 50px;
      line-height: 50px;
      background-color: #FFFFFF;
      border: 1px solid #333333 !important;
      margin: 0 5px;
      &:hover{
        color: #E6E6E6;
      }
    }
    .active{
      color: #FFFFFF;
      background-color: #333333;
    }
  }
  .notfound{
    display: flex;
    justify-content: center;
    img{
      width: 300px;
      height: 256px;
    }
  }
  .delDia{
    .noBottom{
     margin: 45px 0 0 0 !important;
    }
    .diaContent{
        margin: 45px 0 70px 0;
        text-align: center;
        .diaTitle{
            font-size: 24px;
            color: #333333;
        }
        .diaText{
            margin-top: 24px;
            font-size: 20px;
            color: #666666;
        }
        .qrcode{
            margin-top: 30px;
        }
        .pink{
          width: 310px;
          height: 600px;
        }
    }
    .diaBtn{
        border-top: 1px solid #000;
        padding-top: 40px;
        button{
            width: 190px;
            height: 60px;
            cursor: pointer;
            font-size: 20px;
            margin-bottom: 10px;
            color: #333333;
            background-color: #FFFFFF;
            border: 1px solid #333333;
        }
        .el-button{
          border-radius: 0;
        }
        .yes{
            float: right;
            color: #FFFFFF;
            background-color: #333333;
            border: 1px solid #333333;
        }
        .shareTxt{
          display: inline-block;
          width: 340px;
          height: 60px;
          line-height: 60px;
          padding: 0 82px;
          font-size: 20px;
          color: #000000;
        }
        .cpqr{
          display: inline-block;
          width: 100%;
          &:hover{
            color: #FFFFFF;
            background-color: #333333;
            border: 1px solid #333333;
          }
        }
    }
  }
}
</style>
