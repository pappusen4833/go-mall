/**
    首页默认内容
 */
<template>
  <div class="indexAll">
    <!-- 轮播图 -->
    <nav class="carousel">
      <el-carousel
      height="755px"
      arrow="never">
          <el-carousel-item  key="1">
          <div carContant>
            <el-image class="bcimg" src="https://image.dayouqiantu.cn/log1.jpg"></el-image>
            <span>https://image.dayouqiantu.cn/log1.jpg</span>
          </div>
        </el-carousel-item>
          <el-carousel-item  key="2">
          <div carContant>
            <el-image class="bcimg" src="https://image.dayouqiantu.cn/log2.jpg"></el-image>
            <span>https://image.dayouqiantu.cn/log2.jpg</span>
          </div>
        </el-carousel-item>
      </el-carousel>
      <TypeSelectBox class="typeBox" :data="this.types" :typeSelect="this.typeSelect"/>
    </nav>
    <!-- 分区内容 -->
    <section class="indexContent" v-loading="loading">
      <!-- 新款上新 -->
      <IndexListSlot
      v-if="layoutData.indexOf('ProductNew') !== -1"
      class="blockItem"
      :obj="productNew.header">
        <div class="common">
          <ProductItem
            class="newPro"
            v-for="(i, index2) in productNew.section"
            :key="index2"
            :obj="i"
            :type="0"/>
        </div>
      </IndexListSlot>
      <!-- 人气爆款 -->
      <IndexListSlot
      v-if="layoutData.indexOf('ProductHot') !== -1"
      class="blockItem"
      :obj="productHot.header">
        <div class="common hotItem">
        <div class="left">
          <div class="title">{{ productHot.section[0].storeName }}</div>
          <div class="price">
            {{ `￥${productHot.section[0].otPrice}` }}
            <div class="putCart" @click="toCart(productHot.section[0].id)">
            <svg
              t="1618572221481" class="icon" viewBox="0 0 1170 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1604" width="16" height="16"><path d="M1119.085714 138.313143H269.165714L239.177143 42.861714a60.928 60.928 0 0 0-58.514286-42.715428H40.96a40.594286 40.594286 0 1 0 0 81.115428h124.342857l229.668572 727.04h656.822857a40.521143 40.521143 0 1 0 0-81.115428H455.68l-33.645714-105.398857 705.828571-118.637715a51.2 51.2 0 0 0 42.422857-50.029714V189.001143a51.2 51.2 0 0 0-51.2-50.688z m-587.337143 707.876571a88.795429 88.795429 0 1 0 89.965715 88.795429 89.380571 89.380571 0 0 0-89.965715-88.795429z m430.08 0a88.795429 88.795429 0 1 0 89.965715 88.795429 89.380571 89.380571 0 0 0-89.965715-88.795429z" fill="#333333" p-id="1605"></path></svg>
              </div>
          </div>
          <el-image class="imgPro"
            :src="productHot.section[0].image"
            alt="商品图片"
            @click="toDetail(productHot.section[0].id,0)"></el-image>
        </div>
        <!-- 人气爆款右边部分 -->
        <div class="right" v-if="productHot.section.length > 1">
          <div v-for="(i, index2) in productHot.section" :key="index2">
            <ProductItem class="wrapItem" :obj="i" :type="0"/>
          </div>
        </div>
      </div>
      </IndexListSlot>
    </section>
  </div>
</template>

<script>
/**
 * 导入组件
 * IndexListSlot  插槽容器
 * ProductItem   单项渲染组件
 * CDItem     倒计时组件
 * CombItem  拼团组件
 */
import IndexListSlot from './components/indexSlot'
import ProductItem from './components/indexItem'
import CDItem from './components/countdownItem'
import NotFoundInfo from '@/components/NotFound/NotFoundInfo.vue'
import TypeSelectBox from './components/typeSelectBox.vue'
/**
 * 引入api
 */
import {
  getIndexPageList,
  getCanvasData,
  getCategory
} from '@/api/indexpage.js'

export default {
  components: {
    ProductItem,
    IndexListSlot,
    CDItem,
    NotFoundInfo,
    TypeSelectBox
  },
  data () {
    return {
      productNew: {
        'header': {
            'title': '新品首发',
            'bcText': 'THE NEW START',
            'interface': '/indexpage/productlist'
        },
        'section': [{'image': '', 'storeName': '', 'otPrice': ''}]
      },
      productHot: {
        'header': {
          'title': '人气爆款',
          'bcText': 'POPULAR STYLE',
          'interface': '/indexpage/productlist'
        },
        'section': [
          {'image': '', 'storeName': '', 'otPrice': ''}
        ]
      },
    
      banner: [],
      types: [],
      typeSelect: '',
      layoutData: [],
      loading: false,
      loading2: false
    }
  },
  created () {
    this.getIndexData()
    this.getCanvas()
    this.getCateList()
  },
  methods: {
    async getCanvas () {
      var res = await getCanvasData({ terminal: '4' })
      this.layoutData = JSON.parse(res.data.content.json)
      if (Array.isArray(this.layoutData)) {
        this.layoutData = this.layoutData.map(item => {
        return item.type
      })
      } else {
        this.$message.error('返回数据有误，请检查返回的数据格式！')
      }

        console.log("aa"+this.layoutData.indexOf('PcBanner'))
    },
    // 分类查询
    async getCateList () {
        var response = await getCategory()
         console.log('a:' + response.status)
        if (response.status === 200) {
            var res = response.data.content
            console.log('a:' + response.data.content)
            this.types = this.types.concat(
                res.map(item => {
                    return {
                        'value': `${item.id}`,
                        'label': `${item.cateName}`,
                        'children': item.children.map(i => {
                          return {
                            value: i.id,
                            label: i.cateName,
                            pic: i.pic
                          }
                        })
                    }
                })
            )
            this.typeSelect = this.types[0].value
        } else {
          this.$message.error(res.msg || '无分类信息！')
        }
    },
    showSelectPage () {},
    // 查询首页数据
    getIndexData () {
      this.loading = true
      getIndexPageList().then(res => {
          if (res.status === 200) {
            this.seletcVal(res.data)
            this.loading = false
          } else {
            this.$message.error(res.msg)
          }
        })
    },
    
    // 数据筛选
    seletcVal (obj) {
      // 新品首发
      this.productNew.section = obj.firstList.map((item, index) => {
        if (index < 4) {
          return item
        }
      }).filter(r => r)
      // 人气爆款
      this.productHot.section = this.filterRes(obj.likeInfo, 4)
    },
    // 只留下固定数量的数据
    filterRes (arr, num) {
      return arr.map((item, index) => {
        if (index < num) {
          return item
        }
      }).filter(r => r)
    },
    // 详情跳转
    toDetail (id, type) {
      this.$router.push({
        path: '/productDetail',
        query: {
          productId: id,
          type: type
        }
      })
    }
    
  }
}
</script>

