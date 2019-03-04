import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    shop: []
  },
  getters: {
    foodItems (state) {
      return state.shop
    }
  },
  mutations: {
    addFoodShop (state, payload) {
      state.shop = [
        ...state.shop,
        payload
      ]
    }
  },
  actions: {
    addFoodShop ({ commit, state }, payload) {
      commit('addFoodShop', payload)
    }
    // todo 更改值,如果添加的时候id存在,那么count+1
  }
})
