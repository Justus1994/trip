<template>
<v-card class="animationcard card" ref="mycard" v-on:click="google">
  <v-img :src=node.urls.regular height="250px">
    <div>
      <div class="headline">{{node.location.city}}</div>
    </div>
  </v-img>
</v-card>
</template>

<script>
export default {
  name: "TripCard",
  props: ['node'],
  methods: {
    google() {
      let lat = this.node.location.position.latitude;
      let long = this.node.location.position.longitude;
      let url = 'http://maps.google.com/maps?t=k&q=loc:' + lat + '+' + long;
      //let url = 'http://maps.google.com/?q='+ lat + ',' + long;
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
