<template>
  <v-card class="animationcard card" ref="mycard" @mousedown="down" @mouseup="up">
    <v-img :src=node.urls.regular height="250px">
      <div>
        <div class="headline">{{cityOrCountry}}</div>
      </div>
    </v-img>
    <v-layout align-center justify-end>
      <div v-if="deleteButton == true">
        <v-btn color="accent" flat @click="deleteButton= false">Cancel</v-btn>
        <v-btn color="accent" flat @click="deleteNode">Delete</v-btn>
      </div>
    </v-layout>
  </v-card>
</template>

<script>
import fetch from '../fetchData.js'

export default {
  name: "TripCard",
  props: ['node', 'index', 'i'],
  computed:{
    cityOrCountry(){
      if (!this.node.location.hasOwnProperty("city")){
            return this.node.location.country
      }
        return this.node.location.city
    }
  },
  data() {
    return {
      delay: null,
      deleteButton: false
    }
  },
  methods: {
    deleteNode() {
      this.deleteButton = false
      fetch('trip/' + 0 + '/nodes/' + this.node.id, 'DELETE').then(json => {
        if (json.Nodes.length <= 0) {
          fetch('trip/' + this.$route.params.id, 'DELETE').then(this.$router.push('/'));
        } else {
          this.$root.$data.sharedState.trips[this.$route.params.id].Nodes = json.Nodes
        }
      })
    },
    down() {
      let that = this
      this.delay = window.setTimeout(function() { 
          that.deleteButton = true
      },1000);
    },
    up() {
        clearTimeout(this.delay)
        if(!this.deleteButton) {
          this.google()
          this.deleteButton = false
        }
    },
    google() {
      let url;
      /*
      gmaps URL scheme
      Open latitude & Longitude : http://maps.google.com/?q='+ lat + ',' + long
      Search a Place :            https://www.google.com/maps/search/?api=1&query=place
      */
      let lat = this.node.location.position.latitude;
      let long = this.node.location.position.longitude;
      if(lat === undefined &&  this.node.location.hasOwnProperty("city")){
        url = 'https://www.google.com/maps/search/?api=1&query=' + this.node.location.city
      }
      else{
        url = 'https://www.google.com/maps/search/?api=1&query=' + this.node.location.country
      }
      if(lat){
        url = 'http://maps.google.com/maps?t=k&q=loc:' + lat + '+' + long;
      }
      window.open(url)
    }
  },

}
</script>

<style scoped>
.card {
  margin: 25px;
}

.headline {
  margin: 5%;
  color: white;
  font: 900 20px Montserrat;
  letter-spacing: 2px;
  text-shadow: 0 10px 25px #000000;
  text-transform: uppercase;
}
</style>
