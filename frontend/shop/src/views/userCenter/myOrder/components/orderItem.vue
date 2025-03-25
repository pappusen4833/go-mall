/**
    订单内容组件
*/
<template>
  <div class="orderListItem">
        <ul class="ListItem" v-for="(item,index1) in this.list" :key="index1">
            <div class="ordertitle" >
                <div class="title">
                    <div v-if="item.status === 11">
                        <div class="container"></div>
                        <svg t="1618397979614" class="icon" viewBox="0 0 1804 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1604" width="40" height="40"><path d="M24.380952 24.380952h1755.428572v975.238096H24.380952z" fill="#F26E47" opacity=".1" p-id="1605"></path><path d="M1804.190476 1024H0V0h1804.190476v1024zM48.761905 975.238095h1706.666666V48.761905H48.761905v926.47619z" fill="#F26E47" p-id="1606"></path><path d="M843.580952 326.704762l-29.257142 29.257143h73.142857v39.009524h-180.419048c4.87619 73.142857 9.752381 131.657143 24.380952 170.666666 4.87619 14.628571 9.752381 29.257143 9.752381 39.009524 29.257143-48.761905 53.638095-102.4 73.142858-165.790476l39.009523 14.628571c-24.380952 78.019048-53.638095 141.409524-92.647619 195.047619 9.752381 19.504762 14.628571 34.133333 24.380953 48.761905 19.504762 29.257143 34.133333 43.885714 48.761904 43.885714 9.752381 0 19.504762-29.257143 29.257143-92.647619l39.009524 19.504762c-14.628571 78.019048-34.133333 117.028571-58.514286 117.028572s-48.761905-19.504762-78.019047-53.638096c-9.752381-14.628571-19.504762-29.257143-29.257143-48.761904-39.009524 48.761905-87.771429 82.895238-146.285714 107.27619l-24.380953-34.133333c58.514286-24.380952 107.27619-63.390476 146.285715-117.028572-9.752381-19.504762-14.628571-43.885714-24.380953-68.266666-14.628571-48.761905-19.504762-112.152381-24.380952-185.295238H463.238095v97.523809h141.409524c0 97.52381-4.87619 160.914286-14.628571 190.171429-9.752381 24.380952-24.380952 39.009524-53.638096 39.009524h-48.761904l-14.628572-39.009524h58.514286c14.628571 0 19.504762-9.752381 24.380952-24.380953 4.87619-19.504762 4.87619-63.390476 4.876191-126.780952h-97.52381v19.504762c0 102.4-24.380952 190.171429-68.266666 253.561905l-34.133334-29.257143c39.009524-53.638095 58.514286-126.780952 58.514286-224.304762v-195.047619h243.809524V253.561905h43.885714v92.647619h107.276191c-19.504762-19.504762-39.009524-43.885714-73.142858-63.390476l29.257143-24.380953c29.257143 24.380952 53.638095 43.885714 73.142857 68.266667zM1453.104762 799.695238H1414.095238v-24.380952h-390.095238v24.380952h-39.009524V287.695238h472.990476v512h-4.87619z m-43.885714-63.390476V326.704762h-390.095238v409.6h390.095238z m-151.161905-307.2V355.961905h39.009524v73.142857h73.142857v39.009524h-73.142857v175.542857c0 34.133333-14.628571 53.638095-48.761905 53.638095h-68.266667l-9.752381-39.009524c19.504762 0 43.885714 4.87619 63.390476 4.876191 14.628571 0 24.380952-9.752381 24.380953-24.380953V472.990476c-43.885714 78.019048-107.27619 141.409524-185.295238 185.295238l-24.380953-39.009524c73.142857-39.009524 131.657143-87.771429 175.542858-151.161904h-165.790477v-39.009524h199.92381z" fill="#F26E47" p-id="1607"></path></svg></div>
                    <div class="titleTime">下单时间：{{ item.createTime | comverTime('YYYY-MM-DD HH:mm:ss')}}</div>
                    <div class="titleTime">订单号：{{ item.orderId }}</div>
                    <div class="titleTime" v-if="item.status === 10">付款：<span>{{ item.resetTime }}</span></div>
                </div>
                <div class="deleBtn" @click="deleOrderBtn(item.unique, item.orderId)"><svg t="1618218277357" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1480" width="14" height="14"><path d="M798.1 872.6c0 34.3-27.9 62.2-62.1 62.2H288.1c-34.3-0.1-62.1-27.9-62.2-62.2V212.8h572.2v659.8zM350.2 101.2c0-7.2 5.6-12.8 12.8-12.8h298.8c7.2 0 12.7 5.6 12.7 12.8v37.5H350.2v-37.5z m634.3 37.5H748.7v-37.5c0-47.8-39-86.9-86.9-86.9H363c-47.9 0.1-86.8 38.9-86.9 86.9v37.5H39.5C18.7 138.7 2 155.4 2 176.1s16.7 37.5 37.5 37.5H151v659c0 75.7 61.4 137.1 137.1 137.1h447.8c75.7 0 137.1-61.4 137.1-137.1V212.8h111.6c20.7 0 37.5-16.7 37.5-37.5s-16.8-36.6-37.6-36.6zM512 822.4c20.7 0 37.5-16.7 37.5-37.5V386.5c0-20.7-16.7-37.5-37.5-37.5-20.7 0-37.5 16.7-37.5 37.5v398.4c0 20.7 16.8 37.5 37.5 37.5m-174.5 0c20.7 0 37.5-16.7 37.5-37.5V386.5c0-20.7-16.7-37.5-37.5-37.5-20.7 0-37.5 16.7-37.5 37.5v398.4c0.8 20.7 17.6 37.5 37.5 37.5m349 0c20.7 0 37.5-16.7 37.5-37.5V386.5c0-20.7-16.7-37.5-37.5-37.5-20.7 0-37.5 16.7-37.5 37.5v398.4c0.1 20.7 16.8 37.5 37.5 37.5" p-id="1481"></path></svg>删除</div>
            </div>
            <div class="orderContent">
                <div class="commoditys">
                    <div class="orderCommodity" v-for="(i,index2) in item.cartInfo" :key="index2">
                        <div class="orderPre">
                            <div class="img">
                                <!-- <img :src="i.img" /> -->
                                <el-image :src="i.img"></el-image>
                            </div>
                            <div class="info">
                                <div>{{ i.storeName }}</div>
                                <div class="wrap">
                                    <!-- <div><span>颜色：</span>{{ i.color }}</div>
                                    <div><span>尺码：</span>{{ i.size }}</div>
                                    <div><span>数量：</span>{{ i.sum }}</div> -->
                                    <div>{{ i.sku }}</div>
                                </div>
                            </div>
                            <div v-if="item.status !== -1 && item.status !== -2" class="buyAgain" @click="toPageId('/', i.attrProductId)">再次购买</div>
                            <div class="buyAgain" v-if="item.status === -1" >售后中</div>
                            <div v-if="item.status === -2" class="buyAgain">售后完成</div>
                            <div class="itemPrice">￥{{ i.truePrice }}</div>
                        </div>
                    </div>
                </div>
                <div class="orderStatusBtnList">
                    <div class="column">
                        <!-- -
                        -1 : 申请退款
                        -2 : 退货成功
                        0：待发货
                        1：待收货
                        2：已收货
                        3：待评价
                        -1：已退款 -->
                        <button
                        class="orderBtn remind"
                        v-if="item.status === 11"
                        @click="invitBtn(item.pinkId)">邀请好友</button>
                        <!-- <button
                        class="orderBtn remind"
                        v-if="item.status === 0">提醒发货</button> -->
                        <button
                        class="orderBtn remind"
                        v-if="item.status === 1"
                        @click="takeBtn(item.id, item.orderId)">确认收货</button>
                        <!-- <button
                        class="orderBtn normal"
                        v-if="item.status === 0 ||item.status === 10"
                        @click="canOrderBtn(item.id, item.orderId)"
                        >取消订单</button> -->
                        <button
                        class="orderBtn normal"
                        v-if="item.status === 0 || item.status === 10"
                        @click="canOrderBtn(item.id, item.orderId)"
                        >取消订单</button>
                        <button
                        class="orderBtn remind"
                        v-if="item.status === 2"
                        @click="toPageId('/evaluate', item.orderId)">评价</button>
                        <button
                        class="orderBtn normal"
                        @click="rebackOrder(item.orderId)"
                        v-if="item.status === -1"
                        >撤销申请</button>
                        <button
                        @click="toPageId('/asRecordDetail', item.id, item.orderId)"
                        v-if="item.status === -1 || item.status == -2"
                        >售后详情</button>
                        <button
                        v-else
                        @click="toPageId('/orderDetail', item.orderId)"
                        >订单详情</button>
                        <!-- <button
                        @click="toPageId('/orderDetail', item.orderId)"
                        >订单详情</button> -->
                    </div>
                </div>
            </div>
        </ul>
  </div>
