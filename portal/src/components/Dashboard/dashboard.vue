<template>
  <div class="dashboard-container">

    <!-- welcome message, settings, notification -->
    <div class="topbar-container">
      <div class="welcome-container">
        <h2>Welcome Back, name</h2>
        <p>Take a look at your progress</p>
      </div>
      <div class="profile-container">

      </div>
    </div>

    <!-- assignments and team list -->
    <div class="tracker-container">
      <div class="this-week-container">

      </div>
      <div class="completed-container">

      </div>
      <div class="team-container">

      </div>
    </div>

    <!-- upcoming things list -->
    <div class="upcoming-container">

    </div>

    <!-- calendar and blast -->
    <div class="right-container">
      <div class="calendar-container">

      </div>
      <div class="blast-container">

      </div>
    </div>
  </div>
</template>

<script setup>
// import { ref } from 'vue'

</script>

<style lang="scss">
@use '@/variables.scss' as *;

.welcome-container {

  margin: 1.88rem 0 0 5rem;
  
  h2 {
    font-size: 48px;
    font-weight: 300;
  }

  p {
    font-size: 20px;
    font-weight: 300;
  }
}

</style>