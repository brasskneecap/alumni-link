<template>
  <div>
    <ALCard
      class="top-container"
      title="Create New Assignment"
      icon="mdi-file-document-multiple-outline"
      height="19.875rem"
      width="96.375rem"
    >
      <template #header-right>
        <router-link to="/assignments-overview">
          <v-btn>back</v-btn>
        </router-link>      
      </template>
  
      <template #content>
        <div style="padding: 24px 12px 12px 12px;">
          <v-text-field 
            v-model="title"
            label="Assignment Title"
            variant="outlined"></v-text-field>
          <v-textarea 
            v-model="description"
            label="Assignment Details / Instructions"
            variant="outlined"></v-textarea>
        </div>
      </template>
    </ALCard>
  
    <ALCard
      class="container"
      title="Due Date"
      icon="mdi-calendar-outline"
      height="10.875rem"
      width="96.375rem"
    >
      <template #content>
        <div class="due-date-container">
          <v-date-input
            v-model="dueDate"
            label="Due Date"
            prepend-icon=""
            append-inner-icon="mdi-calendar-outline"
            variant="outlined"
          ></v-date-input>
          <v-text-field
            v-model="dueTime"
            label="Select Time"
            type="time"
            variant="outlined"
            density="comfortable"
          />
          <v-date-input
            v-model="publishDate"
            label="Publish Date"
            prepend-icon=""
            append-inner-icon="mdi-calendar-outline"
            variant="outlined"
          ></v-date-input>
        </div>
      </template>
    </ALCard>
    <ALCard
      class="container"
      title="Submission Type"
      icon="mdi-checkbox-outline"
      height="10.875rem"
      width="96.375rem"
    >
      <template #content>
        <div class="due-date-container">
          <v-checkbox
            v-model="allowedContent"
            color="#4065DA"
            label="Text"
            value="Text"
          ></v-checkbox>
          <v-checkbox
            v-model="allowedContent"
            color="#4065DA"
            label="File Upload"
            value="Upload"
          ></v-checkbox>
          <v-checkbox
            v-model="allowedContent"
            color="#4065DA"
            label="URL"
            value="URL"
          ></v-checkbox>
        </div>
      </template>
    </ALCard>
    <div class="button-actions">
      <v-btn
        rounded="xl"
        prepend-icon="mdi-pencil"
        color="#E0EAFA"
      >
        Save As Draft
      </v-btn>
      <v-btn
        rounded="xl"
        prepend-icon="mdi-eye-outline"
        color="#E0EAFA"
      >
        Preview
      </v-btn>
      <v-btn
        rounded="xl"
        prepend-icon="mdi-content-save"
        color="#E0EAFA"
        @click="createAssignment"
      >
        Create Assignment
      </v-btn>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import ALCard from '../../reusables/ALCard.vue'

const title = ref("")
const description = ref("")

const dueDate = ref(null)
const dueTime = ref(null)
const publishDate = ref(null)

const allowedContent = ref([])

const createAssignment = () => {
  console.log(title.value)
  console.log(description.value)
  console.log(dueDate.value)
  console.log(dueTime.value)
  console.log(publishDate.value)
  console.log(allowedContent.value)
}
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.top-container {
  margin: 7.12rem 2rem 2.81rem 2rem;
}

.container {
  margin: 2rem 2rem 2.81rem 2rem;
}

.assignments-flex {
  display: flex;
  flex-direction: row;
  gap: 2rem;
  width: 100%;
}

.assignments-flex > *:first-child {
  flex: 0 0 32rem;
}

.assignments-flex > *:last-child {
  flex: 1 1 0;
}

.due-date-container {
  display: flex;
  gap: 24px;
  padding: 24px 12px 12px 12px;
}

.button-actions {
  display: flex;
  justify-content: end;
  gap: 24px;
  padding-right: 32px;
}
</style>