<template>
    <div>
        
        <v-container class="container" fluid>
        <v-img max-height="821px" min-height="821px" :src=trip.image>
            <v-layout pa-2 column fill-height class="lightbox white--text">
            <v-spacer></v-spacer>
                <v-flex >
                <div class="imageHeading">{{trip.title}}</div>
                <div class="imageSubHeading"> {{trip.subtitle}}</div>
                </v-flex>
            <v-spacer></v-spacer>
            <v-spacer></v-spacer>
          
                <v-flex class="text-xs-center">
                <v-btn
                fab
                right
                outline
                color="white"
                >
                    <v-icon>close</v-icon>
                </v-btn>
                <v-btn
                fab
                left
                outline
                color="white"
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
export default {
    data() {
        return {
            trip: {
                
                id: 1,
                title: "Black Forest",
                subtitle: "Germany",
                image: "https://images.unsplash.com/photo-1506663728000-f68915a9e635?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1950&q=80"
              
            }
        }
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
