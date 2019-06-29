<template lang="html">
  <!-- Content:  Header, Placeholder, List of TripCard.vue -->
  <div v-darkmode:[background]="darkmode" class="content">
    <!-- Header -->
    <div v-darkmode="darkmode" class="header">
      <v-btn v-darkmode="darkmode" flat fab @click="$emit('closeMenu')">
        <v-icon>menu</v-icon>
      </v-btn>
      <div class="header_text">
        trip
      </div>
     </div>
     <!-- Placeholder -->
     <div v-bind:class="[tripExist ? 'displayNone':'placeholderImg']">
       <p>no trips yet</p>
       <p>try creating one</p>
       <span class="arrowBtn"></span>
     </div>
      <!-- List -->
     <div class="somespace">
     </div>
      <div class="scrollSnapHome" v-bind:key="index" v-for="(trip, index) in getTrips">
        <TripCard
        :darkmode="darkmode"
        :trip="trip"
        :index="index"
        v-on:shareTrip="$emit('share',index)"
        v-on:showDetails="$emit('show',index)"
        />
      </div>
      <div class="somespace">
      </div>
   </div>
</template>

<script>
import TripCard from '../components/TripCard.vue'
export default {
  name: "Content",
  components: {
    TripCard
  },
  props: ['darkmode'],
  computed: {
    getTrips: function() {
      return this.$root.$data.store.trips;
    },
    tripExist(){
      return this.$root.$data.store.trips ? true: false;
    }
  },
  data(){
    return{
      background: 'background',
    }
  }
}
</script>

<style lang="css" scoped>
  .content{
    transition: all 0.3s ease;
    box-shadow: 0 5px 20px rgba(0, 0, 0, 0.3);
    min-height: 100vh;
    max-height: 105vh;
    overflow: scroll;
    position: absolute;
    width: 100%;
    scroll-snap-type: y mandatory;
    top: 0;
  }
  .header{
    display:flex;
    box-shadow: 2px 0 20px rgba(0,0,0,0.2);
    padding-top: -20px;
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

  .placeholderImg{
    opacity: 0;
    background: url('../assets/camera.svg');
    position: fixed;
    top: 40%;
    left: 50%;
    transform: translateY(-50%) translateX(-50%);
    height: 5em;
    width: 5em;
    animation: fadeIn 2s cubic-bezier(0.94, -0.01, 0.32, 0.99) 0s 1 normal forwards;
  }
  .placeholderImg svg{
    fill: #E0E0E0;
  }
  .displayNone{
    display: none;
  }
  .placeholderImg p:first-child{
    color: #E0E0E0;
    position: fixed;
    top: 150%;
    left: 50%;
    width: 200px;
    text-align: center;
    font: 400 22px Montserrat;
    transform: translateY(-50%) translateX(-50%);
  }
  .placeholderImg p:nth-child(2){
    color: #E0E0E0;
    position: fixed;
    top: 200%;
    left: 50%;
    width: 200px;
    text-align: center;
    font: 400 12px Montserrat;
    transform: translateY(-50%) translateX(-50%);
  }
  .somespace{
    height: 5em;
  }
  .scrollSnapHome{
    scroll-snap-align: center;

  }
  .arrowBtn {
    display: block;
    width: 48px;
    height: 48px;
    border-radius: 50%;
    position: absolute;
    top: 280%;
    left: 50%;
    transform: translateX(-50%);
  }
  .arrowBtn::before {animation: arrowMove 2s 1s ease-in-out infinite;}
  .arrowBtn::after {animation: arrowMove 2s ease-in-out infinite;}

  .arrowBtn::before,
  .arrowBtn::after {
    content: "";
    display: block;
    box-sizing: border-box;
    border-right: 3px solid #E0E0E0;
    border-bottom: 3px solid #E0E0E0;
    border-radius: 2px;
    width: 20px;
    height: 20px;
    transform-origin: top left;
    transform: rotate(45deg);
    position: absolute;
    top: 0;
    left: 50%;
    opacity: 0;
  }
  @keyframes fadeIn {
    from{
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }
  @keyframes arrowMove {
    0% {
      top: 0;
      opacity: 0;
    }
    50% {opacity: 1;}
    100% {
      top: 60%;
      opacity: 0;
    }
  }
</style>
