<template>
  <div class="dashboard-container">

    <div class="notif-profile-container">

    </div>

    <div class="left-side">
      <!-- welcome message, settings, notification -->
      <div class="header-container">
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
        
        <div v-for="team in teams" class="team-container al-card">
          <TeamCard
            :icon="team.icon"
            :title="team.title"
            :isHeader="true"
          />
          <div class="team-members">
            <div v-for="member in team.members">
              <TeamCard
                :name="member.name"
                :role="member.role"
                :description="member.description"
                :isHeader="false"
              />
            </div>
          </div>
        </div>
      </div>

      <!-- upcoming things list -->
      <div v-for="list in upcoming" class="upcoming-container al-card">
        <UpcomingCard
          :icon="list.icon"
          :title="list.title"
          :isHeader="true"
        />
        <div class="list-items">
          <div v-for="item in list.items">
            <UpcomingCard 
              :date="item.date"
              :itemTitle="item.itemTitle"
              :itemTag="item.itemTag"
              :itemSubtitle="item.itemSubtitle"
              :isHeader="false"
            />
          </div>
        </div>
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
import TeamCard from './inner-cards/TeamCard.vue';
import UpcomingCard from './inner-cards/UpcomingCard.vue';
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
    components: {
      AssignmentCard,
      TeamCard,
      UpcomingCard,

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

    const teams = [
      {
        icon:"mdi-account-group-outline",
        title:"Your Team",
        members: [
          {
            src:"#",
            name:"Mike Harper",
            role:"Faculty",
            description:"Professor at Utah Valley University",
          },
          {
            src:"#",
            name:"Mike Harper",
            role:"Faculty",
            description:"Professor at Utah Valley University",
          }
        ]
      }
    ]

    const upcoming = [
      {
        icon:"mdi-calendar-month-outline",
        title:"UPCOMING",
        items: [
          {
            date:"APRIL 7",
            itemTitle:"Create LinkedIn Profile and Upload",
            itemTag:"Assignment",
            itemSubtitle:"Create your LinkedIn profile and upload the link",
          },
          {
            date:"APRIL 7",
            itemTitle:"Create LinkedIn Profile and Upload",
            itemTag:"Assignment",
            itemSubtitle:"Create your LinkedIn profile and upload the link",
          },
          {
            date:"APRIL 7",
            itemTitle:"Create LinkedIn Profile and Upload",
            itemTag:"Assignment",
            itemSubtitle:"Create your LinkedIn profile and upload the link",
          },
          {
            date:"APRIL 7",
            itemTitle:"Create LinkedIn Profile and Upload",
            itemTag:"Assignment",
            itemSubtitle:"Create your LinkedIn profile and upload the link",
          },
          {
            date:"APRIL 7",
            itemTitle:"Create LinkedIn Profile and Upload",
            itemTag:"Assignment",
            itemSubtitle:"Create your LinkedIn profile and upload the link",
          },
        ]
      }
    ]

    return {
      user,
      assignments,
      teams,
      upcoming,
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
  }
}

.upcoming-container {
  width: 63rem;
  height: 40.625rem;
  margin: 2.66rem 0 0 5rem;
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