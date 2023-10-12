<template>
    <div class="payment banxin">
        <!-- 选择支付方式 -->
        <div class="paymentBox" v-if="payStatus===false">
            <div class="top fs24 mar-bot-20">请选择支付方式</div>
            <div class="amount display-between col-666">
                <p class="fs20">交易号：{{orderId}}</p>
                <p class="fs16">总额： <span class="fs20 col-F26">¥{{money}}</span></p>
            </div>
            <div class="payMethod flex">
                <div class="wechat flex" @click="select(1)" :class="active==1?'active':''">
                    <el-image src="static/images/wxicon@2x.png" ></el-image>
                    <div class="mar-left-20">
                        <p class="fs24">微信支付</p>
                        <p class="fs16">Wechat Pay</p>
                    </div>
                </div>
                 <div class="wechat flex" @click="select(2)" :class="active==2?'active':''">
                    <el-image src="static/images/yue.jpeg" ></el-image>
                    <div class="mar-left-20">
                        <p class="fs24">余额支付</p>
                        <p class="fs16">Yue Pay</p>
                    </div>
                </div>
                <!-- <div class="alipay mar-left-30 flex" @click="select(2)" :class="active==2?'active':''">
                    <img src="static/images/zfbicon@2x.png" alt="">
                    <div class="mar-left-20">
                        <p class="fs24">支付宝支付</p>
                        <p class="fs16">Alipay Pay</p>
                    </div>
                </div> -->
            </div>
            <div class="pay tar">
                <el-button :loading="false" @click="submit">立即付款</el-button>
                <p class="mar-top-10 mar-right-10">剩余付款时间：<span class="col-F26">{{time}}</span></p>
            </div>
        </div>
        <!-- 支付成功 -->
        <div class="paySuccess display-center column" v-else>
            <el-image src="static/images/zfcgicon@2x.png" ></el-image>
            <p class="fs28 mar-bot-20">恭喜您，支付成功</p>
            <p class="col-666 fs18 mar-bot-50">请耐心等待发货,保持手机畅通</p>
            <div>
                <el-button
                class="col-fff bgc-333"
                @click="$router.push({'path': '/orderDetail', 'query': {'id': orderId}})"
                >查看订单</el-button>
                <el-button class="col-333 mar-left-70" @click="$router.push('/')">继续购物</el-button>
            </div>
        </div>
        <!-- 支付二维码 -->
        <el-dialog
          title="微信支付"
          :visible.sync="payVisible"
          width="30%"
          center
          :before-close="handleClose">
          <div class="display-column display-center">
              <div class="qr mar-bot-20">
                   <VueQrcode :value="orderKey"/> 
              </div>
              <p class="mar-bot-30 fs20">应付金额：¥{{ money }}</p>
              <div class="sys flex">
                  <el-image src="static/images/sys@2x.png" ></el-image>
                  <div class="mar-left-10">
                      <p class="fs14">请使用微信扫一扫</p>
                      <p class="fs14">扫二维码支付</p>
                  </div>
              </div>
          </div>
        </el-dialog>
    </div>
</template>
<script>
import VueQrcode from 'vue-qrcode'
import moment from 'moment'

