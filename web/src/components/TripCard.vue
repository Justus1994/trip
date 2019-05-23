<template>
    <v-card class="card" v-on:click="toggleLike">
          <v-img
            :src=trip.Nodes[0].urls.regular
            height="250px"
          >
          </v-img>

          <v-card-title primary-title>
            <div>
              <div class="headline">{{trip.Nodes[0].location.city}}</div>
              <span class="grey--text">{{trip.Nodes[0].location.country}}</span>
            </div>
            <v-layout
                align-center
                justify-end
              >
              <v-btn  icon>

                <v-icon color="accent">favorite</v-icon>
              </v-btn>

              <!--v-btn icon>
                <v-icon>share</v-icon>
              </v-btn-->
            </v-layout>
          </v-card-title>
    </v-card>
</template>

<script>
import {store} from '../main.js'
export default {
    name: "TripCard",
    props: ['trip', 'index'],
    methods: {
      toggleLike() {
        console.log("toogle");
        this.$root.$data.sharedState.trips =  [{
            Nodes : [{
                location: {
                city : "Schwede",
                country: "deutschland",
                title : "astra <3"
              },
              urls: {
                regular: "https://upload.wikimedia.org/wikipedia/commons/3/31/Hamburg_Jungfernstieg_%281890-1900%29.jpg"
              }
            }]

        }]

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
</style>
