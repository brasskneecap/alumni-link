<template>
  <div class="signin-left">
    <div class="al-logo-signin">
      <v-img
        src="/al-logo.svg"
        alt="Logo"
        height="60"
        width="142"
        class="ml-4"
      />
    </div>
    <div class="left-signin-content">
      <div class="graphic-container">
        <v-img
          src="/thinking-graduate.png"
          alt="Logo"
          height="457"
          width="457"
        />
      </div>
      <div class="graphic-text">
        <h1>Beyond Graduation</h1>
        <p>Connect with alumni to prepare for life after graduation.</p>
      </div>
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
  height: 100vh;
  display: flex;
  flex-direction: column;
  position: relative;
}

.al-logo-signin {
  margin-top: 32px;
  margin-left: 32px;

  img {
    display: block;
    width: 100%;
    height: auto;
  }
}

.left-signin-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  flex: 1;
  position: relative;
}

.graphic-container {
  display: flex;
  flex-direction: column;
  align-items: center;

  img {
    display: block;
    width: 100%;
    height: auto;
  }
}

.graphic-text {
  text-align: center;

  h1 {
    font-size: 2.25rem;
    margin-bottom: 1.38rem;
  }

  p {
    font-size: 1.25rem;
  }
}
</style>
