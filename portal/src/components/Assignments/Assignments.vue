<template>
  <ALCard
    class="assignments-container"
    title="ASSIGNMENTS"
    icon="mdi-calendar-month-outline"
    height="59.875rem"
    width="96.375rem"
  >
  
    <template #content>
      <div class="assignments-flex">
        <AssignmentsList @show-assignment-view="handleShowAssignmentView" />
        <AssignmentsView v-if="isVisible" :assignment="selectedAssignment" :teamMember="selectedTeamMember"/>
      </div>
    </template>
  </ALCard>
</template>

<script setup>
import { ref, computed } from 'vue'
import ALCard from '../reusables/ALCard.vue'
import AssignmentsList from './AssignmentsList.vue'
import AssignmentsView from './AssignmentsView.vue'

const isVisible = ref(false)
const selectedAssignment = ref(null)
const selectedTeamMember = ref({
  profilePicture: '/path/to/image.jpg',
  name: 'John Doe',
  role: 'Developer',
  school: 'Utah Valley University'
});

function handleShowAssignmentView(assignment) {
  selectedAssignment.value = assignment
  isVisible.value = true
}
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.assignments-container {
  margin: 7.12rem 2rem 2.81rem 2rem;
}

.assignments-flex {
  display: flex;
  flex-direction: row;
  gap: 2rem;
  width: 100%;
}

.assignments-flex > *:first-child {
  flex: 0 0 32rem;
}

.assignments-flex > *:last-child {
  flex: 1 1 0;
}
</style>