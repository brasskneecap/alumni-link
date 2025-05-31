<template>
  <div>

  </div>
  <div>
    <div>
      {{ assignments[0].name }}
      Due Date: {{ assignments[0].dueDate }}
    </div>
    <div>
      Profile Card (profile.information)
    </div>
    <div>
      Description (description)
    </div>
    <div>
      submission card (type)
    </div>
  </div>
</template>

<script>
import { computed } from 'vue';
import { useStore } from 'vuex'
export default {
  setup() {
    const store = useStore()
    
    const assignments = computed(() => store.getters["assignments/assignments"]);

    console.log(assignments)
    return {
      assignments
    }
  }
}
</script>