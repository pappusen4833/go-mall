/**
    订单状态组件
 */
<template>
    <div class="orderStatus">
        <div class="title">
            <div class="status"><span>{{ this.title }}</span></div>
            <div v-if="status === 10">剩余时间：<span>{{  '00:55:55' }}</span></div>
        </div>
        <ul class="section" v-if="status !== 12">
            <li v-for="item in this.layoutlist" :key="item.val" :class="item.type">{{ item.des }}</li>
        </ul>
        <ul class="section" v-else>
            <li class="statusfail">提交订单</li>
            <li class="statusfail">付款</li>
            <li class="statusfail">拼团失败</li>
            <li>发货</li>
            <li>收货</li>
            <li>交易成功</li>
        </ul>
    </div>
</template>

<script>
export default {
    name: 'OrderStatus',
    props: {
        status: Number
    },
    data () {
        return {
            allStatus: [
                {
                    'des': '提交订单',
                    'status': [-1, -2, 0, 1, 2, 3, 4, 10, 11, 12, 13],
                    'type': '',
                    'val': 0
                },
                {
                    'des': '付款',
                    'status': [-1, -2, 0, 1, 2, 3, 4, 10, 11, 12, 13],
                    'type': '',
                    'val': 1
                },
                {
                    'des': '拼团/待发货',
                    'status': [-1, -2, 0, 1, 2, 3, 4, 11, 12],
                    'type': '',
                    'val': 2
                },
                {
                    'des': '发货',
                    'status': [-1, -2, 1, 2, 3, 4],
                    'type': '',
                    'val': 3
                },
                {
                    'des': '收货',
                    'status': [-1, -2, 2, 3, 4],
                    'type': '',
                    'val': 4
                },
                {
                    'des': '交易成功',
                    'status': [-1, -2, 2, 3, 4],
                    'type': '',
                    'val': 5
                }
            ],
            layoutlist: [],
            title: '',
            titleList: ['待付款', '拼团中', '待发货', '待收货', '交易成功']
        }
    },
    watch: {
        status: {
            handler (nVal, oVal) {
                this.statusLayout(nVal)
                if (nVal === 12) {
                    this.title = '拼团失败'
                }
            }
        }
    },
    methods: {
        statusLayout (sta) {
            // 4
            if (!sta && sta !== 0) { return }
            this.layoutlist = this.allStatus.map(item => {
                if (item.status.indexOf(sta) !== -1) {
                    item.type = 'statusSuccess'
                }
                return item
            })
            for (var i = this.layoutlist.length - 1; i > 0; i--) {
                if (this.layoutlist[i].type !== '') {
                    if (i >= this.titleList.length - 1) {
                        this.title = this.titleList[this.titleList.length - 1]
                        return
                    }
                    this.title = this.titleList[i]
                    return
                }
            }
        }
    }
}
</script>

<style lang="scss">
.orderStatus{
    margin: 30px;
    .title{
        div{
            display: inline-block;
        }
        .status{
            font-size: 24px;
            margin-right: 20px;
        }
        span{
            color: #F26E47
        }
    }
    .section,.footer{
        height: 20px;
        line-height: 20px;
        display: flex;
        cursor: default;
        li{
            flex: 1;
            text-align: center;
            font-size: 12px;
        }
    }
    .section{
        margin: 20px 0;
        border-radius: 10px;
        background-color: #F5F5F5;
        li{
            color: #333333;
        }
        .statusSuccess{
            color: #FFFFFF;
            background-color: #91C25F;
            border-radius: 10px;
        }
        .statusfail{
            color: #FFFFFF;
            background-color: #C25F5F;
            border-radius: 10px;
        }
    }
}
</style>
