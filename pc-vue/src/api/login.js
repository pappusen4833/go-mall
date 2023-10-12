import request from '@/axios/request.js'

// 创建账号
export function registerID (data) {
    // {
    //     "account": "", 手机号码
    //     "captcha": "", 验证码
    //     "inviteCode": "", 邀请码
    //     "password": "", 密码
    //     "spread": "" 分销绑定关系的ID
    // }
    return request({
        url: '/register',
        method: 'post',
        data
    })
}
export function getVerify (data) {
    // {
    //     "phone": "", 手机号码
    //     "type": "" 短信类型 bind绑定手机短信 login登陆短信 register注册短信
    // }
    return request({
        url: '/register/verify',
        method: 'post',
        data
    })
}

// H5登录授权
export function loginWeb (data) {
    // {
    //     "password": "",必填
    //     "spread": "",
    //     "username": ""
    // }
    return request({
        url: '/login',
        method: 'post',
        data
    })
}

// 退出登录
export function logoutWeb () {
    return request({
        url: '/auth/logout',
        method: 'post'
    })
}

// 获取用户信息
export function getUserInfo () {
    return request({
        url: '/userinfo',
        method: 'get'
    })
}

// 修改用户信息
/**
    avatar: this.formData.avatar,
    nickname: this.formData.nickname
*/
export function editUserInfo (data) {
    return request({
        url: '/user/edit',
        method: 'post',
        data
    })
}