</template>

<script>
import moment from 'moment'
export default {
    props: {
        list: Array
    },
    data () {
        return {
            leftTime: '15:20:06',
            timer: null
        }
    },
    created () {
        this.timeLeft()
    },
    methods: {
        timeLeft () {
            this.list.map(item => {
                if (item.paid === 0) {
                    this.$set(item, 'resetTime', '00:00:00')
                }
                return item
            })
            this.timer = setInterval(() => {
                    this.list.map(item => {
                       var  mytime = this.$moment(item.createTime).format('YYYY-MM-DD HH:mm:ss')
                    //console.log(mytime)
                        item.resetTime = this.countDownTime(new Date(mytime).getTime() + 900000)
                
                    })
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
        },
        // 页面跳转 传递查询参数
        toPageId (path, id, extra) {
            this.$router.push({
                'path': `${path}`,
                'query': {
                    'id': id,
                    'ID': extra || ''
                }
            })
        },
        deleOrderBtn (id, orderid) { this.$emit('del', {'show': true, 'id': id, 'orderid': orderid}) },
        canOrderBtn (id, orderid) { this.$emit('cancel', {'show': true, 'id': id, 'orderid': orderid}) },
        invitBtn (id) { this.$emit('invit', {'show': true, 'id': id}) },
        takeBtn (id, orderid) { this.$emit('take', {'show': true, 'id': id, 'orderid': orderid}) },
        rebackOrder (id) { this.$emit('rebackAfterSeals', id) }
    },
    destroyed () {
        clearTimeout(this.timer)
    }
}
</script>

<style lang="scss">
.orderListItem{
    .ListItem{
        border: 1px solid #CCCCCC;
        margin: 20px 0;
    }
    .ordertitle{
        font-size: 16px;
        width: 890px;
        height: 50px;
        line-height: 50px;
        background-color: #E6E6E6;
        padding-left: 29px;
        margin-bottom: 24px;
        .title{
            height: 50px;
            float: left;
            display: flex;
            align-items: center;
            .container{
                display: block;
                height: 20px;
            }
            .titleTime{
                font-size: 16px;
                display: inline-block;
                color: #333333;
                margin-right: 20px;
                span{
                    color: #F26E47;
                }
            }
        }
        .deleBtn{
            float: right;
            margin-right: 30px;
            cursor: pointer;
            display: flex;
            justify-content: center;
            align-items: center;
            .icon{
                margin-right: 10px;
            }
        }
    }
    .orderContent{
        display: flex;
        .commoditys{
            flex: 8;
            display: flex;
            flex-direction: column;
            .orderCommodity{
                margin-bottom: 30px;
                padding-left: 30px;
                display: flex;
                flex-direction: column;
                .orderPre{
                    flex: 10;
                    display: flex;
                    .img{
                        width: 130px;
                        height: 150px;
                        background-color: #F5F5F5;
                        img{
                            width: 130px;
                            height: 150px;
                        }
                    }
                    .info{
                        width: 170px;
                        height: 150px;
                        margin: 0 20px;
                        padding-top: 7px;
                        .wrap{
                            margin-top: 20px;
                            flex-direction: row;
                            div{
                                display: inline-block;
                                margin: 13px 3px 0 0;
                            }
                            span{
                                color: #999999;
                            }
                        }
                    }
                    .buyAgain,.itemPrice{
                        width: 190px;
                        height: 150px;
                        text-align: center;
                        line-height: 150px;
                    }
                    .buyAgain{
                        cursor: pointer;
                        color: #F26E47;
                    }
                    .itemPrice{
                        cursor: default;
                    }
                }
            }
        }
        .orderStatusBtnList{
            margin-bottom: 30px;
            flex: 4;
            display: flex;
            justify-items: center;
            align-items: center;
            .column{
                flex-direction: column;
                padding-left: 20px;
                button{
                    color: #F29C46;
                    background-color: transparent;
                    cursor: pointer;
                    width: 86px;
                    height: 40px;
                    line-height: 40px;
                }
                .orderBtn{
                    margin: 5px 0;
                }
                .remind{
                    color: #FFFFFF;
                    background-color: #F26E47;
                }
                .normal{
                    color: #333333;
                    border: 1px solid #333333;
                }
            }
        }
    }
}
</style>
