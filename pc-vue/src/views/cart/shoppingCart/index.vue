<template>
    <div class="shoppingCart banxin">
        <div class="top">
            <span class="fs24 col-333 bold">全部商品</span>
            <span class="fs16 col-999">共{{cartNum}}件</span>
        </div>
        <div class="content flex">
            <!-- 商品列表 -->
            <div class="left mar-right-20" v-loading="loading">
                <div class="productBox">
                    <div class="label fs18 col-fff flex tac">
                        <div class="selectAll display-center flex1">
                            <img :src="all?xz1:wz1" @click="selectAll" />
                            <span class="mar-left-10">全选</span>
                        </div>
                        <div class="productInfo flex3">商品信息</div>
                        <div class="number flex1">数量</div>
                        <div class="price flex1">价格</div>
                    </div>
                    <div v-if="!noGoods">
                        <div class="productList mar-bot-20" v-if="cartList.length">
                            <div class="product flex" v-for="(item) in cartList" :key="item.id">
                                <div class="select display-center flex1">
                                    <el-image :src="item.selected?xz:wz" @click="select(item)"></el-image>
                                </div>
                                <div class="productInfo flex flex3">
                                    <el-image class="cur-poi" :src="item.productInfo.attrInfo.image" @click="toDetail(item.productId)"></el-image>
                                    <div class="mar-left-20 fs16">
                                        <p class="productName mar-bot-20">{{item.productInfo.storeName}}</p>
                                        <div class="size">
                                            <p class="col-999">{{item.productInfo.attrInfo.sku}}</p>
                                        </div>
                                    </div>
                                </div>
                                <div class="number flex1">
                                    <el-input-number v-model="item.cartNum" @change="handleChange(item.id,item.cartNum)" :min="1" :max="item.productInfo.attrInfo.stock" size="mini"></el-input-number>
                                </div>
                                <div class="price flex flex1 tac">
                                    <div>
                                        <p class="currentPrice mar-bot-10 fs18">¥{{item.productInfo.attrInfo.price}}</p>
                                        <p class="originalPrice line-through col-999 fs14">¥{{item.productInfo.attrInfo.otPrice}}</p>
                                    </div>
                                    <p class="fs18 cur-poi" @click="delCart(item.id)">删除</p>
                                </div>
                            </div>
                        </div>
                        <!-- 失效商品 -->
                        <div class="productList invalid mar-bot-20" v-if="invalidGoods.length">
                            <div class="product flex" v-for="(item) in invalidGoods" :key="item.id">
                                <div class="mc display-center fs24 col-fff">已失效</div>
                                <div class="select display-center flex1">
                                    <!-- <img src="static/images/wz@2x.png" alt=""> -->
                                </div>
                                <div class="productInfo flex flex3">
                                    <el-image class="cur-poi" :src="item.productInfo.image" @click="toDetail(item.productId)"></el-image>
                                    <div class="mar-left-20 fs16">
                                        <p class="productName mar-bot-20">{{item.productInfo.storeName}}</p>
                                        <div class="size">
                                            <!-- <p class="col-999">{{item.productInfo.sku}}</p> -->
                                        </div>
                                    </div>
                                </div>
                                <div class="number flex1">
                                    <el-input-number v-model="item.cartNum" @change="handleChange(item.id,item.cartNum)" :min="1" :max="item.productInfo.stock" size="mini"></el-input-number>
                                </div>
                                <div class="price flex flex1 tac">
                                    <div>
                                        <p class="currentPrice mar-bot-10 fs18">¥{{item.productInfo.price}}</p>
                                        <p class="originalPrice line-through col-999 fs14">¥{{item.productInfo.otPrice}}</p>
                                    </div>
                                    <p class="del fs18 cur-poi" @click="delCart(item.id)">删除</p>
                                </div>
                            </div>
                        </div>
                        <div class="batchDel flex">
                            <div class="selectAll display-center mar-left-30">
                                <el-image :src="all?xz:wz" alt="" @click="selectAll"></el-image>
                                <span class="mar-left-10">全选</span>
                            </div>
                            <div class="mar-left-30 cur-poi" @click="delCarts">批量删除</div>
                        </div>
                    </div>
                    <div class="nothing display-column display-center" v-else>
                        <el-image src="static/images/wsj@2x.png"></el-image>
                        <p class="fs18 mar-top-20">暂无数据哦～</p>
                    </div>
                </div>
            </div>
            <!-- 价格明细  结算 -->
            <div class="right">
                <div class="label fs20">商品合计</div>
                <div class="pad-30 fs18">
                    <div class="totalPrice display-between">
                        <p>商品总价</p>
                        <p>¥{{totalPrice}}</p>
                    </div>
                    <div class="fare display-between">
                        <p>运费</p>
                        <p>¥{{postage}}</p>
                    </div>
                    <div class="settlement flex mar-top-30">
                        <div class="total flex1">
                            <p class="fs16">总计</p>
                            <p class="fs26">¥{{totalPrice + postage}}</p>
                        </div>
                        <div class="btn flex1">
                            <el-button :loading="flag" @click="sett">立刻结算</el-button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import {mapGetters, mapMutations} from 'vuex'
