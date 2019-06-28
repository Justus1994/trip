import Vue from 'vue'
import App from './App'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import router from './router.js'
import fetch from './fetch-intercept'

Vue.use(Vuetify);

Vue.directive('darkmode',function(el, binding, vnode){

     var background = (binding.arg === 'background' ? '#1c1d21' : '#333');
     el.style.background  = binding.value ? background : '#fcfcfc';
     el.style.color = binding.value ? '#fcfcfc': '#131313';

 });

const store = Vue.observable({
    trips: null,
    snack: {
      msg: "",
      show: false,
      color: "red",
    }
  });

new Vue({
  el: '#app',
  router,
  data: {
    store
  },
  render: h => h(App)
});
