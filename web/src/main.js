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

Vue.directive('darkmode',function(el, binding, vnode){

     var background = (binding.arg === 'background' ? '#1c1d21' : '#333');
     el.style.background  = binding.value ? background : '#fcfcfc';
     el.style.color = binding.value ? '#fcfcfc': '#131313';

 })

export var store = {
  debug: true,
  state: {
      trips: {},
      darkmode: false,

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
