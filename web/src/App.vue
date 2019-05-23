<template>
  <v-app>
    <v-content>
     <div id="app">
      <router-view/>
      </div>
    </v-content>
    <BottomNavigation>
  </v-app>
</template>

<script>
import BottomNavigation from './components/BottomNavigation.vue'
import store from '../src/store.js'
import fetchTrips from './fetchData.js'

export default {
  beforeMount() {
     this.load(1050, this.auth);
     fetchTrips('trip','GET').then(data => {
       this.$root.$data.sharedState.trips = data;
     });
   },
  name: 'App',
  components: {
    BottomNavigation
  },
  methods: {
      load(ms, auth) {
        const startPoint = new Date().getTime();
        auth().then( token => {
          window.localStorage.setItem('Authorization-Token', token);
        });
        while (new Date().getTime() - startPoint <= ms) {
        }
      },
      async auth () {
        let response = await fetch('api/auth',{
          headers:{
            'Authorization' : window.localStorage.getItem('Authorization-Token')
          }
        });
        let token = await response.text();
        console.log(token)
        return token;
      },
  }
}
</script>

<style >
 h1 {
    margin: 10px
  }

</style>
