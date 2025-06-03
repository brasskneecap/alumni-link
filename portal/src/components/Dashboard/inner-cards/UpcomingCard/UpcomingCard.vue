<template>
  <div class="upcoming-container al-card upcoming-card">
    <div class="upcoming-header">
      <div class="upcoming-icon">
        <v-icon 
          icon="mdi-calendar-month-outline" 
          class="al-icon" 
          size="22">
        </v-icon>
      </div>
      <div>
        <div class="title-container">
          <h3 class="upcoming-title">UPCOMING</h3>
        </div>
      </div>
    </div>

    <div class="upcoming-items">
      <UpcomingItem v-for="item in items"
        :title="item.title"
        :date="item.date"
        :tag="item.tag"
        :description="item.description"
      />
      <!-- <UpcomingItem v-for="assignment in assignments"
        :title="item.title"
        :date="item.date"
        :tag="item.tag"
        :description="item.description"
      /> -->
    </div>
  </div>
</template>

<script>
// import { defineProps } from 'vue'
import UpcomingItem from './UpcomingItem.vue';
  
export default {
  components: {
      UpcomingItem,

  },

  setup() {
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

    return {
      items,
    }
  }
}
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;


.upcoming-container {
  width: 63rem;
  height: 40.625rem;
  margin: 2.66rem 0 0 5rem;
}

.upcoming-card {
  padding: 2rem 3.12rem;
  display: flex;
  flex-direction: column;
}

.upcoming-header {
    display: flex;
    flex-direction: row;
    align-items: center;

  .upcoming-icon {
    display: flex;
    align-items: center;
  }

  .title-container {
    display: flex;
    align-items: center;

    .upcoming-title {
      margin-left: 1rem;
      font-size: 16px;
      font-weight: 700;
      padding: 0;
    }
  }
}

.upcoming-items {
  margin-top: 1rem;
}
</style>