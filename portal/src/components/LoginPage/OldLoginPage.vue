<template>
  <div class="login-page-container">
    <GraphicFiller />
    <div>
      <SigninPage />
      <button @click="handleLogin(credentials)"></button>
    </div>
  </div>
</template>


<script>
// import { ref } from 'vue'
import GraphicFiller from './GraphicFiller.vue';
import SigninPage from './SigninPage.vue';
import { useStore } from 'vuex'

export default {
  components: { GraphicFiller, SigninPage },
  setup () {
    const store = useStore()
    const credentials = {email: "test@example.com", password: "password123"}
    
    return {
      // access a mutation
      handleLogin: (credentials) => store.dispatch('user/login', credentials),
      credentials
    }
  }
}

</script>


<style lang="scss">
@use '@/variables.scss' as *;

.login-page-container {
  display: grid;
  grid-template-columns: 1fr 1fr;
}
</style>
