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
          </v-card-title>
          
          <v-layout align-center justify-end>
            <v-btn flat color="accent"><v-icon>favorite</v-icon></v-btn>
            <v-btn flat><v-icon>share</v-icon></v-btn>
          </v-layout>
        </v-card>
</template>

<script>
export default {
    name: "TripCard",
    props: ["trip"]
}
</script>

<style scoped>
 .card {
   margin: 10px;
 }
</style>
