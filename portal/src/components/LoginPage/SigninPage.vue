<template>
  <div class="signin-container">
    <h3>Sign in</h3>
    <TextField 
      class="input-field"
      label="Email"
      placeholder="Enter your email address"
      icon="mdi-email-outline"
    />
    <TextField 
      class="input-field"
      label="Password"
      placeholder="Enter your password"
      icon="mdi-lock-outline"
    />
    <div class="signin-functions">
      <v-checkbox label="Remember Me"></v-checkbox>
      <a href="#">Forgot Password?</a>
    </div>
    <v-btn rounded="xl" size="x-large" class="text-subtitle-1 font-weight-medium">Login</v-btn>
  </div>
</template>

<script setup>
import TextField from '../input-fields/TextField.vue';
</script>   

<style lang="scss">
@use '@/variables.scss' as *;

.signin-container {

  // margin: 8rem 13rem 14rem 8rem;

  .input-field {
    padding: 0;
  }

  .signin-functions {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
}
</style>