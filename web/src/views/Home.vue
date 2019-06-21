<template>
<div id="zone" v-bind:class="[darkmode ? 'darkmode' : 'lightmode', 'homeContainer']">
  <div id="filter" v-bind:class="[darkmode ? 'darkmode' : 'lightmode', 'swipeMenu']">
    <v-btn v-bind:class="[darkmode ? 'darkmode' : 'lightmode']" flat @click="toggleDarkmode">{{darkmode ? 'lightmode' : 'darkmode'}}</v-btn>
    <v-btn v-bind:class="[darkmode ? 'darkmode' : 'lightmode']" flat @click="closeFilter">Close</v-btn>
  </div>

  <div class="controls ">
    <v-dialog id="NewTripDialog" content-class='animationCard' v-model="dialog" max-width="600px">
      <template v-slot:activator="{ on }">
          <v-btn v-bind:class="[darkmode ? 'darkmode' : 'lightmode','fab']" fab v-on="on">
            <v-icon>add</v-icon>
          </v-btn>
      </template>

      <v-card v-bind:class="[darkmode ? 'darkmode' : 'lightmode']">
        <v-card-title>
          <span class="headline_dialog">create a new trip</span>
        </v-card-title>
        <v-card-text>
          <v-container grid-list-md>
              <v-text-field color="darkmode? #fcfcfc : #333" :dark='darkmode' autofocus v-on:keyup.enter="getNodes"
              class="textfield" prepend-inner-icon="search"
                placeholder="enter a Country, City or Place..."
                v-model="place" required>
            </v-text-field>
          </v-container>
        </v-card-text>
        <v-card-actions class="spaceBetween">
          <v-btn v-bind:class="[darkmode ? 'darkmode' : 'lightmode']" flat @click="dialog= false">Close</v-btn>
          <v-btn v-bind:class="[darkmode ? 'darkmode' : 'lightmode']" flat v-on:click="getNodes">Show places</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-snackbar class="snackbar" v-model="snackbar" top :timeout="4000">
      You need to enter something
      <v-btn flat @click="snackbar = false">
        Close
      </v-btn>
    </v-snackbar>
  </div>

  <div id="content" v-bind:class="[darkmode ? 'darkmodebg' : 'lightmode','swipeMenu']">
    <div v-bind:class="[darkmode ? 'darkmode' : 'lightmode','header']">
    <v-btn v-bind:class="[darkmode ? 'darkmode' : 'lightmode','menu']" flat fab @click="toggleFilter">
       <v-icon >menu</v-icon>
     </v-btn>
     <div class="header_text">
       trip
     </div>
   </div>
   <div class="somespace">
   </div>
    <div  v-bind:key="index" v-for="(trip, index) in getTrips">
      <TripCard :darkmode="darkmode" :trip="trip" :index="index" />
    </div>
    <div class="somespace">
    </div>
  </div>
</div>
</template>

<script>
import TripCard from '../components/TripCard.vue'
import fetch from '../fetchData'
export default {
  name: "NewTrip",
  components: {
    TripCard
  },
  created() {
    this.fetchTrips();

    if(window.localStorage.getItem('darkmode') === 'true'){
      this.darkmode = true;
    }

  },
  methods: {
    fetchTrips() {
      fetch('trip', 'GET').then(data => {
        this.$set(this.$root.$data.sharedState, 'trips', data)
      });
    },
    getNodes() {
      if (this.place.length == 0) {
        document.getElementsByClassName('animationCard')[0].style.animation = 'wobble 0.8s';
        let that = this
        setTimeout(function() {
          document.getElementsByClassName('animationCard')[0].style.removeProperty('animation');
          that.snackbar = true
        }, 800);
      } else {
        fetch('trip/' + this.place, 'POST').then(json => {
          this.$root.$data.sharedState.pendingTrip = json;
          this.$router.push('/tripnodes');
        });
      }
    },
    toggleFilter(){
      var DOMnodes = document.getElementsByClassName('swipeMenu');
      DOMnodes.filter.style.transform = 'translateX(100%)';
      DOMnodes.content.style.transform = 'translateX(33%) translateY(-5em)';

    },
    closeFilter(){
      var DOMnodes = document.getElementsByClassName('swipeMenu');
      for(let item of DOMnodes){
        item.style.removeProperty('transform');
      }
    },
    toggleDarkmode(){
      this.darkmode = this.darkmode ? false: true;
      this.closeFilter();
      window.localStorage.setItem('darkmode',this.darkmode);
    }
  },
  watch: {
    '$route': 'fetchTrips'
  },
  computed: {
    getTrips: function() {
      return this.$root.$data.sharedState.trips
    }
  },
  data: function() {
    return {
      trips: this.$root.$data.sharedState.trips,
      dialog: false,
      place: '',
      isValid: true,
      snackbar: false,
      darkmode: false,


    }
  },
}
</script>

