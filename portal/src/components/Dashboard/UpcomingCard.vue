<template>
  <div class="upcoming-card">
    <div v-if="isHeader" class="upcoming-header">
      <div class="upcoming-icon">
        <v-icon 
          :icon="icon" 
          class="al-icon" 
          size="22">
        </v-icon>
      </div>
      <div>
        <div class="title-container">
          <h3 class="upcoming-title">{{ title }}</h3>
        </div>
      </div>
    </div>

    <div v-else class="upcoming-items">
      <div class="upcoming-date">{{ date }}</div>
      <div class="upcoming-item">
        <div class="title-tag-container">
          <span class="item-title">{{ itemTitle }}</span>
          <span class="item-tag">{{ itemTag }}</span>
        </div>
        <div class="item-subtitle">
          <span class="subtitle-text">{{ itemSubtitle }}</span>
        </div>
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
  date: {
    type: String,
    default: ''
  },
  itemTitle: {
    type: String,
    default: ''
  },
  itemTag: {
    type: String,
    default: ''
  },
  itemSubtitle: {
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
  display: flex;
  flex-direction: row;
  gap: 1.25rem;

  .upcoming-date {
    display: flex;
    align-items: center;
    color: $al-primary-blue;
    font-size: 16px;
    font-weight: 500;
  }

  .title-tag-container {
    gap: 0.6rem;
    display: flex;

    .item-title {
      font-size: 16px;
      font-weight: 500;
    }

    .item-tag {
      font-size: 10px;
    }
  }
}
</style>