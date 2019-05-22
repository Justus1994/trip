<template>
    <v-card class="card" v-on:click="showDetails">
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
    props: ['trip', 'index'],
    methods: {
      toggleLike() {
        console.log("toggle like")
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
