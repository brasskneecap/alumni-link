<template>
  <div>
    <div v-for="assignment in assignments">{{ assignment.id }}</div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

defineProps({
    studentId: {
      type: String,
      required: true
    },
    assignments: {
      type: Array,
      required: true
    },
  })

</script>