<template>
  <div v-darkmode:[background]="darkmode" class="maxcontainer">
    <div class='container'>
      <v-img max-height="100vh" min-height="100vh" class="scrollSnap" :src="getHeaderImg">
        <v-layout pa-2 column fill-height class="lightbox white--text">
          <v-spacer></v-spacer>
          <v-flex >
            <div class="imageHeading">{{nodes[0].location.city}}</div>
            <div class="imageSubHeading">{{nodes[0].location.country}}</div>
            <div class="imageDaysHeading">{{nodes.length}} Places</div>
          </v-flex>
        </v-layout>
      </v-img>
      <transition-group name="nodes-complete">
        <NodeCard
          v-for="(node, i) in getNodes"
          class="scrollSnap nodes-complete-item"
          :darkmode="darkmode"
          :key="node.id"
          :node="node"
          v-on:deleteNode="deleteNode(node)"
        />
      </transition-group>
      <div class="center">
          <v-btn v-darkmode="darkmode" fab v-on:click="deleteTrip">
            <v-icon>delete</v-icon>
          </v-btn>
      </div>
  </div>
</template>

<script>

import NodeCard from '../components/NodeCard.vue'

export default {
  name: 'TripDetails',
  components: {
      NodeCard
  },
  computed: {
    getNodes: function(){
      return this.$root.$data.store.trips[this.index].Nodes
    },
    getHeaderImg(){
      return this.$root.$data.store.trips[this.index].Nodes[0].urls.regular;
    }
  },
  data() {
    return {
        index: this.$route.params.id,
        nodes: this.$root.$data.store.trips[this.$route.params.id].Nodes,
        darkmode: false,
        background: true,
    }
  },
  created(){
    if(!this.nodes){
       this.$router.push('/');
    }
    if(window.localStorage.getItem('darkmode') === 'true'){
      this.darkmode = true;
    }
  },
  methods: {
    deleteTrip() {
      fetch('trip/' + this.index, 'DELETE').then(json =>this.$router.push('/'));
    },
    deleteNode(node){
      if(this.getNodes.length > 1){
        fetch('trip/' + this.index + '/nodes/' + node.id, 'DELETE').then(json =>{
          this.$root.$data.store.trips[this.index].Nodes = json.Nodes;
        });
      }else{this.deleteTrip()}

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
    z-index: 9;
  }
  .trans{
    background: transparent;
  }
  .container{
    padding: 0;
    scroll-snap-type: y mandatory;
    overflow: scroll;
    height:100vh;
  }
  .scrollSnap{
    scroll-snap-align: start;
  }
  .container > div:first-child{
    scroll-snap-align: center;
  }

  .maxcontainer{
      width: 100%;
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
  .nodes-complete-item {
    transition: all 0.9s cubic-bezier(0.22, 0.61, 0.36, 1);
    display: block;
  }
  .nodes-complete-enter, .nodes-complete-leave-to
  {
    opacity: 0;
    position: absolute;
    transform: translateY(-2px);
  }
  .nodes-complete-leave-active {
    position: absolute;
    top: 80%;
    left: 0;
    right: 0;
  }
  .sure{
    margin-right: 1em;
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
