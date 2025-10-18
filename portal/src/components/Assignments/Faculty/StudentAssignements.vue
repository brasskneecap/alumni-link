<template>
  <div class="student-assignment-items">
    <UpcomingItem v-for="assignment in studentAssignments"
      :key="assignment.id"
      :title="assignment.name"
      :date="formatDate(assignment.dueDate)"
      :tag="assignment.submission ? assignment.submission.status : 'Assignment'"
      :description="assignment.description"
    />
  </div>
</template>

<script>
import UpcomingItem from '../../Dashboard/inner-cards/UpcomingCard/UpcomingItem.vue';
import { computed } from 'vue';
import { useStore } from 'vuex';
import { formatDate } from '@/utils/formatters';

export default {
  components: {
      UpcomingItem,
  },
  setup() {
    const store = useStore()
    const studentAssignments = computed(() => {
      const assignments = store.getters["assignments/assignments"]
      return assignments;
    });

    
    return {
      studentAssignments,
      formatDate,
    }
  }
}
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.student-assignment-items {
  margin-top: 1rem;
}
</style>
