<template>
  <div class="team-card al-card">
    <div class="team-header">
      <div class="team-icon">
        <v-icon 
          icon="mdi-account-group-outline" 
          class="al-icon" 
          size="22">
        </v-icon>
      </div>
      <div class="title-container">
        <h3 class="team-title">Your Team</h3>
      </div>
    </div>
    <div class="team-members">
      <TeamMember 
        v-for="member in team"
        :src="member.src"
        :name="member.name"
        :role="member.role"
        :description="member.description"
      />
    </div>
  </div>
</template>

<script>
import TeamMember from './TeamMember.vue';
import { computed } from 'vue';
import { useStore } from 'vuex';

export default {
    components: {
      TeamMember,
    },
    setup () {
      const store = useStore()
      const team = [
        {
          src:"/vite.svg",
          name:"Mike Harper",
          role:"Faculty",
          description:"Professor at Utah Valley University",
        },
        {
          src:"uvu-logo.svg",
          name:"John Doe",
          role:"Mentor",
          description:"Professor at Utah Valley University",
        }
      ]

      return {
        team,
      }
    }
}
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.team-card {
  padding: 0.625rem 0.875rem;
  display: flex;
  flex-direction: column;
  width: 25rem;
  height: 12.5rem;
}

.team-header {
  display: flex;
  flex-direction: row;
  align-items: center;
  height: 3rem;

  .team-icon {
    display: flex;
    align-items: center;
  }

  .title-container {
    display: flex;
    align-items: center;

    .team-title {
      margin-left: 1rem;
      font-size: 16px;
      font-weight: 700;
    }
  }
}

.team-members {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}
</style>