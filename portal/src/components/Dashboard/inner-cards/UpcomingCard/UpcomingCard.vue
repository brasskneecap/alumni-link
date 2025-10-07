<template>
  <ALCard
    class="upcoming-container"
    title="UPCOMING"
    icon="mdi-calendar-month-outline"
    height="40.875rem"
    width="63rem"
  >

    <template #content>
      <div class="upcoming-items">
        <UpcomingItem v-for="assignment in upcomingAssignments"
          :title="assignment.name"
          :date="formatDate(assignment.dueDate)"
          :tag="assignment.submission ? assignment.submission.status : 'Assignment'"
          :description="assignment.description"
        />

      </div>
    </template>
  </ALCard>
</template>

<script>
// import { defineProps } from 'vue'
import UpcomingItem from './UpcomingItem.vue';
import ALCard from '../../../reusables/ALCard.vue';
import { computed } from 'vue';
import { useStore } from 'vuex';
import { formatDate } from '@/utils/formatters';

export default {
  components: {
      UpcomingItem,
      ALCard,
  },
  setup() {
    const store = useStore()
    const upcomingAssignments = computed(() => {
      const assignments = store.getters["assignments/assignments"]
      return assignments.slice(-6);
    });

    
    return {
      upcomingAssignments,
      formatDate,
    }
  }
}
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;


.upcoming-container {
  margin-left: 5rem;
}

.upcoming-items {
  margin-top: 1rem;
}
</style>