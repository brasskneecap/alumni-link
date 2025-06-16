<template>
  <div class="blast-item">
    <div class="blast-date">{{ date }}</div>
    <div class="container">
        <span class="title">{{ title }}</span>
        <span class="description">{{ description }}</span>
    </div>

  </div>
</template>

<script setup>
  defineProps({
    date: {
      type: String,
      default: ''
    },
    title: {
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

.blast-item {
    display: flex;
    gap: 1.25rem;
    margin-bottom: 3.19rem;
    border-bottom: 1px solid $al-secondary-gray;
  
  .blast-date {
    display: flex;
    align-items: center;
    color: $al-primary-blue;
    font-size: 16px;
    font-weight: 500;
    white-space: pre;
  }

  .container {
    display: flex;
    flex-direction: column;
    overflow: hidden;
    white-space: nowrap;


    .title {
      font-size: 16px;
      font-weight: 500;
      text-overflow: ellipsis;
      overflow: hidden;
    }

    .description {
      font-size: 14px;
      font-weight: 400;
      color: $al-secondary-gray;
      text-overflow: ellipsis;
      overflow: hidden;
    }
  }
}

</style>