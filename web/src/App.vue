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
    fetch('auth','GET').then(response => {
          window.localStorage.setItem('Authorization-Token', response);
          /**
          * Simulate loading task :D
          */
          setTimeout(() =>{
            this.loading = false;
          },0);

    }).catch( err => {
      this.loading = false;
      this.$set(this.$root.$data.store.snack,'msg','Upps, Something went wrong.');
      this.$set(this.$root.$data.store.snack, 'show', true)
    });

   },
}
</script>

<style >
</style>
