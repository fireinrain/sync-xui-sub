import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/xui-node',
      name: 'xui-node',
      component: () => import('../views/XuiNodeView.vue')
    },
    {
      path: '/user-sub',
      name: 'user-sub',
      component: () => import('../views/UserSubView.vue')
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue')

    },
    {
      path: '/logout',
      name: 'logout',
      component: () => import('../views/LogoutView.vue')

    }
  ]
})

export default router
