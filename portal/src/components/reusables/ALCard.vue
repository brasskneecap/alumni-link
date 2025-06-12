<template>
  <div 
    class="al-card-container al-card"
    :style="`width: ${width}; height: ${height}`">
    <div class="al-header">
      <div class="al-header-icon"> 
        <v-icon 
        :icon="icon"
        class="al-icon"
        size="22">
        </v-icon>
      </div>
      <div class="title-container">
        <h3 class="title">{{ title }}</h3>
      </div>
    </div>

    <slot name="content"></slot>
  </div>
</template>

<script setup>
import { defineProps } from 'vue'

defineProps({
  title: {
    type: String,
    required: true
  },
  icon: {
    type: String,
    required: true
  },
  width: {
    type: String,
    default: '15.25rem'
  },
  height: {
    type: String,
    default: '12.5rem'
  },
})
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.al-card-container {
  padding: 0.5rem 0.75rem;
  display: flex;
  flex-direction: column;
}

.al-header {
    display: flex;
    flex-direction: row;
    align-items: center;
    height: 3rem;

  .al-header-icon {
    display: flex;
    align-items: center;
  }

  .title-container {
    display: flex;
    align-items: center;
    
    .title {
      margin-left: 1rem;
      font-size: 16px;
      font-weight: 700;
      padding: 0;
    }
  }
}

.bottom {
  margin-top: auto;
  margin-bottom: 1rem;

  .content {
    font-size: 32px;
    font-weight: 700;
    margin-top: 2.44rem;
  }

  .subtitle {
    font-size: 16px;
    font-weight: 400;
    margin-top: 0.38rem;
    color: $al-secondary-gray;
  }
}

</style>