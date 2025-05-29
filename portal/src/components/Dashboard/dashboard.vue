<template>
  <div class="dashboard-container">

    <div class="notif-profile-container">

    </div>

    <div class="left-side">
      <!-- welcome message, settings, notification -->
      <div class="topbar-container">
        <div class="welcome-container">
          <h2>Welcome Back, {{ user.name }}</h2>
          <p>Take a look at your progress</p>
        </div>
        <div class="profile-container">

        </div>
      </div>

      <!-- assignments and team list -->
      <div class="tracker-container">
        <div v-for="assignment in assignments" class="this-week-container al-card">
          <AssignmentCard
            :icon="assignment.icon"
            :title="assignment.title"
            :content="assignment.content"
          />
        </div>
        <div class="team-container">

        </div>
      </div>

      <!-- upcoming things list -->
      <div class="upcoming-container"></div>
    </div>

    <div class="right-side">
      <!-- calendar and blast -->
      <div class="right-container">
        <div class="calendar-container">

        </div>
        <div class="blast-container">

        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AssignmentCard from './AssignmentCard.vue';
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
    components: {
      AssignmentCard,
  },
  setup () {
    const store = useStore()

    const user = computed(() => store.getters["user/user"]);
    const assignments = [
      {
        icon:"mdi-file-document-multiple-outline",
        title:"This Week",
        content:"1/3"},
      {
        icon:"mdi-file-document-multiple-outline",
        title:"Total Completed",
        content:"8/16",
      }
    ]

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

.notif-profile-container {
  background-color: grey;
  width: 7rem;
  height: 3.75rem;
  position: absolute;
  right: 0;
  margin-top: 1.88rem;
  margin-right: 4.5rem;
}

.topbar-container {

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

.tracker-container {
  display: flex;
  margin: 2.62rem 0 0 5rem;
  gap: 3.75rem;

  .this-week-container, .completed-container {
    width: 15.25rem;
    height: 12.5rem;
  }
  
  .team-container {
    width: 25rem;
    height: 12.5rem;
    background-color: aquamarine;
  }
}

.upcoming-container {
  width: 63rem;
  height: 40.625rem;
  background-color: lightcoral;
  margin: 2.66rem 0 0 5rem;
}

.right-container {
  display: flex;
  flex-direction: column;
  margin: 10.90rem 0 0 3.75rem;
  gap: 2.5rem;

  .calendar-container {
    background-color: lightblue;
    width: 25rem;
    height: 19.8125rem;
  }

  .blast-container {
    background-color: green;
    width: 25rem;
    height: 33.3125rem;
  }
}

</style>