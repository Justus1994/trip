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
     if(binding.value){
       console.log(binding.arg);
       var background = (binding.arg === 'background' ? '#1c1d21' : '#333');
       el.style.background  = background;
       el.style.color = '#fcfcfc';
     }else{
       el.style.background = '#fcfcfc';
       el.style.color = '#131313';
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
});
