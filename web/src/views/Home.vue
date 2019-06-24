<template>
  <div v-darkmode="darkmode" class="homeContainer">
    <Menu
    :darkmode="darkmode"
    v-bind:class="{'menuActive': menuActive}"
    v-on:closeMenu="menuActive = false"
    v-on:toggleDarkmode="toggleDarkmode"
    >
    </Menu>
    <Controls
    :darkmode="darkmode"
    :loading="loading"
    v-on:newTrip="newTrip"
    >
    </Controls>
    <Content
    :darkmode="darkmode"
    v-bind:class="{'contentActive': menuActive}"
    v-on:closeMenu="menuActive = true"
    >
    </Content>
    <v-snackbar class="snackbar" v-model="snackbar" top :timeout="4000">
      {{msg}}
    <v-btn flat @click="snackbar = false">
      Close
    </v-btn>
  </v-snackbar>
  </div>
</template>

<script>

import Menu from '../components/Menu'
import Content from '../components/Content'
import Controls from '../components/Controls'

import fetch from '../fetchData'

export default {
  name: "Home",
  components: {
    Menu,
    Content,
    Controls,
  },
  data() {
    return {
      trips: this.$root.$data.sharedState.trips,
      darkmode: false,
      menuActive: false,
      loading: false,
      wobble: true,
      snackbar: false,
      msg:  "Sorry, we couldn't find any pictures.",
    }
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
    newTrip(place){
      if(place.length === 0){
        this.snackbar = true;
        this.wobbleActivator();
      }else{
        this.loading = true;
        this.fetchNewTrip(place);
      }
    },
    fetchNewTrip(place){
      fetch('trip/' + place, 'POST').then(response => {
          this.loading = false;
          if(response === 404){
            this.snackbar = true;
            this.wobbleActivator();
          }else{
            this.$root.$data.sharedState.pendingTrip = response;
            this.$router.push('/tripnodes');
          }


      }).catch(err => console.log(err));
    },
    toggleDarkmode(){
      this.darkmode = this.darkmode? false: true;
      this.menuActive = false;
      window.localStorage.setItem('darkmode',this.darkmode);
    },
    wobbleActivator(){
      /**
      * use document.getElement because v-dialog cannot bind dynamic style
      */
      document.getElementsByClassName('dialog')[0].style.animation = 'wobble 0.8s linear';
      let that = this;
      setTimeout(function() {
        document.getElementsByClassName('dialog')[0].style.removeProperty('animation');
        that.snackbar = true
      }, 800);
    },
  },
  watch: {
    '$route': 'fetchTrips'
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

body{
  margin: 0;
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
.modifyload #Group-7{
  display: none;
}
.wobble{
  animation: wobble 0.8s linear;
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
