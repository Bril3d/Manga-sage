import { createRouter, createWebHistory } from 'vue-router'
import App from "../App.vue"
import LoginView from "../pages/LoginView.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'App',
      component: App
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
  ]
})

export default router
