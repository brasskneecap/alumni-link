<template>
  <div class="login-page-container">
    <GraphicFiller />
    <SigninPage />
  </div>
</template>


<script setup>
// import { ref } from 'vue'
import GraphicFiller from './GraphicFiller.vue';
import SigninPage from './SigninPage.vue'; 

const props = defineProps({

})

</script>

<style lang="scss">
@use '@/variables.scss' as *;

.login-page-container {
  display: grid;
  grid-template-columns: 1fr 1fr;
}
</style>
