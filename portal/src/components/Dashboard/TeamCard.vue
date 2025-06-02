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
      <div>
        <div class="title-container">
          <h3 class="team-title">{{ title }}</h3>
        </div>
      </div>
    </div>

    <!-- member view -->
    <div v-else class="team-member">
      <div class="member-avatar">
        <v-avatar class="avatar">
          <v-img
            alt="member-avatar"
            :src="src"
            size="42"
          ></v-img>
        </v-avatar>
      </div>
      <div class="title-container">
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
  src: {
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
    align-items: center;
    height: 3rem;

  .team-icon {
    display: flex;
    align-items: center;
    
    .icon {
      background-color: $al-primary-color;
      padding: 1.15rem;
      border-radius: 0.625rem;
      color: $al-primary-blue;
    }
  }

  .title-container {
    display: flex;
    align-items: center;

    .team-title {
      margin-left: 1rem;
      font-size: 16px;
      font-weight: 700;
      padding: 0;
    }
  }
}

.team-member {
  display: flex;
  flex-direction: row;
  gap: 1.25rem;

  .member-avatar {

    .avatar {
      background-color: $al-primary-color;
      color: $al-primary-blue;
    }
  }

  .title-container {

    .member-title {

      h4 {
        display: flex;
        flex-direction: row;
        gap: 0.5rem;
        align-items: center;

        .member-name, .bullet, .member-role {
          font-size: 16px;
          font-weight: 700;
        }
        
        .member-role {
          color: $al-primary-blue;
        }
      }
    }

    .description {
      font-size: 12px;
      font-weight: 400;
      color: $al-secondary-gray;
    }
  }
}

</style>