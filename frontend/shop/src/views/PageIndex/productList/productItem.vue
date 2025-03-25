<template>
  <div class="productItem">
      <el-image class="imgTit" :src="obj.image" alt="商品图片" @click="toDetail(obj.id,0)"></el-image>
      <div class="name">{{ obj.storeName }}</div>
      <div class="buyIt">
          <div class="price">￥{{ obj.price }}</div>
          <div class="cart" @click="toDetail(obj.id,0)">
            <svg
          t="1618572221481" class="icon" viewBox="0 0 1170 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1604" width="16" height="16"><path d="M1119.085714 138.313143H269.165714L239.177143 42.861714a60.928 60.928 0 0 0-58.514286-42.715428H40.96a40.594286 40.594286 0 1 0 0 81.115428h124.342857l229.668572 727.04h656.822857a40.521143 40.521143 0 1 0 0-81.115428H455.68l-33.645714-105.398857 705.828571-118.637715a51.2 51.2 0 0 0 42.422857-50.029714V189.001143a51.2 51.2 0 0 0-51.2-50.688z m-587.337143 707.876571a88.795429 88.795429 0 1 0 89.965715 88.795429 89.380571 89.380571 0 0 0-89.965715-88.795429z m430.08 0a88.795429 88.795429 0 1 0 89.965715 88.795429 89.380571 89.380571 0 0 0-89.965715-88.795429z" fill="#333333" p-id="1605"></path></svg>
          </div>
      </div>
  </div>
</template>

<script>
import {
  cartAdd
} from '@/api/cart.js'
export default {
    props: {obj: Object},
    methods: {
        toDetail (id, type) {
            this.$router.push({
                path: '/productDetail',
                query: {
                productId: id,
                type: type
                }
            })
        },
        toCart (id) {
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
                } else {
                    this.$message.error(res.msg)
                }
            })
        }
    }
}
</script>

<style lang="scss">
.productItem{
    .imgTit{
        // width: 100%;
        width: 278px;
        height: 278px;
        min-width: 278px;
        min-height: 278px;
        cursor: pointer;
    }
    .name{
        margin: 20px;
        font-size: 24px;
        font-family: PingFang SC;
        font-weight: 400;
        color: #666666;
        overflow: hidden;
        white-space:nowrap;
        text-overflow: ellipsis;
    }
    .buyIt{
        margin: 0 20px;
        font-size: 22px;
        font-family: PingFang SC;
        font-weight: 500;
        color: #333333;
        .price{
            display: inline-block;
        }
        .cart{
            float: right;
            cursor: pointer;
        }
    }
}
</style>
