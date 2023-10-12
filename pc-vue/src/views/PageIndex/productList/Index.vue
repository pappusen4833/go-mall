/**
    首页-全部商品
    用来放在首页做展示商品列表
 */
<template>
    <div class="classify" @click.stop="hideFatherItem">
        <nav class="searchNav">
            <div class="breadTitle">
                <div class="bread">
                    <span>首页</span>|
                    <span>{{ this.bread }}</span>
                    |<span>{{ this.breadSon }}</span>
                </div>
                <!-- 筛选 -->
                <ul class="searchType">
                    <li class="searchItem">
                        <DropDown
                        :title="'价格排序'"
                        :arr="this.priceSorts"
                        @selectVal="getPriceSort"/>
                    </li>
                    <li class="searchItem">
                        <DropDown
                        :title="'销量排序'"
                        :arr="this.sealsSorts"
                        @selectVal="getSealsSort"/>
                    </li>
                    <li class="searchItem">
                        <DropDown
                        :title="'新品排序'"
                        :arr="this.newsSorts"
                        @selectVal="getNewsSort"/>
                    </li>
                </ul>
            </div>
            <ul class="fatherType commType">
                <li
                    class="fatherTypeItem"
                    v-for="item in this.types.slice(0, 10)"
                    :key="item.value"
                    :class="{beSelected : item.value === fatherType}"
                    @click="fatherSelected(item)">
                    {{item.label}}
                </li>
                <li
                v-if="this.types.length > 10"
                class="fatherTypeItem moreItem"
                @click.stop="showFatherItem">更多分类</li>
            </ul>
            <!-- 溢出部分 -->
            <ul
            class="secondFatherType commType"
            ref="secondFatherType"
            :class="isShow ? 'showIt' : ''">
                <li
                    class="fatherTypeItem"
                    v-for="item in this.types.slice(10, 21)"
                    :key="item.value"
                    :class="{beSelected : item.value === fatherType}"
                    @click="fatherSelected(item)">
                    {{item.label}}
                </li>
            </ul>
            <ul class="sonType commType" v-loading="loadingType">
                <li
                class="fatherTypeItem sonTypeItem"
                v-for="item in this.filters"
                :key="item.value"
                :class="{beSelected : item.value === sonType}"
                @click="sonSelected(item)">
                    {{ item.label }}
                </li>
            </ul>
        </nav>
        <!-- 列表部分 -->
        <section
        class="productList"
        v-if="productList"
        v-loading="loading">
            <ProductItem
            class="productItem"
            v-for="(item, index) in this.productList"
            :key="index"
            :obj="item"/>
        </section>
        <section class="productList" v-else>
            <NotFoundInfo class="nodata" />
        </section>
        <button class="moreProduct" v-if="productList"  @click="getMore">加载更多</button>
    </div>
