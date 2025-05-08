<template>
  <div class="signin-left">
    <v-img
      src="/al-logo.svg"
      alt="Logo"
      height="60"
      width="142"
      class="ml-4 al-logo-signin"
    />
    <div class="filler-container">
      <v-img
        src="/thinking-graduate.png"
        alt="Logo"
        height="457"
        width="457"
        class="ml-4"
      />
      <div>3</div>
      <div>4</div>
    </div>
  </div>
</template>


<script setup>
// import { ref } from 'vue'
const props = defineProps({
})
</script>


<style lang="scss">
@use '@/variables.scss' as *;

.signin-left {
  background-color: #fafafa;
}

.filler-container {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.al-logo-signin {
  margin-top: 32px;
  margin-left: 32px;
}
</style>
