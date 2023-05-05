import { createRouter, createWebHistory } from 'vue-router'
import LoginView from "../pages/LoginView.vue"
import FeedView from "../pages/FeedView.vue"
import RegisterView from "../pages/RegisterView.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'feed',
      component: FeedView
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView
    }
  ]
})

export default router
