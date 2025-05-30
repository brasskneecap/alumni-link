<template>
  <div class="assignments-tracker">
    <div class="tracker-topbar">
      <div 
      class="tracker-icon"> 
        <v-icon 
        :icon="icon"
        class="icon"
        size="22">
        </v-icon>
      </div>
      <h3 
      class="tracker-title"> 
      {{ title }}
    </h3>
    </div>
    <div class="tracker-bottom">
      <div 
      class="tracker-content">
      {{ content }}
      </div>

      <div class="tracker-subtitle">
        <p>Assignments Completed</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineProps } from 'vue'

defineProps({
  icon: {
    type: String,
    default: () => ""
  },
  title: {
    type: String,
    required: true
  },
  content: {
    type: String,
    required: true
  }
})
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.assignments-tracker {
  padding: 0.625rem 0.875rem;
  display: flex;
  flex-direction: column;
}

.tracker-topbar {
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    align-items: center;
    height: 3rem;

  .tracker-icon {
    display: flex;
    
    .icon {
      background-color: $al-primary-color;
      padding: 1.15rem;
      border-radius: 0.625rem;
      color: $al-primary-blue;
    }
  }

  .tracker-title {
    margin-left: 1rem;
    font-size: 16px;
    font-weight: 700;
    padding-top: .5rem;
  }
}

.tracker-bottom {
  margin-top: auto;
  margin-bottom: 1rem;

  .tracker-content {
    font-size: 32px;
    font-weight: 700;
    margin-top: 2.44rem;
  }

  .tracker-subtitle {
    font-size: 16px;
    font-weight: 400;
    margin-top: 0.38rem;
    color: $al-secondary-gray;
  }
}

</style>