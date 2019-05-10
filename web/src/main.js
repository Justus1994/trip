import Vue from 'vue'
import App from './App'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import VueRouter from 'vue-router'

Vue.use(Vuetify)
Vue.use(VueRouter)

new Vue({
  el: '#app',
  render: h => h(App)
})
