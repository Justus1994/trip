<template>
    <div>
        <div v-bind:key="index" v-for="(trip, index) in getTrips">
            <TripCard :trip="trip" :index="index"/>
        </div>

      <div class='centerBottom'>
        <v-btn
            fab
            color="accent"
            to='/NewTrip'
        >
            <v-icon>add</v-icon>
        </v-btn>
      </div>
    </div>
</template>

<script>
import TripCard from '../components/TripCard.vue'
import fetch from '../fetchData'
export default {
    components: {
        TripCard
    },
    created(){
      this.fetchTrips();
    },
    methods: {
      fetchTrips(){
        fetch('trip','GET').then(data => {
          this.$root.$data.sharedState.trips = data;
        });
      }
    },
    watch: {
    '$route': 'fetchTrips'
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
.centerBottom{
  position: fixed;
  left: 50%;
  transform: translateX(-50%);
  top: 90%;
}
</style>