</template>
<script>
import ProductItem from './productItem'
import NotFoundInfo from '@/components/NotFound/NotFoundInfo.vue'
import DropDown from '@/components/common/dropDown.vue'
import {
    getProducts,
    getCategory
} from '@/api/indexpage.js'
export default {
    name: 'classify',
    components: {ProductItem, NotFoundInfo, DropDown},
    data () {
        return {
            listquery: {
                'isIntegral': 0, // 是否积分兑换商品 0否  1是
                'keyword': '', // 关键
                'news': '', // 是否新品  0否  1是
                'priceOrder': '', // 价格排序
                'salesOrder': '', // 销量排序
                'sid': '', // 分类ID
                'type': '', // 类别 不传
                'limit': 10,
                'page': 1
            },
            productList: [],
            types: [
                {'value': '', 'label': '全部'}
            ],
            filters: [
                {'value': '', 'label': '全部'}
            ],
            // 筛选数据
            newsSorts: [
                { 'value': '', 'label': '默认' },
                {'value': 1, 'label': '是'},
                {'value': 0, 'label': '否'}
            ],
            priceSorts: [
                {'value': 'desc', 'label': '价格正序'},
                {'value': 'asc', 'label': '价格倒序'}
            ],
            sealsSorts: [
                {'value': 'desc', 'label': '销量正序'},
                {'value': 'asc', 'label': '销量倒序'}
            ],
            fatherType: '',
            isShow: false,
            showLock: 0,
            sonType: '',
            bread: '全部宝贝',
            breadSon: '',
            loading: false,
            loadingType: false
        }
    },
    created () {
        this.listquery.keyword = this.$route.query.keyword || ''
        this.fatherType = this.$route.query.type1 || ''
        this.sonType = this.$route.query.type2 || ''
        this.getCateList()
        this.getProList()
    },
    watch: {
        $route: {
            handler (nVal, oVal) {
                this.listquery.keyword = nVal.query.keyword || ''
                this.getProList()
            }
        }
    },
    methods: {
        filterRoute (id) {
            this.types.forEach(item => {
                if (item.value === id) {
                    this.filters = item.children
                }
            })
        },
        // 分类查询
        async getCateList () {
            var response = await getCategory()
            if (response.status === 200) {
                var res = response.data.content
                this.types = this.types.concat(
                    res.map(item => {
                        return {
                            'value': `${item.id}`,
                            'label': `${item.cateName}`,
                            'children': item.children.map(i => {
                                return {
                                    'value': `${i.id}`,
                                    'label': `${i.cateName}`,
                                    'children': `${i.children}`
                                }
                            })
                        }
                    })
                )
            } else {
                this.$message.error(res.msg || '网络异常！')
            }
            if (this.$route.query.type1 !== '' && this.$route.query.type2 !== '') {
                this.filterRoute(this.$route.query.type1)
            }
        },
        // 商品查询
        async getProList () {
            this.loading = true
            var res = await getProducts(this.listquery)
            if (res.status === 200) {
                //this.productList = this.productList.concat(res.data)
                this.productList = res.data
            }
            this.loading = false
        },
        // 选中父级
        fatherSelected (item) {
            this.loadingType = true
            this.bread = item.label
            this.breadSon = ''
            this.fatherType = item.value
            if (item.value === '') {
                this.filters = [{'value': '', 'label': '全部'}]
            } else {
                if (Array.isArray(item.children)) {
                    this.filters = item.children.map(i => {
                        return {
                            'value': `${i.value}`,
                            'label': `${i.label}`,
                            'children': i.children
                        }
                    })
                } else {
                    this.filters = [{'value': '', 'label': '全部'}]
                }
            }
            this.loadingType = false
            this.initQuery(item)
            this.listquery.sid = item.value
            getProducts(this.listquery).then(res => {
                if (res.status === 200) {
                    this.productList = res.data
                    this.loading = false
                } else {
                    this.$message.error(res.msg)
                    this.loading = false
                }
            })
        },
        // 展示溢出父级分类
        showFatherItem () {
            this.showLock++
            if (this.showLock % 2) {
                this.$refs.secondFatherType.style = 'display: none'
            } else {
                this.$refs.secondFatherType.style = 'display: inline-flex'
                this.isShow = true
            }
        },
        // 隐藏溢出部分
        hideFatherItem () {
            if (this.isShow) {
                this.showLock++
                this.$refs.secondFatherType.style = 'display: none'
            }
        },
        // 选中子元素
        sonSelected (item) {
            this.sonType = item.value
            this.breadSon = item.label
            this.loading = true
            this.initQuery(item)
            this.listquery.sid = item.value
            getProducts(this.listquery).then(res => {
                if (res.status === 200) {
                    this.productList = res.data
                }
                this.loading = false
            })
        },
        initQuery (item) {
            this.listquery.sid = item.value
            this.listquery.page = 1
        },
        // 新品排序
        getNewsSort (item) {
            this.listquery.news = item.value
            this.getProList()
        },
        getSealsSort (item) {
            this.listquery.salesOrder = item.value
            this.getProList()
        },
        getPriceSort (item) {
            this.listquery.priceOrder = item.value
            this.getProList()
        },
        // 加载更多
        getMore () {
            this.listquery.page++
            this.getProList()
        }
    }
}
</script>
<style lang="scss">
.classify{
    padding: 0 0 60px 0;
    .searchNav{
        max-width: 1170px;
        line-height: 56px;
        margin: 10px auto;
        position: relative;
        .breadTitle{
            .bread{
                height: 56px;
                display: inline-block;
                cursor: default;
                span{
                    margin: 5px;
                    font-size: 16px;
                    font-family: PingFang SC;
                    font-weight: 400;
                    color: #333333;
                }
            }
            .searchType{
                float: right;
                display: flex;
                .searchItem{
                    margin-left: 20px;
                    cursor: pointer;
                }
            }
        }
        .commType{
            display: flex;
            // justify-content: space-around ;
            font-size: 16px;
            font-family: PingFang SC;
            font-weight: 500;
            color: #333333;
            li{
                border-bottom: 3px solid transparent;
                cursor: pointer;
            }
        }
        // 分类
        .fatherType{
            background-color: #FFFFFF;
        }
        // 分类溢出部分
        .secondFatherType{
            display: none;
            background-color: #FFFFFF;
            box-shadow: 0px 5px 15px 0px rgba(0, 0, 0, 0.15);
            position: absolute;
            top: 118px;
            z-index: 1;
        }
        .fatherTypeItem{
            width: 65px;
            margin: 0 20px;
            text-align: center;
        }
        .moreItem{
            color: #F26E47;
            font-weight: bold;
        }
        .showIt{
            animation: showOn .5s;
            animation-fill-mode : forwards;
            line-height: 80px;
            width: 100%;
            justify-content: start;
            flex-wrap: wrap;
        }
        @keyframes showOn {
            from{
                height: 56px;
                opacity: 0;
            }
            to{
                height: 80px;
                opacity: 1;
            }
        }
        .sonType{
            margin: 10px 0;
            background-color: #FFFFFF;
            display: flex;
            flex-wrap: wrap;
        }
        .beSelected{
            color: #F26E47;
            border-bottom: 3px solid #F26E47 !important;
        }
    }
    .productList{
        max-width: 1170px;
        margin: auto;
        display: flex;
        flex-wrap: wrap;
        .productItem{
            margin: 60px 6px 30px;
            width: 278px;
        }
        .nodata{
            margin: 80px auto;
        }
    }
    .moreProduct{
        width: 160px;
        height: 50px;
        display: block;
        margin: 60px auto;
        border: 1px solid #E6E6E6;
        background-color: #FFF;
        cursor: pointer;
    }
}
</style>
