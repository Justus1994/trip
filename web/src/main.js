import Vue from 'vue'
import App from './App'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import router from './router.js'

Vue.use(Vuetify, {
  theme: {
    primary: '#f5f5f5',
    dark: '#333',
    light: '#f5f5f5',
    error: '#b71c1c'
  }
});


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
});
