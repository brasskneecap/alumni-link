<template>
  <ALCard
    title="Your Team"
    icon="mdi-account-group-outline"
    width="25rem"
    height="12.5rem"
  >
    <template #content>
      <div class="team-members">
        <TeamMember 
          v-for="member in getTeam()"
          :src="member.profilePicture"
          :name="member.name"
          :role="member.role"
          :description="`${member.role} at ${member.school}`"
        />
      </div>
    </template>
  </ALCard>
</template>

<script>
import TeamMember from './TeamMember.vue';
import ALCard from '../../../reusables/ALCard.vue';
import { computed, toRefs } from 'vue';
import { useStore } from 'vuex';

export default {
    components: {
      ALCard,
      TeamMember,
    },

    methods: {
      getTeam() {
        const facultyRoles = ['Mentor', 'Faculty']
        const team = this.users.filter(user => facultyRoles.includes(user.role))
        return team
      },  
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

      const users = computed(() => store.getters["users/users"]);
      
      console.log('USERS', users)
      return {
        team,
        users,
      }
    }
}
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

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