<template lang="html">
  <!--
    Controls: Dialog, FAB,,
  -->
  <div class="controls ">
    <!--
      dialog
    -->

    <v-dialog  content-class="dialog" v-model="dialog" max-width="600px">
      <template v-slot:activator="{ on }">
        <!--
          FAB
        -->
        <v-btn v-darkmode="darkmode" class="fab" fab v-on="on">
          <v-icon>add</v-icon>
        </v-btn>
      </template>
      <v-card v-darkmode="darkmode">
        <Loading class="modifyload" v-if="loading"> </Loading>
        <v-card-title>
          <span class="headline_dialog">create a new trip</span>
        </v-card-title>
        <v-card-text>
          <v-container grid-list-md>
            <v-text-field color="darkmode? #fcfcfc : #333" :dark='darkmode' autofocus v-on:keyup.enter="$emit('newTrip',place)"
                class="textfield" prepend-inner-icon="search"
                placeholder="enter a Country, City or Place..."
                v-model="place" required>
            </v-text-field>
          </v-container>
        </v-card-text>
        <v-card-actions class="spaceBetween">
          <v-btn v-darkmode="darkmode" flat @click="dialog = false">Close</v-btn>
          <v-btn v-darkmode="darkmode" flat @click="$emit('newTrip',place)">Show places</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import Loading from './Loading'
export default {
  name: "Controls",
  props: ['darkmode','loading'],
  components: {
    Loading
  },
  data(){
    return{
      dialog: false,
      place: '',
    }
  },
}
</script>

<style lang="css">
  .controls{
    height: 0;
  }
  .fab{
    position: fixed;
    top: 90%;
    left: 50%;
    z-index: 99;
    transform: translateX(-62%);
  }
  .fab:hover{
      position: fixed;
  }
  .headline_dialog {
    font: 900 40px 'Great Vibes', cursive;
    margin: auto;
  }
  .textfield {
    font: 400 16px Montserrat !important;
  }
  .spaceBetween{
    display: flex;
    justify-content: space-between;
  }
  .modifyload #Group-7{
    display: none;
  }
</style>
