<template>
    <div class="productDetail banxin">
        <!-- 导航 -->
        <div class="nav">
            <router-link to="/">
                <span>首页</span>
            </router-link>
            <span>|</span>
            <span>所有宝贝</span>
            <span>|</span>
            <span>{{
                activityType==0?'普通商品'
                :activityType==1?'秒杀商品'
                :activityType==2?'拼团商品'
                :activityType==3?'砍价商品':''
                }}
            </span>
        </div>
        <div class="product flex">
            <!-- 商品图片 -->
            <div class="left flex mar-right-70">
                <div class="thumbnail mar-right-20" v-if="storeInfo !== undefined && storeInfo.length!==0">
                    <el-image class="img cur-poi mar-bot-20"
                    :src="item"
                    v-for="(item, index) in storeInfo.sliderImageArr"
                    :key="index"
                    @click="changeImg(item)"></el-image>
                </div>
                <div class="thumbnail mar-right-20" v-else>
                    暂无信息~
                </div>
                <div class="preview">
                    <vue-photo-zoom-pro :width="250" :url="bigImg"></vue-photo-zoom-pro>
                    <!-- <img class="img" :src="imgUrl" alt=""> -->
                </div>
            </div>
            <!-- 商品信息 -->
            <div class="right flex1">
                <div class="productName display-between">
                    <p class="fs28">
                        {{
                            activityType==0?`${storeInfo.storeName}`
                            :activityType==1?`${storeInfo.title}`
                            :activityType==2?`${storeInfo.title}`
                            :activityType==3?`${storeInfo.title}`:''
                        }}
                    </p>
                    <el-image v-if="activityType!=3" class="collect cur-poi" :src="storeInfo.userCollect?starUrl:star1Url" @click="collectClick"></el-image>
                </div>
                <p class="productInfo abstract col-999 fs16 mar-top-20">
                    {{
                        activityType==0?`${storeInfo.storeInfo}`
                        :activityType==1?`${storeInfo.info}`
                        :activityType==2?`${storeInfo.info}`
                        :activityType==3?`${storeInfo.info}`:''
                    }}
                </p>
                <p v-if="activityType!=3" class="fs26 price mar-top-30">
                    ¥{{currentProduct.price}}<span class="stock col-666 fs16 mar-left-10">仅剩{{currentProduct.stock}}件</span>
                </p>
                <p v-else class="fs18 price mar-top-30">
                    可砍至
                    <span class="fs26 col-F26 mar-left-10 mar-right-20">¥{{storeInfo.bargainMinPrice}}</span>
                    <span class="fs16 col-999 line-through">¥{{storeInfo.price}}</span><span class="bargainNum col-F26 fs16 mar-left-10 tar">已有{{storeInfo.sales}}人已砍价成功</span>
                </p>
                <div class="activity display-between mar-top-30" v-if="activityType!=0">
                    <p class="title fs20">{{activityType==1?'秒杀':activityType==2?'拼团':activityType==3?'砍价':''}}</p>
                    <!--  -->
                    <p class="fs14" v-if="activityType==1">距结束 <span class="fs16">{{ storeInfo.resetTime }}</span></p>
                    <p class="fs16" v-if="activityType==2">{{ storeInfo.people }}人拼团</p>
                    <p class="fs16" v-if="activityType==3">邀请好友一起砍价</p>
                </div>
                <!-- 优惠券 -->
                <!-- <div class="promotion flex mar-top-20">
                    <div class="title fs18">优惠券：</div>
                    <div class="list flex1" :class="zoom?'lists':''">
                        <div class="listBox display-align mar-bot-10" v-for="(item,index) in couponList" :key="item.id">
                            <div class="label col-F26 mar-right-10">满{{item.useMinPrice}}减{{item.couponPrice}}</div>
                            <!-- <div class="listInfo mar-right-10">满{{item.useMinPrice}}减{{item.couponPrice}}</div> -->
                            <!-- <div v-if="!item.isUse" class="underline cur-poi" @click="receive(item)">领取</div>
                            <div v-else class="underline">已领取</div>
                            <span class="mar-left-30" v-if="index===0">
                                <el-image v-if="zoom" class="xl cur-poi" src="static/images/xl@2x.png" @click="zoomClick1"></el-image>
                                <el-image v-else class="sl cur-poi" src="static/images/sl@2x.png" @click="zoomClick1"></el-image>
                            </span>
                        </div>
                    </div>
                </div> --> 
                <el-divider></el-divider>
                <!-- 商品规格 -->
                <div class="specs" v-for="(item,index) in productAttr" :key="index">
                    <div v-if="activityType!==3" class="color" :class="item.unfold?'colors':''">
                        <p class="fs18 display-align">
                            {{item.attrName}}：<span class="mar-left-5">{{item.activeAttr?item.activeAttr:'请选择'}}</span>
                            <el-image v-if="!item.unfold"
                            class="xl mar-left-5 cur-poi"
                            src="static/images/xl@2x.png"
                            @click="foldClick(item)"></el-image>
                            <el-image v-else
                            class="sl mar-left-5 cur-poi"
                            src="static/images/sl@2x.png"
                            @click="foldClick(item)"></el-image>
                        </p>
                        <div class="specList flex mar-top-20">
                            <div
                            class="spec fs16 mar-right-20 cur-poi"
                            :class="indexs==item.activeIndex?'active':''"
                            v-for="(items, indexs) in item.attrValue"
                            :key="indexs" @click="selectAttr(item,index,indexs)">
                            {{items.attr}}
                            </div>
                        </div>
                    </div>
                    <el-divider></el-divider>
                </div>
                <div class="buy flex mar-top-40"
                v-if="activityType!=3">
                    <el-input-number v-model="num" @change="handleChange" :min="1" :max="10" label=""></el-input-number>
                    <div class="flex flex1" v-if="activityType!=2">
                        <el-button
                        class="addCart flex1 fs18"
                        :loading="addLoading"
                        @click="addCart">加入购物车</el-button>
                        <el-button
                        class="buyNow flex1 fs18"
                        :loading="addLoading"
                        @click="buyIt">立即购买</el-button>
                    </div>
                    <div class="flex flex1" v-else>
                        <el-button
                        class="addCart flex1 fs18"
                        @click="buyAlone">单独购买</el-button>
                        <el-button
                        class="buyNow flex1 fs18"
                        @click="buyTeam">我要开团</el-button>
                    </div>
                </div>
                <div class="bargain flex mar-top-40" v-else>
                    <div class="bargainBtn fs18 tac cur-poi" @mouseenter="bargain" @mouseout="bargain1">参与砍价</div>
                    &nbsp;
                    <div class="bargainBtn fs18 tac cur-poi" @mouseenter="bargain" @mouseout="bargain1">邀请好友砍价</div>
                </div>
            </div>
        </div>
        <!-- 拼单列表 -->
        <div class="groupList mar-top-70" v-if="activityType==2">
            <p class="title fs26">这些人正在拼单</p>
            <div class="groupBox display-between mar-top-30">
                <div class="group display-align mar-bot-30" v-if="groupList.length === 0">
                    暂无数据哦~
                </div>
                <div
                v-else
                class="group display-align mar-bot-30"
                v-for="(item,index) in groupList"
                :key="index">
                    <el-image :src="item.avatar"></el-image>
                    <div class="mar-left-20 flex1">
                        <p>{{item.nickname}}</p>
                        <p>还差<span class="col-F26">{{ item.count }}人</span>拼成</p>
                        <p>剩余{{ item.resetTime }}</p>
                    </div>
                    <el-button class="join">去拼团</el-button>
                </div>
            </div>
        </div>
        <!-- 砍价帮信息 -->
        <div class="groupList mar-top-70" v-if="activityType==3">
            <p class="title fs26">砍价帮</p>
            <div class="groupBox display-between mar-top-30">
                <div class="group display-align mar-bot-30" v-if="groupList.length === 0">
                    暂无数据哦~
                </div>
                <div
                v-else
                class="group display-align mar-bot-30"
                v-for="(item,index) in groupList"
                :key="index">
                    <el-image :src="item.avatar"></el-image>
                    <div class="mar-left-20 flex1">
                        <p>{{item.nickname}}</p>
                        <p>还差<span class="col-F26">{{ item.count }}人</span>拼成</p>
                        <p>剩余{{ item.resetTime }}</p>
                    </div>
                    <el-button class="join">去拼团</el-button>
                </div>
            </div>
        </div>
        <div class="info flex">
            <el-tabs v-model="activeName" @tab-click="handleClick">
                <el-tab-pane label="商品详情" name="first">
                    <!-- 商品介绍 -->
                    <NotFoundInfo class="tableNull" v-if="!storeInfo.description"/>
                    <div :style="{width: '860px'}" v-html="storeInfo.description" v-else></div>
                    <!-- <div :style="{width: '860px'}" v-html="storeInfo.description" v-if="activityType==3"></div>
                    <div v-else>
                        <NotFoundInfo class="tableNull" v-if="!poster"/>
                        <img v-else class="productInfo" :src="poster" alt="">
                    </div> -->
                </el-tab-pane>
                <el-tab-pane label="商品评价" name="second">
                    <!-- 评价 -->
                    <div class="evaluate">
                        <ProductEvaluate :productId="productId"></ProductEvaluate>
                    </div>
                </el-tab-pane>
            </el-tabs>
            <!-- 推荐商品 -->
            <div class="recommendBox">
                <Recommend :list="hotProductList" @reload="reload" />
            </div>
        </div>
    </div>
