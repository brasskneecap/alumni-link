<template>
  <v-list lines="two">  
    <v-list-header>WEEK 1</v-list-header>
    <v-list-item
      v-for="(item, i) in assignments"
      :key="i"
      :value="item"
      :title="item.name"
      :subtitle="item.description"
      :date="item.dueDate"
      color="primary"
    >
      <template v-slot:prepend>
        <v-icon icon="mdi-file-document-multiple-outline"></v-icon>
      </template>
      <template v-slot:append>
        <span>{{ item.dueDate }}</span>
      </template>
    </v-list-item>
  </v-list>
</template>

<!-- :date="formatDate(item.dueDate)" for later -->

<script setup>
import { computed } from 'vue';
import { useStore } from 'vuex';
// import { formatDate } from '@/utils/formatters' //will work once the formatDate file is merged

const store = useStore()

const assignments = computed(() => store.getters["assignments/assignments"])

</script>