<style lang="scss">
.indexAll{
  width: 100%;
  min-width: 1440px;
  .el-carousel__button{
    height: 2px;
    opacity: 1;
    background-color: #CCCCCC;
  }
  .is-active{
    .el-carousel__button{
      background-color: #333333;
    }
  }
  .carousel{
    max-width: 1172px;
    margin: auto;
    position: relative;
    .typeBox{
      display: flex;
      position: absolute;
      top: 0;
      left: 0;
      z-index: 10;
      height: 100%;
    }
  }
  .bcimg{
    width: 100%;
    height: 755px;
    float: left;
  }
  .indexContent{
    width: 1170px;
    margin: auto;
    padding-bottom: 70px;
    .blockItem{
      margin: 30px 0;
        .common{
          margin: 30px 0;
          display: flex;
          justify-content: space-between;
          .newPro{
            width: 278px;
          }
        }
        // 人气爆款 独有部分
        .hotItem{
          .left{
            width: 500px;
            height: 695px;
            background-color: #FFF;
            display: flex;
            flex-direction: column;
            .title{
                margin: 50px 0 0 50px;
                text-align: left;
                font-size: 30px;
                overflow: hidden;
                white-space:nowrap;
                text-overflow: ellipsis;
            }
            .price{
                margin: 30px 50px;
                color: #333333;
                font-size: 22px;
                font-weight: 500;
                .putCart{
                    float: right;
                    cursor: pointer;
                }
            }
            .imgPro{
              width: 375px;
              height: 375px;
              margin: auto;
              cursor: pointer;
            }
          }
          .right{
            width: 650px;
            height: 675px;
            display: flex;
            flex-wrap: wrap;
            justify-content: space-between;
            .wrapItem{
              width: 315px;
              background-color: #FFFFFF;
              padding-bottom: 10px;
              margin-bottom: 20px;
            }
          }
        }
        // 拼团
        .combItem{
          display: flex;
          flex-wrap: wrap;
          .comb{
            width: 550px;
            height: 290px;
          }
          .NoInfo{
              display: block;
              width: 100%;
            }
        }
        // 砍价专区
        .bargainBlock{
          display: flex;
          justify-content: space-between;
          img{
            cursor: pointer;
          }
          .left{
            width: 574px;
            height: 574px;
            background-color: #FFFFFF;
            img{
              width: 100%;
              padding: 49px 115px;
            }
            .des{
              margin: 60px 40px 0;
              color: #666666;
              font-size: 24px;
              font-family: PingFang SC;
              .title{
                display: inline-block;
                width: 312px;
                height: 30px;
                overflow: hidden;
                white-space:nowrap;
                text-overflow: ellipsis;
              }
              .pprice{
                float: right;
                text-align: center;
                color: #333333;
                font-size: 30px;
                font-family: PingFang SC;
                font-weight: 500;
                display: flex;
                flex-direction: column;
              }
            }
          }
          .right{
            width: 574px;
            height: 574px;
            // 砍价区右方可抽象内容样式
            .item{
              width: 574px;
              height: 277px;
              padding: 25px 29px;
              background-color: #FFFFFF;
              display: flex;
              img{
                width: 230px;
                height: 230px;
                margin-right: 58px;
              }
              .itemright{
                margin: 45px 0;
                .title{
                  font-size: 24px;
                  font-family: PingFang SC;
                  font-weight: 400;
                  color: #666666;
                  margin-bottom: 40px;
                }
                .pprice{
                  font-size: 22px;
                  font-family: PingFang SC;
                  font-weight: 500;
                  color: #333333;
                }
              }
            }
          }
          .prisp{
            font-size: 18px;
            font-weight: 400;
            color: #999999;
            text-decoration: line-through;
          }
        }
        // 秒杀
        .SeckItem{
          .seckillItem{
            display: flex;
            .seckProduct{
              width: 278px;
              margin: 10px;
            }
          }
          .seckillNoInfo{
            width: 100%;
            .NoInfo{
              margin: auto;
            }
          }
        }
    }
  }
}
</style>
