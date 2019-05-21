<template>
    <div>
        <h1>New Trip</h1>
        <v-flex xs12 sm6 md3 class="search">
            <v-text-field 
              label="Solo"
              prepend-inner-icon="search"
              placeholder="Search for a country"
              solo
              color="accent"
              v-model="city"
            ></v-text-field>
          </v-flex>
          <div class="text-xs-center">
            <v-btn v-on:click="getNodes" round class="text-xs-center" color="accent" >Show Places</v-btn>
          </div>
    </div>
</template>

<script>
import fetch from '../fetchData.js'

export default {
    name: "NewTrip",
    data () {
      return {
        city: '',
        trips: {}
      }
    },    
    methods: {
      getNodes() {
        console.log(this.city)
        fetch('trip/germany', 'POST').then(json => this.$router.push('/tripnodes'));
      }
    }
}
</script>

<style scoped>
  .search {
    margin: 10px
  }
</style>
