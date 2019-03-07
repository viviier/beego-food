import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/element.js'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import 'normalize.css'

// component
import Header from '@/components/Header.vue'

Vue.config.productionTip = false
Vue.use(ElementUI)
Vue.component('beego-header', Header)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
