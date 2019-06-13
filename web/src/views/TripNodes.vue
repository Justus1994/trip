<template>
<div class="max" id="prevent">
  <v-img class="img" :src=currentNode.urls.regular>
    <div class="mid">
      <div class="imageHeading">{{currentNode.location.title}}</div>
      <div class="imageSubHeading"> {{currentNode.location.country}}</div>
    </div>
    <div class="left">
      <v-btn fab class="btn" outline color="white" v-on:click="nextNode(false)">
        <v-icon>close</v-icon>
      </v-btn>
    </div>
    <div class="right">
      <v-btn class="btn" fab left outline color="white" v-on:click="nextNode(true)">
        <v-icon>favorite</v-icon>
      </v-btn>
    </div>
  </v-img>
</div>
</template>

<script>
import store from '../store.js'
import fetch from '../fetchData'

export default {
  data() {
    return {
      currentNode: {},
      trips: this.$root.$data.sharedState.pendingTrip,
      counter: 0
    }
  },
  methods: {
    nextNode(like) {
      if (like) {
        console.log("Node like and added")
      } else {
        this.deleteNode(this.currentNode.id)
      }
      this.trips.Nodes.length > this.counter ? this.currentNode = this.trips.Nodes[this.counter++] : this.$router.push('/')
    },
    deleteNode(id) {
      console.log("node deleted")
      console.log(id)
      fetch('trip/' + 0 + '/nodes/' + id, 'DELETE').then(json => {
        console.log(json)
      })
    },
  },
  created() {
    this.currentNode = this.trips.Nodes[this.counter++]
    console.log(this.trips)
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

#app {
  height: 100%;
}

.max {
  overflow: hidden;
  height: 100%;
}

.btn {
  margin: 0;
}

.bottomNavBar {
  display: none;
}

.left {
  margin: 0;
  position: fixed;
  left: 25%;
  transform-origin: 50% 50%;
  transform: translateX(-50%);
  top: 88%;
}

.right {
  margin: 0;
  position: fixed;
  right: 15%;
  transform-origin: 50% 50%;
  transform: translateX(-50%);
  top: 88%;
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
  text-shadow: 0 10px 20px black;
}

.imageSubHeading {
  color: white;
  text-align: center;
  font: 900 28px Montserrat;
  letter-spacing: 8px;
  text-shadow: 0 10px 25px #000000;
  text-transform: uppercase;
}

.lightbox {
  box-shadow: 0 0 20px inset rgba(0, 0, 0, 0.2);
  background-image: linear-gradient(to top, rgba(0, 0, 0, 0.4) 0%, transparent 72px);
}
</style>
