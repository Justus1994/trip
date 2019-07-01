import Vue from "vue";
import { mount } from "@vue/test-utils";
//import VueRouter from 'vue-router';
import Vuetify from 'vuetify';
import NodeCard from '../NodeCard.vue';
import TestTrip from '../../TestTrip.js'

describe('Components/NodeCard', function() {
  let wrapper;

  const mockedDirective = function(el, binding, vnode) {

    var background = (binding.arg === 'background' ? '#1c1d21' : '#333');
    el.style.background = binding.value ? background : '#fcfcfc';
    el.style.color = binding.value ? '#fcfcfc' : '#131313';
  }

  beforeAll(() => {
    Vue.use(Vuetify);
    wrapper = mount(NodeCard, {
      propsData: {
        darkmode: false,
        node: TestTrip.Nodes[0],

      },
      directives: {
        darkmode: mockedDirective,
      },
    });
  });

  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
  });

  it('Display all elements if there is a Node', () => {
    expect(wrapper.find('.card').exists()).toBe(true);
    expect(wrapper.find('.v-image').exists()).toBe(true);
    expect(wrapper.find('.headline').exists()).toBe(true);
  });

});
