import { createWebHashHistory, createRouter } from 'vue-router'

import Dashboard from '../components/Dashboard/dashboard.vue'

const routes = [
  { path: '/', component: Dashboard },
  // { path: '/about', component: AboutView },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router