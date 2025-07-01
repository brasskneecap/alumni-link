<template>
  <div class="dashboard-container">

    <!-- Left Side Content -->
    <div class="left-side">
      <!-- Header section -->
      <div class="header-container">
         <div class="welcome-container">
          <h2>Welcome Back, {{ user.name }}</h2>
          <p>Take a look at your progress</p>
        </div>
      </div>

      <!-- Tracker section -->
      <div class="tracker-container">
        <AssignmentCard 
          title="This Week"          
          :content="weeklyCompleted()"
        />
        <AssignmentCard 
          title="Total Completed"
          :content="totalCompleted()"
        />
        <!-- Container for Your Team Section-->
        <TeamCard />
      </div>

      <!-- Upcoming section -->
      <div class="upcoming-wrapper">
        <UpcomingCard />
      </div>
    </div>

      <div class="right-side">
        <!-- calendar and blast -->
        <div class="right-container">
          <div class="calendar-container al-card">

          </div>
          <BlastCard />
        </div>
      </div>
    </div>
</template>

<script>
import AssignmentCard from './inner-cards/AssignmentCard.vue';
import TeamCard from './inner-cards/TeamCard/TeamCard.vue';
import { computed } from 'vue'
import { useStore } from 'vuex'
import UpcomingCard from './inner-cards/UpcomingCard/UpcomingCard.vue';
import BlastCard from './inner-cards/BlastCard/BlastCard.vue';

export default {
  components: {
    AssignmentCard,
    TeamCard,
    UpcomingCard,
    BlastCard
  },
  methods: {
    totalCompleted() {
      let completed = 0
      this.assignments.forEach((assignment) => {
        if (assignment.submission) {
          completed++
        }
      });
      
      return `${completed}/${this.assignments.length}`
    },
    weeklyCompleted() {
      const now = new Date()
      const currentDay = now.getUTCDay()
      
      const daysSinceMonday = currentDay === 0 ? 6 : currentDay - 1

      const startOfWeek = new Date(Date.UTC(
        now.getUTCFullYear(),
        now.getUTCMonth(),
        now.getUTCDate() - daysSinceMonday,
        0, 0, 0, 0
      ))
      const endOfWeek = new Date(Date.UTC(
        now.getUTCFullYear(),
        now.getUTCMonth(),
        now.getUTCDate() - daysSinceMonday + 6,
        23, 59, 59, 999
      ))

      const weeklyAssignments = this.assignments.filter((assignment) => {
        const dueDate = new Date(assignment.dueDate)
        return dueDate >= startOfWeek && dueDate <= endOfWeek
      })

      let completed = 0

      weeklyAssignments.forEach((assignment) => {
        if (assignment.submission) {
          completed++
        }
      });

      return `${completed}/${weeklyAssignments.length}`  
    },
  },
  setup () {
    const store = useStore()
    const user = computed(() => store.getters["user/user"]);
    const assignments = computed(() => store.getters["assignments/assignments"]);

    return {
      user,
      assignments,
    }
  }
}
</script>

<style lang="scss">
@use '@/variables.scss' as *;

.dashboard-container {
  display: flex;
  flex-direction: row;
  min-height: 69.8125rem;
}

.header-container {

    .welcome-container {
    margin: 1.88rem 0 0 5rem;
    
    h2 {
      font-size: 48px;
      font-weight: 300;
    }

    p {
      font-size: 20px;
      font-weight: 300;
    }
  }
}

.left-side {
  display: flex;
  flex-direction: column;
  gap: 2.5rem;
}

.tracker-container {
  display: flex;
  margin-left: 5rem;
  gap: 3.75rem;
}

.right-container {
  display: flex;
  flex-direction: column;
  margin: 10.90rem 0 0 3.75rem;
  gap: 2.5rem;

  .calendar-container {
    width: 25rem;
    height: 19.8125rem;
  }
}

</style>