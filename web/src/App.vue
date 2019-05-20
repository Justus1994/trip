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
    this.delay(1200);
  },
  name: 'App',
  components: {
    BottomNavigation
  },
  methods: {
    newTrip() {
      console.log("new Trip clicked")

    },
    delay(ms) {
      const startPoint = new Date().getTime();
      while (new Date().getTime() - startPoint <= ms) {
        /* wait */
      }
    }
  }
}
</script>

<style >
 h1 {
    margin: 10px
  }

</style>
