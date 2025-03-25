<template>
    <div class="userCenter banxin">
      <div class="leftMenu">
        <div class="leftMenuBox">
          <el-menu :default-active="defaultActive" active-text-color="#333333" text-color="#999999" router>
            <div class="mar-bot-40">
                <p class="fs20">个人中心</p>
                <el-divider></el-divider>
                <el-menu-item class="flex-start" index="userInfo"><span class="title">基本资料</span></el-menu-item>
                <el-menu-item class="flex-start" index="address"><span class="title">地址管理</span></el-menu-item>
            </div>
            <div>
                <p class="fs20">交易管理</p>
                <el-divider></el-divider>
                <el-menu-item class="flex-start" index="myOrder"><span class="title">我的订单</span></el-menu-item>
                <el-menu-item class="flex-start" index="myEvaluate"><span class="title">评价晒单</span></el-menu-item>
                <el-menu-item class="flex-start" index="favorites"><span class="title">我的收藏</span></el-menu-item>
            </div>
          </el-menu>
        </div>
      </div>
      <div class="userCenterInfo flex1">
        <router-view></router-view>
      </div>
    </div>
</template>
<script>
export default {
    name: 'userCenter',
    data () {
        return {}
    },
    computed: {
        defaultActive: function () {
        return this.$route.path.replace('/', '')
        }
    },
    // 修改背景色
    beforeCreate () {
        document.querySelector('body').setAttribute('style', 'background-color:#F6F6F8')
    },
    beforeDestroy () {
        document.querySelector('body').removeAttribute('style')
    }
}
</script>
<style lang="scss" scoped>
.userCenter{
    display: flex;
    margin-top: 30px;
    margin-bottom: 100px;
    .leftMenu{
        width: 200px;
        margin-right: 20px;
        .el-menu{
            padding: 30px;
            border: 0;
            .el-menu-item{
                font-size: 18px;
                padding: 0 !important;
            }
            >>> .el-menu-item:focus, .el-menu-item:hover {
                background: none;
            }
            .is-active{
                font-weight: bold;
                color: #333333;
            }
        }
    }
}
</style>
