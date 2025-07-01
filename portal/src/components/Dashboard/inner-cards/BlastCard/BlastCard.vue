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
        <BlastItem v-for="blast in blasts"
          :date="formatDate(blast.createdAt)"
          :title="blast.title"
          :description="blast.description"
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
import { formatDate } from '@/utils/formatters';

export default {
  components: {
    BlastItem,
    ALCard,
  },
  setup() {
    const store = useStore();
    const blasts = computed(() => store.getters["blasts/blasts"]);

    return {
      blasts,
      formatDate,
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