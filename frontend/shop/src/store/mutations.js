import * as types from './mutation-types'

const mutations = {
  setLoginVisible (state, bool) {
    state.loginVisible = bool
  },
  reset (state) {
    state.userInfo = {}
    state.cartNum = 0
  },
  [types.SET_USERINFO] (state, obj) {
    state.userInfo = obj
  },
  [types.SET_CARTNUM] (state, num) {
    state.cartNum = num
  }
}

export default mutations
