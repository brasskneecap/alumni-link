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

export default {
  components: {
      UpcomingItem,
      ALCard,
  },
  methods: {
    formatDate(upcomingDate) {
      const date = new Date(upcomingDate);
      const month = new Intl.DateTimeFormat('en-US', { month: 'long' }).format(date);
      const day = date.getUTCDate()
      
      const formattedDate = `${month} ${day}`
      return formattedDate.toUpperCase()
    }
  },
  setup() {
    const store = useStore()
    const assignments = computed(() => store.getters["assignments/assignments"]);

    return {
      assignments,
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