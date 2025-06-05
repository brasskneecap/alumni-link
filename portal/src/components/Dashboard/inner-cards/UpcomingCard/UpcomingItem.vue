<template>
    <div class="upcoming-item">
      <div class="upcoming-date">{{ date }}</div>
      <div>
        <div class="container">
          <span class="title">{{ title }}</span>
          <span class="tag">{{ tag }}</span>
        </div>
        <div>
          <span class="description">{{ description }}</span>
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
  margin-bottom: 3.19rem;

  .upcoming-date {
    display: flex;
    align-items: center;
    color: $al-primary-blue;
    font-size: 16px;
    font-weight: 500;
  }

  .container {
    gap: 1rem;
    display: flex;

    .title {
      font-size: 16px;
      font-weight: 500;
    }

    .tag {
      display: flex;
      align-items: center;
      padding: 0 0.7rem;
      border-radius: 15px;
      font-size: 10px;
      background-color: $al-tag-assignment;
    }
  }

  .description {
    font-size: 14px;
    font-weight: 400;
    color: $al-secondary-gray;
  }
}
</style>