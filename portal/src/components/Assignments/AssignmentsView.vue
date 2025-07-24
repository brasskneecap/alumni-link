<template>
  <div class="container">
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
          <v-tab v-for="v in assignment.allowedContent" 
            class="tab text-none text-capitalize"
            v-ripple="false" 
            :value="v">{{ v }}
          </v-tab>
        </v-tabs>

        <p class="chooseText">Choose a submission type</p>

        <v-tabs-window class="tabWindow" v-model="tab">
          <v-tabs-window-item
            v-for="(v, k) in assignment.allowedContent"
            :key="k"
            :value="v"
            :transition="false"
            :reverse-transition="false"
          >
            <v-textarea v-if="v === 'text'" variant="solo" class="textArea"></v-textarea>
            <v-file-upload v-else-if="v === 'upload'" 
              density="comfortable" 
              variant="comfortable"
              title="Drag and drop your files here, or click to browse"
              icon="mdi-tray-arrow-up"
            ></v-file-upload>
            <v-text-field v-else-if="v === 'url'" label="Enter URL" variant="outlined" class="url"></v-text-field>
          </v-tabs-window-item>
        </v-tabs-window>
      </v-card>
      <v-btn 
        variant="flat" 
        v-ripple="false" 
        class="submit 
        text-none" 
        color="#DEEAFC"
        rounded="xl"
        size="large"
      >
        Submit Assignment
      </v-btn>
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
import { ref } from 'vue'
import { VFileUpload } from 'vuetify/labs/VFileUpload'
import { formatDateMDY } from '@/utils/formatters';

const tab = ref(null)
const props = defineProps({
  assignment: Object,
  teamMember: Object
});
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

:deep(.v-textarea) {
  background: transparent;
  box-shadow: none;
  padding: 0;
  overflow: visible;
  margin-bottom: -2rem;
}

:deep(.v-textarea .v-field) {
  box-shadow: 0 0 24px rgba(0,0,0,0.12);
  width: 40rem;
  max-width: 100%;
  margin-left: auto;
  margin-right: auto;
  display: block;
}

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

.subtitle {
  font-size: 16px;
  font-weight: 500;
  margin-top: 1.5rem;
}

.submit {
  margin-top: 2rem;
}

.feedback {
  display: flex;
  justify-content: flex-end;
  margin-right: 4rem;
  font-family: 'Poppins', sans-serif !important;
}
</style>

