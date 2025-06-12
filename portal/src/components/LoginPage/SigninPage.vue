<template>
  <div class="signin-container">
    <h3>Sign in</h3>
    <TextField
      v-model="email" 
      class="email-field"
      label="Email"
      placeholder="Enter your email address"
      icon="mdi-email-outline"
    />
    <PasswordField
      v-model="password"
      class="password-field"
      label="Password"
      placeholder="Enter your password"
      icon="mdi-lock-outline"
    />
    <div class="signin-functions">
      <v-checkbox        
        label="Remember Me"
        density="compact"
        :ripple="false"
        class="signin-checkbox">
      </v-checkbox>
      <a 
        href="#" 
        class="pw-forgot">Forgot Password?</a>
    </div>
    <v-btn 
      rounded="xl"
      class=" login-btn text-subtitle-1 font-weight-medium"
      @click="handleLogin()">Login</v-btn>
  </div>
</template>

<script setup>
import TextField from '../reusables/input-fields/TextField.vue';
import PasswordField from '../reusables/input-fields/PasswordField.vue';
import { ref } from 'vue'
import { useStore } from 'vuex'

const store = useStore()

const email = ref('')
const password = ref('')

function handleLogin() {
  const credentials = {
    email: email.value,
    password: password.value
  }
  store.dispatch('user/login', credentials)
}
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
      font-weight: 300;
    }

    a {
      padding: 0;
      margin: 0;
      color: black;
    }
  }

  .pw-forgot {
    font-size: 12px;
    font-weight: 300;
  }

  .login-btn {
    width: 26.8rem;
    background-color: black;
    color: white;
    font-size: 17px !important;
    height: 53px !important;
    margin-top: 3.75rem;
  }

}
</style>