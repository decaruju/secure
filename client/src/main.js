import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import VueRouter from 'vue-router'

import LoginPage from './components/LoginPage.vue'
import SignupPage from './components/SignupPage.vue'

Vue.config.productionTip = false
Vue.use(VueRouter)
Vue.use(VueAxios, axios)

/* const router = new VueRouter({ */
/* mode: 'history', */
/* base: __dirname, */
/* routes: [ */
/* { path: '/', component: App }, */
/* { path: '/login', component: LoginPage }, */
/* { path: '/signup', component: SignupPage }, */
/* ] */
/* }) */

new Vue({
    el: '#app',
    components: { App },
    template: '<App/>'
}).$mount("#app")

