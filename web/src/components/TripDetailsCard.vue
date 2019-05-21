<template>
<div>
  <v-card class="card" color="white">
    <v-card-title 
      class=title 
      color="black"
      >
        {{node.location.title}}
        <v-layout
                align-center
                justify-end
              >
              <v-btn small icon v-on:click="deleteNode">
                <v-icon>delete</v-icon>
              </v-btn>
        </v-layout>
    </v-card-title>
    <v-img
      :src=node.urls.regular
      height="150px"
    >
    </v-img>
  </v-card>
</div>
</template>

<script>
export default {
    name: "TripDetailsCard",
    props: ["node"],
    methods: {
      deleteNode() {
        console.log("node deleted")
      }
    },
}
</script>

<style scoped>
 .card {
   padding: 0px;
 }
 .title {
   padding: 5px 10px;   
 }
</style>
