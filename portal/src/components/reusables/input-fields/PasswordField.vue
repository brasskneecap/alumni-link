<template>
  <v-container>
    <v-text-field
      v-model="internalValue"
      :type="isPasswordVisible ? 'text' : 'password'"
      :label="label"
      :placeholder="placeholder"
      :persistent-placeholder="true"
      :prepend-inner-icon="icon"
      variant="underlined"
      :append-inner-icon="isPasswordVisible ? 'mdi-eye-off' : 'mdi-eye'"
      @click:append-inner="togglePasswordVisibility"
      @input="updateValue"
    />
  </v-container>
</template>

<script setup>
import { ref, watch, defineProps, defineEmits } from 'vue'

const emit = defineEmits()

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
    },
    modelValue: {
      type: String,
      default: ''
    }
  })

// Password visibility state
const isPasswordVisible = ref(false)

function togglePasswordVisibility() {
  isPasswordVisible.value = !isPasswordVisible.value
}


// Internal state
const internalValue = ref(props.modelValue)

// Sync the internal value with the parent modelValue
watch(() => props.modelValue, (newValue) => {
  internalValue.value = newValue
})

// Emit the updated value when the internal value changes
function updateValue() {
  emit('update:modelValue', internalValue.value)
}
</script>