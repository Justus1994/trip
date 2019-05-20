<template>
    <v-card class="card" to="/tripdetails">
          <v-img
            :src=trip.image
            height="250px"
          >
          </v-img>
  
          <v-card-title primary-title>
            <div>
              <div class="headline">{{trip.title}}</div>
              <span class="grey--text">{{trip.subtitle}}</span>
            </div>
            <v-layout
                align-center
                justify-end
              >
              <v-btn v-on:click="toggleLike" icon>
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
export default {
    name: "TripCard",
    props: ["trip"],
    methods: {
      toggleLike() {
        console.log("toggle like")
      }
    },  
}
</script>

<style scoped>
 .card {
   margin: 10px;
 }
</style>
