<template>
  <div class="countdownItem">
    <span>倒计时中</span>
    <ul class="cdTime">
        <li class="timeItem">12</li>
        <span>:</span>
        <li class="timeItem">12</li>
        <span>:</span>
        <li class="timeItem">12</li>
    </ul>
  </div>
</template>

<script>
export default {
    props: {}
}
</script>

<style lang="scss">
.countdownItem{
    background-color: #333333;
    text-align: center;
    color: #FFFFFF;
    .cdTime{
        display: flex;
        span{
            line-height: 33px;
        }
        .timeItem{
            width: 33px;
            height: 33px;
            line-height: 33px;
            color: #000;
            background: #FFFFFF;
        }
    }
}
</style>
