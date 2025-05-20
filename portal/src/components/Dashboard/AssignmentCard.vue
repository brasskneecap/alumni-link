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
      <div 
      class="tracker-title"> 
      {{ title }}
      </div>
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
  padding: 1vh 1vw;
  width: 15vw;
  height: 19vh;
  display: flex;
  flex-direction: column;
}

.tracker-topbar {
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    align-items: center;

  .tracker-icon {
    display: flex;
    width: clamp(2rem, 2vw, 2.5rem);
    height: clamp(2rem, 2vw, 2.5rem);
    
    .icon {
      background-color: $al-primary-color;
      padding: clamp(1.2rem, 1vw, 1.25rem);
      border-radius: 0.625rem;
      color: $al-primary-blue;
    }
  }

  .tracker-title {
    margin-left: clamp(0.5rem, 1vw, 1rem);
    font-size: clamp(14px, 1vw, 20px);
    font-weight: 700;
  }
}

.tracker-bottom {
  margin-top: auto;
  margin-bottom: 1rem;

  .tracker-content {
    font-size: clamp(24px, 2vw, 48px);
    font-weight: 700;
    margin-top: clamp(1.5rem, 2.5vh, 2.44rem);
  }

  .tracker-subtitle {
    font-size: clamp(14px, 1vw, 20px);
    font-weight: 400;
    margin-top: clamp(0.25rem, 0.5vh, 0.38rem);
    color: $al-secondary-gray;
  }
}

</style>