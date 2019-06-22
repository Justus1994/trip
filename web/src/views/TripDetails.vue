<template>
  <div v-bind:class="[darkmode ? 'darkmodebg' : 'lightmode','container']">
    <v-img max-height="100vh" min-height="100vh" class="scrollSnap" :src=getHeaderImg>
      <v-layout pa-2 column fill-height class="lightbox white--text">
        <v-spacer></v-spacer>
        <v-flex >
          <div class="imageHeading">{{nodes[0].location.city}}</div>
          <div class="imageSubHeading">{{nodes[0].location.country}}</div>
          <div class="imageDaysHeading">{{nodes.length}} Places</div>
        </v-flex>
      </v-layout>
    </v-img>

    <div class="scrollSnap" v-for="(node, i) in getNodes" :key="i">
        <NodeCard :darkmode="darkmode" :node="node"/>
    </div>
    <div class="center">
        <v-btn v-bind:class="[darkmode ? 'darkmode' : 'lightmode']" fab v-on:click="deleteTrip">
          <v-icon>delete</v-icon>
        </v-btn>
    </div>
  </div>
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
    },
    getHeaderImg(){
      return this.$root.$data.sharedState.trips[this.index].Nodes[0].urls.regular;
    }
  },
  data() {
    return {
        index: this.$route.params.id,
        nodes: this.$root.$data.sharedState.trips[this.$route.params.id].Nodes,
        darkmode: false
    }
  },
  created(){
    if(!this.nodes){
       this.$router.push('/');
    }

    if(window.localStorage.getItem('darkmode') === 'true'){
      this.darkmode = true;
      document.getElementById('app').style.background = '#1c1d21';
    }
  },
  methods: {
    deleteTrip() {
      console.log("trip deleted")
        fetch('trip/' + this.index, 'DELETE').then(  this.$router.push('/'));
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
    text-shadow: 0 10px 20px rgba(0,0,0,0.5);
  }
  .imageDaysHeading{
    text-align: center;
    font: 900 20px Montserrat;
    letter-spacing: 2px;
    text-shadow: 0 10px 25px rgba(0,0,0,0.5);
    text-transform: uppercase;
    -webkit-text-stroke: 1px rgba(255,255,255,0.9);
    letter-spacing: 8px;
    color: transparent;
  }
  .scrollSnap{
    scroll-snap-align: start;
    padding-top:0.25em;
  }
  .scrollSnap:last-child(2){
    scroll-snap-align: end;
  }
  .imageSubHeading{
    transition: all 0.5s ease;
    text-align: center;
    font: 900 30px Montserrat;
    letter-spacing: 8px;
    text-shadow: 0 10px 25px rgba(0,0,0,0.5);
    text-transform: uppercase;
    color: rgba(255,255,255,0.7);
  }

  .lightbox {
      box-shadow: 0 0 20px inset rgba(0, 0, 0, 0.2);
      background-image: linear-gradient(to top, rgba(0, 0, 0, 0.4) 0%, transparent 72px);
  }

  .arrow {
    transform: rotate(90deg) scale(2);
    margin: 20px;
  }
  @media only screen and (min-width: 600px) {
    .imageSubHeading{
        font-size: 80px;

    }
    .imageDaysHeading{
      font-size: 40px;
      letter-spacing: 5px;

    }
  }
</style>
