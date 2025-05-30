<template>
  <div class="team-card">
    <!-- header view -->
    <div v-if="isHeader" class="team-header">
      <div class="team-icon">
        <v-icon 
          :icon="icon" 
          class="icon" 
          size="22">
        </v-icon>
      </div>
      <h3>{{ title }}</h3>
    </div>

    <!-- member view -->
    <div v-else class="team-member">
    <div class="member-title">
      <h4>
        <span class="member-name">{{ name }}</span>
        <span class="bullet">&bull;</span>
        <span class="member-role">{{ role }}</span>
      </h4>
    </div>
      <p class="description">{{ description }}</p>
    </div>
  </div>
</template>

<script setup>
import { defineProps } from 'vue'

defineProps({
  icon: {
    type: String,
    default: ''
  },
  title: {
    type: String,
    default: ''
  },
  name: {
    type: String,
    default: ''
  },
  role: {
    type: String,
    default: ''
  },
  description: {
    type: String,
    default: ''
  },
  isHeader: {
    type: Boolean,
    default: false
  }
})
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.team-card {
  padding: 0.625rem 0.875rem;
  display: flex;
  flex-direction: column;
}

.team-header {
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    align-items: center;
    height: 3rem;

  .team-icon {
    display: flex;
    
    .icon {
      background-color: $al-primary-color;
      padding: 1.15rem;
      border-radius: 0.625rem;
      color: $al-primary-blue;
    }
  }

  h3 {
    margin-left: 1rem;
    font-size: 16px;
    font-weight: 700;
    padding-top: .5rem;
  }
}

.member-title {
  h4 {
    display: flex;
    gap: 0.5rem;
    align-items: center;
    
    .member-role {
      color: $al-primary-blue;
    }
  }
}

</style>