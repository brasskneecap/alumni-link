import { createWebHashHistory, createRouter } from 'vue-router'

import Dashboard from '../components/Dashboard/Dashboard.vue'
import Assignments from '../components/Assignments/Assignments.vue'
import Calendar from '../components/Calendar/Calendar.vue'
import Messages from '../components/Messages/Messages.vue'
import Blast from '../components/Blast/Blast.vue'
import Groups from '../components/Groups/Groups.vue'

const routes = [
  { path: '/', component: Dashboard },
  { path: '/assignments', component: Assignments },
  { path: '/calendar', component: Calendar },
  { path: '/messages', component: Messages },
  { path: '/blast', component: Blast },
  { path: '/groups', component: Groups },
]

const router = createRouter({
  linkActiveClass: 'border-indigo-100',
  history: createWebHashHistory(),
  routes,
})

export default router