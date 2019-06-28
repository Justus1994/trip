import Vue from "vue";
import { mount } from "@vue/test-utils";
//import VueRouter from 'vue-router';
import Vuetify from 'vuetify';
import Menu from '../Menu.vue';

describe('Components/Menu', function() {


  const mockedDirective = function(el, binding, vnode) {
    var background = (binding.arg === 'background' ? '#1c1d21' : '#333');
    el.style.background = binding.value ? background : '#fcfcfc';
    el.style.color = binding.value ? '#fcfcfc' : '#131313';
  }

  const factory = (propsData) => {
    return mount(Menu, {
      propsData: {
        ...propsData
      },
      directives: {
        darkmode: mockedDirective,
      },
    })
  }

  beforeAll(() => {
    Vue.use(Vuetify);
  });

  test('is a Vue instance', () => {
    const wrapper = factory();
    expect(wrapper.isVueInstance()).toBeTruthy()
  })

  it('diplay darkmode css if darkmode is true', () => {
    const wrapper = factory({
      darkmode: true
    });
    const actualCss = wrapper.find('.menu').attributes().style
    expect(actualCss).toEqual("background: rgb(51, 51, 51); color: rgb(252, 252, 252);");
  });


  it('diplay lightmode css background if darkmode is false', () => {
    const wrapper = factory({
      darkmode: false
    });
    const actualCss = wrapper.find('.menu').attributes().style
    expect(actualCss).toEqual("background: rgb(252, 252, 252); color: rgb(19, 19, 19);");
  });
});
