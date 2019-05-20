<template>
  <v-container class="container" fluid >
    <v-img max-height="821px" min-height="821px" :src=trip.image>
      <v-layout pa-2 column fill-height class="lightbox white--text">
        <v-spacer></v-spacer>
        <v-flex >
          <div class="imageHeading">{{trip.title}}</div>
          <div class="imageSubHeading"> {{trip.subtitle}}</div>
          <div class="imageSubHeading">{{trip.nodes.length}} Places</div>
        </v-flex>
        <v-layout justify-end align-end>
          <v-btn class="arrow" v-on:click="scrollDown" dark icon><v-icon>arrow_right_alt</v-icon></v-btn>
        </v-layout>
      </v-layout>
    </v-img>

    <v-card-text class="py-0">
      <v-timeline
      color="white"
      align-top
      dense
      >
        <v-timeline-item
            color="accent"
            small
            icon="place"
            v-for="(node, i) in trip.nodes"
            :key="i"
            fill-dot
          > 
            <TripDetailsCard v-bind:node="node"/>
        </v-timeline-item>          
      </v-timeline>
    </v-card-text>
    <div class="text-xs-center">
        <v-btn class="fab" fab color="accent" dark v-on:click="deleteTrip"><v-icon>delete</v-icon></v-btn>
    </div>
  <v-container>
</template>

<script>
import TripDetailsCard from '../components/TripDetailsCard.vue'

export default {
  name: 'TripDetails',
  components: {
      TripDetailsCard
  },
  data() {
    return {
        trip: {
            id: 1,
            title: "Moraine Lake",
            subtitle: "Canada",
            image: "https://images.unsplash.com/photo-1516471835429-167f83503f4b?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=975&q=80",
            nodes: [{
                id: 1,
                title: "Moraine Lake",
                image: "https://images.unsplash.com/photo-1516471835429-167f83503f4b?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=975&q=80"
            },
            {
                id: 2,
                title: "Moraine Lake",
                image: "https://images.unsplash.com/photo-1516471835429-167f83503f4b?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=975&q=80"
            }]
        }
    }
  },
  methods: {
    newPic: function(event) {
        var xhr = new XMLHttpRequest();
        xhr.onload = function () {

	         if (xhr.status >= 200 && xhr.status < 300) {
		           console.log('success!', xhr);
	         } else {
		           console.log('The request failed!');
	         }

	      console.log('This always runs...');
        };
        xhr.open('GET', 'http://reverseproxy:8080/');
        xhr.send();
    },
    deleteTrip() {
      console.log("trip deleted")
    },
    scrollDown() {
      console.log("scroll")
    }
  },
}
</script>

<style scoped>
    .container{
        padding: 0;
    }
    .imageHeading{
        font-size: 40px;
    }
    .imageSubHeading{
        font-size: 20px;
        font-weight: 500;
    }
    .lightbox {
        box-shadow: 0 0 20px inset rgba(0, 0, 0, 0.2);
        background-image: linear-gradient(to top, rgba(0, 0, 0, 0.4) 0%, transparent 72px);
    }
    .fab {
      margin-bottom: 20px
    }
    .arrow {
      transform: rotate(90deg) scale(2);
      margin: 20px;
    }
</style>
