<template>
    <div>
        <h1>New Trip</h1>
        <v-flex xs12 sm6 md3 class="search">
            <v-text-field 
              label="Solo"
              prepend-inner-icon="search"
              placeholder="Search for a country"
              solo
            ></v-text-field>
          </v-flex>
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
