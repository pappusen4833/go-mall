<template>
    <div class="recommend">
        <div class="label">
            <div class="line"></div>
            <div class="title fs24">热门推荐</div>
        </div>
        <div class="goodsList">
            <div class="goodsBox mar-bot-20" v-for="(item) in list" :key="item.id">
                <el-image class="goodsImg cur-poi" :src="item.image" @click="lookProductDetail(item.id)"></el-image>
                <div class="goodsInfo">
                    <p class="goodsName overflow fs22 col-666">{{item.storeName}}</p>
                    <div class="mar-top-20 display-between">
                        <p class="goodsPrice fs20">¥{{item.price}}</p>
                        <el-image class="cart cur-poi" src="static/images/gwc.png" @click="addToCart(item.id)"></el-image>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import {cartAdd} from '@/api/cart.js'
export default {
    name: 'recommend',
    props: ['list'],
    methods: {
        lookProductDetail (id) {
            console.log(this.$route.query.productId, id)
            if (this.$route.query.productId === `${id}`) {
                this.$message.warning('该内容为当前商品')
                return
            }
            this.$router.push({
                path: '/productDetail',
                query: {
                    productId: id,
                    type: 0
                }
            })
            this.$emit('reload', id)
        },
        addToCart (id) {
            cartAdd({
                'bargainId': 0,
                'cartNum': 1,
                'combinationId': 0,
                'new': 0,
                'productId': id,
                'secKillId': 0,
                'uniqueId': ''
            }).then(res => {
                if (res.status === 200) {
                    this.$message.success('成功添加到购物车')
                } else {
                    this.$message.warning(res.msg)
                }
            })
        }
    }
}
</script>

<style lang="scss" scoped>
.recommend{
    padding: 30px 0 30px 30px;
    box-sizing: border-box;
    .label{
        height: 40px;
        position: relative;
        .line{
            position: absolute;
            top: 50%;
            left: 50%;
            transform: translate(-50%,-50%);
            width: 200px;
            height: 2px;
            background-color: #333333;
        }
        .title{
            width: 140px;
            text-align: center;
            background-color: #FFFFFF;
            position: absolute;
            top: 50%;
            left: 50%;
            transform: translate(-50%,-50%);
        }
    }
    .goodsList{
        .goodsBox{
            border: 1px solid ;
            .goodsImg{
                width: 280px;
                height: 280px;
                display: block;
            }
            .goodsInfo{
                width: 100%;
                padding: 20px;
                box-sizing: border-box;
                div{
                    align-items: center;
                }
                .cart{
                    width: 20px;
                    height: 18px;
                }
            }
        }
    }
}
</style>
