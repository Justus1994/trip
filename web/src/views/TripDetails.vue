<template>
  <v-container class="container" fluid >
    <v-img max-height="821px" min-height="821px" :src=trip.image>
          <v-layout pa-2 column fill-height class="lightbox white--text">
            <v-spacer></v-spacer>
            <v-flex >
              <div class="imageHeading">{{trip.title}}</div>
              <div class="imageSubHeading"> {{trip.subtitle}}</div>
              <div class="imageSubHeading">6 Places</div>
            </v-flex>

            <!--v-btn  round outline color="white" v-on:click="newPic"><router-link to="/newtrip">new Pic</router-link></v-btn-->
          </v-layout>
        </v-img>

        <h3 class="text-xs-center">Place 1</h3>
        <h3 class="text-xs-center">Place 2</h3>
  <v-container>
</template>

<script>
export default {
  name: 'TripDetails',
  data() {
    return {
        trip: {
            id: 1,
            title: "Moraine Lake",
            subtitle: "Canada",
            image: "https://images.unsplash.com/photo-1516471835429-167f83503f4b?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=975&q=80"
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
</style>
