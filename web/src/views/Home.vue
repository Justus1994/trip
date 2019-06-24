<template>
  <div v-darkmode="darkmode" class="homeContainer">
    <Menu
    :darkmode="darkmode"
    v-bind:class="{'menuActive': menuActive}"
    v-on:closeMenu="menuActive = false"
    v-on:toggleDarkmode="toggleDarkmode"
    >
    </Menu>
    <!--
      Controls:  FAB, snackbar, dialog
    -->
    <div class="controls ">
      <!--
        dialog
      -->
      <v-dialog id="NewTripDialog" content-class='animationCard' v-model="dialog" max-width="600px">
        <template v-slot:activator="{ on }">
          <!--
            FAB
          -->
          <v-btn v-darkmode="darkmode" class="fab" fab v-on="on">
            <v-icon>add</v-icon>
          </v-btn>
        </template>
        <v-card v-darkmode="darkmode">
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
            <v-btn v-darkmode="darkmode" flat @click="dialog= false">Close</v-btn>
            <v-btn v-darkmode="darkmode" flat v-on:click="getNodes">Show places</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
      <!--
        snackbar
      -->
      <v-snackbar class="snackbar" v-model="snackbar" top :timeout="4000">
        Sorry, we couldn't find any pictures.
        <v-btn flat @click="snackbar = false">
          Close
        </v-btn>
      </v-snackbar>
    </div>
    <Content
    :darkmode="darkmode"
    v-bind:class="{'contentActive': menuActive}"
    v-on:closeMenu="menuActive = true"

    >
    </Content>
   </div>
</template>

<script>

import Menu from '../components/Menu'
import Content from '../components/Content'
import fetch from '../fetchData'
export default {
  name: "Home",
  components: {
    Menu,
    Content
  },
  beforeCreate(){
    if(window.localStorage.getItem('darkmode') === 'true'){
      this.darkmode = true;
    }
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
        this.triggerAnimation();
      }else{
        fetch('trip/' + this.place, 'POST').then(json => {
          this.$root.$data.sharedState.pendingTrip = json;
          if(typeof(json) === "string"){
            this.snackbar = true;
            this.triggerAnimation();
          }
          else {
            this.$router.push('/tripnodes');
          }
        });
      }

    },
    toggleDarkmode(){
      this.darkmode = this.darkmode? false: true;
      this.menuActive = false;
      window.localStorage.setItem('darkmode',this.darkmode);
    },
    triggerAnimation(){
      document.getElementsByClassName('animationCard')[0].style.animation = 'wobble 0.8s';
      let that = this
      setTimeout(function() {
        document.getElementsByClassName('animationCard')[0].style.removeProperty('animation');
        that.snackbar = true
      }, 800);
    }
  },
  watch: {
    '$route': 'fetchTrips'
  },
  data() {
    return {
      trips: this.$root.$data.sharedState.trips,
      dialog: false,
      place: '',
      snackbar: false,
      darkmode: false,
      menuActive: false,
      loading: true
    }
  },
}
</script>

<style>
.menuActive{
  transform: translateX(100%);
}

.contentActive{
    transform: translateX(33%) translateY(-5em);
}

.headline_dialog {
  font: 900 40px 'Great Vibes', cursive;
  margin: auto;
}
.textfield {
  font: 400 16px Montserrat !important;
}

input{
  letter-spacing: 1px;
  font-size: 11px;
  margin-left: -1em;
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
  font: 400 12px Montserrat;
  padding: 1rem;
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

button{
  font: 400 12px Montserrat;
}
::-webkit-scrollbar {
    display: none;
}
*{
  -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
            user-select: none;
}
@media only screen and (min-width: 600px) {
    input{
      font-size: 16px;
    }
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
