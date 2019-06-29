<template>
  <v-card class="card" @mousedown="down" @touchstart="down" @touchend="touchup" @mouseup="up">
    <v-img  class="img" :src=node.urls.regular height="250px">
      <div>
        <div class="headline">{{cityOrCountry}}</div>
      </div>
    </v-img>
    <v-layout v-darkmode="darkmode" align-center justify-end>
        <div v-darkmode="darkmode"  v-bind:class="[deleteButton? 'removeCardActive': 'removeCardDisable','removeCard']" >
          <v-btn v-darkmode="darkmode" flat @click="deleteButton=false">Cancel</v-btn>
          <v-btn v-darkmode="darkmode" flat @click="$emit('deleteNode',node)">Delete</v-btn>
        </div>
    </v-layout>
  </v-card>
</template>

<script>
export default {
  name: "NodeCard",
  props: ['darkmode','node'],
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
      deleteButton: false,
    }
  },
  methods: {
    down() {
      let that = this;
      this.delay = window.setTimeout(function() {
          that.deleteButton = true
      },800);
    },
    up(){
      clearTimeout(this.delay);
      if(!this.deleteButton) {
        this.deleteButton = false;
      }
    },
    touchup(){
      clearTimeout(this.delay);
    },
    google() {
      let url;
      /**
      * Gmaps URL scheme
      * Open latitude & Longitude : http://maps.google.com/?q='+ lat + ',' + long
      * Search a Place :            https://www.google.com/maps/search/?api=1&query=place
      */
      let lat = this.node.location.position.latitude;
      let long = this.node.location.position.longitude;
      if(!lat && this.node.location.hasOwnProperty("city")){
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
.removeCard{
  transition: all 0.4s cubic-bezier(0.18, 0.89, 0.32, 1.28);
  width: 100%;
  display: flex;
  justify-content: space-between;
  z-index: 0;
}
.removeCardActive{
  margin-top: 0;
}
.removeCardDisable{
  margin-top: -48px;
}
.card {
  margin: 1em;
}
.card .img {
  z-index: 8;
}
.headline {
  margin: 5%;
  color: rgba(255,255,255,0.9);
  font: 900 18px Montserrat;
  letter-spacing: 2px !important;
  text-shadow: 0 10px 25px rgba(0,0,0,0.5);
  text-transform: uppercase;
}
@media only screen and (min-width: 600px) {
  .card {
    margin: auto;
    max-width: 500px;
    margin-top: 1em;
    margin-bottom: 1em;
  }
}
</style>
