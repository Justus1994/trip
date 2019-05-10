<template>
  <v-app>
    <v-content>
     <div id="app">
      <router-view/>
      </div>
    </v-content>
  </v-app>
</template>

<script>

export default {
  name: 'App',
  components: {
  },
  methods: {
    newTrip() {
      console.log("new Trip clicked")
    }
  }
}
</script>

<style scoped>


</style>
