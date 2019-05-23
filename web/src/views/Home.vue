<template>
    <div>
        <div v-bind:key="index" v-for="(trip, index) in getTrips">
            <TripCard :trip="trip" :index="index"/>
        </div>
        <div class="text-xs-center">
            <div class='right'>
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
            v-on:click="drop"
        >
            <v-icon>add</v-icon>
        </v-btn>
      </div>
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
      render() {
        fetchTrip('trip/Malaysia','POST').then(json =>{
          console.log(json);
          fetchTrip('trip','GET').then(data => {
            this.$root.$data.sharedState.trips = data;
          });
        });
      },
      drop(){
        fetchTrip('trip/0','DELETE').then(data => {
          this.$root.$data.sharedState.trips = data;
        });
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
.text-xs-center{
  position: fixed;
  top: 90%;

}
.text-xs-center .right{
  float:right;
  margin: auto;
}
</style>
