<template>
    <div class="addressManage">
        <h2 class="fs30">地址管理
            <span class="fr fs18 mar-top-10 cur-poi" @click="add">+新增地址</span>
        </h2>
        <el-divider></el-divider>
        <!-- 地址列表 -->
        <div class="addressList" v-loading="loading" v-if="!noAddress">
            <div class="addressBox flex mar-bot-20" v-for="(item) in addressList" :key="item.id">
                <div class="flex1">
                    <p class="receiver">{{item.realName}}</p>
                    <p class="default col-999" v-if="item.isDefault">默认地址</p>
                    <p class="default col-999" v-else><span class="cur-poi" @click="setDefault(item.id)">设为默认地址</span></p>
                </div>
                <div class="flex3">
                    <p class="phoneNumber">手机号码：{{item.phone}}</p>
                    <p class="detail">详细地址：{{item.province + item.city + item.district + item.detail}}</p>
                </div>
                <div class="flex1">
                    <p class="edit display-center">
                        <span class="cur-poi" @click="edit(item)"><img src="static/images/bj@2x.png" alt="" />编辑</span>
                    </p>
                    <p class="del display-center">
                        <span class="cur-poi" @click="del(item.id)"><img src="static/images/sc@2x.png" alt="">删除</span>
                    </p>
                </div>
            </div>
        </div>
        <div class="nothing display-column display-center" v-else>
            <img src="static/images/wsj@2x.png" alt="">
            <p class="fs18 mar-top-20">暂无数据哦～</p>
        </div>
        <!-- 添加/编辑收货地址 -->
        <AddAddress :addVisible="addVisible" :options="options" :newAddress="newAddress" :value="value" @addClose="addClose"></AddAddress>
    </div>
</template>
<script>
import AddAddress from '@/components/common/addAddress'

import {
    getCityData,
    getAddressData,
    setDefaultAddress,
    deleteAddress
} from '@/api/user.js'
export default {
    name: 'addressManage',
    components: {AddAddress},
    data () {
        return {
            addressList: [],
            newAddress: {
                real_name: null,
                phone: null,
                address: {
                    city: null,
                    cityId: null,
                    district: null,
                    province: null
                },
                id: null,
                detail: null,
                is_default: false
            },
            addVisible: false,
            loading: false,
            options: [],
            value: {
                label: []
            },
            noAddress: false
        }
    },
    mounted () {
        this.getAddressList()
        this.getCityList()
    },
    methods: {
        // 获取地址列表
        async getAddressList () {
            this.loading = true
            var res = await getAddressData()
            if (res.status === 200) {
                this.addressList = res.data
                if (res.data.length) {
                    this.noAddress = false
                } else {
                    this.noAddress = true
                }
            } else {
                this.$message.warning(res.msg)
            }
            this.loading = false
        },
        // 获取城市列表
        async getCityList () {
            var res = await getCityData()
            if (res.status === 200) {
                const json1 = JSON.parse(JSON.stringify(res.data).replace(/v/g, 'value'))
                const json2 = JSON.parse(JSON.stringify(json1).replace(/n/g, 'label'))
                const json3 = JSON.parse(JSON.stringify(json2).replace(/c/g, 'children'))
                json3.forEach(i => {
                    i.children.forEach(j => {
                        j.children.forEach(k => {
                            delete k.children
                        })
                    })
                })
                this.options = json3
            } else {
                this.$message.warning(res.msg)
            }
        },
        // 删除地址
        del (id) {
            this.$confirm('此操作将永久删除该地址, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                deleteAddress({id}).then(res => {
                    if (res.status === 200) {
                        this.$message.success('删除成功')
                        this.getAddressList()
                    } else {
                        this.$message.warning(res.msg)
                    }
                })
            }).catch(() => {})
        },
        // 设为默认地址
        async setDefault (id) {
            var res = await setDefaultAddress({id})
            if (res.status === 200) {
                this.getAddressList()
            } else {
                this.$message.warning(res.msg)
            }
        },
        // 编辑
        edit (item) {
            this.value.label = []
            this.options.forEach(i => {
                if (i.label === item.province) {
                    this.newAddress.address.province = i.label
                    this.value.label.push(i.value)
                    i.children.forEach(j => {
                        if (j.label === item.city) {
                            this.newAddress.address.city = j.label
                            this.newAddress.address.cityId = j.value
                            this.value.label.push(j.value)
                            j.children.forEach(k => {
                                if (k.label === item.district) {
                                    this.newAddress.address.district = k.label
                                    this.value.label.push(k.value)
                                }
                            })
                        }
                    })
                }
            })
            this.newAddress.real_name = item.realName
            this.newAddress.phone = item.phone
            this.newAddress.detail = item.detail
            this.newAddress.id = item.id
            this.newAddress.is_default = !(!item.isDefault)
            this.addVisible = true
        },
        // 重置数据
        reset () {
            this.newAddress = {
                real_name: null,
                phone: null,
                address: {
                    city: null,
                    cityId: null,
                    district: null,
                    province: null
                },
                detail: null,
                is_default: false
            }
            this.value.label = []
        },
        // 打开添加收货地址弹窗
        add () {
            this.addVisible = true
        },

        // 关闭添加收货地址弹窗
        addClose (x) {
            if (x) {
                this.getAddressList()
            }
            this.addVisible = false
            this.reset()
        }
    }
}
</script>
<style lang="scss" scoped>
.addressManage{
    padding: 30px;
    background-color: #fff;
    .addressList{
        min-height: 330px;
        .addressBox{
            padding: 35px 40px;
            line-height: 2;
            background-color: #F6F6F8;
            .edit,.del{
                img{
                    width: 14px;
                    height: 15px;
                    margin-right: 5px;
                }
            }
        }
    }
    .nothing{
        padding: 80px 0;
        img{
            width: 150px;
            height: 130px;
        }
    }
    >>>.el-dialog{
        .el-dialog__header{
            padding-top: 60px;
            .el-dialog__title{
                font-size: 26px;
            }
        }
        .el-dialog__body{
            padding: 25px 60px;
            box-sizing: border-box;
            p{
                font-size: 18px;
            }
            .default{
                img{
                    width: 18px;
                    height: 18px;
                    cursor: pointer;
                }
                span{
                    font-size: 18px;
                }
            }
            .addressList{
                .addressBox{
                    padding: 20px;
                    border: 1px solid #CCCCCC;
                    margin-bottom: 20px;
                    font-size: 18px;
                    p{
                        line-height: 1.7;
                    }
                }
                .active{
                    border-color: #F26E47;
                    background-color: #FFFBFA;
                }
            }
        }
        .el-dialog__footer{
            padding-bottom: 60px;
            .el-button{
                font-size: 18px;
                padding: 20px 75px;
                border-radius: 0;
                border: 1px solid #333333;
            }
            .close{
                color: #333333;
            }
            .submit{
                background-color: #333333;
                color: #FFFFFF;
            }
        }
    }
    >>>.el-input__inner{
        border-radius: 0;
        border-color: #333333;
    }
    >>>.distpicker-address-wrapper{
        select{
            width: 32.7%;
            border-radius: 0;
            border-color: #333333;
        }
    }
    >>>.el-cascader{
        .el-input .el-input__inner:focus, .el-cascader .el-input.is-focus .el-input__inner{
            border-color: #333333;
        }
        .el-input__inner:hover{
            border-color: #333333;
        }
    }
}
</style>
