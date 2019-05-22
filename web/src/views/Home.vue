<template>
    <div>
        <div v-bind:key="index" v-for="(trip, index) in trips">
            <TripCard :trip="trip" :index="index"/>
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
