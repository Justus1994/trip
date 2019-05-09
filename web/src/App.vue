<template>
  <v-container class="container" fluid >
    <v-img max-height="821px" min-height="821px" src="https://images.unsplash.com/photo-1552733407-5d5c46c3bb3b?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1014&q=80">
          <v-layout pa-2 column fill-height class="lightbox white--text">
            <v-spacer></v-spacer>
            <v-flex >
              <div class="imageHeading">Siargao island</div>
              <div class="imageSubHeading"> Philippines</div>
              <div class="imageSubHeading"> {{aspect}}</div>
            </v-flex>


                <v-btn  round outline color="white" v-on:click="newPic">new Pic</v-btn>
          </v-layout>


        </v-img>
  <v-container>
</template>
<script>
export default {
  name: 'App',
  data() {
    return {
      aspect: '9/19.5',
    };
  },
  methods: {
    newPic: function (event) {
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
};
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