</template>

<script>
// 插件
import moment from 'moment'
// 组件
import Recommend from './recommend/index'
import ProductEvaluate from './productEvaluate/index'
import NotFoundInfo from '@/components/NotFound/NotFoundInfo.vue'
// API
import {mapMutations} from 'vuex'

import {
    getNormalProductDetail,
    getProductPoster,
    getHotProduct
} from '@/api/indexpage.js'
import {
    cartAdd,
    cartCount
} from '@/api/cart.js'
import {
    addCollection,
    deleteCollection
} from '@/api/user.js'

export default {
    name: 'productDetail',
    components: { Recommend, ProductEvaluate, NotFoundInfo },
    data () {
        return {
            productData: {},
            productAttr: [],
            productValue: {},
            storeInfo: [],
            currentProduct: {},
            currentAttr: [],
            addLoading: false,
            bigImg: '',
            hotProductList: [],
            couponList: [],
            page: 1,
            limit: 999,
            type: 1,
            starUrl: 'static/images/star.png',
            star1Url: 'static/images/star1.png',
            activeName: 'first',
            num: 1,
            activityType: this.$route.query.type,
            qrCode: false,
            productId: this.$route.query.productId,
            poster: '',
            imgList: [{
                id: 1,
                url: 'static/images/cpxt@2x.png'
            },
            {
                id: 2,
                url: 'static/images/cpxt1@2x.png'
            },
            {
                id: 3,
                url: 'static/images/cpxt2@2x.png'
            },
            {
                id: 4,
                url: 'static/images/cpxt3@2x.png'
            }],
            zoom: 0,
            groupList: [],
            timer: null
        }
    },
    watch: {
        productId: {
            handler (newVal, oldVal) {
                switch (this.$route.query.type) {
                    case '0':
                        this.getProductDetail()
                        break
                }
                this.getHotProduct()
             
            },
            deep: true
        }
    },
    mounted () {
        switch (this.$route.query.type) {
            // 0普通商品 1秒杀 2拼团商品 3砍价
            case '0':
                this.getProductDetail()
                break
            default:
                this.getProductDetail()
                break
        }
        this.getHotProduct()
        //this.getPoster()
        //this.getCoupons()
    },
    methods: {
        ...mapMutations({
          setCartNum: 'SET_CARTNUM' // 修改vuex购物车数量
        }),

        // 刷新页面
        reload (id) { this.productId = id },

        // 普通商品详情
        async getProductDetail () {
            // const dataLoading = this.openLoading()
            var res = await getNormalProductDetail(this.productId)
            if (res.status === 200) {
                this.productData = res.data
                this.storeInfo = res.data.storeInfo
                this.productAttr = res.data.productAttr
                this.productValue = res.data.productValue
                this.currentProduct = res.data.storeInfo
                this.bigImg = res.data.storeInfo.image
                this.productAttr.forEach(i => {
                    this.$set(i, 'unfold', false)
                    this.$set(i, 'activeAttr', null)
                    this.$set(i, 'activeIndex', null)
                })
                window.scrollTo(0, 0)
                // dataLoading.close()
            } else {
                this.$message.warning(res.msg)
            }
        },

     
        // 倒计时方法
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
        // 推荐商品——热门推荐
        async getHotProduct () {
            var res = await getHotProduct()
            if (res.status === 200) {
                this.hotProductList = res.data
            } else {
                this.$message.warning(res.msg)
            }
        },
       
        // 获取优惠券列表
        async getCoupons () {
            var res = await couponList({
                productId: `${this.productId}`,
                limit: `${this.limit}`,
                page: `${this.page}`
            })
            if (res.status === 200) {
                this.couponList = res.data
            } else {
                this.$message.warning(res.msg)
            }
        },
        // 领取优惠券
        async receive (item) {
            var res = await getCoupon({couponId: item.id})
            if (res.status === 200) {
                this.$message.success('领取成功')
            } else {
                this.$message.warning(res.msg)
            }
        },
        handleChange () {},
        handleClick (tab) {
            console.log(tab.index)
        },
        zoomClick1 () { this.zoom = this.zoom ? 0 : 1 },

        // 商品规格展开
        foldClick (item) { item.unfold = !item.unfold },

        // 收藏
        collectClick () {
            let params = {
                id: this.storeInfo.id,
                category: 'collect'
            }
            if (this.storeInfo.userCollect) {
                deleteCollection(params).then(res => this.collectionRes(res))
            } else {
                addCollection(params).then(res => this.collectionRes(res))
            }
        },
        collectionRes (res) {
            if (res.status === 200) {
                this.storeInfo.userCollect = !this.storeInfo.userCollect
                this.$message.success(this.storeInfo.userCollect ? '收藏成功' : '取消收藏')
            } else {
                this.$message.warning(res.msg)
            }
        },

        // 选中规格
        selectAttr (item, index, indexs) {
            item.activeIndex = indexs
            item.activeAttr = item.attrValue[indexs].attr
            this.currentAttr[index] = item.attrValue[indexs].attr
            if (this.productValue[this.currentAttr] !== undefined) {
                this.currentProduct = this.productValue[this.currentAttr]
                this.bigImg = this.currentProduct.image
                console.log(this.currentProduct)
            }
        },

        // 切换图片
        changeImg (item) { this.bigImg = item },

        // 加入购物车
        addCart () {
            if (this.currentProduct.unique === undefined) {
                return this.$message.warning('请先选择商品规格')
            }
            this.addLoading = true
            let params = {
                    productId: this.storeInfo.id,
                    cartNum: this.num,
                    uniqueId: this.currentProduct.unique
                }
            // 秒杀商品处理
            if (this.$route.query.type === '1') {
                params = {
                    productId: this.currentProduct.productId,
                    secKillId: this.storeInfo.id,
                    cartNum: this.num,
                    uniqueId: this.currentProduct.unique
                }
            }
            cartAdd(params).then(res => {
                if (res.status === 200) {
                    this.getCartNum()
                    this.$message.success('加入购物车成功')
                } else {
                    this.$message.warning(res.msg)
                }
            })
            this.addLoading = false
        },
        // 立即购买
        buyIt () {
            if (this.currentProduct.unique === undefined) {
                return this.$message.warning('请先选择商品规格')
            }
            this.addLoading = true
            let params = {
                isNew: 1,
                productId: this.storeInfo.id,
                cartNum: this.num,
                uniqueId: this.currentProduct.unique
            }
            // 秒杀商品处理
            if (this.$route.query.type === '1') {
                params = {
                    productId: this.currentProduct.productId,
                    secKillId: this.storeInfo.id,
                    cartNum: this.num,
                    uniqueId: this.currentProduct.unique
                }
            }
            cartAdd(params).then(res => {
                if (res.status === 200) {
                    // 添加到购物车，拿到购物id 跳转到结算页面并传递id
                    this.$router.push({
                        path: '/cart/settlement',
                        query: {
                            cartId: res.data.cartId
                        }
                    })
                } else {
                    this.$message.warning(res.msg)
                }
            })
            this.addLoading = false
        },
        // 单独购买
        buyAlone () {
            // 跳转到商品详情（普通商品）
            // console.log(this.$route.query.normalId)
            this.$router.push({
                path: '/productDetail',
                query: {
                productId: this.$route.query.normalId,
                type: 0
                }
            })
            location.reload()
        },
        // 发起拼团
        buyTeam () {
            if (!this.currentProduct.sku) {
                this.$message.warning('请选择规格')
                return
            }
            if (this.num > 1) {
                this.$message.warning('拼团商品限购1件')
                return
            }
            // 立即购买
            cartAdd({
                productId: this.currentProduct.productId,
                combinationId: this.storeInfo.id,
                cartNum: this.num,
                uniqueId: this.currentProduct.unique
            }).then(res => {
                if (res.status === 200) {
                    // 添加到购物车，拿到购物id 跳转到结算页面并传递id
                    this.$router.push({
                        path: '/cart/settlement',
                        query: {
                            cartId: res.data.cartId
                        }
                    })
                } else {
                    this.$message.warning(res.msg)
                }
            })
        },
        // 获取购物车商品数量
        async getCartNum () {
            var res = await cartCount()
            if (res.status === 200) {
                this.setCartNum(res.data.count)
            } else {
                this.$message.warning(res.msg)
            }
        },

        // 邀请砍价
        bargain () { this.qrCode = true },
        bargain1 () { this.qrCode = false }
    },
    destroyed () {
        clearTimeout(this.timer)
    }
}
</script>
<style lang="scss" scoped>
@font-face {
    font-family: 'Mfbanhei';
    src: url('../../assets/font/MFbanhei.otf');
}
.productDetail{
    .nav{
        height: 56px;
        line-height: 56px;
    }
    .product{
        .left{
            .thumbnail{
                width: 130px;
                height: 580px;
                overflow: hidden;
                overflow-y: auto;
                img{
                    width: 130px;
                    height: 130px;
                    display: block;
                }
                .img:last-child{
                    margin-bottom: 0;
                }
            }
            .preview{
                width: 500px;
                >>>.magnifier-box{
                    img{
                        width: 500px !important;
                        display: block;
                    }
                }
            }
        }
        .right{
            .productName{
                align-items: center;
                .collect{
                    width: 25px;
                    height: 24px;
                }
            }
            .productInfo{
                overflow: hidden;
                white-space:nowrap;
                text-overflow: ellipsis;
            }
            .price{
                position: relative;
                .bargainNum{
                    position: absolute;
                    right: 0;
                    bottom: 2px;
                }
            }
            .activity{
                height: 40px;
                line-height: 40px;
                background-color: #F26e47;
                padding: 0 20px;
                color: #FFFFFF;
                .title{
                    font-family: 'Mfbanhei';
                }
            }
            .promotion{
                .title{
                    line-height: 28px;
                }
                .list{
                    height: 28px;
                    overflow: hidden;
                    .listBox{
                        height: 28px;
                        .label{
                            border: 1px solid #F26E47;
                            padding: 4px 12px;
                        }
                        .xl,.sl{
                            width: 19px;
                            height: 10px;
                        }
                    }
                }
                .lists{
                    height: 190px;
                    overflow-y: auto;
                }
            }
            .color{
                height: auto;
                overflow: hidden;
                .xl,.sl{
                    width: 19px;
                    height: 10px;
                }
                .specList{
                    .spec{
                        padding: 10px;
                        border: 1px solid #333333;
                    }
                    .active{
                        color: #FFFFFF;
                        background-color: #333333;
                    }
                }
            }
            .colors{
                height: 20px;
            }
            .buy{
                >>>.el-input-number{
                    margin-right: 20px;
                    width: 140px;
                    height: 50px;
                    border-color: #CCCCCC;
                    .el-input-number__decrease,.el-input-number__increase{
                        width: 30px;
                        box-sizing: border-box;
                        border-radius: 0;
                        height: 48px;
                        line-height: 48px;
                        .el-icon-minus,.el-icon-plus{
                            color: #666666;
                        }
                    }
                    .el-input__inner{
                        border-radius: 0;
                        font-size: 18px;
                        height: 50px;
                        line-height: 50px;
                        border-color: #CCCCCC;
                    }
                }
                .el-button{
                    margin-left: 0;
                    height: 50px;
                    border-radius: 0;
                    box-sizing: border-box;
                }
                .addCart{
                    border: 1px solid #333333;
                    color: #333333;
                }
                .addCart:hover{
                    background-color: #FFFFFF;
                }
                .buyNow{
                    background-color: #F26E47;
                    color: #FFFFFF;
                }
            }
            .bargain{
                .bargainBtn{
                    width: 140px;
                    height: 50px;
                    line-height: 50px;
                    background-color: #F26E47;
                    border-radius: 0;
                    color: #FFFFFF;
                }
                .qrCode{
                    width: 170px;
                    height: 170px;
                    box-shadow: 0px 0px 29px 1px #bbbbbb;
                }
            }
        }
    }
    .groupList{
        .groupBox{
            border: 1px solid #F26E47;
            padding: 30px 30px 0 30px;
            box-sizing: border-box;
            height: 138px;
            overflow-y: auto;
            flex-wrap: wrap;
            .group{
                width: 30%;
                img{
                    width: 76px;
                    height: 76px;
                }
                p{
                    line-height: 1.5;
                }
                .join{
                    width: 100px;
                    height: 40px;
                    border: 1px solid #F26E47;
                    color: #F26E47;
                    border-radius: 0;
                    background-color: #FFFFFF;
                }
            }
        }
    }
    .info{
        margin-top: 70px;
        margin-bottom: 100px;
        .tableNull{
            width: 860px;
        }
        >>>.el-tabs{
            .el-tabs__item{
                font-size: 24px;
                color: #333333;
                font-weight: normal;
            }
            .is-active{
                font-weight: 600;
            }
            .el-tabs__active-bar{
                background-color: #333333;
            }
            .el-tabs__nav-wrap::after{
                display: none;
            }
            .productInfo{
                border: 1px solid #E6E6E6;
                width: 860px;
                box-sizing: border-box;
            }
            .evaluate{
                width: 860px;
            }
        }
        .recommendBox{
            margin-top: 54px;
            width: 310px;
            border-top: 1px solid #E6E6E6;
        }
    }
}

</style>
