<template>
<div>
  <v-card class="card" color="white">
    <v-card-title 
      class=title 
      color="black"
      >
        {{node.location.title}}
        <v-layout
                align-center
                justify-end
              >
              <v-btn small icon v-on:click="deleteNode">
                <v-icon>delete</v-icon>
              </v-btn>
        </v-layout>
    </v-card-title>
    <v-img
      :src=node.urls.regular
      height="150px"
    >
    </v-img>
  </v-card>
</div>
</template>

<script>
import fetch from '../fetchData.js'
import store from '../store.js'

export default {
    name: "TripDetailsCard",
    props: ["node", "index", "i"],
    methods: {
      deleteNode() {
        console.log("node deleted")
        fetch('trip/'+ this.index +'/nodes/' + this.i, 'DELETE').then(json => store.data.trips[this.index] = json)
      }
    },
}
</script>

<style scoped>
 .card {
   padding: 0px;
 }
 .title {
   padding: 5px 10px;   
 }
</style>
