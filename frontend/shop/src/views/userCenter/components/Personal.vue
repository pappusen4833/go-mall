/**
    订单页面 左侧功能导航组件：个人中心、交易管理。。。
    渲染示例：
    list: [
        {
            'title': '', //导航部分的标题
            'list': [ //导航路由部分
                {
                    'des': '基本资料', //导航部分描述
                    'route': '/information' //跳转的路由
                }
                ...如上结构
            ]
        }
        ...如上结构
    ]
*/
<template>
  <div class="personalLeft">
      <ul class="personalNavCom" v-for='( item,index1 ) in this.personList' :key='index1'>
          <li class='personalTitle'>{{ item.title }}</li>
          <ul v-for="(i,index2) in item.list" :key="index2">
              <li class="personalList" :class="{person_active: isActive === (`${index1}`+`${index2}`)}" @click="NavTo(i.route,index1,index2)">{{ i.des }}</li>
          </ul>
      </ul>
  </div>
</template>

<script>
export default {
    name: 'personal',
    data () {
        return {
            personList: [
                {
                    'title': '个人中心',
                    'list': [
                        {
                            'des': '基本资料',
                            'route': '/personal/information'
                        },
                        {
                            'des': '地址管理',
                            'route': '/personal/address'
                        },
                        {
                            'des': '我的足迹',
                            'route': '/personal/history'
                        },
                        {
                            'des': '优惠卷',
                            'route': '/personal/coupon'
                        }
                    ]
                },
                {
                    'title': '交易中心',
                    'list': [
                        {
                            'des': '我的订单',
                            'route': '/personal/order'
                        },
                        {
                            'des': '评价晒单',
                            'route': '/personal/appraisal'
                        },
                        {
                            'des': '售后记录',
                            'route': '/personal/aftersales/record'
                        }
                    ]
                }
            ],
            isActive: '00'
        }
    },
    created () {
        this.findIndex(this.personList, this.$route.path)
    },
    methods: {
        findIndex (arr, route) {
            /**
             * 初始化导航栏
             * arr：导航栏数据
             * route：目标路由
            */
            for (var i = 0; i < arr.length; i++) {
                for (var j = 0; j < arr[i].list.length; j++) {
                    if (arr[i].list[j].route === route) {
                        this.isActive = `${i}${j}`
                    }
                }
            }
        },
        NavTo (arg, index1, index2) {
            /**
             * index1父级下标
             * index2子集下标
             */
            if (!arg || arg === this.$route.path) return
            this.$router.push(arg)
            this.isActive = `${index1}` + `${index2}`
        }
    }
}
</script>

<style lang="scss">
.personalLeft{
    max-width: 400px;
    min-width: 200px;
    height: 540px;
    padding: 10px 30px 0 30px;
    background-color: #F6F6F8;
    .personalNavCom{
        margin-bottom: 30px;
        .personalTitle{
            color: #333333;
            font-size: 20px;
            border-bottom: 1px solid #999;
            padding: 19px 0;
            cursor: default;
        }
        .personalList{
            padding-top: 30px;
            color: #999999;
            font-size: 18px;
            cursor: pointer;
        }
        .person_active{
            color: #333333;
        }
    }
}
</style>
