<template>
  <v-container class="container">
    <v-img max-height="100vh" min-height="100vh" class="scrollSnap" :src=nodes[0].urls.regular>
      <v-layout pa-2 column fill-height class="lightbox white--text">
        <v-spacer></v-spacer>
        <v-flex >
          <div class="imageHeading">{{nodes[0].location.city}}</div>
          <div class="imageSubHeading">{{nodes[0].location.country}}</div>
          <div class="imageDaysHeading">{{nodes.length}} Places</div>
        </v-flex>
        <v-layout justify-end align-end>
          <v-btn class="arrow" v-on:click="scrollDown" dark icon><v-icon>arrow_right_alt</v-icon></v-btn>
        </v-layout>
      </v-layout>
    </v-img>

    <div class="scrollSnap" v-for="(node, i) in getNodes" :key="i">
        <NodeCard :node="node"/>
    </div>
    <div class="center">
        <v-btn class="fab" fab color="accent" dark v-on:click="deleteTrip"><v-icon>delete</v-icon></v-btn>
    </div>
  <v-container>
</template>

<script>

import NodeCard from '../components/NodeCard.vue'
import fetch from '../fetchData.js'
import store from '../store.js'

export default {
  name: 'TripDetails',
  components: {
      NodeCard
  },
  computed: {
    getNodes: function(){
      return this.$root.$data.sharedState.trips[this.index].Nodes
    }
  },
  data() {
    return {
        index: this.$route.params.id,
        nodes: this.$root.$data.sharedState.trips[this.$route.params.id].Nodes
    }
  },
  methods: {
    deleteTrip() {
      console.log("trip deleted")
        fetch('trip/' + this.index, 'DELETE').then(  this.$router.push('/'));
    },
    scrollDown() {
      console.log("scroll")
    }
  },
}
</script>

<style scoped>
  .center{
    position: fixed;
    left: 50%;
    transform: translateX(-50%);
    top: 90%;
  }
    .container{
        padding: 0;
        scroll-snap-type: y mandatory;
        overflow: scroll;
        height:100vh;

    }
    .imageHeading{
      font: 900 80px 'Great Vibes', cursive;
      text-align: center;
      text-shadow: 0 10px 20px black;
    }
    .imageDaysHeading{
      text-align: center;
      font: 900 20px Montserrat;
      letter-spacing: 2px;
      text-shadow: 0 10px 25px #000000;
      text-transform: uppercase;

    }
.scrollSnap{
  scroll-snap-align: start;
}
    .imageSubHeading{
      text-align: center;
        font: 900 30px Montserrat;
        letter-spacing: 8px;
        text-shadow: 0 10px 25px #000000;
        text-transform: uppercase;
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
