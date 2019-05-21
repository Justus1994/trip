<template>
    <div>
        <div v-bind:key="i" v-for="(trip, i) in trips">
            <TripCard v-bind:trip="trip"/>
        </div>
        <div class="text-xs-center">
        <v-btn
            fab
            color="accent"
            to="/newTrip"
        >
            <v-icon>add</v-icon>
        </v-btn>
        </div>
    </div>
</template>

<script>
import TripCard from '../components/TripCard.vue'
import store from '../store.js'

export default {
    components: {
        TripCard
    },
    data() {
      return {
          trips: store.data.trips
      }
    },
}
</script>

<style>

</style>
