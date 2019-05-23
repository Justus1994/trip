<template>
    <div>

        <v-container class="container" fluid>
        <v-img max-height="821px" min-height="821px" :src=currentNode.urls.regular>
            <v-layout pa-2 column fill-height class="lightbox white--text">
            <v-spacer></v-spacer>
                <v-flex >
                <div class="imageHeading">{{currentNode.location.title}}</div>
                <div class="imageSubHeading"> {{currentNode.location.country}}</div>
                </v-flex>
            <v-spacer></v-spacer>
            <v-spacer></v-spacer>

                <v-flex class="text-xs-center">
                <v-btn
                fab
                right
                outline
                color="white"
                v-on:click="nextNode(false)"
                >
                    <v-icon>close</v-icon>
                </v-btn>
                <v-btn
                fab
                left
                outline
                color="white"
                v-on:click="nextNode(true)"
                >
                    <v-icon>favorite</v-icon>
                </v-btn>
                </v-flex>

            </v-layout>
            </v-img>

    <v-container>
    </div>
</template>

<script>
import store from '../store.js'

export default {
    data() {
        return {
            currentNode: {},
            trips: this.$root.$data.sharedState.pendingTrip,
            counter: 0
        }
    },
    methods: {
        nextNode(like) {
            like ? console.log("Node like and added") : console.log("Node dislike and skipped")
            this.trips.Nodes.length > this.counter ? this.currentNode = this.trips.Nodes[this.counter++] : this.$router.push('/tripdetails')
        }
    },
    created() {
        console.log(this.trips)
        this.currentNode = this.trips.Nodes[this.counter++]
    },
}
</script>

<style >
    .bottomNavBar {
        display: none;
    }
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
