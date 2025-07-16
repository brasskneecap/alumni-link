<template>
  <div>
    <div>
      <h1 class="title">{{ assignment.name }}</h1>
      <div class="dueDate">
        <v-icon icon="mdi-calendar-month-outline" size="24"/>
        <p>Due {{ formatDateMDY(assignment.dueDate) }}</p>
      </div>
      <TeamMember
        v-if="teamMember"
        :src="teamMember.profilePicture"
        :name="teamMember.name"
        :role="teamMember.role"
        :description="`${teamMember.role} at ${teamMember.school}`"
      />
      <p class="description">{{ assignment.description }}</p>
    </div>
    
    <div>
      <h2 class="subtitle">Submit Assignment</h2>
      <v-card class="elevation-0">
        <v-tabs
          v-model="tab"
          align-tabs="start"
        >
          <v-tab class="text-none" v-ripple="false" :value="1">Text</v-tab>
          <v-tab class="text-none" v-ripple="false" :value="2">Upload Document</v-tab>
          <v-tab class="text-none" v-ripple="false" :value="3">URL</v-tab>
        </v-tabs>
        <v-tabs-window v-model="tab">
          <v-tabs-window-item
            v-for="n in 3"
            :key="n"
            :value="n"
          >
            <v-textarea v-if="n === 1" label="Upload Document" variant="outlined"></v-textarea>
            <v-file-upload v-else-if="n === 2" 
              density="comfortable" 
              variant="comfortable"
              title="Drag and drop your files here, or click to browse"
              icon="mdi-tray-arrow-up"
            ></v-file-upload>
            <v-text-field v-else-if="n === 3" label="Enter URL" variant="outlined"></v-text-field>
          </v-tabs-window-item>
        </v-tabs-window>
      </v-card>
    </div>
  </div>

  <div class="feedback">
    <v-btn variant="outlined" v-ripple="false" class="text-none">
      Feedback
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

<style lang="scss" scoped>
@use '@/variables.scss' as *;
.title {
  font-size: 24px;
  font-weight: 700;
}
.dueDate {
  display: flex;
  align-items: center;
  font-size: 24px;
  color: $al-secondary-gray;
  gap: 0.3rem;
  margin: 0.5rem 0 1.8rem 0;
}
.description {
  margin-top: 2.3rem;
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