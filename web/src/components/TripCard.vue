<template>
<div>
  <v-card class="card" @click="showDetails">
    <v-img :src=trip.Nodes[0].urls.regular height="250px">
    </v-img>

    <v-card-title primary-title>
      <div>
        <div class="headline_tripcard">{{trip.Nodes[0].location.country}}</div>
        <span class="subhead_tripcard">{{trip.Nodes.length}} Places</span>
      </div>
      <v-layout align-center justify-end>
        <v-btn @click.stop="shareTrip" icon>
          <v-icon color="accent">share</v-icon>
        </v-btn>
      </v-layout>
    </v-card-title>
  </v-card>

  <v-snackbar class="snackbar" v-model="snackbar" top :timeout="4000">
      Trip was copied to clipboard
      <v-btn color="white" flat @click="snackbar = false">
        Close
      </v-btn>
    </v-snackbar>
</div>
</template>

<script>
import copyTripToClipboard from '../share.js'
export default {
  name: "TripCard",
  props: ['trip', 'index'],
  data() {
    return {
      snackbar: false
    }
  },
  methods: {
    shareTrip() {
      copyTripToClipboard(this.trip)
      this.snackbar = true
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
.snackbar {
  padding: 2rem;
}
</style>
