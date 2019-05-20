import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import NewTrip from './views/NewTrip.vue'
import TripDetails from './views/TripDetails.vue'
import TripNodes from './views/TripNodes.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/newtrip',
      name: 'newtrip',
      component: NewTrip
    }, 
    {
        path: '/tripdetails',
        name: 'tripdetails',
        component: TripDetails
    },
    {
        path: '/tripnodes',
        name: 'tripnodes',
        component: TripNodes
    }
  ]
})
