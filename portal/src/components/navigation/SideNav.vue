<template>
  <div class="side-nav">
    <div class="nav-container">
      <div v-for="item in items" :key="item.id"
        class="nav-item">
        <v-icon
          :icon="item.icon"
          size="40"
          class="nav-icon"
          end
        ></v-icon>
        <span>{{item.name}}</span>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
export default {
  setup() {
    const items = ref([
      { id: 1, name: 'Dashboard', icon: 'mdi-view-dashboard-outline' },
      { id: 2, name: 'Assignments', icon: 'mdi-file-document-multiple-outline' },
      { id: 3, name: 'Calendar', icon: 'mdi-calendar-month-outline' },
      { id: 4, name: 'Messages', icon: 'mdi-message-text-outline' },
      { id: 5, name: 'Blast', icon: 'mdi-bullhorn-outline' },
      { id: 6, name: 'Groups', icon: 'mdi-account-group-outline' },
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
  width: 7.25rem !important;
  background-color: $al-primary-neutral;
  box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.12);
  box-shadow: 0px 0px 10px 0px color(display-p3 0 0 0 / 0.12);
}

.nav-container {
  margin-top: 1.25rem;
}

.nav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 1.25rem 0;
  font-size: 14px;
  font-weight: 300;

  .nav-icon {
    margin: 0;
    padding-bottom: 0.6rem;
  }

  &:hover {
    cursor: pointer;
    background-color: $al-primary-color;
  }
}
</style>