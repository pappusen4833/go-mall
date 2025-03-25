<template>
    <div class="userInfo" v-loading="loading">
        <h2 class="fs30">基本资料</h2>
        <el-divider></el-divider>
        <el-form :ref="formData" :model="formData" label-width="80px">
            <el-form-item label="头像">
                <el-upload
                  class="avatar-uploader"
                  :action="action"
                  :show-file-list="false"
                  :on-success="handleAvatarSuccess">
                  <img v-if="formData.avatar" :src="formData.avatar" class="avatar">
                  <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                  <img class="add" src="../../../../static/images/xjicon@2x.png" alt="">
                </el-upload>
            </el-form-item>
            <el-form-item label="用户ID">
                <span>{{formData.uid}}</span>
            </el-form-item>
            <el-form-item label="昵称">
                <el-input v-model="formData.nickname"></el-input>
            </el-form-item>
            <!-- <el-form-item label="性别">
                <el-radio-group v-model="formData.gender">
                    <el-radio-button label="男"></el-radio-button>
                    <el-radio-button label="女"></el-radio-button>
                </el-radio-group>
            </el-form-item>
            <el-form-item label="出生日期">
                <el-date-picker
                    v-model="formData.birthday"
                    type="date"
                    placeholder="选择日期">
                </el-date-picker>
            </el-form-item> -->
            <el-form-item label="手机号">
                <span>{{formData.phone}}</span>
                <!-- <el-input v-model="formData.phone"></el-input> -->
            </el-form-item>
            <el-form-item class="mar-top-60">
                <el-button class="btn close" plain @click="getUserInfo">取消</el-button>
                <el-button class="btn submit mar-left-30" plain :loading="btnLoading" @click="onSubmit">保存</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>
<script>
import {uploadURL} from '@/api/user.js'
import {
    getUserInfo,
    editUserInfo
} from '@/api/login.js'
export default {
    name: 'userInfo',
    data () {
        return {
            formData: {},
            action: uploadURL,
            imageUrl: '',
            loading: false,
            btnLoading: false
        }
    },
    mounted () {
        this.getUserInfo()
    },
    methods: {
        // 获取用户信息
        async getUserInfo () {
            this.loading = true
            var res = await getUserInfo()
            if (res.status === 200) {
                this.formData = res.data
                this.loading = false
            } else {
                this.$message.warning(res.msg)
            }
        },
        // 上传图片
        handleAvatarSuccess (res) {
            this.formData.avatar = res.link
        },
        // 修改
        onSubmit () {
            this.btnLoading = true
            editUserInfo({
                avatar: this.formData.avatar,
                nickname: this.formData.nickname
            }).then(res => {
                if (res.status === 200) {
                    this.$message.success('修改成功')
                } else {
                    this.$message.warning(res.msg)
                }
            })
            this.btnLoading = false
        }
    }
}
</script>
<style lang="scss" scoped>
.userInfo{
    padding: 30px;
    background-color: #fff;
    >>>.el-form{
        .el-form-item{
            .el-form-item__label{
                text-align-last: left;
            }
            .el-input{
                width: 240px;
            }
            .el-input__inner{
                border-radius: 0;
                border-color: #CCCCCC;
            }
            .el-radio-button__inner:hover{
                color: #606266;
            }
            .el-radio-button__inner{
                width: 120px;
                border-radius: 0;
                box-shadow: -1px 0 0 0 #DCDFE6;
            }
            .el-radio-button__orig-radio:checked+.el-radio-button__inner{
                background-color: #CCCCCC;
                border-color: #DCDFE6;
            }
            .el-radio-button__orig-radio:checked+.el-radio-button__inner:hover{
                color: #FFFFFF;
            }
            .el-upload {
                width: 60px;
                height: 60px;
                border: 1px dashed #d9d9d9;
                border-radius: 6px;
                cursor: pointer;
                position: relative;
                border-radius: 30px;
                .avatar {
                    width: 60px;
                    height: 60px;
                    display: block;
                    border-radius: 30px;
                }
                .avatar-uploader-icon {
                  font-size: 20px;
                  color: #8c939d;
                  width: 60px;
                  height: 60px;
                  line-height: 60px;
                  text-align: center;
                }
                .add{
                    position: absolute;
                    bottom: -5px;
                    right: -5px;
                    width: 24px;
                    height: 22px;
                }
            }
            .btn{
                width: 130px;
                height: 50px;
                border: 1px solid #333333;
                box-sizing: border-box;
                border-radius: 0;
                font-size: 20px;
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
}
</style>
