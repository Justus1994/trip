<template>
    <div>
        <div v-bind:key="index" v-for="(trip, index) in getTrips">
            <TripCard :trip="trip" :index="index"/>
        </div>
        <div class="text-xs-center">
        <v-btn
            fab
            color="accent"
            to='/NewTrip'
        >
            <v-icon>add</v-icon>
        </v-btn>

        <v-btn
            fab
            color="black"
            v-on:click="render"
        >
            <v-icon>add</v-icon>
        </v-btn>
        <v-btn
            fab
            color="black"
            v-on:click="logging"
        >
            <v-icon>add</v-icon>
        </v-btn>
        </div>
    </div>
</template>

<script>
import TripCard from '../components/TripCard.vue'
import {store} from '../main.js'
import fetchTrip from '../fetchData.js'
export default {
    components: {
        TripCard
    },
    methods: {
      render(){
        fetchTrip('trip','GET').then(data => {
          store.setTrips(data);
        });
      },
      logging(){
        console.log("this.trips",this.trips);
        console.log("store",this.$root.$data.sharedState.trips);
      }
    },
    computed:{
      getTrips: function(){
        return this.$root.$data.sharedState.trips
      }
    },
    data: function() {
      return {
        trips : this.$root.$data.sharedState.trips
      }
    },
}
</script>

<style>

</style>