<style>
.headline_dialog {
  font: 900 40px 'Great Vibes', cursive;
  margin: auto;
}
.my-text-style >>> .v-text-field__slot input {
    color: red
  }
.textfield {
  font: 400 16px Montserrat !important;
}

input{
  letter-spacing: 2px;
  font-weight: 400;
  text-align: center;
  text-transform: uppercase;
}
.text-xs-center {
  position: fixed;
  top: 90%;
}
.animation {
  animation: wobble 0.8s;
}

.text-xs-center .right {
  float: right;
  margin: auto;
}

.snackbar {
  padding: 1rem;
}

.header{
  display:flex;
  box-shadow: 2px 0 20px rgba(0,0,0,0.2);
  position: fixed;
  width: 100%;
  z-index: 99;
}
.header_text{
  width: 100%;
  text-align: center;
  margin: auto;
  font: 900 40px 'Great Vibes', cursive;
  transform: translateX(-30px);
}
.spaceBetween{
  display: flex;
  justify-content: space-between;
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
#content{
  transition: all 0.3s ease;
  box-shadow: 0 5px 20px rgba(0, 0, 0, 0.3);
  min-height: 100vh;
  max-height: 105vh;
  overflow: scroll;
  position: absolute;
  width: 100%;
  top: 0;
}
#filter{
  transition: all 0.5s cubic-bezier(0.25, 0.1, 0, 1.2);
  position: absolute;
  background: white;
  height: 100vh;
  top: 0;
  left: -33%;
  width: 33%;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
body{
  margin: 0;
}
.controls{
  height: 0;
}
.homeContainer{
  min-height: 100vh;
  max-height: 100vh;
  overflow: hidden;
}
.somespace{
  height: 5em;
}
.darkmode{
  color: #fcfcfc !important;
  background: #333 !important;
}
.darkmodebg{
  background: #1c1d21 !important;
}
.lightmode{
  color: #131313 !important;
  background: #fcfcfc !important;
}
button{
  font: 400 12px Montserrat;
}
::-webkit-scrollbar {
    display: none;
}
@-webkit-keyframes wobble {
  0% {
    -webkit-transform: none;
    transform: none;
  }

  15% {
    -webkit-transform: translate3d(-25%, 0, 0) rotate3d(0, 0, 1, -5deg);
    transform: translate3d(-25%, 0, 0) rotate3d(0, 0, 1, -5deg);
  }

  30% {
    -webkit-transform: translate3d(20%, 0, 0) rotate3d(0, 0, 1, 3deg);
    transform: translate3d(20%, 0, 0) rotate3d(0, 0, 1, 3deg);
  }

  45% {
    -webkit-transform: translate3d(-15%, 0, 0) rotate3d(0, 0, 1, -3deg);
    transform: translate3d(-15%, 0, 0) rotate3d(0, 0, 1, -3deg);
  }

  60% {
    -webkit-transform: translate3d(10%, 0, 0) rotate3d(0, 0, 1, 2deg);
    transform: translate3d(10%, 0, 0) rotate3d(0, 0, 1, 2deg);
  }

  75% {
    -webkit-transform: translate3d(-5%, 0, 0) rotate3d(0, 0, 1, -1deg);
    transform: translate3d(-5%, 0, 0) rotate3d(0, 0, 1, -1deg);
  }

  100% {
    -webkit-transform: none;
    transform: none;
  }
}

</style>
