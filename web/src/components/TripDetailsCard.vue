<template>
<div>
  <v-card dark class="card" color="accent">
    <v-card-title 
      class=title 
      color="white"
      >
        {{node.title}}
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
      :src=node.image
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
