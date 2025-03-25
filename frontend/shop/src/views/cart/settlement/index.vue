<template>
    <div class="settlement banxin" v-loading="loading">
        <div class="top fs24">结算</div>
        <div class="content flex">
            <div class="left mar-right-20">
                <div class="address mar-bot-10">
                    <!-- 收货方式 -->
                    <div class="receivingMethod flex">
                        <div class="label flex1 display-between col-333 fs20" :class="activeMethod===1?'active':''" @click="selectMethod(1)">
                            <span>配送到家</span>
                            <!-- <div class="add display-center">
                                <img v-if="activeMethod===1" class="cur-poi" src="../../../../static/images/tjdz@2x.png" alt="" @click="addAddress">
                                <img v-else class="cur-poi" src="../../../../static/images/tjdz1@2x.png" alt="" @click="addAddress">
                                <span class="mar-left-10 fs18">新增收货地址</span>
                            </div> -->
                        </div>
                        <!-- <div class="label flex1 display-between col-333 fs20" :class="activeMethod===2?'active':''" @click="selectMethod(2)">
                            <span>门店自取</span>
                        </div> -->
                    </div>
                    <!-- 联系方式 -->
                    <div class="contactDetails">
                        <div class="selectAddress display-between" v-if="activeMethod===1">
                            <div class="addressInfo">
                                <div class="mar-bot-10">
                                    <span class="title col-999">收货人：</span>
                                    <span>{{defaultAddress.realName}} {{defaultAddress.phone}}</span>
                                </div>
                                <div>
                                    <span class="title col-999">地址：</span>
                                    <span>{{defaultAddress.province + defaultAddress.city + defaultAddress.district + defaultAddress.detail}}</span>
                                </div>
                            </div>
                            <div class="btn col-999">
                                <span class="cur-poi" @click="switchAddress">切换</span>
                                <span>|</span>
                                <span class="cur-poi" @click="addAddress">新增</span>
                            </div>
                        </div>
                        <div class="pickupInfo" v-if="activeMethod===2">
                            <div class="store">
                                <span class="col-999">提货门店：</span>
                                <span class="cur-poi" @click="mapOpen">
                                    <span>{{defaultStore.name}}</span>
                                    <img src="static/images/yl@2x.png" />
                                </span>
                            </div>
                            <el-divider></el-divider>
                            <div class="contactInfo flex">
                                <div class="flex1">
                                    <div class="mar-bot-10">
                                        <span class="title col-999">门店地址：</span>
                                        <span>{{defaultStore.address + defaultStore.detailedAddress}}</span>
                                    </div>
                                    <div class="mar-bot-10">
                                        <span class="title col-999">营业时间：</span>
                                        <span>{{defaultStore.dayTime}}</span>
                                    </div>
                                    <div>
                                        <span class="title col-999">联系电话：</span>
                                        <span>{{defaultStore.phone}}</span>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <!-- 商品列表 -->
                <div class="productBox">
                    <div class="label col-fff">
                        <span class="fs20">商品清单</span>
                        <span class="fs16">共 {{ amount }} 件</span>
                    </div>
                    <div class="productList">
                        <div class="product flex fs16" v-for="(item) in settlementData.cartInfo" :key="item.id">
                            <el-image class="mar-right-20" :src="item.productInfo.attrInfo.image" ></el-image>
                            <div class="productInfo display-column mar-right-20 flex2">
                                <p class="productName mar-bot-20">{{item.productInfo.storeName}}</p>
                                    <div class="size">
                                        <p>{{item.productInfo.attrInfo.sku}}</p>
                                    </div>
                            </div>
                            <div class="num display-center flex1">
                                <p class="fs18">x{{item.cartNum}}</p>
                            </div>
                            <div class="price flex1 display-center">
                                <p class="currentPrice mar-bot-10 fs20">¥{{item.productInfo.attrInfo.price}}</p>
                                <p class="originalPrice line-through col-999">¥{{item.productInfo.attrInfo.otPrice}}</p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="right">
                <!-- 优惠券 -->
                <div class="couponBox mar-bot-20">
                    <div class="label fs20">
                        优惠券<span v-if="couponList.length>0">(可用{{couponList.length}}张)</span><span v-else>(暂无可用)</span>
                    </div>
                    <div class="couponList" v-if="couponList.length>0">
                        <div class="coupon flex cur-poi"
                        v-for="(item) in couponList"
                        :key="item.id"
                        @click="selectCoupon(item)">
                            <div class="flex2 tac">
                                <span class="fs20">¥</span>
                                <span class="fs50 bold">{{item.couponPrice}}</span>
                            </div>
                            <div class="flex2 display-column endTime">
                                <p class="fs18">满{{item.useMinPrice}}才可使用</p>
                                <p class="fs16 col-666">有效期至:{{item.endTime}}</p>
                            </div>
                            <div class="select flex1 display-center">
                                <el-image v-if="item.id===curCoupon.id" src="static/images/xr@2x.png" ></el-image>
                                <el-image v-else src="static/images/mr@2x.png" ></el-image>
                            </div>
                        </div>
                        <p class="fs18 col-999">共使用{{curCoupon.id!==null?1:0}}张优惠券，合计优惠金额{{curCoupon.money || 0}}元</p>
                    </div>
                </div>
                <!-- 价格明细  提交订单 -->
                <div class="sett">
                    <div class="label fs20">商品合计</div>
                    <div class="total">
                        <div class="totalPrice display-between">
                            <p>商品总价</p>
                            <p>¥{{orderPrice.totalPrice}}</p>
                        </div>
                        <div class="discountAmount display-between">
                            <p>优惠</p>
                            <p>-¥{{orderPrice.deductionPrice}}</p>
                        </div>
                        <div class="fare display-between">
                            <p>运费</p>
                            <p>¥{{orderPrice.payPostage}}</p>
                        </div>
                        <div class="submitOrder flex mar-top-30">
                            <div class="sum flex1">
                                <p class="fs16">总计</p>
                                <p class="fs26">¥{{orderPrice.webPrice}}</p>
                            </div>
                            <div class="btn flex1">
                                <el-button :loading="btnLoading" @click.once="submit">提交订单</el-button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <!-- 添加/编辑收货地址 -->
        <AddAddress
            :addVisible="addVisible"
            :options="options"
            :newAddress="newAddress"
            :value="value"
            @addClose="addClose" />
        <!-- 更换收货地址 -->
        <el-dialog
          title="请选择您的收货地址"
          :visible.sync="switchVisible"
          width="45%"
          center
          :before-close="switchClose">
          <div class="addressList">
              <div class="addressBox flex" :class="addressIndex===index?'active':''" v-for="(item,index) in addressList" :key="item.id" @click="selectAddress(index)">
                  <div class="flex3 col-333">
                      <p><span class="col-999" style="width:70px;display: inline-block;">收货人:</span>{{item.realName}} {{item.phone}}</p>
                      <p><span class="col-999" style="width:70px;display: inline-block;">地址:</span>{{item.province + item.city + item.district + item.detail}}</p>
                  </div>
                  <div class="flex1 tar col-999">
                      <span class="edit cur-poi" @click="edit(item)">编辑</span>
                      <br>
                      <span v-if="item.isDefault" class="col-333 defaultAddress">默认地址</span>
                      <span v-else class="cur-poi setDefault" @click.stop="setDefault(item.id)">设为默认地址</span>
                  </div>
              </div>
          </div>
          <span slot="footer" class="dialog-footer">
            <el-button class="close" plain @click="switchClose">取消</el-button>
            <el-button class="submit" plain @click="switchSave">保存</el-button>
          </span>
        </el-dialog>
        <!-- 选择门店 -->
        <el-dialog
          class="map-dialog"
          title="选择门店"
          :visible.sync="mapVisible"
          width="45%"
          center
          :before-close="mapClose">
          <div>
              <div class="mapBox mar-top-30">
                  <TencentMap :storeList="storeList" :storeIndex="storeIndex"></TencentMap>
              </div>
          </div>
          <span slot="footer" class="dialog-footer">
            <el-button class="close" plain @click="mapClose">取消</el-button>
            <el-button class="submit" plain @click="mapSave">确认</el-button>
          </span>
        </el-dialog>
    </div>
