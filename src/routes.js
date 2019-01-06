import Vue from 'vue';
import VueRouter from 'vue-router';

import Artists from './components/Artists';
import Albums from './components/Albums';
import Settings from './components/Settings';

Vue.use(VueRouter);

export default new VueRouter({
  mode: 'history',
  routes: [
    {path: '/', name: 'Artists', component: Artists},
    {path: '/albums', name: 'Albums', component: Albums},
    {path: '/settings', name: 'Settings', component: Settings}
  ]
});
