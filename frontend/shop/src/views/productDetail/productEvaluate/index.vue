<template>
    <div class="productEvaluate">
        <div class="label">
            <el-radio-group v-model="radio" @change="handleChange">
                <el-radio :label="0">全部({{evaluateData.sumCount}})</el-radio>
                <el-radio :label="1">好评({{evaluateData.goodCount}})</el-radio>
                <el-radio :label="2">中评({{evaluateData.inCount}})</el-radio>
                <el-radio :label="3">差评({{evaluateData.poorCount}})</el-radio>
            </el-radio-group>
        </div>
        <div class="evaluateList" v-loading="loading">
            <div class="evaluateBox" v-for="(item) in evaluateList" :key="item.id">
                <div class="user flex">
                    <el-image class="avatar mar-right-20" :src="item.avatar"></el-image>
                    <div class="flex1">
                        <p class="mar-bot-10">{{item.nickname}}</p>
                        <el-rate :colors="color" v-model="item.productScore" disabled></el-rate>
                    </div>
                    <p class="col-999">{{item.createTime | comverTime('YYYY-MM-DD HH:mm')}}</p>
                </div>
                <div class="content mar-top-20">
                    <p>{{item.comment}}</p>
                </div>
                <div class="imgList mar-top-40">
                    <el-image
                        v-for="(img,index) in item.pics"
                        :key="index"
                        style="width: 70px; height: 70px; margin-right: 20px;"
                        :src="img"
                        :preview-src-list="item.pics">
                    </el-image>
                </div>
                <div class="spec mar-top-30 col-666">规格：{{item.sku}}</div>
                <div class="reply mar-top-20" v-if="item.merchantReplyTime">
                    <p>商家回复：{{item.merchantReplyContent}}</p>
                    <p class="mar-top-20 tar">{{item.merchantReplyTime | comverTime('YYYY-MM-DD HH:mm')}}</p>
                </div>
            </div>
        </div>
        <!-- <div v-else>
            <img src="../../../../static/images/wsj@2x.png" alt="">
        </div> -->
        <button class="getMore" @click="getMore">加载更多</button>
        <!-- <el-pagination
          class="mar-top-40 tar"
          @current-change="handleCurrentChange"
          :current-page.sync="page"
          background
          :page-size="limit"
          layout="prev, pager, next, jumper"
          :total="total">
        </el-pagination> -->
    </div>
</template>

<script>
import {
    getReplyData,
    getReplyList
} from '@/api/reply.js'
export default {
    name: 'productEvaluate',
    props: ['productId'],
    data () {
        return {
            radio: 0,
            color: ['#F26E47'],
            page: 1,
            limit: 10,
            total: 20,
            type: 0,
            loading: false,
            evaluateData: {},
            evaluateList: []
        }
    },
    mounted () {
        //this.getReply(this.productId)
        this.getReplyList(this.productId)
    },
    methods: {
        // 获取商品评论数据
        async getReply (id) {
            var res = await getReplyData(id)
            if (res.status === 200) {
                this.evaluateData = res.data
                // Object.assign(this.evaluateData, res.data)
            } else {
                this.$message.warning(res.msg)
            }
        },

        // 获取商品评论列表
        getReplyList (id) {
            this.loading = true
            getReplyList(id, {
                limit: this.limit,
                page: this.page,
                type: this.type
            }).then(res => {
                if (res.status === 200) {
                    // 往数据中累加
                    if (res.data) {
                        this.evaluateList = [...this.evaluateList, ...res.data]
                    }
                    
                } else {
                    this.$message.warning(res.msg)
                }
                this.loading = false
            })
        },

        handleChange (val) {
            // 初始化查询条件
            this.type = val
            this.total = val === 0 ? this.evaluateData.sumCount : val === 1 ? this.evaluateData.goodCount : val === 2 ? this.evaluateData.inCount : this.evaluateData.poorCount
            this.evaluateList = []
            this.page = 1
            this.getReplyList(this.productId)
        },

        // 翻页
        getMore () {
            this.page += 1
            this.getReplyList(this.productId)
        }
    }
}
</script>
<style lang="scss">
.productEvaluate{
    .label{
        padding: 0 20px;
        box-sizing: border-box;
        height: 56px;
        line-height: 56px;
        background-color: #F5F5F5;
        border: 1px solid #E6E6E6;
        .el-radio{
            .el-radio__input.is-checked{
                .el-radio__inner{
                    background-color: #666666;
                    border-color: #666666;
                }
            }
            .el-radio__input.is-checked+.el-radio__label{
                color: #666666;
            }
        }
    }
    .evaluateList{
        border: 1px solid #E6E6E6;
        border-top: 0;
        .evaluateBox{
            padding: 30px;
            border-bottom: 1px solid #E6E6E6;
            box-sizing: border-box;
            .user{
                align-items: center;
                .avatar{
                    width: 60px;
                    height: 60px;
                    border-radius: 30px;
                }
                .el-rate__icon{
                    font-size: 24px;
                    color: #F26E47;
                }
            }
            .content{
                line-height: 1.5;
            }
            .imgList{
                .small{
                    width: 70px;
                    height: 70px;
                }
            }
            .reply{
                padding: 20px;
                background-color: #F6F6F6;
            }
        }
        .evaluateBox:last-child{
            border-bottom: 0;
        }
    }
    .getMore{
        display: block;
        width: 200px;
        height: 40px;
        margin: 20px auto;
        cursor: pointer;
        color: #FFF;
        background-color: #F26E47;
        &:hover {
            color: #F26E47;
            border: 1px solid #F26E47;
            background-color: #FFF;
        }
    }
}
</style>
