<template>
  <div>
    <div>
        <h1 class="title">{{ assignment.name }}</h1>
        <p>Due {{ formatDateMDY(assignment.dueDate) }}</p>
        <p>{{ assignment.description }}</p>
    </div>
    <div>
      <h2>Submit Assignment</h2>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  assignment: Object
})

function formatDateMDY(dateStr) {
  const date = new Date(dateStr)
  const month = String(date.getMonth() + 1)
  const day = String(date.getDate())
  const year = date.getFullYear().toString().slice(-2)
  return `${month}/${day}/${year}`
}
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.title {
  font-size: 24px;
  font-weight: 700;
}

</style>