</template>
<script>
import TencentMap from '@/components/common/TencentMap'
import AddAddress from '@/components/common/addAddress'

import {
    confirmOrder,
    getStoreList,
    computedOrder,
    createOrder
} from '@/api/order.js'
import {
    getCityData,
    getAddressData,
    setDefaultAddress
} from '@/api/user.js'
export default {
    name: 'settlement',
    components: {TencentMap, AddAddress},
    data () {
        return {
            loading: false,
            activeMethod: 1,
            settlementData: {},
            storeList: [],
            amount: 0,
            defaultAddress: {'realName': ''},
            defaultStore: {},
            storeIndex: {
                index: null
            },
            contact: {
                receiver: '李八八',
                phoneNumber: 13456789101
            },
            addressList: [],
            newAddress: {
                real_name: null,
                phone: null,
                address: {
                    city: null,
                    cityId: null,
                    district: null,
                    province: null
                },
                id: null,
                detail: null,
                is_default: false
            },
            addressIndex: null,
            options: [],
            value: {
                label: []
            },
            couponList: [],
            curCoupon: {
                id: null,
                money: null
            },
            orderPrice: {},
            btnLoading: false,
            addVisible: false,
            switchVisible: false,
            mapVisible: false
        }
    },
    mounted () {
        this.loading = true
        // Promise.all([this.getAddressList(), this.getOrderCoupons(), this.getCityList()])
        // .then(this.getSettlement())
        Promise.all([this.getAddressList(), this.getCityList()])
        .then(this.getSettlement())
    },
    methods: {
        // 获取结算详情
        getSettlement () {
            confirmOrder({
                cartId: this.$route.query.cartId
            }).then(res => {
                if (res.status === 200) {
                    this.settlementData = res.data
                    this.settlementData.cartInfo.forEach(item => {
                        this.amount += parseInt(item.cartNum)
                    })
                    this.defaultAddress = res.data.addressInfo
                    this.defaultStore = res.data.systemStore
                    if (this.defaultAddress === null) {
                        this.defaultAddress = this.addressList[0]
                    }
                   // this.getStoreList()
                    this.computedAmount()
                } else {
                    this.$message.warning(res.msg)
                }
            })
        },
        // 计算订单金额
        computedAmount () {
            computedOrder(this.settlementData.orderKey, {
                    addressId: this.defaultAddress.id,
                    bargainId: 0,
                    combinationId: 0,
                    couponId: 0,
                    payType: 1,
                    pinkId: 0,
                    shipping_type: this.activeMethod,
                    useIntegral: 0
            }).then(res => {
                if (res.status === 200) {
                    this.orderPrice = res.data.result
                    this.orderPrice.webPrice = this.orderPrice.payPrice
                    this.loading = false
                } else {
                    this.$message.warning(res.msg)
                }
            })
        },
        // 订单可使用优惠券
        async getOrderCoupons () {
            var res = await couponOrder(this.$route.query.cartId)
            if (res.status === 200) {
                this.couponList = res.data
            } else {
                this.$message.warning(res.msg)
            }
        },
        // 获取地址列表
        async getAddressList () {
            var res = await getAddressData()
            if (res.data.length === 0) {
                this.$message.warning('您当前没有任何地址，请添加一个地址！')
                this.$router.push({path: '/address'})
                // this.addVisible = true
            }
            if (res.status === 200) {
                this.addressList = res.data
                this.addressList.forEach(i => {
                    this.$set(i, 'selected', false)
                })
            } else {
                this.$message.warning(res.msg)
            }
        },
        // 获取门店列表
        async getStoreList () {
            var res = await getStoreList()
            if (res.status === 200) {
                this.storeList = res.data.list
                for (var i in this.storeList) {
                    if (this.storeList[i].id === this.defaultStore.id) {
                        this.storeIndex.index = Number(i)
                    }
                }
            } else {
                this.$message.error(res.msg)
            }
        },
        // 获取城市列表
        async getCityList () {
            var res = await getCityData()
            if (res.status === 200) {
                const json1 = JSON.parse(JSON.stringify(res.data).replace(/v/g, 'value'))
                const json2 = JSON.parse(JSON.stringify(json1).replace(/n/g, 'label'))
                const json3 = JSON.parse(JSON.stringify(json2).replace(/c/g, 'children'))
                json3.forEach(i => {
                    i.children.forEach(j => {
                        j.children.forEach(k => {
                            delete k.children
                        })
                    })
                })
                this.options = json3
                // console.log(this.options)
            } else {
                this.$message.warning(res.msg)
            }
        },
        // 设为默认地址
        async setDefault (id) {
            var res = await setDefaultAddress({id})
            if (res.status === 200) {
                this.getAddressList()
            } else {
                this.$message.warning(res.msg)
            }
        },
        // 编辑
        edit (item) {
            this.value.label = []
            this.options.forEach(i => {
                if (i.label === item.province) {
                    this.newAddress.address.province = i.label
                    this.value.label.push(i.value)
                    i.children.forEach(j => {
                        if (j.label === item.city) {
                            this.newAddress.address.city = j.label
                            this.newAddress.address.cityId = j.value
                            this.value.label.push(j.value)
                            j.children.forEach(k => {
                                if (k.label === item.district) {
                                    this.newAddress.address.district = k.label
                                    this.value.label.push(k.value)
                                }
                            })
                        }
                    })
                }
            })
            this.newAddress.real_name = item.realName
            this.newAddress.phone = item.phone
            this.newAddress.detail = item.detail
            this.newAddress.id = item.id
            this.newAddress.is_default = !(!item.isDefault)
            this.addVisible = true
        },
        // 重置数据
        reset () {
            this.newAddress = {
                real_name: null,
                phone: null,
                address: {
                    city: null,
                    cityId: null,
                    district: null,
                    province: null
                },
                detail: null,
                is_default: false
            }
            this.value.label = []
        },
        // 切换地址
        selectAddress (index) {
            this.addressIndex = index
        },
        switchSave () {
            this.defaultAddress = this.addressList[this.addressIndex]
            this.computedAmount()
            this.switchVisible = false
            location.reload()
        },
        // 选择收货方式
        selectMethod (val) {
            this.activeMethod = val
            this.computedAmount()
        },

        // 打开选择门店窗口
        mapOpen () { this.mapVisible = true },

        mapSave () {
            this.defaultStore = this.storeList[this.storeIndex.index]
            this.mapClose()
        },
        // 关闭选择门店窗口
        mapClose () {
            this.mapVisible = false
            this.storeIndex.index = null
        },

        // 打开添加收货地址弹窗
        addAddress () { this.addVisible = true },

        // 打开更换收货地址弹窗
        switchAddress () { this.switchVisible = true },

        // 关闭添加收货地址弹窗
        addClose (x) {
            if (x) {
                this.getAddressList()
            }
            this.addVisible = false
            this.reset()
        },

        // 关闭更换收货地址弹窗
        switchClose () {
            this.switchVisible = false
        },

        // 是否默认地址
        selectDefault () { this.newAddress.default = !this.newAddress.default },

        // 选择优惠券
        selectCoupon (item) {
            if (this.curCoupon.id !== item.id) {
                this.curCoupon.id = item.id
                this.curCoupon.money = item.couponPrice
                this.orderPrice.webPrice = this.orderPrice.payPrice - item.couponPrice
            } else {
                this.curCoupon = {
                    id: null,
                    money: null
                }
            }
        },

        // 提交订单
        submit () {
            this.btnLoading = true
            createOrder(this.settlementData.orderKey, {
                addressId: this.defaultAddress.id,
                //bargainId: 0,
                //combinationId: 0,
                //couponId: this.curCoupon.id,
                from: 'pc',
                //isChannel: '',
                mark: '',
                payType: 'weixin', // 'weixin'/'yue'
                phone: '',
               // pinkId: 0,
                realName: '',
               // seckillId: 0,
                shippingType: this.activeMethod,
                //storeId: this.defaultStore.id,
                useIntegral: 0
            }).then(res => {
                if (res.status === 200) {
                    res.data.result.price = this.orderPrice.payPrice
                    res.data.result.createTime = res.data.createTime
                    this.$router.push({
                        path: '/cart/payment',
                        query: res.data.result
                    })
                } else {
                    console.log(res)
                    this.$message.warning(res.msg || '提交订单失败！')
                }
                this.btnLoading = false
            })
        }
    }
}
</script>
<style lang="scss" scoped>
.settlement{
    margin-top: 30px;
    margin-bottom: 80px;
    .top{
        margin-bottom: 24px;
    }
    .content{
        .left{
            width: 705px;
            .address{
                .label{
                    background-color: #E6E6E6;
                    padding: 0 20px;
                    height: 57px;
                    line-height: 57px;
                    box-sizing: border-box;
                    .add{
                        img{
                            width: 22px;
                            height: 22px;
                        }
                    }
                }
                .active{
                    background-color: #333333;
                    color: #FFFFFF;
                }
                .selectAddress{
                    border: 1px solid #CCCCCC;
                    padding: 30px;
                    .addressInfo{
                        .title{
                            display: inline-block;
                            width: 70px;
                        }
                    }
                }
                .pickupInfo{
                    border: 1px solid #CCCCCC;
                    padding: 30px;
                    .store{
                        img{
                            width: 8px;
                            height: 13px;
                        }
                    }
                    .contactInfo{
                        .title{
                            display: inline-block;
                            width: 80px;
                        }
                    }
                }
            }
            .productBox{
                .label{
                    height: 57px;
                    line-height: 57px;
                    background-color: #333333;
                    padding: 0 20px;
                }
                .productList{
                    border: 1px solid #CCCCCC;
                    padding: 30px;
                    padding-bottom: 0;
                    .product{
                        margin-bottom: 30px;
                        >>>.el-image__inner{
                            width: 130px;
                            height: 150px;
                        }
                        // img{
                        //     width: 130px;
                        //     height: 150px;
                        // }
                        .price{
                            flex-direction: column;
                        }
                        .productInfo{
                            padding: 10px 0;
                            justify-content: space-between;
                        }
                    }
                }
            }
        }
        .right{
            width: 445px;
            .couponBox{
                width: 100%;
                .label{
                    background-color: #E6E6E6;
                    height: 57px;
                    line-height: 57px;
                    padding: 0 20px;
                    border: 1px solid #CCCCCC;
                    box-sizing: border-box;
                    border-bottom: 0;
                }
                .couponList{
                    border: 1px solid #CCCCCC;
                    box-sizing: border-box;
                    padding: 20px;
                    .coupon{
                        width: 100%;
                        height: 85px;
                        line-height: 85px;
                        background-image: url('../../../../static/images/yhq1@2x.png');
                        background-size: 100% 100%;
                        margin-bottom: 20px;
                        .endTime{
                            line-height: 1.5;
                            justify-content: center;
                        }
                        .select{
                            img{
                                width: 26px;
                                height: 26px;
                            }
                        }
                    }
                }
            }
            .sett{
                .label{
                    height: 57px;
                    line-height: 57px;
                    background-color: #E6E6E6;
                    border: 1px solid #CCCCCC;
                    box-sizing: border-box;
                    padding: 0 20px;
                    border-bottom: 0;
                }
                .total{
                    border: 1px solid #CCCCCC;
                    box-sizing: border-box;
                    padding: 30px;
                    .totalPrice,.discountAmount,.fare{
                        border-bottom: 1px solid #CCCCCC;
                        line-height: 57px;
                    }
                    .submitOrder{
                        .sum{
                            display: flex;
                            flex-direction: column;
                            justify-content: space-around;
                        }
                        .btn{
                            height: 60px;
                            line-height: 60px;
                            background-color: #F26E47;
                            .el-button{
                                width: 100%;
                                height: 100%;
                                background-color: #F26E47;
                                color: #FFFFFF;
                                font-size: 20px;
                                border-radius: 0;
                            }
                        }
                    }
                }
            }
        }
    }
    >>>.el-dialog{
        .el-dialog__header{
            padding-top: 40px;
            .el-dialog__title{
                font-size: 26px;
            }
        }
        .el-dialog__body{
            padding: 25px 60px;
            box-sizing: border-box;
            p{
                font-size: 18px;
            }
            .default{
                img{
                    width: 18px;
                    height: 18px;
                    cursor: pointer;
                }
                span{
                    font-size: 18px;
                }
            }
            .addressList{
                .addressBox{
                    padding: 20px;
                    border: 1px solid #CCCCCC;
                    margin-bottom: 20px;
                    font-size: 18px;
                    p{
                        line-height: 1.7;
                    }
                    .edit,.defaultAddress,.setDefault{
                        line-height: 1.7;
                    }
                }
                .active{
                    border-color: #F26E47;
                    background-color: #FFFBFA;
                }
            }
        }
        .el-dialog__footer{
            padding-bottom: 60px;
            .el-button{
                font-size: 18px;
                padding: 20px 75px;
                border-radius: 0;
                border: 1px solid #333333;
            }
            .close{
                color: #333333;
            }
            .submit{
                background-color: #333333;
                color: #FFFFFF;
            }
        }
    }
    >>>.el-input__inner{
        border-radius: 0;
        border-color: #333333;
    }
    >>>.distpicker-address-wrapper{
        select{
            width: 32.7%;
            border-radius: 0;
            border-color: #333333;
        }
    }
    >>>.map-dialog{
        min-width: 1640px;
        .el-dialog__body{
            padding: 0;
            .amap-wrapper {
                width: 100%;
                height: 387px;
            }
            .storeList{
                padding: 60px 60px 30px 60px;
                .store{
                    height: 135px;
                    border: 1px solid #E6E6E6;
                    padding: 20px;
                    .label{
                        font-size: 22px;
                    }
                    .storeAddress{
                        img{
                            width: 15px;
                            height: 15px;
                        }
                    }
                    .businessHours{
                        img{
                            width: 15px;
                            height: 15px;
                        }
                    }
                }
                .active{
                    border-color: #F26E47;
                    background-color: #FFFBFA;
                }
            }
        }
    }
}
</style>
