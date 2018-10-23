import Vue from 'vue'
import Router from 'vue-router'
import LoginComponent from '../views/login.vue'
import SecureComponent from '../views/secure.vue'
import VueCookies from 'vue-cookies'

Vue.use(Router)
Vue.use(VueCookies)

export default new Router({
  routes: [
    {
      path: '/',
      redirect: {
        name: 'login'
      }
    },
    {
      path: '/login',
      name: 'login',
      component: LoginComponent
    },
    {
      path: '/secure',
      name: 'secure',
      component: SecureComponent
    }
  ]
})
