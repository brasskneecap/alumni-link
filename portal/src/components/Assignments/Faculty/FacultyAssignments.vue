<template>
  <ALCard
    class="assignments-container"
    title="ASSIGNMENTS OVERVIEW"
    icon="mdi-calendar-month-outline"
    height="59.875rem"
    width="96.375rem"
  >
    <template #header-right>
      <router-link to="/assignments-overview/create">
        <v-btn color="#4065DA">Add Assignment</v-btn>
      </router-link>
    </template>
  </ALCard>
</template>

<script setup>
import { ref, computed } from 'vue'
import ALCard from '../../reusables/ALCard.vue'

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