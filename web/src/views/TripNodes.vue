<template>
<div class="max" id="prevent">
  <v-img class="img" v-bind:src="currentNode.id? currentNode.urls.regular : ''" >
    <div class="mid">
        <div class="imageHeading">{{currentNode.id ? currentNode.location.title : ""}}</div>
        <div class="imageSubHeadingChar">
          <div class="letter" v-for="(val) in currentNode.id ? currentNode.location.country: 'Loading'">{{val}}</div>
        </div>
    </div>
    <div class="buttonActions">
        <v-btn fab class="btn" outline color="white" v-on:click="nextNode(false)">
          <v-icon>close</v-icon>
        </v-btn>
        <v-btn class="btn" fab left outline color="white" v-on:click="nextNode(true)">
          <v-icon>favorite</v-icon>
        </v-btn>
    </div>
  </v-img>
</div>
</template>

<script>
import fetch from '../fetchData'

export default {
  data() {
    return {
      currentNode: {},
      trips: this.$root.$data.sharedState.pendingTrip,
      counter: 0,
      lastNode: false

    }
  },
  methods: {
    nextNode(like) {
      if (!like) {
        console.log('currentNode',this.currentNode);
        fetch('trip/' + 0 + '/nodes/' + this.currentNode.id, 'DELETE').then(json => {
          console.log(json);
        });
      }
      this.trips.Nodes.length > this.counter ? this.currentNode = this.trips.Nodes[this.counter++] : this.$router.push('/');
    },
  },
  mounted(){
    this.currentNode = this.trips.Nodes[this.counter++];
  },
  created() {
    if(!this.trips){
       this.$router.push('/');
    }


  },
}
</script>

<style>
.mid {
  position: fixed;
  left: 50%;
  transform-origin: 50% 50%;
  transform: translateX(-50%) translateY(-50%);
  top: 40%;
}

.max {
  overflow: hidden;
  height: 100vh;
  cursor: pointer;
}

.btn {
  transition: all 0.5s ease;
  margin: 0;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.3) !important;
  height: 4em;
  width: 4em;
}

.buttonActions{
  width: 100%;
  display: flex;
  justify-content: space-around;
  position: fixed;
  bottom: 5%;

}
.container {
  padding: 0;
}

.img {
  height: 100%;
  background-position: center;
  background-repeat: no-repeat;
  background-size: cover;
}

.imageHeading {
  color: white;
  font: 900 65px 'Great Vibes', cursive;
  text-align: center;
  text-shadow: 0px 1px 20px rgba(0,0,0,0.5);
}
.letter{
  margin: auto;
  transition: 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}
.letter:hover{
  -webkit-text-stroke: 2px #fff;
  color: transparent;
  transform: scale(1.5);
}
.imageSubHeadingChar {
  color: white;
  text-align: center;
  font: 900 24px Montserrat;
  letter-spacing: 10px;
  text-shadow: 2px 2px 20px rgba(0,0,0,0.3);
  text-transform: uppercase;
  display: flex;
}

.lightbox {
  box-shadow: 0 0 20px inset rgba(0, 0, 0, 0.2);
  background-image: linear-gradient(to top, rgba(0, 0, 0, 0.4) 0%, transparent 72px);
}

@media only screen and (min-width: 600px) {
  .imageSubHeadingChar{
      font-size: 40px;
  }
  .imageHeading{
    font-size: 80px;
    margin-bottom: 10px;
  }
}
</style>
