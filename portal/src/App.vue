<template>
  <v-app class="app-main">
    <LoginPage />
  </v-app>
</template>


<script setup>
  // import LandingPage from './components/LandingPage.vue'
  import LoginPage from './components/LoginPage/LoginPage.vue'
</script>


<style lang="scss">
@use '@/variables.scss' as *;

.app-main {
  width: 100vw;
  height: 100vh;
}
</style>
