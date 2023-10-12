/**
    付款金额组件
    传参arr
    示例:
    arr: [
        ...
        'count': 100,   //总计
        'postage': 10,  //邮费
        'activity': 40,
        'ticket': 40,
        ...
    ]
 */
<template>
  <div class="shouldPay">
      <div class="priceCom">商品合计：<span>￥{{ obj.totalPrice }}</span></div>
      <div class="priceCom">邮费：<span>￥{{ obj.totalPostage }}</span></div>
      <!-- <div class="priceCom">活动优惠：<span>-￥{{ obj.deductionPrice }}</span></div> -->
      <div class="priceCom">优惠卷：<span>-￥{{ obj.couponPrice }}</span></div>
      <div class="priceCom count">最终金额：<span>￥{{ shouldPay }}</span></div>
  </div>
</template>

<script>
export default {
    name: 'shouldPay',
    props: {
        obj: Object
    },
    computed: {
        shouldPay () {
            return this.obj.totalPrice + this.obj.totalPostage - this.obj.deductionPrice - this.obj.couponPrice
        }
    }
}
</script>

<style lang="scss">
.shouldPay{
    text-align: right;
    margin: 30px;
    .priceCom{
        font-size: 16px;
        color: #999999;
        margin-bottom: 18px;
        span{
            color: #F26E47;
        }
    }
    .count{
        padding: 15px 0;
        font-size: 18px;
        span{
            font-size: 30px;
        }
    }
}
</style>
