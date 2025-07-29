<template>
  <ALCard
    class="assignments-container"
    title="ASSIGNMENTS OVERVIEW"
    icon="mdi-calendar-month-outline"
    height="59.875rem"
    width="96.375rem"
  >
    <template #header-right>
      <router-link to="/assignments-overview/create">
        <v-btn color="#4065DA">Add Assignment</v-btn>
      </router-link>
    </template>

    <template #content>
      <v-list lines="two" class="container">
        <v-list-subheader class="header">ASSIGNMENTS</v-list-subheader>
        <template v-for="(item, i) in assignments" :key="i">
          <v-list-item v-ripple="false"
            @click="toggleAssignment(item)"
            :value="item"
            :date="item.dueDate"
            color="primary"
            rounded="lg"
            class="border-b-sm"
          >
            <template #prepend>
              <v-icon class="al-icon icon" icon="mdi-file-document-multiple-outline" size="20" ></v-icon>
            </template>
            <div>
              <v-list-item-title class="assignment-title">{{ item.name }}</v-list-item-title>
              <span class="dueDate">Due {{ formatDateMDY(item.dueDate) }}</span>
            </div>
          </v-list-item>
        </template>
      </v-list>
    </template>
  </ALCard>
</template>

<script setup>
import { ref } from 'vue'
import { computed } from 'vue'
import { useStore } from 'vuex'
import ALCard from '../../reusables/ALCard.vue'
import { formatDateMDY } from '@/utils/formatters'

const store = useStore()
const assignments = computed(() => store.getters["assignments/assignments"])

function toggleAssignment(item) {
  emit('show-assignment-view', item); 
}
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.assignments-container {
  margin: 7.12rem 2rem 2.81rem 2rem;
}

.assignments-flex {
  display: flex;
  flex-direction: row;
  gap: 2rem;
  width: 100%;
}

.assignments-flex > *:first-child {
  flex: 0 0 32rem;
}

.assignments-flex > *:last-child {
  flex: 1 1 0;
}

.container {
  max-width: 20rem;
  margin: 1.25rem 0 0 0.5rem;
}

.header {
  font-size: 16px;
  font-weight: 600;  
  color: black;
}

:deep(.al-icon.icon) {
  opacity: 1;
}

:deep(.assignment-title) {
  font-size: 16px;
  font-weight: 700;
  margin-right: 1rem;
  margin-top: 0.5rem;
}

:deep(.dueDate) {
  color: $al-secondary-gray;
}

:deep(.v-list-item__append) {
  align-items: flex-start !important;
  justify-content: flex-start !important;
  display: flex !important;
  flex-direction: column !important;
  height: 100%;
  color: $al-secondary-gray;
  font-size: 16px;
  font-weight: 400;
  margin-top: 0.5rem;
  align-self: flex-start;
}
</style>