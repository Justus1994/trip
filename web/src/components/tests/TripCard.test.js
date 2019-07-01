import Vue from "vue";
import { mount } from "@vue/test-utils";
//import VueRouter from 'vue-router';
import Vuetify from 'vuetify';
import TripCard from '../TripCard.vue';
import TestTrip from '../../TestTrip.js'

describe('Components/TripCard', function() {
  let wrapper;

  const mockedDirective = function(el, binding, vnode) {

    var background = (binding.arg === 'background' ? '#1c1d21' : '#333');
    el.style.background = binding.value ? background : '#fcfcfc';
    el.style.color = binding.value ? '#fcfcfc' : '#131313';
  }

  beforeAll(() => {
    Vue.use(Vuetify);

    wrapper = mount(TripCard, {
      propsData: {
        darkmode: false,
        trip: TestTrip,
        index: 0
      },
      directives: {
        darkmode: mockedDirective,
      },
    });
  });


  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
  })

  it('Display all elements if there is a trip', () => {
    expect(wrapper.find('.card').exists()).toBe(true);
    expect(wrapper.find('.v-image').exists()).toBe(true);
    expect(wrapper.find('button').exists()).toBe(true);
    expect(wrapper.find('.v-card__title').exists()).toBe(true);
    expect(wrapper.find('.headline_tripcard').exists()).toBe(true);
    expect(wrapper.find('.subhead_tripcard').exists()).toBe(true);

  });
});
