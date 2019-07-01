<template lang="html">
  <!-- Menu: -->
  <div v-darkmode="darkmode" class="menu">
    <v-btn v-darkmode="darkmode" flat @click="$emit('toggleDarkmode')">{{darkmode? 'lightmode': 'darkmode'}}</v-btn>
    <v-btn v-darkmode="darkmode" flat @click="$emit('closeMenu')">Close</v-btn>
  </div>
</template>

<script>
export default {
    name: "Menu",
    props: ['darkmode'],
}
</script>

<style lang="css" scoped>
  .menu{
    transition: all 0.5s cubic-bezier(0.25, 0.1, 0, 1.2);
    position: absolute;
    height: 100vh;
    top: 0;
    left: -33%;
    width: 33%;
    display: flex;
    flex-direction: column;
    justify-content: center;
  }
</style>
