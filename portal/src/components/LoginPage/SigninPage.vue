<template>
  <div>
    <h1>Sign in</h1>
    <TextField 
      label="Email"
      icon="mdi-email-outline"
    />
    <TextField 
      label="Password"
      icon="mdi-lock-outline"
    />
  </div>
</template>

<script setup>
import TextField from '../input-fields/TextField.vue';
</script>   