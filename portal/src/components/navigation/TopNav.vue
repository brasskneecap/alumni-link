<template>
  <v-app-bar 
    app 
    class="app-bar-custom" 
    height="87"
    dark>
    <v-img
        src="/al-logo.svg"
        alt="Logo"
        max-height="38"
        max-width="91"
        class="ml-4 nav-logo"
      />
    <v-toolbar-title></v-toolbar-title>
  </v-app-bar>
</template>

<script setup>
</script>

<style lang="scss">
@use '@/variables.scss' as *;

.app-bar-custom {
  background-color: $al-primary-color !important;
  box-shadow: none !important;
}

.nav-logo {
  border-right: 1px solid black;
}
</style>