<template>
  <div class="student-assignment-items">
    <UpcomingItem v-for="assignment in studentAssignments"
      :key="assignment.id"
      :title="assignment.name"
      :date="formatDate(assignment.dueDate)"
      :tag="assignment.status"
      :tagClass="getStatusClass(assignment.status)"
      :description="assignment.description"
      :class="getStatusClass(assignment.status)"
    />
  </div>
</template>

<script>
import UpcomingItem from '../../Dashboard/inner-cards/UpcomingCard/UpcomingItem.vue';
import { computed } from 'vue';
import { useStore } from 'vuex';
import { formatDate } from '@/utils/formatters';

export default {
  components: {
      UpcomingItem,
  },
  setup() {
    const store = useStore()
    const studentAssignments = computed(() => {
      const assignments = store.getters["assignments/assignments"]
      return assignments.map(assignment => ({
        ...assignment,
        status: getAssignmentStatus(assignment)
      }));
    });

    const getAssignmentStatus = (assignment) => {
      const now = new Date();
      const dueDate = new Date(assignment.dueDate);
      
      // Has submission
      if (assignment.submission) {
        if (assignment.submission.status === 'completed') {
          return 'Completed';
        }
        const submittedDate = new Date(assignment.submission.submittedAt);
        return submittedDate > dueDate ? 'Late' : 'Completed';
      }
      
      return 'Pending';
    }

    const getStatusClass = (status) => {
      const statusMap = {
        'Completed': 'status-completed',
        'Late': 'status-late',
        'Pending': 'status-pending'
      };
      return statusMap[status] || '';
    }

    
    return {
      studentAssignments,
      formatDate,
      getStatusClass
    }
  }
}
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.student-assignment-items {
  margin-top: 1rem;
}

.status-completed {
  :deep(.tag) {
    background-color: #6DD87F !important;
    color: #115402 !important;
  }
}

.status-late {
  :deep(.tag) {
    background-color: #FF3B4A !important;
    color: #fff !important;
  }
}

.status-pending {
  :deep(.tag) {
    background-color: #A5A5A5 !important;
    color: #000 !important;
  }
}
</style>
