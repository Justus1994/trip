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
            ></v-text-field>
          </v-flex>
          <div class="text-xs-center">
            <router-link to="/tripnodes"><v-btn round class="text-xs-center" color="accent" >Show Places</v-btn></router-link>
          </div>
    </div>
</template>

<script>
export default {
    name: "NewTrip",
    methods: {
    }
}
</script>

<style scoped>
  .search {
    margin: 10px
  }
</style>
