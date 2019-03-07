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
    },
    uploadFoodShop (state, payload) {
      state.shop = [
        ...payload
      ]
    }
  },
  actions: {
    addFoodShop ({ commit, state }, payload) {
      let reState = state.shop
      let hasItem = false
      reState.map(item => {
        if (item.id === payload.id) {
          item.count += 1
          hasItem = true
        }
        return item
      })

      if (hasItem) {
        commit('uploadFoodShop', reState)
      } else {
        payload = {
          ...payload,
          count: 1
        }
        commit('addFoodShop', payload)
      }
    },
    changeFoodShop ({ commit, state }, payload) {
      let reState = state.shop
      reState.map(item => {
        if (item.id === payload.id) {
          item.count = payload.count
        }
        return item
      })
      commit('uploadFoodShop', reState)
    },
    clearFoodShop ({ commit }) {
      commit('uploadFoodShop', [])
    }
  }
})
