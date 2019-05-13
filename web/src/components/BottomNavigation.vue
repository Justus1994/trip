<template>
    <v-bottom-nav
        :active.sync="bottomNav"
        :value="true"
        app
        fixed
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
            bottomNav: 'home'
        }
    }
}
</script>

<style scoped>

</style>
