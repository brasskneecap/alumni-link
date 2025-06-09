<template>
  <div class="team-member">
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
</template>

<script setup>
import { defineProps } from 'vue'

defineProps({
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
})
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.team-member {
  display: flex;
  flex-direction: row;
  gap: 1.25rem;
  margin-top: 1rem;

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