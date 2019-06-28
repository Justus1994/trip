<template>
<div >
  <v-card v-darkmode="darkmode" class='card' @click="showDetails">
    <v-img :src=trip.Nodes[0].urls.regular height="250px">
    </v-img>
    <v-card-title primary-title>
      <div>
        <div class="headline_tripcard">{{trip.Nodes[0].location.country}}</div>
        <span class="subhead_tripcard">{{trip.Nodes.length}} Places</span>
      </div>
      <v-layout align-center justify-end>
        <v-btn v-darkmode="darkmode" @click="$emit('shareTrip',index)" icon>
          <v-icon>share</v-icon>
        </v-btn>
      </v-layout>
    </v-card-title>
  </v-card>
</div>
</template>

<script>
import copyTripToClipboard from '../share.js'
export default {
  name: "TripCard",
  props: ['darkmode','trip', 'index'],
  methods: {
    shareTrip() {
      copyTripToClipboard(this.trip)

    },
    showDetails() {
      this.$router.push('/tripdetails/' + this.index);
    }
  },
}
</script>

<style scoped>
  .card {
    margin: 1em;
  }
  .headline_tripcard {
    font: 22px Montserrat;
    font-weight: 300;
  }
  .subhead_tripcard {
    font: 15px Montserrat;
  }
  @media only screen and (min-width: 600px) {
    .card {
      margin: auto;
      max-width: 500px;
      margin-top: 1em;
      margin-bottom: 1em;
      position: relative;
    }
  }
</style>
