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
        <!-- left column: stacked lists -->
        <div class="left-column">
          <!-- STUDENTS group -->
          <v-list lines="two" class="container">
            <v-list-group value="Students">
              <template v-slot:activator="{ props, isOpen }">
                <v-list-item
                  v-bind="props"
                  class="list-header"
                  v-ripple="false"
                  rounded="lg"
                >
                  <template #prepend>
                    <span class="header-title">STUDENTS</span>
                    <v-icon size="24" class="pl-2">
                      {{ isOpen ? 'mdi-chevron-up' : 'mdi-chevron-down' }}
                    </v-icon>
                  </template>
                </v-list-item>
              </template>

              <template v-for="(item, i) in assignments" :key="'student-'+i">
                <v-list-item
                  v-ripple="false"
                  @click="toggleAssignment(item)"
                  :value="item"
                  :title="item.name"
                  rounded="lg"
                  class="border-b-sm"
                >
                  <!-- keep content minimal or add slots as needed -->
                </v-list-item>
              </template>
            </v-list-group>
          </v-list>

          <!-- ASSIGNMENTS group -->
          <v-list lines="two" class="container">
            <v-list-group value="Assignments">
              <template v-slot:activator="{ props, isOpen }">
                <v-list-item
                  v-bind="props"
                  class="list-header"
                  v-ripple="false"
                  rounded="lg"
                >
                  <template #prepend>
                    <span class="header-title">ASSIGNMENTS</span>
                    <v-icon size="24" class="pl-2">
                      {{ isOpen ? 'mdi-chevron-up' : 'mdi-chevron-down' }}
                    </v-icon>
                  </template>

                  <template #append>
                    <router-link to="/assignments-overview/create" class="add-button">
                      <v-icon
                        icon="mdi-plus"
                        size="24"
                        color="#000000"
                        class="pr-4"
                      />
                    </router-link>
                  </template>
                </v-list-item>
              </template>

              <template v-for="(item, i) in assignments" :key="'assign-'+i">
                <v-list-item
                  v-ripple="false"
                  @click="toggleAssignment(item)"
                  :value="item"
                  :title="item.name"
                  rounded="lg"
                  class="border-b-sm pl-4"
                >
                  <template #prepend>
                    <v-icon class="al-icon icon" icon="mdi-file-document-multiple-outline" size="20"></v-icon>
                  </template>
                  <template #subtitle>
                    <span class="dueDate">Due {{ formatDateMDY(item.dueDate) }}</span>
                  </template>
                </v-list-item>
              </template>
            </v-list-group>
          </v-list>
        </div>

        <!-- right column: details view -->
        <FacultyAssignmentsView
          v-if="isVisible"
          :assignment="selectedAssignment"
          :teamMember="selectedTeamMember"
          class="faculty-view"
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
  height: 59.875rem;
  // padding-top: 3rem;
}

.assignments-flex {
  display: flex;
  flex-direction: row;
  gap: 2rem;
  width: 100%;
  align-items: flex-start; /* align top */
  height: 100%;
}

/* left column holds stacked lists */
.left-column {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  width: 20rem;
  margin-top: 1.25rem;
  align-self: flex-start;

  max-height: calc(100% - 4rem); /* adjust the subtraction if the top offset is different */
  overflow-y: auto;
  overflow-x: hidden;
  padding-right: 0.5rem; /* room for scrollbar */
}

/* Prevent inner v-list from introducing its own scroll or extra height */
:deep(.left-column > .v-list),
:deep(.left-column .v-list) {
  max-height: none !important;
  overflow: visible !important;
  margin: 0 !important;
  padding: 0 !important;
}

/* Tweak inner item padding so content aligns when scrolled */
:deep(.left-column .v-list .v-list-item) {
  padding-left: 8px !important;
}

/* keep view pinned to top and allow it to grow */
.faculty-view {
  flex: 1 1 0;
  align-self: flex-start;
  margin-top: 1.25rem;
  height: calc(100% - 1.25rem);
  overflow: auto; 
}

/* optional: prettier scrollbars (WebKit) */
.left-column::-webkit-scrollbar {
  width: 10px;
}
.left-column::-webkit-scrollbar-track {
  background: transparent;
}
.left-column::-webkit-scrollbar-thumb {
  background-color: rgba(0,0,0,0.12);
  border-radius: 8px;
  border: 2px solid transparent;
}

/* header and item styles (keep existing overrides) */
:deep(.list-header) {
  padding: 0.5rem 1rem;

  .header-title {
    font-size: 16px;
    font-weight: 600;
    color: black;
  }

  .add-button {
    text-decoration: none;
  }
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

/* Vuetify overrides (fine-tune as needed) */
:deep(.v-list-group) {
  .v-list-group__header.v-list-item {
    padding: 8px 0 !important;

    .v-list-item__prepend {
      padding-left: 8px !important;
    }
  }

  .v-list-group__items {
    .v-list-item {
      padding: 8px 0 !important;

      .v-list-item__content {
        padding-left: 0 !important;
      }

      .v-list-item__prepend {
        padding-left: 8px !important;
      }
    }
  }
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