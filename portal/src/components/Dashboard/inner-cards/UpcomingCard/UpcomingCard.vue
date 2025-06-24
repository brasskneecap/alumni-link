<template>
  <ALCard
    class="upcoming-container"
    title="UPCOMING"
    icon="mdi-calendar-month-outline"
    height="40.625rem"
    width="63rem"
  >

    <template #content>
      <div class="upcoming-items">
        <UpcomingItem v-for="assignment in assignments"
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
    const assignments = computed(() => store.getters["assignments/assignments"]);

    return {
      assignments,
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