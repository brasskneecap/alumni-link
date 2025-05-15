<template>
  <div class="main">
    <SideNav></SideNav>
  </div>
</template>

<script setup>
  import TopNav from './navigation/TopNav.vue'
  import SideNav from './navigation/SideNav.vue'
</script>


<style lang="scss">
.main {
  display: flex;
  width: 100vw;
  height: 100vh;
  background-color: #F9FAFB;
}


</style>