<template>
  <ALCard
    class="assignments-container"
    title="UPCOMING ASSIGNMENTS"
    icon="mdi-calendar-month-outline"
    height="59.875rem"
    width="96.375rem"
  >
  
    <template #content>
      <AssignmentsList />
    </template>
  </ALCard>
</template>

<script>
import { computed } from 'vue';
import { useStore } from 'vuex'
import ALCard from '../reusables/ALCard.vue';
import AssignmentsList from './AssignmentsList.vue'; 

export default {
  components: {
  ALCard,
  AssignmentsList,
  
  },
  setup() {
    const store = useStore()
    
    const assignments = computed(() => store.getters["assignments/assignments"]);
    
    return {
      assignments
    }
  }
}
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.assignments-container {
  margin: 7.12rem 2rem 2.81rem 2rem;
}

</style>