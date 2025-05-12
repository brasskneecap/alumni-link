<template>
  <div class="signin-container">
    <h3>Sign in</h3>
    <TextField 
      class="email-field"
      label="Email"
      placeholder="Enter your email address"
      icon="mdi-email-outline"
    />
    <TextField 
      class="password-field"
      label="Password"
      placeholder="Enter your password"
      icon="mdi-lock-outline"
    />
    <div class="signin-functions">
      <v-checkbox label="Remember Me" density="compact" class="signin-checkbox"></v-checkbox>
      <a href="#">Forgot Password?</a>
    </div>
    <v-btn rounded="xl" size="x-large" class=" login-btn text-subtitle-1 font-weight-medium">Login</v-btn>
  </div>
</template>

<script setup>
import TextField from '../input-fields/TextField.vue';
</script>   

<style lang="scss">
@use '@/variables.scss' as *;

.signin-container {
  width: 26.8rem;
  margin: 8.25rem auto;

  h3 {
    padding-bottom: 8rem;
    font-size: 1.875rem;
  }

  .email-field, .password-field {
    padding: 0;
  }

  .password-field {
    margin-top: 3rem;
  }

  .signin-functions {
    display: flex;
    justify-content: space-between;
    align-items: center;

    .signin-checkbox {
      display: flex;
      margin: 0;
      padding: 0;
    }

    a {
      padding: 0;
      margin: 0;
    }
  }

  .login-btn {
    width: 26.8rem;
    background-color: black;
    color: white;
    margin-top: 3.75rem;
  }
}
</style>