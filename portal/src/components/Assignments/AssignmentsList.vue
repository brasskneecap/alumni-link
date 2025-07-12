<template>
  <v-list lines="two" class="container">
    <v-list-subheader class="header">WEEK 1</v-list-subheader>
    <template v-for="(item, i) in assignments" :key="i">
      <v-list-item
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

        <v-list-item-title class="assignment-title">{{ item.name }}</v-list-item-title>
        <v-list-item-subtitle class="assignment-subtitle">{{ item.description }}</v-list-item-subtitle>

        <template #append>
          <span>Due {{ formatDateMDY(item.dueDate) }}</span>
        </template>
      </v-list-item>
    </template>
  </v-list>
</template>

<script setup>
import { computed } from 'vue'
import { useStore } from 'vuex'
const emit = defineEmits(['show-assignment-view'])

const store = useStore()
const assignments = computed(() => store.getters["assignments/assignments"])

function toggleAssignment(item) {
  emit('show-assignment-view', item);
}

function formatDateMDY(dateStr) {
  const date = new Date(dateStr);
  const month = String(date.getMonth() + 1);
  const day = String(date.getDate());
  const year = date.getFullYear().toString().slice(-2);
  return `${month}/${day}/${year}`;
}

</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.container {
  max-width: 30.1875rem;
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

:deep(.assignment-subtitle) {
  color: $al-secondary-gray;
  font-size: 14px;
  font-weight: 400;
  margin-right: 1rem;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  display: block;
  max-width: 100%;
  opacity: 1;
  margin-bottom: 1rem;
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