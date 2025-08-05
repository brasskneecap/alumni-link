import { createWebHashHistory, createRouter } from 'vue-router'

import Dashboard from '../components/Dashboard/dashboard.vue'
import Assignments from '../components/Assignments/Assignments.vue'
import Calendar from '../components/Calendar/Calendar.vue'
import Messages from '../components/Messages/Messages.vue'
import Blast from '../components/Blast/Blast.vue'
import Groups from '../components/Groups/Groups.vue'
import FacultyAssignments from '../components/Assignments/Faculty/FacultyAssignments.vue'
import CreateAssignment from '../components/Assignments/Faculty/CreateAssignment.vue'

const routes = [
  { path: '/', component: Dashboard },
  { path: '/assignments', component: Assignments },
  { path: '/assignments-overview', component: FacultyAssignments },
  { path: '/assignments-overview/create', component: CreateAssignment },
  { path: '/calendar', component: Calendar },
  { path: '/messages', component: Messages },
  { path: '/blast', component: Blast },
  { path: '/groups', component: Groups },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router