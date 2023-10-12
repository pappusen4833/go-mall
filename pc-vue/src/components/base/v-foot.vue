<template>
  <div class="foot-con">
    <div class="content">
        <div class="lab-group">
            <label class="fs24">客服电话</label>
            <label class="fs50 mar-top-40">{{ phone }}</label>
            <label class="col-ccc fs28 mar-top-10">{{ time }}</label>
            <div class="consult mar-top-40">咨询客服</div>
        </div>
        <div class="verticalline"></div>
        <div class="icon-group">
            <el-image class="gzh-icon" :src="img1.pic"></el-image>
            <label class="mar-top-20">{{img1.des}}</label>
        </div>
        <div class="icon-group">
            <el-image class="xcxm-icon" :src="img2.pic"></el-image>
            <label class="mar-top-20">{{img2.des}}</label>
        </div>
    </div>
    <div class="foot-box"><span style="font-family:arial;">&copy; </span>2016-2020YSHOP-GO版权所有，官网https://www.yixiang.co 并保留所有权利 </div>
  </div>
</template>
<script>
    export default {
        name: 'v-foot',
        props: {
            obj: Object
        },
        data () {
            return {
                phone: '400 800 2100',
                time: '9:00-23:00',
                img1: {
                    des: '微信公众号',
                    pic: 'static/images/gzh.png'
                },
                img2: {
                    des: '小程序码',
                    pic: 'static/images/xcxm.png'
                }
            }
        },
        watch: {
            obj: {
                handler (nVal, oVal) {
                    var data = nVal.componentContent.publicList
                    for (var i = 0; i < data.length; i++) {
                        if (data[i].title === '客服电话') { this.phone = data[i].des }
                        if (data[i].title === '营业时间') { this.time = data[i].des }
                        if (data[i].title === '微信公众号') {
                            this.img1.des = data[i].des
                            this.img1.pic = data[i].pic
                        }
                        if (data[i].title === '小程序码') {
                            this.img2.des = data[i].des
                            this.img2.pic = data[i].pic
                        }
                    }
                }
            }
        }
    }
</script>
<style scoped lang="scss">
.foot-con{
    height: 454px;
    min-width: 1440px;
    background-color: #333333;
    color: #ffffff;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    .content{
        display: flex;
        // width: 60%;
        // justify-content: space-between;
        width: 1240px;
        justify-content: space-around;
        .lab-group{
            display: flex;
            flex-direction: column;
            label{
                font-weight: 400;
            }
            .consult{
                width: 164px;
                height: 50px;
                color: #cccccc;
                border: 1px #cccccc solid;
                line-height: 50px;
                text-align: center;
                cursor: pointer;
            }
        }
        .icon-group{
            display: flex;
            flex-direction: column;
            align-items: center;
            label{
                font-size: 20px;
                font-weight: 400;
            }
        }
        .gzh-icon{
            height: 196px;
            width: 196px;
        }
        .xcxm-icon{
            height: 196px;
            width: 196px;
        }
        .verticalline{
            height: 238px;
            width: 1px;
            background-color:#666666;
        }
    }
}
.foot-box{
    margin-top: 80px;
}
</style>
