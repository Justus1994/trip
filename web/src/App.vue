<template>
  <v-app>
    <v-content v-if="!loading">
     <div id="app">
      <router-view/>
      </div>
    </v-content>
    <Loading v-else></Loading>
  </v-app>
</template>

<script>
import store from '../src/store.js'
import Loading from './components/Loading'

export default {
  name: 'App',
  components: {
    Loading
  },
  data() {
    return {
      loading: true
    }
  },
  beforeCreate() {
    fetch('api/auth',{
      headers:{
        'Authorization' : window.localStorage.getItem('Authorization-Token')
      }
    }).then(response => {
        response.text().then(token => {
          window.localStorage.setItem('Authorization-Token', token);
          /**
          * Simulate loading task :D
          */
          setTimeout(() =>{
            this.loading = false;
          },1050);
        });
    });

   },
}
</script>

<style >
</style>
