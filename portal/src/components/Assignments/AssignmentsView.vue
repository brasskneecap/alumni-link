<template>
  <div>
    <div>
      <h1 class="title">{{ assignment.name }}</h1>
      <TeamMember
        v-if="teamMember"
        :src="teamMember.profilePicture"
        :name="teamMember.name"
        :role="teamMember.role"
        :description="`${teamMember.role} at ${teamMember.school}`"
      />
      <p class="description">{{ assignment.description }}</p>
    </div>

    <div class="infoContainer">
      <div class="info">
        <v-icon class="infoIcon al-icon" icon="mdi-calendar-month-outline" size="24"/>
        <p>Due {{ formatDateMDY(assignment.dueDate) }}</p>
      </div>
      <div class="info">
        <v-icon class="infoIcon al-icon" icon="mdi-folder-outline" size="24"/>
        <div>
          <p class="submitPar">Submission Type</p>
          <p class="anyPar">Any</p>
        </div>
      </div>
    </div>
    
    <div>
      <h2 class="subtitle">Submit Assignment</h2>
      <v-card class="elevation-0">
        <v-tabs
          v-model="tab"
          align-tabs="start"
          slider-color="#B8D3FF"
          class="tabs"
        >
          <v-tab class="text-none" v-ripple="false" :value="1">Text</v-tab>
          <v-tab class="text-none" v-ripple="false" :value="2">Upload Document</v-tab>
          <v-tab class="text-none" v-ripple="false" :value="3">URL</v-tab>
        </v-tabs>

        <p class="chooseText">Choose a submission type</p>

        <v-tabs-window class="tabWindow" v-model="tab">
          <v-tabs-window-item
            v-for="n in 3"
            :key="n"
            :value="n"
          >
            <v-textarea v-if="n === 1" variant="solo" class="textArea"></v-textarea>
            <v-file-upload v-else-if="n === 2" 
              density="comfortable" 
              variant="comfortable"
              title="Drag and drop your files here, or click to browse"
              icon="mdi-tray-arrow-up"
            ></v-file-upload>
            <v-text-field v-else-if="n === 3" label="Enter URL" variant="outlined"></v-text-field>
          </v-tabs-window-item>
        </v-tabs-window>

        <v-btn 
          variant="flat" 
          v-ripple="false" 
          class="submit 
          text-none" 
          color="#DEEAFC"
          rounded="xl"
        >
          Submit Assignment
        </v-btn>

      </v-card>
    </div>
  </div>

  <div class="feedback">
    <v-btn variant="outlined" v-ripple="false" class="text-none">
      Message Alumni
    </v-btn>
  </div>
</template>


<script setup>
import TeamMember from '../Dashboard/inner-cards/TeamCard/TeamMember.vue';
import { ref, shallowRef } from 'vue'
import { VFileUpload } from 'vuetify/labs/VFileUpload'

const density = shallowRef('default')
const tab = ref(null)
const props = defineProps({
  assignment: Object,
  teamMember: Object
});

function formatDateMDY(dateStr) {
  const date = new Date(dateStr)
  const month = String(date.getMonth() + 1)
  const day = String(date.getDate())
  const year = date.getFullYear().toString().slice(-2)
  return `${month}/${day}/${year}`
}
</script>

<style lang="scss">
@use '@/variables.scss' as *;

:deep(.v-btn__content),
:deep(.v-tab__content) {
  font-family: 'Poppins', sans-serif !important;
}

.v-btn:focus,
.v-tab:focus {
  outline: none !important;
  box-shadow: none !important;
}

.title {
  font-size: 24px;
  font-weight: 700;
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
    margin: 3rem 0 4rem 0;

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

.tabs {
  margin: 1.5rem 0 2rem 0;
  border-bottom: 1px solid $al-gray;
  max-width: 22.5rem;
}

.chooseText {
  font-style: italic;
  font-weight: 300;
}

.textArea {
  margin: 0.1rem;
}

.tabWindow {
  margin-top: 2rem;
  width: 40rem;
}

.subtitle {
  font-size: 16px;
  font-weight: 500;
  margin-top: 1.5rem;
}

.feedback {
  display: flex;
  justify-content: flex-end;
  margin-right: 4rem;
  font-family: 'Poppins', sans-serif !important;
}
</style>