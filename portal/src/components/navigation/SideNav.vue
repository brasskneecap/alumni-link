<template>
  <div class="side-nav">
    <v-img
        src="/uvu-logo.svg"
        alt="School Logo"
        class="school-logo"
      />
    <div class="spacer"></div>
    <div v-for="item in items" :key="item.id"
      class="nav-item">
      <v-icon
        :icon="item.icon"
        size="25"
        end
      ></v-icon>
      <span>{{item.name}}</span>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
export default {
  setup() {
    const items = ref([
      { id: 1, name: 'Dashboard', icon: 'mdi-view-dashboard-outline' },
      { id: 2, name: 'Activity', icon: 'mdi-bell-outline' },
      { id: 3, name: 'Calendar', icon: 'mdi-calendar-month-outline' },
      { id: 4, name: 'Chat', icon: 'mdi-chat-outline' },
      { id: 5, name: 'Blast', icon: 'mdi-bullhorn-outline' },
      { id: 6, name: 'Groups', icon: 'mdi-account-multiple-outline' },
      { id: 7, name: 'Assignments', icon: 'mdi-notebook-outline' },
    ])

    return {
      items
    }
  },
}
</script>

<style lang="scss">
@use '@/variables.scss' as *;

.side-nav {
  height: calc(100% - 87px);
  margin-top: 87px;
  width: 139px !important;
  background-color: $al-primary-neutral;
}

.school-logo {
  margin: 0 auto;
  height: 100px;
  width: 100px;
}

.spacer {
  border-bottom: 1px solid $al-gray;
  margin: 0px 12px;
}

.nav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 10px;
  width: 139px;

  &:hover {
    color: #275D38;
    padding-left: 7px;
    border-left: 3px solid #275D38;
    cursor: pointer;
  }
}
</style>