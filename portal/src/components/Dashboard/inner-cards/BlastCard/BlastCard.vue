<template>
  <ALCard
    class="blast-container"
    title="Blast"
    icon="mdi-bullhorn-outline"
    height="33.3125rem"
    width="25rem"
  >

    <template #content>
      <div>
        <BlastItem v-for="item in items"
          :date="item.date"
          :title="item.title"
          :description="item.description"
        />
      </div>
     </template> 

  </ALCard>
</template>

<script>
import BlastItem from './BlastItem.vue';
import ALCard from '../../../reusables/ALCard.vue';

import { computed } from 'vue';
import { useStore } from 'vuex';

export default {
  components: {
    BlastItem,
    ALCard,
  },

  setup() {
    const store = useStore();
    const blasts = computed(() => store.getters["blasts/blasts"]);
    const items = [
      {
        date: "JUNE 1",
        title: "Review Session Extended",
        description: "For those of you that haven't signed up for the review session,",
      },
      {
        date: "MAY 16",
        title: "A Cool Ted Talk that is worth watching",
        description: "Create at least 3 user flows related to the video.",
      },
      {
        date: "MAY 3",
        title: "Calendar Update",
        description: "Please check the calendar for the latest updates on upcoming events.",
      },
      {
        date: "APRIL 10",
        title: "Check out this article!",
        description: "Design thinking is not a linear process, but rather a flexible and iterative approach to problem-solving.",
      },
      {
        date: "APRIL 6",
        title: "Welcome Message",
        description: "I'm your assigned Alumni mentor for this semester.",
      },
    ]

    return {
      items,
    };
  },
};

</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.blast-container {
  gap: 1rem;
}


</style>