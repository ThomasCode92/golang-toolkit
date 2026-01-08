import { createRouter, createWebHistory, type RouterOptions } from 'vue-router'

import AppBody from '@/components/AppBody.vue'
import LoginForm from '@/components/LoginForm.vue'

const routes: RouterOptions['routes'] = [
  { path: '/', name: 'home', component: AppBody },
  { path: '/login', name: 'login', component: LoginForm },
]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})
