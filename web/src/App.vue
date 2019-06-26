<template>
  <v-app>
    <v-content v-if="!loading">
    <!--<v-content>-->
     <div id="app">
      <router-view/>
      </div>
    </v-content>
    <Loading v-else></Loading>
    <Snackbar></Snackbar>
  </v-app>
</template>

<script>
import Loading from './components/Loading'
import Snackbar from './components/Snackbar'
export default {
  name: 'App',
  components: {
    Loading,
    Snackbar,
  },
  data() {
    return {
      loading: true,
    }
  },
  beforeCreate() {
    this.$root.$data.store.error = true;
    fetch('api/auth',{
      headers:{
        'Authorization' : window.localStorage.getItem('Authorization-Token')
      }
    }).then(response => {
        if(response.status >= 300){
          this.loading = false;
          this.$set(this.$root.$data.store.snack,'msg','Upps, Something went wrong.');
          this.$set(this.$root.$data.store.snack, 'show', true);
        }
        response.text().then(token => {
          window.localStorage.setItem('Authorization-Token', token);
          /**
          * Simulate loading task :D
          */
          setTimeout(() =>{
            this.loading = false;
          },0);
        });
    });

   },
}
</script>

<style >
</style>
