<template>
    <div>
        <h1>Your upcoming trips</h1>
        <div v-bind:key="trip.id" v-for="trip in trips">
            <TripCard v-bind:trip="trip"/>
        </div>
        <v-btn
            fab
            color="accent"
            to="/newTrip"
        >
            <v-icon>add</v-icon>
        </v-btn>
    </div>
</template>

<script>
import TripCard from '../components/TripCard.vue'
export default {
    components: {
        TripCard
    },
    data() {
      return {
          trips: [
              {
                id: 1,
                title: "Moraine Lake",
                subtitle: "Canada",
                image: "https://images.unsplash.com/photo-1516471835429-167f83503f4b?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=975&q=80"
              },
              {
                id: 2,
                title: "Roys Peak",
                subtitle: "New Zealand",
                image: "https://images.unsplash.com/photo-1496458590512-56d2688442b1?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1950&q=80"
              },
              {
                id: 3,
                title: "Mijas",
                subtitle: "Spain",
                image: "https://images.unsplash.com/photo-1507415953574-2aadbf10e38a?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=958&q=80"
              }
          ]
      }
    },
}
</script>

<style>

</style>
