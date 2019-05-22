<template>
    <v-bottom-nav
        :active.sync="bottomNav"
        :value="showNavBar"
        class="bottomNavBar"
        app
        color="white"
        >
        <v-btn
            color="accent"
            flat
            value="home"
            to="/"
        >
            <v-icon>home</v-icon>
        </v-btn>
    
        <v-btn
            color="accent"
            flat
            value="favorites"
            to="/newtrip"
        >
            <v-icon>favorite</v-icon>
        </v-btn>
    
        <!--v-btn
            color="accent"
            flat
            value="nearby"
        >
            <v-icon>history</v-icon>
        </v-btn-->
      </v-bottom-nav>
</template>

<<script>
export default {
    name: "BottomNavigation",
    data () {
        return {
            bottomNav: 'home',
            showNavBar: false
        }
    }
}
</script>

<style scoped>

</style>
