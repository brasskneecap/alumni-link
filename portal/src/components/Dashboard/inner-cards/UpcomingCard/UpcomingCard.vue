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
        <UpcomingItem v-for="item in items"
          :title="item.title"
          :date="item.date"
          :tag="item.tag"
          :description="item.description"
        />

        <!-- <UpcomingItem v-for="assignment in assignments"
          :title="assignment.name"
          :date="assignment.dueDate"
          :tag="assignment.submission.status"
          :description="assignment.description"
        /> -->

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

  setup() {
    const store = useStore()
    const assignments = computed(() => store.getters["assignments/assignments"]);
    const items = [
      {
        title:"Create LinkedIn Profile and Upload",
        date:"APRIL 7",
        tag:"Assignment",
        description:"Create your LinkedIn profile and upload the link",
      },
      {
        title:"Create LinkedIn Profile and Upload",
        date:"APRIL 7",
        tag:"Assignment",
        description:"Create your LinkedIn profile and upload the link",
      },
      {
        title:"Create LinkedIn Profile and Upload",
        date:"APRIL 7",
        tag:"Assignment",
        description:"Create your LinkedIn profile and upload the link",
      },
      {
        title:"Create LinkedIn Profile and Upload",
        date:"APRIL 7",
        tag:"Assignment",
        description:"Create your LinkedIn profile and upload the link",
      },
      {
        title:"Create LinkedIn Profile and Upload",
        date:"APRIL 7",
        tag:"Assignment",
        description:"Create your LinkedIn profile and upload the link",
      },
    ]
    
    console.log(assignments)
    return {
      items,
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