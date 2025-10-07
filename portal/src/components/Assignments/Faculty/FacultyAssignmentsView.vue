<template>
  <div class="container">
    <div class="header">
      <div v-if="!isEditing">
        <h1 class="title">{{ assignment.name }}</h1>
      </div>
      <div v-else>
        <v-text-field 
          class="edit-title" 
          v-model="editableAssignment.name" 
          label="Assignment Name"
          variant="outlined" />
      </div>

      <v-btn v-if="!isEditing" @click="startEditing">Edit</v-btn>
      <div v-else class="d-flex gap-2">
        <v-btn color="primary" @click="saveChanges">Save</v-btn>
        <v-btn variant="text" @click="cancelEditing">Cancel</v-btn>
      </div>
    </div>

    <TeamMember
      v-if="teamMember"
      :src="teamMember.profilePicture"
      :name="teamMember.name"
      :role="teamMember.role"
      :description="`${teamMember.role} at ${teamMember.school}`"
    />


    <p v-if="!isEditing" class="description">{{ assignment.description }}</p>
    <v-textarea
      v-else
      class="edit-description"
      v-model="editableAssignment.description"
      label="Description"
      rows="3"
      variant="outlined"
    >
    </v-textarea>


    <div class="infoContainer">
      <div class="info mt-12">
        <v-icon class="infoIcon al-icon" icon="mdi-calendar-month-outline" size="24"/>
        <p v-if="!isEditing">Due {{ formatDateMDY(assignment.dueDate) }}</p>
        <v-date-input
          v-else
          class="mt-5"
          width="250"
          v-model="editableAssignment.dueDate"
          label="Due Date"
          prepend-icon=""
          append-inner-icon="mdi-calendar-outline"
          variant="outlined"
        ></v-date-input>
      </div>

      <div class="info mt-12">
        <v-icon class="infoIcon al-icon" icon="mdi-folder-outline" size="24"/>
        <div>
          <p class="submitPar">Submission Type</p>
          <p v-if="!isEditing" class="anyPar">{{ assignment.allowedContent.join(', ') }}</p>
          <v-select
            v-else            
            v-model="editableAssignment.allowedContent"
            :items="['text', 'upload', 'url']"
            multiple
            variant="outlined"
          />
        </div>
      </div>
    </div>
  </div>
</template>


<script setup>
import { ref, reactive } from 'vue'
import TeamMember from '../../Dashboard/inner-cards/TeamCard/TeamMember.vue'
import { formatDateMDY } from '@/utils/formatters'

const props = defineProps({
  assignment: Object,
  teamMember: Object
})

// local editing state
const isEditing = ref(false)

// create a reactive copy for editing
const editableAssignment = reactive({})

// create a helper to start editing
const startEditing = () => {
  Object.assign(editableAssignment, props.assignment)
  isEditing.value = true
}

// helper to cancel edits
const cancelEditing = () => {
  isEditing.value = false
  Object.keys(editableAssignment).forEach(k => delete editableAssignment[k])
}

// save handler (later you can connect this to your API)
const saveChanges = async () => {
  console.log('Submitting updated assignment:', editableAssignment)

  // example: call your API here
  // await api.updateAssignment(editableAssignment.id, editableAssignment)
  
  isEditing.value = false
}
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.v-btn:focus,
.v-tab:focus {
  outline: none !important;
  box-shadow: none !important;
}

.container {
  padding: 0 4rem;
  min-width: 60rem;
}

.header {
  display: flex;
  justify-content: space-between ;
}

.title {
  font-size: 24px;
  font-weight: 700;
}

.edit-title {
  font-size: 24px;
  width: 500px;
}

.infoContainer {
  display: flex;
  flex-direction: row;
  gap: 4rem;

    .info {
    display: flex;
    align-items: center;
    font-size: 24px;
    color: $al-secondary-gray;
    gap: 0.3rem;
    margin: 2rem 0;

    .infoIcon {
      padding: 1.5rem;
      margin-right: 1rem;
    }

    .submitPar {
      font-size: 16px;
      font-weight: 400;
    }

    .typePar {
      font-size: 18px;
      font-weight: 500;
      color: black;
    }
  }
}

.description {
  margin-top: 1rem;
}

.edit-description {
  margin-top: 3rem;
}

// tabs

.tabs {
  margin: 1.5rem 0 2rem 0;
}

.tab {
  border-bottom: 1px solid $al-gray;
}

:deep(.v-tabs-window),
:deep(.v-tabs-window-item),
.tabWindow {
  width: 40rem;
  max-width: 100%;
  overflow: visible;
}

:deep(.v-tab--selected) {
  font-weight: 700 !important;
}

.chooseText {
  font-style: italic;
  font-weight: 300;
  margin-bottom: 2rem;
}

// tab window

:deep(.v-card) {
  width: 40rem;
  max-width: 100%;
  overflow: visible;
  box-shadow: none;
}

.tabWindow {
  width: 40rem;

  .url {
    margin-bottom: -2rem;
  }
}

.allowedContent {
  text-transform: capitalize;
}

.subtitle {
  font-size: 16px;
  font-weight: 500;
  margin-top: 1.5rem;
}
</style>