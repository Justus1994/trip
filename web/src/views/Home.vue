<template>
  <div v-darkmode="darkmode" class="homeContainer">

    <Menu
    :darkmode="darkmode"
    v-bind:class="{'menuActive': menuActive}"
    v-on:closeMenu="menuActive = false"
    v-on:toggleDarkmode="toggleDarkmode"
    ></Menu>

    <Controls
    :darkmode="darkmode"
    :loading="loading"
    v-on:newTrip="newTrip"
    ></Controls>

    <Content
    :darkmode="darkmode"
    v-bind:class="{'contentActive': menuActive}"
    v-on:closeMenu="menuActive = true"
    v-on:share="share"
    v-on:show="show"
    ></Content>

  </div>
</template>

<script>

import Menu from '../components/Menu'
import Content from '../components/Content'
import Controls from '../components/Controls'
import copyTripToClipboard from '../share.js'

export default {
  name: "Home",
  components: {
    Menu,
    Content,
    Controls,
  },
  data() {
    return {
      darkmode: false,
      menuActive: false,
      loading: false,
    }
  },
  created() {
    this.fetchTrips();
    if(window.localStorage.getItem('darkmode') === 'true'){
      this.darkmode = true;
    }
  },
  watch: {
    '$route': 'fetchTrips'
  },
  methods: {
    fetchTrips() {
      fetch('trip', 'GET').then(data => {
          this.$set(this.$root.$data.store, 'trips',data)
        }).catch(err => this.showSnackbar('Upps, Something went wrong.','red'));
    },
    newTrip(place){
      if(place.length === 0){
        this.showSnackbar("Some more information please...",'red')
        this.wobbleActivator();
      }else{
        this.loading = true;
        this.fetchNewTrip(place);
      }
    },
    fetchNewTrip(place){
      fetch('trip/' + place, 'POST').then(response => {
          this.loading = false;
          this.$root.$data.store.pendingTrip = response;
          this.$router.push('/tripnodes');
        }).catch(err => {
          this.loading = false;
          this.showSnackbar("Sorry, we can't find pictures of " + place,'red');
          this.wobbleActivator();
        })
    },
    toggleDarkmode(){
      this.darkmode = this.darkmode? false : true;
      this.menuActive = false;
      window.localStorage.setItem('darkmode',this.darkmode);
    },
    share(i){
      this.showSnackbar("Trip was copied to clipboard",'green');
      copyTripToClipboard(this.$root.$data.store.trips[i]);
    },
    show(i) {
      this.$router.push('/tripdetails/' + i);
    },
    showSnackbar(msg,color){
      this.$set(this.$root.$data.store.snack, 'color', color);
      this.$set(this.$root.$data.store.snack, 'msg', msg);
      this.$set(this.$root.$data.store.snack,'show',true);
    },
    wobbleActivator(){
      /**
      * use document.getElement because v-dialog cannot bind dynamic style
      * see docs vuetify: "the content is moved to the end of the app and is not targettable by classes passed directly on the component."
      */
      document.getElementsByClassName('dialog')[0].style.animation = 'wobble 0.8s linear';
      let that = this;
      setTimeout(function() {
        document.getElementsByClassName('dialog')[0].style.removeProperty('animation');
        that.snackbar = true
      }, 800);
    },
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
  .homeContainer{
    min-height: 100vh;
    max-height: 100vh;
    overflow: hidden;
  }
  .snackbar {
    padding: 2rem;
  }
  /**
  *Global definition
  */

  .list-item {
    transition: all 1s ease;
  }
  .list-enter, .list-leave-to
  {
    opacity: 0;
  }
  .list-leave-active {
    opacity: 0
  }

  html,body{
    margin: 0;
    padding: 0;
  }
  input{
    letter-spacing: 1px;
    font-size: 11px;
    margin-left: -1em;
    font-weight: 400;
    text-align: center;
    text-transform: uppercase;
  }
  button{
    font: 400 12px Montserrat;
  }
  ::-webkit-scrollbar {
      display: none; /*disable scrollbar */
  }
  *{
    -webkit-touch-callout: none; /* iOS Safari */
      -webkit-user-select: none; /* Safari */
              user-select: none;
  }

  /*Desktop view */
  @media only screen and (min-width: 600px) {
      input{
        font-size: 16px;
      }
  }
  /* animation classes */
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
/*  in Global sheet because  dialog will appear in dom on different position */
  .lds-ellipsis {
    display: inline-block;
    position: relative;
    width: 64px;
    height: 64px;
  }
  .lds-ellipsis div {
    position: absolute;
    top: 27px;
    width: 11px;
    height: 11px;
    border-radius: 50%;
    animation-timing-function: cubic-bezier(0, 1, 1, 0);
  }
  .lds-ellipsis div:nth-child(1) {
    left: 6px;
    animation: lds-ellipsis1 0.6s infinite;
  }
  .lds-ellipsis div:nth-child(2) {
    left: 6px;
    animation: lds-ellipsis2 0.6s infinite;
  }
  .lds-ellipsis div:nth-child(3) {
    left: 26px;
    animation: lds-ellipsis2 0.6s infinite;
  }
  .lds-ellipsis div:nth-child(4) {
    left: 45px;
    animation: lds-ellipsis3 0.6s infinite;
  }
  @keyframes lds-ellipsis1 {
    0% {
      transform: scale(0);
    }
    100% {
      transform: scale(1);
    }
  }
  @keyframes lds-ellipsis3 {
    0% {
      transform: scale(1);
    }
    100% {
      transform: scale(0);
    }
  }
  @keyframes lds-ellipsis2 {
    0% {
      transform: translate(0, 0);
    }
    100% {
      transform: translate(19px, 0);
    }
  }
</style>
