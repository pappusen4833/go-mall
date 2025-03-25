/**
    收件人信息组件
 */
<template>
<div class="recipient">
    <div class="title">收货信息</div>
    <div class="content">收货人: <span>{{ obj.realName }}</span></div>
    <div class="content">手机号: <span>{{ obj.userPhone }}</span></div>
    <div class="content">地&nbsp;&nbsp;&nbsp;&nbsp;址: <span>{{ obj.userAddress }}</span></div>
</div>
</template>

<script>
export default {
    name: 'recipient',
    props: {
        obj: Object
    }
}
</script>
