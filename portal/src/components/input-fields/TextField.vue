<template>
  <v-container>
    <v-text-field
      label="Email"
      placeholder="Enter your email address"
      :persistent-placeholder="true"
      :prepend-inner-icon="icon"
      variant="underlined"
    ></v-text-field>
  </v-container>
</template>

<script setup>
const props = defineProps({
  label: {
    type: String,
    required: true,
  },
  icon: {
    type: String,
    default: () => ""
  }
})
</script>