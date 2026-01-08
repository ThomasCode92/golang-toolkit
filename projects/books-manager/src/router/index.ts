import { createRouter, createWebHistory, type RouterOptions } from 'vue-router'

import AppBody from '@/components/AppBody.vue'

const routes: RouterOptions['routes'] = [{ path: '/', name: 'home', component: AppBody }]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})
