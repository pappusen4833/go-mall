<template>
    <el-dialog
        title="添加新地址"
        :visible.sync="addVisible"
        width="45%"
        center
        :before-close="addClose">
        <div>
            <div class="flex mar-bot-30 fs20">
                <div class="receiver flex1 mar-right-40">
                    <p class="mar-bot-10">收货人</p>
                    <el-input type="text" v-model="newAddress.real_name"></el-input>
                </div>
                <div class="phoneNumber flex1">
                    <p class="mar-bot-10">联系电话</p>
                    <el-input type="number" v-model="newAddress.phone"></el-input>
                </div>
            </div>
            <div class="mar-bot-30">
                <p class="mar-bot-10">所在区域</p>
                <el-cascader
                v-model="value.label"
                :options="options"
                :props="{ expandTrigger: 'hover' }"
                @change="handleChange"
                clearable
                filterable
                placeholder="搜索">
                </el-cascader>
            </div>
            <div class="detail mar-bot-30">
                <p class="mar-bot-10">详细地址</p>
                <el-input type="text" v-model="newAddress.detail"></el-input>
            </div>
            <div class="default display-align">
                <el-image v-if="newAddress.is_default" src="static/images/xz@2x.png" alt="" @click="selectDefault" />
                <el-image v-else src="static/images/wz@2x.png" alt="" @click="selectDefault" />
                <span class="mar-left-5">设为默认地址</span>
            </div>
        </div>
        <span slot="footer" class="dialog-footer">
        <el-button class="close" plain @click="addClose">取消</el-button>
        <el-button class="submit" :loading="btnLoading" plain @click="editAddress">保存</el-button>
        </span>
    </el-dialog>
</template>

<script>
import {
    editAddressData
} from '@/api/user.js'
export default {
    props: ['addVisible', 'options', 'newAddress', 'value'],
    data () {
        return {
            btnLoading: false
        }
    },
    watch: {
        options: {
            handler (nVal, oVal) {
                console.log(nVal, oVal)
            },
            deep: true
        }
    },
    created () {
        console.log(this.options)
    },
    methods: {
        // 添加/修改地址
        editAddress () {
            if (!this.newAddress.real_name) {
                return this.$message.warning('请输入收货人姓名')
            } else if (!this.newAddress.phone) {
                return this.$message.warning('请输入联系电话')
            } else if (!this.newAddress.address.cityId) {
                return this.$message.warning('请选择省市区')
            } else if (!this.newAddress.detail) {
                return this.$message.warning('请输入详细地址')
            }
            //this.btnLoading = true
            editAddressData(this.newAddress).then(res => {
                if (res.status === 200) {
                    this.$message.success(this.newAddress.id ? '修改成功' : '添加成功')
                    this.$emit('addClose', 1)
                } else {
                    this.$message.warning(res.msg)
                }
                //this.btnLoading = false
            })
        },
        // 选择省市区
        handleChange (val) {
            if (val.length) {
                this.options.forEach(i => {
                    if (i.value === val[0]) {
                        this.newAddress.address.province = i.label
                        i.children.forEach(j => {
                            if (j.value === val[1]) {
                                this.newAddress.address.city = j.label
                                this.newAddress.address.cityId = j.value
                                j.children.forEach(k => {
                                    if (k.value === val[2]) {
                                        this.newAddress.address.district = k.label
                                    }
                                })
                            }
                        })
                    }
                })
            } else {
                this.newAddress.address = {
                    city: null,
                    cityId: null,
                    province: null,
                    district: null
                }
            }
        },
        // 是否默认地址
        selectDefault () {
            this.newAddress.is_default = !this.newAddress.is_default
        },
        // 取消
        addClose () {
            this.$emit('addClose', 0)
        }
    }
}
</script>
