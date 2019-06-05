<template>
<v-card class="card" v-on:click="showDetails">
  <v-img :src=trip.Nodes[0].urls.regular height="250px">
  </v-img>

  <v-card-title primary-title>
    <div>
      <div class="headline_tripcard">{{trip.Nodes[0].location.country}}</div>
      <span class="subhead_tripcard">{{trip.Nodes.length}} Places</span>
    </div>
    <v-layout align-center justify-end>
      <v-btn v-on:click="toggleLike" icon>
        <v-icon color="accent">favorite</v-icon>
      </v-btn>
    </v-layout>
  </v-card-title>
</v-card>
</template>

<script>
export default {
  name: "TripCard",
  props: ['trip', 'index'],
  methods: {
    toggleLike() {
      console.log("toogle");
    },
    showDetails() {
      this.$router.push('/tripdetails/' + this.index)
    }
  },
}
</script>

<style scoped>
.card {
  margin: 10px;
}
.headline_tripcard {
  font: 200 22px Montserrat !important;
}
.subhead_tripcard {
  font: 100 15px Montserrat !important;
  color: black;
}
</style>
