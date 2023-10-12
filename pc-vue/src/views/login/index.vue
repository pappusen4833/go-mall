<template>
    <div class="login">
        <!-- 登录 -->
        <el-dialog
          title="账号登录"
          v-if="show=='login'"
          :visible.sync="loginVisible"
          :before-close="handleClose"
          :close-on-click-modal="false"
          :close-on-press-escape="false"
          width="35%"
          center>
          <div class="mar-top-40">
              <div class="phoneNumber">
                    <div v-if="activeName=='pwdLogin'" class="fs18">
                        <span>手机号
                          <!-- <span class="underline fr cur-poi" @click="codeLogin">短信验证码登录</span> -->
                        </span>
                    </div>
                    <div v-if="activeName=='codeLogin'" class="fs18"><span>手机号<span class="underline fr cur-poi" @click="passwordLogin">账号密码登录</span></span></div>
                    <input class="mar-top-20" v-model="phoneNumber" type="number">
              </div>
              <div v-if="activeName=='pwdLogin'" class="password mar-top-40">
                  <p class="fs18">密码</p>
                  <input class="mar-top-20" v-model="pwd" type="password">
              </div>
              <!-- 验证码登录 -->
              <div v-if="activeName=='codeLogin'" class="password mar-top-40">
                  <p class="fs18">手机验证码</p>
                  <div class="display-between mar-top-20">
                      <input style="width:60%;" v-model="authCode" type="text">
                      <span v-if="showTime" class="underline fs16" style="line-height: 40px;">{{time}}s</span>
                      <span v-else class="underline cur-poi fs16" style="line-height: 40px;" @click="getCode">获取手机验证码</span>
                  </div>
              </div>
              <div class="fs18 mar-top-20 cur-poi" @click="toRegister"><span>立即注册</span><span v-if="activeName=='pwdLogin'" class="fr cur-poi" @click="toForgot">忘记密码?</span></div>
              <div class="btn mar-top-40">
                  <el-button :loading="loading" plain @click="login">登录</el-button>
              </div>
          </div>
        </el-dialog>
        <!-- 注册 -->
        <el-dialog
          title="注册账户"
          v-if="show=='register'"
          :visible.sync="loginVisible"
          :before-close="handleClose"
          :close-on-click-modal="false"
          :close-on-press-escape="false"
          width="35%"
          center>
          <div class="mar-top-40">
              <div class="phoneNumber">
                  <div class="fs18"><span>手机号<span class="fr">已有账号，<span class="underline fr cur-poi" @click="toLogin">去登录</span></span></span></div>
                  <input class="mar-top-20" v-model="phoneNumber1" type="number">
              </div>
              <div class="password mar-top-40">
                  <p class="fs18">手机验证码</p>
                  <div class="display-between mar-top-20">
                      <input style="width:60%;" v-model="authCode1" type="text">
                      <span v-if="showTime1" class="underline fs16" style="line-height: 40px;">{{time1}}s</span>
                      <span v-else class="underline cur-poi fs16" style="line-height: 40px;" @click="getCode1">获取手机验证码</span>
                  </div>
              </div>
              <div class="password mar-top-40">
                  <p class="fs18">密码</p>
                  <input class="mar-top-20" v-model="pwd1" type="password">
              </div>
              <div class="password mar-top-40">
                  <p class="fs18">确认密码</p>
                  <input class="mar-top-20" v-model="pwdAgain" type="password">
              </div>
              <div class="clause mar-top-30 fs16 flex">
                  <div class="select cur-poi" :class="read?'active':''" @click="select"></div>
                  <span class="mar-left-5">创建账号代表您已接受YSHOP</span>
                  <span class="mar-left-5 underline cur-poi">隐私条款</span>
              </div>
              <div class="btn mar-top-40">
                  <el-button :loading="loading1" plain @click="register">立即注册</el-button>
              </div>
          </div>
        </el-dialog>
        <!-- 忘记密码 -->
        <el-dialog
          title="找回密码"
          v-if="show=='forgot'"
          :visible.sync="loginVisible"
          :before-close="handleClose"
          :close-on-click-modal="false"
          :close-on-press-escape="false"
          width="35%"
          center>
          <div class="mar-top-40">
              <div class="phoneNumber">
                  <div class="fs18"><span>手机号<span class="fr">已有账号，<span class="underline fr cur-poi" @click="toLogin">去登录</span></span></span></div>
                  <input class="mar-top-20" v-model="phoneNumber2" type="number">
              </div>
              <div class="password mar-top-40">
                  <p class="fs18">手机验证码</p>
                  <div class="display-between mar-top-20">
                      <input style="width:60%;" v-model="authCode2" type="text">
                      <span v-if="showTime2" class="underline fs16" style="line-height: 40px;">{{time2}}s</span>
                      <span v-else class="underline cur-poi fs16" style="line-height: 40px;" @click="getCode2">获取手机验证码</span>
                  </div>
              </div>
              <div class="password mar-top-40">
                  <p class="fs18">密码</p>
                  <input class="mar-top-20" v-model="pwd2" type="text">
              </div>
              <div class="password mar-top-40">
                  <p class="fs18">确认密码</p>
                  <input class="mar-top-20" v-model="pwdAgain1" type="text">
              </div>
              <div class="btn mar-top-40">
                  <el-button :loading="loading2" plain @click="getBack">确认</el-button>
              </div>
          </div>
        </el-dialog>
    </div>