import {
    cartCount,
    cartList,
    cartUpdate,
    cartDel
} from '@/api/cart.js'
export default {
    name: 'shoppingCart',
    data () {
        return {
            all: false,
            check: false,
            num: 1,
            flag: false,
            cartList: [],
            invalidGoods: [],
            totalPrice: 0,
            postage: 0,
            loading: false,
            noGoods: false,
            cartId: null,
            xz: 'static/images/xz@2x.png',
            wz: 'static/images/wz@2x.png',
            xz1: 'static/images/xz1@2x.png',
            wz1: 'static/images/wz1@2x.png'
        }
    },
    computed: {
        ...mapGetters([
          'cartNum' // 购物车商品数量
        ])
    },
    mounted () {
        this.getCartList()
        this.getCartNum()
    },
    methods: {
        ...mapMutations({
          setCartNum: 'SET_CARTNUM' // 修改vuex购物车数量
        }),
        // 获取购物车商品数量
        async getCartNum () {
            var res = await cartCount()
            if (res.status === 200) {
                this.setCartNum(res.data.count)
            } else {
                this.$message.warning(res.msg)
            }
        },
        // 获取购物车列表
        async getCartList () {
            this.loading = true
            var res = await cartList()
            if (res.status === 200) {
                this.cartList = res.data.valid
                this.invalidGoods = res.data.invalid
                this.totalPrice = 0
                this.postage = 0
                this.noGoods = (this.invalidGoods.length === 0 && this.cartList.length === 0)
                this.cartList.forEach(i => {
                    this.$set(i, 'selected', false)
                })
            } else {
                this.$message.warning(res.msg)
            }
            this.loading = false
        },
        // 跳转到商品详情
        toDetail (id) {
            this.$router.push({
                path: '/productDetail',
                query: {
                    productId: id
                }
            })
        },
        // 全选
        selectAll () {
            this.all = !this.all
            this.cartList.forEach(i => {
                i.selected = this.all
            })
            this.computePrice()
        },

        // 单选
        select (item) {
            item.selected = !item.selected
            this.all = true
            this.cartList.forEach(i => {
                if (!i.selected) {
                    this.all = false
                }
            })
            this.computePrice()
        },

        // 计算价格
        computePrice () {
            this.totalPrice = 0
            this.postage = 0
            this.cartList.forEach(i => {
                if (i.selected) {
                    this.totalPrice += i.cartNum * Number(i.productInfo.attrInfo.price)
                    this.postage += Number(i.productInfo.postage)
                }
            })
        },

        // 修改商品数量
        handleChange (id, num) {
            cartUpdate({
                id: id,
                number: num
            }).then(res => {
                if (res.status === 200) {
                    this.computePrice()
                } else {
                    this.$message.warning(res.msg)
                }
            })
        },
        // 删除请求
        del (ids) {
            console.log('删除购物车产品', ids)
            cartDel({ 'ids': ids }).then(res => {
                if (res.status === 200) {
                    this.$message.success('删除成功')
                    this.getCartList()
                    this.getCartNum()
                } else {
                    this.$message.warning(res.msg)
                }
            })
        },
        // 删除购物车商品
        delCart (id) {
            let ids = []
            ids.push(id)
            this.del(ids)
        },

        // 批量删除
        delCarts () {
            let ids = []
            this.cartList.forEach(i => {
                if (i.selected) {
                    ids.push(i.id)
                }
            })
            if (ids.length === 0) {
                return this.$message.warning('请先选择需要删除的商品')
            }
            this.del(ids)
            this.all = false
        },
        // 结算
        sett () {
            let ids = []
            this.cartList.forEach(i => {
                if (i.selected) {
                    ids.push(i.id)
                }
            })
            if (ids.length === 0) {
                return this.$message.warning('请选择需要结算的商品')
            }
            this.flag = true
            ids = ids.join(',')
            this.$router.push({
                path: '/cart/settlement',
                query: {
                    cartId: ids
                }
            })
        }
    }
}
</script>
<style lang="scss" sopced>
.shoppingCart{
    .top{
        margin-bottom: 24px;
        margin-top: 30px;
    }
    .content{
        margin-bottom: 100px;
        .left{
            width: 785px;
            .productBox{
                .label{
                    background-color: #333333;
                    height: 57px;
                    line-height: 57px;
                    .selectAll{
                        img{
                            width: 18px;
                            height: 17px;
                            cursor: pointer;
                        }
                    }
                }
                .productList{
                    border: 1px solid #CCCCCC;
                    box-sizing: border-box;
                    border-top: 0;
                    .product{
                        height: 190px;
                        padding: 20px 0;
                        box-sizing: border-box;
                        border-top: 1px solid #CCCCCC;
                        position: relative;
                        .mc{
                            position: absolute;
                            left: 0;
                            top: 0;
                            width: 100%;
                            height: 100%;
                            background-color: rgba($color: #000000, $alpha: 0.3);
                            opacity: 0.3;
                            z-index: 10;
                        }
                        .select{
                            img{
                                width: 18px;
                                height: 17px;
                                cursor: pointer;
                            }
                        }
                        .productInfo{
                            padding-right: 20px;
                            img{
                                width: 130px;
                                height: 150px;
                            }
                        }
                        .number{
                            .el-input-number{
                                .el-input-number__decrease,.el-input-number__increase{
                                    border-radius: 0;
                                }
                                .el-input__inner{
                                    font-size: 16px;
                                    border-radius: 0;
                                }
                            }
                        }
                        .price{
                            flex-direction: column;
                            justify-content: space-between;
                        }
                        .del{
                            z-index: 11;
                        }
                    }
                }
            }
            .nothing{
                padding: 80px 0;
                img{
                    width: 150px;
                    height: 130px;
                }
            }
        }
        .right{
            width: 365px;
            height: 325px;
            border: 1px solid #CCCCCC;
            box-sizing: border-box;
            .label{
                background-color: #CCCCCC;
                height: 57px;
                line-height: 57px;
                padding-left: 20px;
            }
            .totalPrice,.fare{
                border-bottom: 1px solid #CCCCCC;
                line-height: 57px;
            }
            .settlement{
                .total{
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
    .batchDel{
        .selectAll{
            img{
                width: 18px;
                height: 17px;
                cursor: pointer;
            }
        }
    }
    >>>.el-checkbox{
        transform : scale(1.5);
        zoom: 200%;
        .el-checkbox__input{
            width: 16px;
            height: 16px;
            .el-checkbox__inner{
                width: 16px;
                height: 16px;
            }
        }
    }
}
</style>
