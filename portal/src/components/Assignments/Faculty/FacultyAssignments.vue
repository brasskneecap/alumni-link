<template>
  <ALCard
    class="assignments-container"
    title="ASSIGNMENTS OVERVIEW"
    icon="mdi-calendar-month-outline"
    height="59.875rem"
    width="96.375rem"
  >

  <template #header-right>

  </template>

    <template #content>
      <div class="assignments-flex">
        <v-list lines="two" class="container">
          <v-list-subheader class="header">
            <div class="header-container">
              <span>ASSIGNMENTS</span>
              <div class="d-flex gap-2">
                <router-link to="/assignments-overview/create">
                  <v-icon 
                    icon="mdi-plus" 
                    size="24"
                    color="#000000"
                  />
                </router-link>
              </div>
            </div>
          </v-list-subheader>
          <template v-for="(item, i) in assignments" :key="i">
            <v-list-item v-ripple="false"
              @click="toggleAssignment(item)"
              :value="item"
              :date="item.dueDate"
              color="primary"
              rounded="lg"
              class="border-b-sm"
            >
              <template #prepend>
                <v-icon class="al-icon icon" icon="mdi-file-document-multiple-outline" size="20" ></v-icon>
              </template>
              <div>
                <v-list-item-title class="assignment-title">{{ item.name }}</v-list-item-title>
                <span class="dueDate">Due {{ formatDateMDY(item.dueDate) }}</span>
              </div>
            </v-list-item>
          </template>
        </v-list>
        <FacultyAssignmentsView 
          v-if="isVisible" 
          :assignment="selectedAssignment" 
          :teamMember="selectedTeamMember"
        />
      </div>
    </template>
  </ALCard>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useStore } from 'vuex'
import ALCard from '../../reusables/ALCard.vue'
import { formatDateMDY } from '@/utils/formatters'
import FacultyAssignmentsView from '../Faculty/FacultyAssignmentsView.vue'

const store = useStore()
const assignments = computed(() => store.getters["assignments/assignments"])

const isVisible = ref(false)
const selectedAssignment = ref(null)
const selectedTeamMember = ref({
  profilePicture: '/path/to/image.jpg',
  name: 'John Doe',
  role: 'Developer',
  school: 'Utah Valley University'
});

function toggleAssignment(item) {
  selectedAssignment.value = item
  isVisible.value = true
}
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.assignments-container {
  margin: 7.12rem 2rem 2.81rem 2rem;
}

.assignments-flex {
  display: flex;
  flex-direction: row;
  gap: 2rem;
  width: 100%;
}

.container {
  width: 20rem;
  margin: 1.25rem 0 0 0.5rem;
}

.header {
  font-size: 16px;
  font-weight: 600;
  color: black;
}

.header-container {
  width: 18rem;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  align-content: center;
  span {
    align-content: center;
  }
}

:deep(.al-icon.icon) {
  opacity: 1;
}

:deep(.assignment-title) {
  font-size: 16px;
  font-weight: 700;
  margin-right: 1rem;
  margin-top: 0.5rem;
}

:deep(.dueDate) {
  color: $al-secondary-gray;
}

:deep(.v-list-item__append) {
  align-items: flex-start !important;
  justify-content: flex-start !important;
  display: flex !important;
  flex-direction: column !important;
  height: 100%;
  color: $al-secondary-gray;
  font-size: 16px;
  font-weight: 400;
  margin-top: 0.5rem;
  align-self: flex-start;
}
</style>