</template>

<script>
import Cookie from 'js-cookie'
import {mapGetters, mapMutations} from 'vuex'
import {cartCount} from '@/api/cart.js'
import {
    getVerify,
    loginWeb,
    registerID,
    getUserInfo
} from '@/api/login.js'
export default {
    name: 'login',
    data () {
        return {
            show: 'login',
            activeName: 'pwdLogin',
            registerVisible: true,
            forgotVisible: true,
            loading: false,
            loading1: false,
            loading2: false,
            phoneNumber: null,
            phoneNumber1: null,
            phoneNumber2: null,
            authCode: null,
            authCode1: null,
            authCode2: null,
            pwd: '',
            pwd1: '',
            pwd2: '',
            pwdAgain: '',
            pwdAgain1: '',
            showTime: false,
            showTime1: false,
            showTime2: false,
            time: 60,
            time1: 60,
            time2: 60,
            timer: null,
            timer1: null,
            timer2: null,
            read: false,
            relus: /^1[13456789]\d{9}$/
        }
    },
    computed: {
        ...mapGetters([
          'loginVisible' // 登录注册弹窗开关
        ])
    },
    methods: {
        ...mapMutations({
          setUserInfo: 'SET_USERINFO', // 修改用户信息
          setCartNum: 'SET_CARTNUM' // 修改购物车商品数量
        }),
        // 倒计时
        countDown () {
            this.showTime = true
            clearInterval(this.timer)
            this.timer = setInterval(() => {
                if (this.time > 0) {
                    this.time--
                } else {
                    this.showTime = false
                    clearInterval(this.timer)
                    this.time = 60
                }
            }, 1000)
        },
        countDown1 () {
            this.showTime1 = true
            clearInterval(this.timer1)
            this.timer1 = setInterval(() => {
                if (this.time1 > 0) {
                    this.time1--
                } else {
                    this.showTime1 = false
                    clearInterval(this.timer1)
                    this.time1 = 60
                }
            }, 1000)
        },
        countDown2 () {
            this.showTime2 = true
            clearInterval(this.timer2)
            this.timer2 = setInterval(() => {
                if (this.time2 > 0) {
                    this.time2--
                } else {
                    this.showTime2 = false
                    clearInterval(this.timer2)
                    this.time2 = 60
                }
            }, 1000)
        },

        // 切换验证码登录
        codeLogin () {
            this.$message.warning('暂未接入短信登录')
            // this.activeName = 'codeLogin'
        },

        // 切换密码登录
        passwordLogin () { this.activeName = 'pwdLogin' },

        // 登录获取验证码
        getCode () {
            if (this.phoneNumber === null) {
                return this.$message.warning('请输入手机号码')
            } else if (!this.relus.test(this.phoneNumber)) {
                return this.$message.warning('手机格式错误')
            }
            getVerify({
                phone: this.phoneNumber,
                type: 'login' // bind绑定手机短信 login登陆短信 register注册短信
            }).then(res => {
                 if (res.status === 200) {
                    this.$message.success('验证码发送成功')
                    this.countDown()
                } else {
                    this.$message.warning(res.msg)
                }
            })
        },
        // 注册获取验证码
        getCode1 () {
            if (this.phoneNumber1 === null) {
                return this.$message.warning('请输入手机号码')
            } else if (!this.relus.test(this.phoneNumber1)) {
                return this.$message.warning('手机格式错误')
            }
            getVerify({
                phone: this.phoneNumber1,
                type: 'register'
            }).then(res => {
                if (res.status === 200) {
                    this.$message.success('验证码发送成功')
                    this.countDown1()
                } else {
                    this.$message.warning(res.msg)
                }
            })
        },

        // 去注册
        toRegister () { this.show = 'register' },

        // 忘记密码
        toForgot () { this.show = 'forgot' },

        // 去登录
        toLogin () { this.show = 'login' },

        // 登录
        login () {
            if (this.pwd.length < 4 || this.pwd.length > 16) {
                return this.$message.warning('请输入4至16位的密码')
            }
            this.loading = true
            loginWeb({
                username: this.phoneNumber,
                password: this.pwd
            }).then(res => {
                if (res.status === 200) {
                    Cookie.set('Yshop_token', res.data.token)
                    this.$message.success('登录成功')
                    this.getUserInfo()
                    this.getCartNum()
                    this.handleClose()
                    location.reload()
                } else {
                    this.$message.warning(res.msg)
                }
            })
            this.loading = false
        },

        // 获取购物车商品数量
        async getCartNum () {
            var res = await cartCount()
            if (res.status === 200) {
                this.setCartNum(res.data.count)
            } else {
                this.$message.warning(res.msg)
            }
        },

        // 隐私协议是否已读
        select () { this.read = !this.read },

        // 注册
        register () {
            if (!this.read) {
                return this.$message.warning('请先阅读YSHOP隐私条款')
            } else if (this.pwd1.length < 4 || this.pwd1.length > 16) {
                return this.$message.warning('请输入4到16位的密码')
            } else if (this.pwd1 !== this.pwdAgain) {
                return this.$message.warning('两次密码不一致')
            }
            this.loading1 = true
            registerID({
                account: this.phoneNumber1,
                captcha: this.authCode1,
                password: this.pwd1
            }).then(res => {
                if (res.status === 200) {
                    this.$message.success('注册成功')
                    setTimeout(() => {
                        this.toLogin()
                    }, 1000)
                } else {
                    this.$message.warning(res.msg)
                }
            })
            this.loading1 = false
        },

        // 修改密码
        getBack () {
            this.loading = true
        },

        // 获取用户信息
        async getUserInfo () {
            var res = await getUserInfo()
            if (res.status === 200) {
                this.$store.commit('setLoginVisible', false)
                this.setUserInfo(res.data)
                this.$router.go(0)
            }
        },

        // 关闭登录注册弹窗
        handleClose () { this.$store.commit('setLoginVisible', false) }
    }
}
</script>
<style lang="scss" scoped>
>>>.el-dialog{
    .el-dialog__header{
        padding: 60px 20px 10px;
        .el-dialog__title{
            font-size: 30px;
        }
        .el-icon-close{
            font-size: 25px;
        }
        .el-icon-close:hover{
            color: #333333;
        }
    }
    .el-dialog__body{
        padding: 0 22% 60px;
        input{
            width: 100%;
            border: 1px solid #333333;
            padding: 10px 15px;
            box-sizing: border-box;
        }
        .btn{
            .el-button{
                width: 100%;
                height: 60px;
                background-color: #333333;
                box-sizing: border-box;
                color: #FFFFFF;
                font-size: 18px;
                border-radius: 0;
            }
        }
        .clause{
            .select{
                width: 16px;
                height: 16px;
                background-image: url('../../../static/images/wz@2x.png');
                background-size: 100% 100%;
            }
            .active{
                background-image: url('../../../static/images/xz@2x.png');
            }
        }
    }
}
</style>
