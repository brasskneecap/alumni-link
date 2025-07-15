<template>
  <div>
    <div>
      <h1 class="title">{{ assignment.name }}</h1>
      <div class="dueDate">
        <v-icon icon="mdi-calendar-month-outline" size="24"/>
        <p>Due {{ formatDateMDY(assignment.dueDate) }}</p>
      </div>
      <TeamMember
        v-if="teamMember"
        :src="teamMember.profilePicture"
        :name="teamMember.name"
        :role="teamMember.role"
        :description="`${teamMember.role} at ${teamMember.school}`"
      />
      <p class="description">{{ assignment.description }}</p>
    </div>
    <div>
      <h2 class="subtitle">Submit Assignment</h2>
    </div>
  </div>
  <div class="feedback">
    <v-btn variant="outlined" v-ripple="false" class="text-none">
      Feedback
    </v-btn>
  </div>
</template>

<script setup>
import TeamMember from '../Dashboard/inner-cards/TeamCard/TeamMember.vue';

const props = defineProps({
  assignment: Object,
  teamMember: Object
});

function formatDateMDY(dateStr) {
  const date = new Date(dateStr)
  const month = String(date.getMonth() + 1)
  const day = String(date.getDate())
  const year = date.getFullYear().toString().slice(-2)
  return `${month}/${day}/${year}`
}
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.title {
  font-size: 24px;
  font-weight: 700;
}

.dueDate {
  display: flex;
  align-items: center;
  font-size: 24px;
  color: $al-secondary-gray;
  gap: 0.3rem;
  margin: 0.5rem 0 1.8rem 0;
}

.description {
  margin-top: 2.3rem;
}

.subtitle {
  font-size: 16px;
  font-weight: 500;
  margin-top: 1.5rem;
}

.feedback {
  display: flex;
  justify-content: flex-end;
  margin-right: 4rem;
  font-family: 'Poppins', sans-serif !important;
}
</style>