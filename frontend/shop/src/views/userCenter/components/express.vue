/**
    物流信息组件
 */
<template>
    <div class="expressInfo">
        <div class="title">物流信息</div>
        <div class="content">物流公司：<span>{{ this.expressCompany }}</span></div>
        <div class="content mar-left">快递单号：<span>{{ this.expressCode }}</span></div>
        <ul class="postStep" v-if="postInfo.length === 0">
            <li class="noData">
                <img src="static/images/NofFound/wsj.webp" alt="暂无数据">
                <span>暂无数据哦~</span>
            </li>
        </ul>
        <ul class="postStep" v-else>
            <li class="steps"
            v-for="(item,index) in this.postInfo"
            :key="index"
            >
                <div class="stepTime">{{ item.time }}</div>
                <div class="stepIcon">
                    <div class="solid"></div>
                    <div class="stepi finalCircle" v-if="item.status === 1"></div>
                    <div class="stepi normalCircle" v-else></div>
                </div>
                <div class="stepDescription">
                    <!-- <div class="tit">{{ item.title }}</div>
                    <br /> -->
                    <div class="des">{{ item.des }}</div>
                </div>
            </li>
            <br />
            <br />
        </ul>
    </div>
</template>

<script>
import {
    expressOder
} from '@/api/order.js'
export default {
    name: 'express',
    props: {
        obj: Object
    },
    data () {
        return {
            expressCompany: '',
            expressCode: '',
            postInfo: [
                {
                    'status': 1,
                    'time': '10:04:56',
                    'title': '已签收',
                    'des': '您的订单已由本人签收。如有疑问您可以联系配送员【温国军，19197832273】确认。欢迎再次光临。'
                },
                {
                    'status': 0,
                    'time': '10:04:56',
                    'title': '已下单',
                    'des': '温馨提示：您的订单预计3月28日09:00-15:00送达您手中'
                },
                {
                    'status': 0,
                    'time': '10:04:56',
                    'title': '已揽件',
                    'des': '包裹等待揽收'
                },
                {
                    'status': 0,
                    'time': '10:04:56',
                    'title': '已发货',
                    'des': '您的订单在【广州九龙分拣中心】分拣完成'
                }
            ]
        }
    },
    watch: {
        obj: {
            handler (nVal, oVal) {
                this.expressCompany = nVal.deliveryName
                this.expressCode = nVal.deliveryId
                this.getExpress(nVal)
            },
            deep: true
        }
    },
    methods: {
        getExpress (obj) {
            let expressInfoDo = {
                'logisticCode': obj.deliveryId,
                // 'logisticCode': 'YT5399329246792', // 1382232177634181120 YTO
                'orderCode': obj.orderCode,
                'shipperCode': obj.deliverySn
            }
            // expressOder(expressInfoDo).then(res => {
            //     if (res.status === 200) {
            //         this.postInfo = res.data.Traces.reverse().map((item, index) => {
            //             return {
            //                 'status': 0,
            //                 'time': item.AcceptTime,
            //                 'des': item.AcceptStation
            //             }
            //         })
            //         if (this.postInfo.length > 0) {
            //             this.postInfo[0].status = 1
            //         }
            //     } else {
            //         this.$message.error(res.msg)
            //     }
            // })
        }
    }
}
</script>

<style lang="scss">
.expressInfo{
    .content{
        display: inline-block;
    }
    .mar-left{
        margin-left: 35px !important;
    }
    .postStep{
        max-width: 888px;
        height: 300px;
        padding-top: 30px;
        background-color: #F5F5F5;
        overflow: auto;
        display: flex;
        // flex-direction: column-reverse;
        flex-direction: column;
        .steps{
            display: flex;
            .stepTime,.stepDescription{
                height: 50px;
                margin: 0 0 30px 0;
            }
            .stepTime{
                flex: 3;
                text-align: right;
                border-right: #000;
            }
            .container{
                width: 24px;
                position: relative;
            }
            .stepIcon{
                width: 24px;
                margin: 0 35px;
                position: relative;
                .solid{
                    width: 2px;
                    position: absolute;
                    top: 0;
                    bottom: 0;
                    left: 11px;
                    background-color: #E6E6E6;
                }
                .stepi{
                    position: relative;
                    z-index: 1;
                }
                .finalCircle{
                    width: 24px;
                    height: 24px;
                    background: #F26E47;
                    border-radius: 50%;
                    position: relative;
                    text-align: center;
                    line-height: 24px;
                    color: #FFFFFF;
                }
                .normalCircle{
                    width: 10px;
                    height: 10px;
                    margin: auto;
                    background: #E6E6E6;
                    border-radius: 50%;
                }
            }
            .stepDescription{
                flex: 9;
                text-align: left;
                .tit{
                    font-size: 14px;
                    color: #000;
                }
                .des{
                    font-size: 12px;
                    color: #666666;
                }
            }
        }
        .noData{
            display: flex;
            flex-direction: column;
            align-items: center;
            img {
                width: 250px;
                height: 200px;
            }
            span{
                margin: 30px 0 10px 0;
                font-size: 20px;
                font-weight: 700;
            }
        }
    }
}
</style>
