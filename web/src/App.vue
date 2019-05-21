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
   this.load(1050, this.auth, this.loadData);
  },
  name: 'App',
  components: {
    BottomNavigation
  },
  methods: {
    newTrip() {
      console.log("new Trip clicked")

    },
    load(ms, auth, loadData) {
      const startPoint = new Date().getTime();
      auth().then( token => {
        window.localStorage.setItem('Authorization-Token', token);
      });
      loadData().then(data => {
        store.data.trips = data;
        console.log(data)
      });
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
    async loadData() {
      let data = await fetchTrips('trip', 'GET');
      return data;
    }

  }
}
</script>

<style >
 h1 {
    margin: 10px
  }

</style>
