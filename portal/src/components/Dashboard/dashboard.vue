<template>
  <div class="dashboard-container">
    <div class="notif-profile-container">
    </div>

    <!-- Left Side Content -->
    <div class="left-side">
      <!-- Header section -->
      <div class="header-container">
         <div class="welcome-container">
          <h2>Welcome Back, {{ user.name }}</h2>
          <p>Take a look at your progress</p>
        </div>
        <div class="profile-container">

        </div>
      </div>

      <!-- Tracker section -->
      <div class="tracker-container">
        <div v-for="assignment in assignments" class="this-week-container al-card">
          <AssignmentCard
            :icon="assignment.icon"
            :title="assignment.title"
            :content="assignment.content"
          />
        </div>      
        <div class="team-container">
          <TeamCard />
        </div>
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
          <div class="blast-container al-card">

          </div>
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

export default {
    components: {
      AssignmentCard,
      TeamCard,
      UpcomingCard

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

  .this-week-container {
    width: 15.25rem;
    height: 12.5rem;
  }
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

  .blast-container {
    width: 25rem;
    height: 33.3125rem;
  }
}

</style>