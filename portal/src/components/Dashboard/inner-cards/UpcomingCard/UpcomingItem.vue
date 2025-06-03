<template>
    <div class="upcoming-item">
      <div class="upcoming-date">{{ date }}</div>
      <div>
        <div class="title-tag-container">
          <span class="item-title">{{ title }}</span>
          <span class="item-tag">{{ tag }}</span>
        </div>
        <div class="item-subtitle">
          <span class="subtitle-text">{{ description }}</span>
        </div>
      </div>
    </div>
</template>

<script setup>
  defineProps({
    title: {
      type: String,
      default: ''
    },
    date: {
      type: String,
      default: ''
    },
    tag: {
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

.upcoming-item {
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