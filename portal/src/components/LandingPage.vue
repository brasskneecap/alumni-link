<template>
  <div class="main">
    <SideNav></SideNav>
    <RouterView />
  </div>
  <div class="profile-container">
    <v-icon
      class="icon"
      icon="mdi-bell-outline"
      size="40"
      @click="notifClick"
    ></v-icon>

    <v-avatar 
      size="60"
      @click="profileClick"
    >
      <v-img
        class="profile-avatar"
        alt="profile icon"
        src="https://cdn.vuetifyjs.com/images/john.jpg"
      ></v-img>
    </v-avatar>
  </div>
</template>

<script setup>
  import SideNav from './navigation/SideNav.vue'

  function notifClick() {
    console.log('Notification icon clicked');
  }

    function profileClick() {
    console.log('Profile icon clicked');
  }
</script>


<style lang="scss">
.main {
  display: flex;
  width: 100vw;
  background-color: #F9FAFB;
}

.profile-container {
  position: absolute;
  right: 0;
  margin-top: 1.88rem;
  margin-right: 4.5rem;

  .icon {
    margin-right: 1.88rem;
  }

  .profile-avatar {
    cursor: pointer;
  }
  
}
</style>