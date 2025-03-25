/**
    商品信息公用组件
    数据示例：
    listarr: [
        {
            'status': 0,
            'img': '',
            'title': 'EMPORIO ARMANI刺绣LOGO连帽卫衣',
            'color': '灰色',
            'size': 'M',
            'sum': '1',
            'price': '500',
            'url': ''
        }
    ],
 */
<template>
    <div class="commoditys">
        <div class="commonOrderCommodity" v-for="(i,index) in this.listarr" :key="index">
            <div class="commonOrderPre">
                <!-- <image :src="item.img" /> -->
                <div class="img">
                    <!-- <img class="img" src="../../../../static/images/orderlist.png" /> -->
                    <el-image class="img" :src="i.img"></el-image>
                </div>
                <div class="info">
                    <div>{{ i.storeName }}</div>
                    <!-- <div class="wrap">
                        <div><span>颜色：</span>{{ i.color }}</div>
                        <div><span>尺码：</span>{{ i.size }}</div>
                    </div>
                    <div class="wrap">
                        <div><span>数量：</span>{{ i.sum }}</div>
                    </div> -->
                    <div class="wrap">
                        <div>{{ i.sku }}</div>
                    </div>
                </div>
                <div class="itemPrice">￥{{ i.truePrice }}</div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'commidity',
    props: {
        listarr: Array
    }
}
</script>

<style lang="scss">
.commoditys{
    display: flex;
    flex-direction: column;
    .commonOrderCommodity{
        margin-bottom: 30px;
        padding-left: 30px;
        display: flex;
        flex-direction: column;
        .commonOrderPre{
            flex: 10;
            display: flex;
            .img{
                width: 130px;
                height: 150px;
                background-color: #000;
            }
            .info{
                width: 280px;
                height: 150px;
                margin: 0 20px;
                padding-top: 7px;
                .wrap{
                    margin-top: 10px;
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
                font-size: 20px;
            }
        }
    }
}
</style>
