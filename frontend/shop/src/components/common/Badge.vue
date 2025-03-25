/**
    右上角角标组件
 */
<template>
  <div class="zxBadge">
      <sup class="badge red" :class="bgwidth" v-if="sum != 0">{{ sum }}</sup>
  </div>
</template>

<script>
export default {
    name: 'zxBadge',
    props: {
        num: Number,
        size: String
    },
    computed: {
        sum () {
            return this.num > 99 ? '99+' : this.num
        },
        bgwidth () {
            return !this.size ? 'small' : this.size
        }
    }
}
</script>

<style lang="scss">
.zxBadge{
    text-align: center;
    .badge{
        border-radius: 48px;
        padding: 0 3px;
    }
    .red{
        color: #FFFFFF;
        background-color: #FF0000;
    }
    .big{
        width: 48px;
        height: 48px;
        line-height: 32px;
    }
    .regular{
        font-size: 14px;
        width: 32px;
        height: 32px;
        line-height: 32px;
    }
    .small{
        font-size: 12px;
        width: 18px;
        height: 18px;
        line-height: 18px;
    }
}
</style>
