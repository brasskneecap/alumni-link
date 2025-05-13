<template>
  <v-container>
    <v-text-field
      :type="isPasswordVisible ? 'text' : 'password'"
      :label="label"
      :placeholder="placeholder"
      :persistent-placeholder="true"
      :prepend-inner-icon="icon"
      variant="underlined"
      :append-inner-icon="isPasswordVisible ? 'mdi-eye-off' : 'mdi-eye'"
      @click:append-inner="togglePasswordVisibility"
    />
  </v-container>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
  label: {
    type: String,
    required: true,
  },
  placeholder: {
    type: String,
    required: true,
  },
  icon: {
    type: String,
    default: () => ""
  }
})

// Password visibility state
const isPasswordVisible = ref(false)

function togglePasswordVisibility() {
  isPasswordVisible.value = !isPasswordVisible.value
}
</script>