import {
    detailOrder,payOrder
} from '@/api/order.js'
export default {
    name: 'payment',
    components: {VueQrcode},
    data () {
        return {
            orderId: '',
            money: 539,
            time: '0分0秒',
            active: '',
            payVisible: false,
            payStatus: false,
            orderKey: null,
            timer: null,
            timer2: null
        }
    },
    created () {
        console.log(this.$route.query)
        this.money = this.$route.query.price
        this.orderId = this.$route.query.orderId
        //this.orderKey = this.$route.query.jsConfig.codeUrl || '暂无有效订单'
        // this.orderKey = this.$route.query.key
        this.resetTimes(this.$route.query.createTime)
    },
    methods: {
        // 选择支付方式
        select (val) { this.active = val },

        // 支付
        submit () {
            if (this.active === '') {
                return this.$message.warning('请选择支付方式')
            }
            //微信支付
            if (this.active == 1) {
                this.payVisible = true

                // 查询付款结果，
                // this.timer = setInterval(() => {
                //     detailOrder(this.orderId).then(res => {
                //         if (res.data.paid === 1) {
                //             this.payVisible = false // 隐藏支付卡
                //             this.payStatus = true // 显示付款成功
                //             clearTimeout(this.timer)
                //         }
                //     })
                // }, 5000)

                payOrder({
                    uni: this.orderId,
                    from: 'pc',
                    payType: 'weixin', // 'weixin'/'yue'     
            }).then(res => {
                
                if (res.status === 200) {
                    this.orderKey = res.data.result.jsConfig.codeUrl
                    //  detailOrder(this.orderId).then(res => {
                         
                    //     if (res.data.paid === 1) {
                    //        // this.payVisible = false // 隐藏支付卡
                    //         this.payStatus = true // 显示付款成功
                    //         //clearTimeout(this.timer)
                    //     }
                    // })
                } else {
                    console.log(res)
                    this.$message.warning(res.msg || '支付失败失败！')
                }
            
            })
            }

            //余额支付
            if (this.active == 2) {
                payOrder({
                    uni: this.orderId,
                    from: 'pc',
                    payType: 'yue', // 'weixin'/'yue'     
            }).then(res => {
                if (res.status === 200) {
                     detailOrder(this.orderId).then(res => {
                        if (res.data.paid === 1) {
                           // this.payVisible = false // 隐藏支付卡
                            this.payStatus = true // 显示付款成功
                            //clearTimeout(this.timer)
                        }
                    })
                } else {
                    console.log(res)
                    this.$message.warning(res.msg || '支付失败失败！')
                }
            
            })


            }
            
        },

        // 关闭支付二维码
        handleClose () { this.payVisible = false },
        // 倒计时
        resetTimes (createTime) {
            if (!createTime) { return }
            // 订单过期时间 = 创建时间 + 15分钟（毫秒）
            var afterBuy = new Date(createTime).getTime() + 900000
            // console.log(afterBuy)
            this.timer2 = setInterval(() => {
                this.time = this.countDownTime(afterBuy)
            }, 1000)
        },
        countDownTime (hms) {
            // hms 为毫秒时间戳
            // 倒计时终点
            var oTim = moment(hms).format('yyyy-MM-DD hh:mm:ss').split(' ')
            var odat = oTim[0]
            var otim = oTim[1].split(':')
            var [oH, oM, oS] = [otim[0], otim[1], otim[2]]

            // 现在时间
            var nTim = moment().format('yyyy-MM-DD hh:mm:ss').split(' ')
            var ndat = nTim[0]
            var nT = nTim[1].split(':')
            var [nH, nM, nS] = [nT[0], nT[1], nT[2]]

            var resetYear = odat.split('-')[0] - ndat.split('-')[0]
            var resetMon = odat.split('-')[1] - ndat.split('-')[1]
            var resetDay = odat.split('-')[2] - ndat.split('-')[2]
            if (resetDay > 0) {
                oH = parseInt(oH) + 24 * resetDay
            }

            var resetHour, resetMin, resetSec
            // 目标小时已小于当前小时
            if (`${parseInt(oH) - parseInt(nH)}` < 0 || ndat > odat || resetYear < 0 || resetMon < 0) {
                resetHour = '0'
                resetMin = '00'
                resetSec = '00'
                return '00: 00 : 00'
            }
            resetHour = parseInt(oH) - parseInt(nH) - 1
            var resM = parseInt(oM) + 60 - parseInt(nM) - 1
            var resS = parseInt(oS) + 60 - parseInt(nS)

            resetSec = resS % 60
            resM += parseInt(resS / 60) // 剩余秒数大于60 剩余分钟加1
            resetMin = resM % 60
            resetHour += parseInt(resM / 60) // 剩余分钟大于60 剩余小时加1

            if (resetHour < 10) { resetHour = `0${resetHour}` }
            if (resetMin < 10) { resetMin = `0${resetMin}` }
            if (resetSec < 10) { resetSec = `0${resetSec}` }
            return `${resetHour}:${resetMin}:${resetSec}`
        }
    },
    destroyed () {
        clearTimeout(this.timer)
        clearTimeout(this.timer2)
    }
}
</script>
<style lang="scss" scoped>
.payment{
    margin-top: 30px;
    margin-bottom: 100px;
    .paymentBox{
        .payMethod{
            padding: 50px 0;
            border-bottom: 1px solid #CCCCCC;
            .wechat,.alipay{
                width: 270px;
                padding: 20px 0;
                justify-content: center;
                border: 1px solid #CCCCCC;
                cursor: pointer;
                >>>.el-image__inner{
                    width: 60px;
                    height: 60px;
                }
                div{
                    display: flex;
                    flex-direction: column;
                    justify-content: space-around;
                }
            }
            .active{
                border-color: #F26E47;
            }
        }
        .pay{
            margin-top: 50px;
            .el-button{
                width: 200px;
                height: 60px;
                box-sizing: border-box;
                background-color: #F26E47;
                color: #FFFFFF;
                font-size: 20px;
                border-radius: 0;
            }
        }
    }
    .paySuccess{
        >>>.el-image__inner{
            width: 100px;
            height: 100px;
            margin-top: 50px;
            margin-bottom: 40px;
        }
        .el-button{
            width: 200px;
            height: 60px;
            box-sizing: border-box;
            border-radius: 0;
            font-size: 20px;
            border: 1px solid #333333;
        }
    }
    >>>.el-dialog{
        min-width: 350px;
        .el-dialog__header{
            padding-top: 50px;
            .el-dialog__title{
                font-size: 26px;
            }
        }
        .el-dialog__body{
            padding-bottom: 50px;
            .qr{
                width: 275px;
                padding: 20px;
                border: 1px solid #333333;
                box-sizing: border-box;
                img{
                    width: 235px;
                    height: 235px;
                }
            }
            .sys{
                img{
                    width: 33px;
                    height: 33px;
                }
            }
        }
    }
}
</style>
