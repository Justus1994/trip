<template>
<div>
  <div v-bind:key="index" v-for="(trip, index) in getTrips">
    <TripCard :trip="trip" :index="index" />
  </div>
  <v-dialog id="NewTripDialog" content-class='animationCard' v-model="dialog" max-width="600px">
    <template v-slot:activator="{ on }">
      <div class='centerBottom'>
        <v-btn fab color="accent" v-on="on">
          <v-icon>add</v-icon>
        </v-btn>
      </div>
    </template>
    <v-card>
      <v-card-title>
        <span class="headline_dialog">CREATE A NEW TRIP</span>
      </v-card-title>
      <v-card-text>
        <v-container grid-list-md>
            <v-text-field class="textfield" prepend-inner-icon="search" placeholder="enter a country" color="accent" v-model="place" required></v-text-field>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="accent" flat @click="dialog= false">Close</v-btn>
        <v-btn color="accent" flat v-on:click="getNodes">show places</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
  <v-snackbar v-model="snackbar" :top='true' :timeout="4000">
    you need to enter something
    <v-btn color="pink" flat @click="snackbar = false">
      Close
    </v-btn>
  </v-snackbar>

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
      } else {
        fetch('trip/' + this.place, 'POST').then(json => {
          this.$root.$data.sharedState.pendingTrip = json;
          this.$router.push('/tripnodes');
        });
      }
      let that = this
      setTimeout(function() {
        document.getElementsByClassName('animationCard')[0].style.removeProperty('animation');
        that.snackbar = true;
      }, 800);
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
    }
  },
}
</script>

<style>
.headline_dialog {
  font: 400 25px Montserrat !important;
  color: #5c1349;
}

.textfield {
  font: 200 22px Montserrat !important;
}

.text-xs-center {
  position: fixed;
  top: 90%;
}

.animation {
  animation: wobble 0.8s;
  background: blue;
}

.text-xs-center .right {
  float: right;
  margin: auto;
}

.centerBottom {
  position: fixed;
  left: 50%;
  transform: translateX(-50%);
  top: 90%;
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
