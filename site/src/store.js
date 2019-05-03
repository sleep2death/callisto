import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

console.log(process.env)

export default new Vuex.Store({
  state: {
    GITHUB_AUTH_URL: `https://github.com/login/oauth/authorize?client_id=${process.env.VUE_APP_CLIENT_ID}&redirect_uri=http://127.0.0.1:8081/auth`,
    resp: ''
  },
  mutations: {
    response (state, payload) {
      state.resp = payload
    }
  },
  actions: {
    response ({ commit }, payload) {
      commit('response', payload)
    }
  }
})
