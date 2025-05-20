<template>
  <v-app class="app-main">
    <LoginPage v-if="!loggedIn"/>
    <LandingPage></LandingPage>
  </v-app>
</template>


<script>
  import LandingPage from './components/LandingPage.vue'
  import LoginPage from './components/LoginPage/LoginPage.vue'
  import { computed } from 'vue'
  import { useStore } from 'vuex'
  
  export default {
    components: { LoginPage, LandingPage },
    setup () {
      const store = useStore()
      const loggedIn = computed(() => store.getters["user/loggedIn"]);
      console.log("Logged In:", loggedIn.value);
      return {
        // access a getter in computed function
        loggedIn,
      }
    }
}
</script>


<style lang="scss">
@use '@/variables.scss' as *;

.app-main {
  width: 100vw;
  height: 100vh;
}
</style>
