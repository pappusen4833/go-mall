<template>
    <div class="evaluate" v-loading="loading">
        <h2 class="fs30">评价晒单
            <el-button class="fr" @click="submitWord">发表</el-button>
        </h2>
        <el-divider></el-divider>
        <!-- 服务评价 -->
        <div class="logistics">
            <div class="flex fs20">
                <div class="flex1 flex">
                    <span class="mar-right-10">服务评分</span>
                    <el-rate :colors="color" v-model="pack"></el-rate>
                </div>
                <!-- <div class="flex1 flex">
                    <span class="mar-right-10">物流速度</span>
                    <el-rate :colors="color" v-model="speed"></el-rate>
                </div>
                <div class="flex1 flex">
                    <span class="mar-right-10">客服服务</span>
                    <el-rate :colors="color" v-model="serve"></el-rate>
                </div> -->
            </div>
            <!-- <el-input
              class="mar-top-20"
              type="textarea"
              :rows="5"
              resize="none"
              placeholder="请填写您遇到的问题,这将帮助我们为您提供更好的服务~"
              v-model="textarea">
            </el-input> -->
        </div>
        <!-- <el-divider></el-divider> -->
        <!-- 商品评价 -->
        <div
        class="product flex mar-top-60"
        v-for="(item, index) in dataList"
        :key="index">
            <div class="left flex2 flex">
                <img :src="item.image" alt="">
                <div class="info mar-left-20">
                    <p>{{ item.title }}</p>
                    <div class="flex mar-top-10">
                        <!-- <p class="col-999 mar-right-10">颜色：<span class="col-333">灰色</span></p>
                        <p class="col-999 mar-right-10">尺码：<span class="col-333">M</span></p>
                        <p class="col-999 mar-right-10">数量：<span class="col-333">1</span></p> -->
                        <p class="col-999 mar-right-10">{{ item.sku}}</p>
                    </div>
                </div>
            </div>
            <div class="right flex3 mar-left-20">
                <div class="flex fs20">
                    <span class="mar-right-10">商品评分</span>
                    <el-rate :colors="color" v-model="item.goods"></el-rate>
                </div>
                <el-input
                  class="mar-top-20"
                  type="textarea"
                  :rows="5"
                  resize="none"
                  placeholder="请填写您遇到的问题,这将帮助我们为您提供更好的服务~"
                  v-model="item.comment">
                </el-input>
                <div @click.once="changed(index)" v-loading="loading2">
                    <el-upload
                    class="mar-top-20"
                    :action="action"
                    list-type="picture-card"
                    :limit="5"
                    :on-success="handSuccess"
                    :on-remove="handleRemove"
                    >
                    <i class="el-icon-plus"></i>
                    </el-upload>
                </div>
                <p class="mar-top-10">最多上传5张图片</p>
                <!-- <div class="imgNum">
                  共<span class="text">{{item.imgNum}}</span>张，还可以上传<span class="text">{{5-(item.imgNum)}}</span>张
                </div> -->
            </div>
        </div>
    </div>
</template>
<script>
import {detailOrder} from '@/api/order.js'
import {commentOrders} from '@/api/reply.js'
import {uploadURL} from '@/api/user.js'
export default {
    name: 'evaluate',
    data () {
        return {
            dataList: [
                {
                    'pics': []
                }
            ],
            loading: false,
            loading2: false,
            pack: 5,
            speed: 5,
            serve: 5,
            goods: 5,
            textarea: '',
            lock: true,
            imgIndex: 0,
            pics: [],
            color: ['#F26E47'],
            dialogVisible: false,
            action: uploadURL
        }
    },
    mounted () {
        this.getProductData(this.$route.query.id)
    },
    methods: {
        async getProductData (id) {
            if (!id) { return }
            var res = await detailOrder(id)
            if (res.status === 200) {
                this.dataList = res.data.cartInfo.map(item => {
                    if (item.isReply === 0) {
                        return {
                            'id': item.id,
                            'productId': item.productId,
                            'combinationId': item.combinationId,
                            'seckillId': item.seckillId,
                            'unique': item.unique,
                            'isReply': item.isReply,
                            'image': item.productInfo.attrInfo.image,
                            'title': item.productInfo.storeName,
                            'sku': item.productInfo.attrInfo.sku,
                            'comment': '',
                            'goods': 5,
                            'pics': ''
                        }
                    }
                }).filter(r => r)
            } else {
                this.$message.error(res.msg)
            }
        },
        changed (index) {
            if (this.lock) {
                this.imgIndex = index
                this.lock = false
            }
            this.lock = true
            this.loading2 = true
        },
        handSuccess (res) {
            this.pics.push({
                'index': this.imgIndex,
                'image': res.data
            })
            this.lock = true
            this.loading2 = false
        },
        handleRemove (res, file) {
            this.pics = this.pics.map(item => {
                if (item.index !== this.imgIndex) {
                    return item
                }
            }).filter(r => r)
            file.forEach(i => {
                this.pics.push({
                    'index': this.imgIndex,
                    'image': i.response.link
                })
            })
        },
        submitWord () {
            if (this.loading2) { return }
            var errs = 0
            this.dataList = this.dataList.map((item, index) => {
                if (item.comment.length === 0) {
                    this.$message.error('评价内容不能为空')
                    errs++
                }
                this.pics.forEach(i => {
                    if (parseInt(i.index) === parseInt(index)) {
                        item.pics += `,${i.image}`
                        console.log(item.pics)
                    }
                })
                if (item.pics[0] === ',') {
                    item.pics = item.pics.substr(1) // 截取从1开始的剩下部分
                }
                item.productScore = item.goods
                item.serviceScore = this.pack
                return item
            }).filter(r => r)
            console.log(this.dataList)
            if (errs !== 0) {
                return
            }
            this.loading = true
            commentOrders(this.dataList,this.$route.query.id).then(res => {
                this.loading = false
                console.log(res)
                this.$message.success('评论成功')
                this.$router.push({path: '/myEvaluate'})
            })
        }
    }
}
</script>
<style lang="scss" scoped>
.evaluate{
    padding: 30px;
    background-color: #fff;
    h2{
        .el-button{
            width: 120px;
            height: 40px;
            box-sizing: border-box;
            color: #FFFFFF;
            background-color: #F26E47;
            border-radius: 0;
        }
    }
    .product{
        .left{
            img{
                width: 130px;
                height: 150px;
            }
            .info{
                line-height: 1.5;
                div{
                    flex-wrap: wrap;
                }
            }
        }
    }
    >>>.el-rate{
        line-height: 24px;
    }
    >>>.el-rate__icon{
        font-size: 24px;
        color: #F26E47;
    }
    >>>.el-upload{
        width: 94px;
        height: 94px;
        line-height: 94px;
    }
}
</style>
