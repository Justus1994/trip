<template>
  <v-container class="container" fluid >
    <v-img max-height="821px" min-height="821px" :src=trip.Nodes[0].urls.regular>
      <v-layout pa-2 column fill-height class="lightbox white--text">
        <v-spacer></v-spacer>
        <v-flex >
          <div class="imageHeading">{{trip.Nodes[index].location.city}}</div>
          <div class="imageSubHeading"> {{trip.Nodes[index].location.country}}</div>
          <div class="imageSubHeading">{{trip.Nodes.length}} Places</div>
        </v-flex>
        <v-layout justify-end align-end>
          <v-btn class="arrow" v-on:click="scrollDown" dark icon><v-icon>arrow_right_alt</v-icon></v-btn>
        </v-layout>
      </v-layout>
    </v-img>

    <v-card-text class="py-0">
      <v-timeline
      color="white"
      align-top
      dense
      >
        <v-timeline-item
            color="accent"
            small
            icon="place"
            v-for="(node, i) in trip.Nodes"
            :key="i"
            fill-dot
          >
            <TripDetailsCard :node="node" :index="index" :i="i"/>
        </v-timeline-item>
      </v-timeline>
    </v-card-text>
    <div class="text-xs-center">
        <v-btn class="fab" fab color="accent" dark v-on:click="deleteTrip"><v-icon>delete</v-icon></v-btn>
    </div>
  <v-container>
</template>

<script>
import TripDetailsCard from '../components/TripDetailsCard.vue'
import fetch from '../fetchData.js'
import store from '../store.js'

export default {
  name: 'TripDetails',
  components: {
      TripDetailsCard
  },
  created() {
    this.trip = this.$root.$data.sharedState.trips[this.index];
  },
  data() {
    return {
        trip: {},
        index: this.$route.params.id
    }
  },
  methods: {
    deleteTrip() {
      console.log("trip deleted")
        fetch('trip/' + this.index, 'DELETE').then(json => console.log("delete: ", json))
    },
    scrollDown() {
      console.log("scroll")
    }
  },
}
</script>

<style scoped>
    .container{
        padding: 0;
    }
    .imageHeading{
        font-size: 40px;
    }
    .imageSubHeading{
        font-size: 20px;
        font-weight: 500;
    }
    .lightbox {
        box-shadow: 0 0 20px inset rgba(0, 0, 0, 0.2);
        background-image: linear-gradient(to top, rgba(0, 0, 0, 0.4) 0%, transparent 72px);
    }
    .fab {
      margin-bottom: 20px
    }
    .arrow {
      transform: rotate(90deg) scale(2);
      margin: 20px;
    }
</style>
