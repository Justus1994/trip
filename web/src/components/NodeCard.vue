<template>
<v-card class="animationcard card" ref="mycard" class="card" v-on:click="google">
  <v-img :src=node.urls.regular height="250px">
    <div>
      <div class="headline">{{cityOrCountry}}</div>
    </div>
  </v-img>
</v-card>
</template>

<script>
export default {
  name: "TripCard",
  props: ['node', 'index', 'i'],
  created(){
    console.log(this.node.location.position.latitude);
  },
  computed:{
    cityOrCountry(){
      if (!this.node.location.hasOwnProperty("city")){
            return this.node.location.country
      }
        return this.node.location.city
    }
  },
  methods: {
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
        url =   url = 'https://www.google.com/maps/search/?api=1&query=' + this.node.location.country
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
