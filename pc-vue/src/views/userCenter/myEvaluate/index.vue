<template>
    <div class="myEvaluate">
        <h2 class="fs30">评价晒单</h2>
        <el-divider></el-divider>
        <section class="evaluteSection" v-loading="loading" v-if="evaluateList.length !==0">
            <ListItem :list="this.evaluateList" />
        </section>
        <section class="noData" v-loading="loading" v-else>
            <NotFoundInfo class="noInfo" />
        </section>
        <el-pagination
            class="elePag"
            layout="prev, pager, next"
            :total="this.totalPage"
            :current-page="this.listquery.page"
            @current-change="currentchange">
        </el-pagination>
        <!-- <div class="evaluateList">
            <div class="evaluateBox">
                <div class="label">
                    <span class="mar-right-20">下单时间：2021-04-06 18:00:00</span>
                    <span>订单号：123456789</span>
                </div>
                <div class="productList">
                    <div class="product flex">
                        <img src="static/images/cp1@2x.png" alt="">
                        <div class="info flex2 mar-left-20">
                            <p>EMPORIO ARMANI刺绣LOGO连帽卫衣</p>
                            <div class="flex mar-top-10">
                                <p class="col-999 mar-right-10">颜色：<span class="col-333">灰色</span></p>
                                <p class="col-999 mar-right-10">尺码：<span class="col-333">M</span></p>
                                <p class="col-999 mar-right-10">数量：<span class="col-333">1</span></p>
                            </div>
                        </div>
                        <div class="price flex3 display-center">
                            <p>¥539</p>
                        </div>
                        <div class="operate flex2 display-column">
                            <el-button @click="toEvaluate">去评价</el-button>
                            <span class="mar-top-10 cur-poi">查看详情</span>
                        </div>
                    </div>
                </div>
            </div>
        </div> -->
    </div>
</template>
<script>
import ListItem from '@/views/userCenter/myOrder/components/orderItem.vue'
import NotFoundInfo from '@/components/NotFound/NotFoundInfo.vue'
import {getOrderList} from '@/api/order.js'
export default {
    name: 'myEvaluate',
    components: {ListItem, NotFoundInfo},
    data () {
        return {
            listquery: {
                'limit': 10,
                'page': 1,
                'type': 3
            },
            totalPage: 10,
            evaluateList: [],
            loading: false
        }
    },
    created () {
        this.getOrderList()
    },
    methods: {
        getOrderList () {
            this.loading = true
            getOrderList(this.listquery).then(res => {
                if (res.status === 200) {
                    this.getItemData(res)
                }
                this.loading = false
            })
        },
        // 解析多个订单数据
        getItemData (res) {
        if (!res.data) { return }
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
        // -1：已退款
        this.evaluateList = res.data.map(item => {
            if (item.paid === 0) { // 已支付为1
            // 未支付
            item.status = 10
            }
            if (item.paid === 1 && item.pinkId !== 0) {
            // 拼团订单
            item.status = 11
            }
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
        // 分页查询
        currentchange (page) {
            this.listquery.page = page
            this.getOrderList(this.listquery)
        }
    }
}
</script>
<style lang="scss" scoped>
.myEvaluate{
    padding: 30px;
    background-color: #fff;
    .evaluateList{
        .evaluateBox{
            border: 1px solid #E6E6E6;
            .label{
                height: 50px;
                line-height: 50px;
                padding: 0 30px;
                background-color: #F5F5F5;
            }
            .productList{
                .product{
                    padding: 25px 30px;
                    img{
                        width: 130px;
                        height: 150px;
                    }
                    .info{
                        p{
                            line-height: 1.5;
                        }
                        div{
                            flex-wrap: wrap;
                        }
                    }
                    .operate{
                        align-items: center;
                        justify-content: center;
                        .el-button{
                            width: 86px;
                            height: 40px;
                            background-color: #F26E47;
                            box-sizing: border-box;
                            color: #FFFFFF;
                            border-radius: 0;
                        }
                        span{
                            color: #F29C46;
                        }
                    }
                }
            }
        }
    }
    .noData{
        display: flex;
        justify-content: center;
    }
}
</style>
<style lang="scss">
.myEvaluate{
    .elePag,.el-pagination{
        margin-bottom: 20px;
        display: flex;
        justify-content: center;
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
        .el-pager{
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
            li,.active{
                color: #FFFFFF !important;
                background-color: #333333 !important;
            }
        }
    }
}
</style>
