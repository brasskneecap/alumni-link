<template>
  <v-container>
    <v-text-field
      v-model="internalValue"
      :label="label"
      :placeholder="placeholder"
      :persistent-placeholder="true"
      :prepend-inner-icon="icon"
      variant="underlined"
      @input="updateValue"
    ></v-text-field>
  </v-container>
</template>

<script setup>
import { ref, watch, defineProps, defineEmits } from 'vue'

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

  // Emit function
  const emit = defineEmits()

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