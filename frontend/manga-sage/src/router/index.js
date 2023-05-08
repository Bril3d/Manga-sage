import { createRouter, createWebHistory } from 'vue-router'
import LoginView from "../pages/LoginView.vue"
import RegisterView from "../pages/RegisterView.vue"
import DashboardLayout from "../pages/dashboard/DashboardLayout.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'feed',
      component: () => import('../pages/FeedView.vue')
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
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: DashboardLayout
    }
  ]
})

export default router
