<template>
  <div class="side-nav">
      <div class="nav-logo-container">
        <v-img
        src="/al-logo.svg"
        alt="Logo"
        max-height="45"
        max-width="108"
        class="nav-logo"
      />
  </div>
    <div class="nav-container">
      <div v-for="item in items" :key="item.id">
        <RouterLink
          :to="item.page"
          class="nav-item"
          active-class="nav-item-active"
          >
            <v-icon
              :icon="item.icon"
              size="40"
              class="nav-icon"
              end
            ></v-icon>
            <span>{{item.name}}</span>
        </RouterLink>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed } from 'vue'
import { RouterLink } from 'vue-router'
import { useStore } from 'vuex'

export default {
  setup() {
    const store = useStore()
    const user = computed(() => store.getters["user/user"]);

    let items = ref([
      { id: 1, name: 'Dashboard', icon: 'mdi-view-dashboard-outline', page: "/", role: "student"},
      { id: 2, name: 'Assignments', icon: 'mdi-file-document-multiple-outline', page: "/assignments", role: "student"},
      { id: 3, name: 'Assignments', icon: 'mdi-file-document-multiple-outline', page: "/assignments-overview", role: "mentor"},
      { id: 4, name: 'Calendar', icon: 'mdi-calendar-month-outline', page: "/calendar", role: "all"},
      { id: 5, name: 'Messages', icon: 'mdi-message-text-outline', page: "/messages", role: "all" },
      { id: 6, name: 'Blast', icon: 'mdi-bullhorn-outline', page: "/blast", role: "student" },
      { id: 7, name: 'Groups', icon: 'mdi-account-group-outline', page: "/groups", role: "student" },
    ])

    console.log(user.value.role)
    items = items.value.filter((item) => item.role === "all" || item.role.toLowerCase() === user.value.role.toLowerCase())
    console.log(items)
    return {
      items
    }
  },
}
</script>

<style lang="scss">
@use '@/variables.scss' as *;

.nav-logo-container {
  display: flex;
  flex-direction: column;
  height: 87px;

  .nav-logo {
    margin: 1.25rem 0 1.25rem 0.25rem;
  }
}


.nav-container {
  height: calc(100% - 87px);
  padding-top: 1.5rem;
  width: 7.25rem !important; 
  min-height: 69.8125rem;
  background-color: $al-primary-neutral;
  box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.12);
  box-shadow: 0px 0px 10px 0px color(display-p3 0 0 0 / 0.12);
}

.nav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 1rem 0;
  font-size: 14px;
  font-weight: 500;
  color: black;
  .nav-icon {
    margin: 0;
    padding-bottom: 0.6rem;
  }

  &:hover {
    cursor: pointer;
    background-color: $al-primary-color;
    color: black;
  }
}

.nav-item-active {
  background-color: $al-primary-color;
}
</style>