import Vue from "vue";
import { mount } from "@vue/test-utils";
//import VueRouter from 'vue-router';
import Vuetify from 'vuetify';
import Controls from '../Controls.vue';

describe('Components/Controls', function () {

  const darkmode = false;

  const mockedDirective = function(el, binding, vnode){
       var background = (binding.arg === 'background' ? '#1c1d21' : '#333');
       el.style.background  = binding.value ? background : '#fcfcfc';
       el.style.color = binding.value ? '#fcfcfc': '#131313';
  }
  addElemWithDataAppToBody();

  const factory = (propsData) => {
  return mount(Controls, {
    propsData: {
      darkmode,
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
  it('Display a loader if loading is true', () => {
    const wrapper = factory({loading : true});
    expect(wrapper.find('.loader').exists()).toBe(true);

  });

  it('Display no loader if loading is false', () => {
    const wrapper = factory({loading : false});
    expect(wrapper.find('.loader').exists()).toBe(false);

  });

  it('Display no loader if loading is false', () => {
    const wrapper = factory({loading : false});
    expect(wrapper.find('.loader').exists()).toBe(false);

  });
});

/**
* Workaorund to elimate waring
* [Vuetify] Unable to locate target [data-app] in "v-dialog"
* https://forum.vuejs.org/t/vuetify-data-app-true-and-problems-rendering-v-dialog-in-unit-tests/27495
*/
function addElemWithDataAppToBody() {
  const app = document.createElement('div');
  app.setAttribute('data-app', true);
  document.body.append(app);
};
