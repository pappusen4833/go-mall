import Cookies from 'js-cookie'
/**
 * Cookies.set('key', 'value')
 * Cookies.set('key', 'value', { expires: 7}) //设置7天过期
 * Cookies.set('key', 'value', { expires: 7, path: ''}) //设置对当前页面的路径
 */

const token = 'Yshop_token'

export function getToken () {
    return Cookies.get(token)
}

export function setToken (value) {
    return Cookies.set(token, value)
}

export function removeToken () {
    return Cookies.remove(token)
}
