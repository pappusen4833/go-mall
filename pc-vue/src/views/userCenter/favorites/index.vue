<template>
    <div class="favorites">
        <h2 class="fs30">我的收藏</h2>
        <el-divider></el-divider>
        <div v-loading="loading" v-if="!noGoods">
            <div class="favoritesList flex">
                <div class="favoritesBox mar-bot-30" v-for="(item) in favoritesList" :key="item.productId">
                    <img class="cur-poi" :src="item.product.image" alt="" @click="toDetail(item.productId)">
                    <div class="pad-15">
                        <p class="col-666 overflow">{{item.product.storeName}}</p>
                        <div class="mar-top-10">
                            <span class="col-333 bold">¥{{item.product.price}}</span>
                            <span class="fr col-666 cur-poi" @click="del(item.productId)">删除</span>
                        </div>
                    </div>
                </div>
                <div class="clearfix"></div>
            </div>
            <el-pagination
              class="mar-top-40 tar"
              @current-change="handleCurrentChange"
              :current-page.sync="page"
              background
              :page-size="limit"
              layout="prev, pager, next, jumper"
              :total="total">
            </el-pagination>
        </div>
        <div class="nothing display-column display-center" v-else>
            <img src="static/images/wsj@2x.png" alt="">
            <p class="fs18 mar-top-20">暂无数据哦～</p>
        </div>
    </div>
</template>
<script>
import {
    getCollectData,
    deleteCollection
} from '@/api/user.js'
export default {
    name: 'favorites',
    data () {
        return {
            favoritesList: [],
            limit: 9,
            page: 1,
            total: 1,
            type: 'collect',
            loading: false,
            noGoods: false
        }
    },
    mounted () {
        this.getCollect()
    },
    methods: {
        // 获取收藏商品列表
        getCollect () {
            this.loading = true
            getCollectData({
                limit: this.limit,
                page: this.page,
                type: `${this.type}`
            }).then(res => {
                if (res.status === 200) {
                    this.favoritesList = res.data
                    this.noGoods = !res.data.length
                } else {
                    this.$message.warning(res.msg)
                }
                this.loading = false
            })
        },
        // 跳转到商品详情
        toDetail (id) {
            console.log("id:"+id)
            this.$router.push({
                path: '/productDetail',
                query: {
                    productId: id
                }
            })
        },
        // 取消收藏
        async del (id) {
            var res = await deleteCollection({
                category: 'collect',
                id
            })
            if (res.status === 200) {
                this.$message.success('删除成功')
                this.getCollect()
            } else {
                this.$message.warning(res.msg)
            }
        },
        // 翻页
        handleCurrentChange (val) {
            this.page = val
            this.getCollect()
        }
    }
}
</script>
<style lang="scss" scoped>
.favorites{
    padding: 30px;
    background-color: #fff;
    .favoritesList{
        min-height: 330px;
        flex-wrap: wrap;
        .favoritesBox{
            margin-right: 58px;
            img{
                display: block;
                width: 258px;
                height: 258px;
            }
        }
        .favoritesBox:nth-child(3n){
            margin-right: 0;
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
</style>
