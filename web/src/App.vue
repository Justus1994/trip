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
  name: 'App',
  components: {
    BottomNavigation
  },
  methods: {
    newTrip() {
      console.log("new Trip clicked")
      
    }
  }
}
</script>

<style >
 h1 {
    margin: 10px
  }

</style>
