<template>
    <div class="notice">
        <div class="noticeContent">
            <div class="nav">
                <router-link to="/">
                    <span>首页</span>
                </router-link>
                <span>|</span>
                <span>图文详情</span>
            </div>
            <section class="section" v-if="this.list.length === 0" v-loading="loading">
                <NotfoundInfo class="NInfo"/>
            </section>
            <section class="section" v-else  v-loading="loading">
                <ul class="listTitle" ref="list">
                    <li
                    class="titleItem"
                    v-for="item in this.list"
                    :key="item.value"
                    @click="getArticleDetial(item.id)"
                    :class="{changeWidths : isAnimation}">
                        <div class="articleList">
                            <div class="title oneLine">{{ item.title }}</div>
                            <div class="sml">{{ item.author }}</div>
                            <div class="sml oneLine"><span>{{ item.synopsis }}</span></div>
                            <div class="footer">
                                <div class="footerTime">{{ item.addTime.split(' ')[0] }}</div>
                                <div class="footerSyn">浏览量:{{ item.visit }}</div>
                            </div>
                        </div>
                    </li>
                </ul>
                <div
                class="content"
                :class="{showWidths : isAnimation}"
                v-if="isAnimation"
                v-loading="loading2"
                >
                    <h2>{{ article.title }}</h2>
                    <div class="info">
                        <div>{{ article.addTime }}</div>
                        <div>作者：{{ article.author }}</div>
                        <div>浏览量：{{ article.visit }}</div>
                    </div>
                    <article v-html="article.content"></article>
                </div>
            </section>
        </div>
    </div>
</template>

<script>
import NotfoundInfo from '@/components/NotFound/NotFoundInfo.vue'
import {articleList, articleDetial} from '@/api/article.js'
export default {
    name: 'notice',
    components: {NotfoundInfo},
    data () {
        return {
            list: [],
            article: {},
            isAnimation: false,
            loading: false,
            loading2: false
        }
    },
    created () {
        this.getarticleList()
    },
    methods: {
        async getarticleList () {
            this.loading = true
            var res = await articleList()
            if (res.status === 200) {
                this.list = res.data.map((item, index) => {
                    return {
                        'value': index,
                        'id': item.id,
                        'title': item.title,
                        'author': item.author,
                        'synopsis': item.synopsis,
                        'addTime': item.addTime,
                        'visit': item.visit
                    }
                })
                this.loading = false
            } else {
                this.$message.error(res.msg || '没有相关信息')
                this.loading = false
            }
        },
        async getArticleDetial (id) {
            this.loading2 = true
            var res = await articleDetial(id)
            if (res.status !== 200) {
                this.$message.error(res.msg)
                return
            }
            this.article = res.data
            this.loading2 = false
            this.isAnimation = true
        }
    }
}
</script>
<style lang="scss" scoped>
.notice{
    width: 1170PX;
    margin: auto;
    min-height: 700px;
    padding: 30px 0;
    .nav{
        margin-bottom: 30px;
    }
    .section{
        display: flex;
        .listTitle{
            min-width: 300px;
            max-height: 900px;
            overflow: auto;
            overflow-x: hidden;
            margin-right: 20px;
            background-color: #FFF;
            .titleItem{
                width: 800px;
                padding: 8px 12px;
                .articleList{
                    padding: 8px 12px;
                    border-radius: 3px;
                    .title{
                        font-size: 20px;
                    }
                    .sml{
                        color: #666;
                        font-size: 14px;
                        margin: 10px 0;
                        span{
                            color: #999;
                        }
                    }
                    .footer{
                        display: flex;
                        justify-content: space-between;
                        .footerSyn{
                            font-size: 14px;
                            color: #999;
                        }
                    }
                    .oneLine{
                        overflow: hidden;
                        white-space:nowrap;
                        text-overflow: ellipsis;
                    }
                    &:hover{
                        background-color: #EEE;
                    }
                }
            }
        }
        .content{
            // width: 850px;
            h2{
                text-align: center;
                font-size: 30px;
                color: #000;
                font-weight: 500;
                margin-top: 20px;
            }
            .info{
                display: flex;
                justify-content: center;
                div{
                    margin: 40px 10px;
                    font-size: 14px;
                    color: #999;
                }
            }
            article{
                text-indent: 2rem;
            }
        }
        .NInfo{
            display: block;
            margin: auto;
        }
    }
    .changeWidths{
        animation: changeWidth 1s;
        animation-fill-mode : forwards;
    }
    @keyframes changeWidth {
        from{
            width: 800px;
        }
        to {
            width: 300px;
        }
    }
    .showWidths{
        animation: showWidth 1s;
        animation-fill-mode : forwards;
        background-color: #FFF;
    }
    @keyframes showWidth {
        from{
            width: 300px;
            opacity: 0;
        }
        to{
            width: 850px;
            opacity: 1;
        }
    }
}
</style>
