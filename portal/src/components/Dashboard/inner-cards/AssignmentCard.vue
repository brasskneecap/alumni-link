<template>
  <ALCard
    :title="title"
    icon="mdi-file-document-multiple-outline"
  >
    <template #content>
      <div class="tracker-bottom">
        <div class="tracker-content">{{ content }}</div>
        <div class="tracker-subtitle">
          <p>Assignments Completed</p>
        </div>
      </div>
    </template>
  </ALCard>
</template>

<script>
import { toRefs  } from 'vue'
import ALCard from '../../reusables/ALCard.vue';

export default {
  components: {
    ALCard,
  },
  props: {
    title: {
      type: String,
      required: true
    },
    content: {
      type: String,
      required: true
    }
  },
  setup (props) {
    const { title, content } = toRefs(props)

    return {
      title,
      content
    }
  }
}
</script>

<style lang="scss" scoped>
@use '@/variables.scss' as *;

.tracker-bottom {
  margin-top: auto;
  margin-bottom: 1rem;

  .tracker-content {
    font-size: 32px;
    font-weight: 700;
    margin-top: 2.44rem;
  }

  .tracker-subtitle {
    font-size: 16px;
    font-weight: 400;
    margin-top: 0.38rem;
    color: $al-secondary-gray;
  }
}

</style>