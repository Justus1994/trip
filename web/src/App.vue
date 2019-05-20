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
export default {
  beforeMount() {
   this.load(1050, this.auth);
  },
  name: 'App',
  components: {
    BottomNavigation
  },
  methods: {
    newTrip() {
      console.log("new Trip clicked")

    },
    load(ms, auth) {
      const startPoint = new Date().getTime();
      auth().then( token => {
        window.localStorage.setItem('Authorization-Token', token);
      });
      while (new Date().getTime() - startPoint <= ms) {
        /* wait */
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

    }

  }
}
</script>

<style >
 h1 {
    margin: 10px
  }

</style>
