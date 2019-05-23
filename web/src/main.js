import Vue from 'vue'
import App from './App'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import router from './router.js'
import "babel-polyfill"

Vue.use(Vuetify, {
  theme: {
    primary: '#3f51b5',
    secondary: '#b0bec5',
    accent: '#5C1349',
    error: '#b71c1c'
  }
})


export var store = {
  debug: true,
  state: {
      trips: {}

  }
}


new Vue({
  el: '#app',
  router,
  data: {
    privateState: {},
    sharedState: store.state
  },
  render: h => h(App)
})
