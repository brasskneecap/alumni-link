<template>
    <div class="calendar-container al-card">
      <div class="calendar-header">
        <div class="calendar-title">
          {{ calendarHeader }}
        </div>
        <div class="calendar-nav">
          <v-icon
            class="icon"
            icon="mdi-chevron-left-circle-outline"
            size="28"
            @click="changeMonth(-1)"
          ></v-icon>
          <v-icon
            class="icon"
            icon="mdi-chevron-right-circle-outline"
            size="28"
            @click="changeMonth(1)"
          ></v-icon>
        </div>
      </div>
      <v-date-picker-month
        :month="calendarMonth.getMonth()"
      >
        <template #day="{props, item}">
          <div
            class="day-cell"
            @click="onDayClick(props)"
          >
            {{ new Date(item.date).getDate() }}
          </div>
        </template>
      </v-date-picker-month>
    </div>
</template>

<script>
import { computed, ref } from 'vue'

export default {
  methods: {
    changeMonth(delta) {
      const newDate = new Date(this.calendarMonth)
      newDate.setMonth(this.calendarMonth.getMonth() + delta)
      newDate.setDate(1)
      this.calendarMonth = newDate
    },
    onDayClick(date) {
      // Handle day click event
      console.log("Clicked on date:", date);
    },
  },
  setup () {
    const selectedDate = ref(new Date().toISOString().slice(0, 10))
    const displayedMonth = ref('2025-05-01')
    const calendarMonth = ref(new Date())

    const calendarHeader = computed(() => {
      let dateMonth = calendarMonth.value.toLocaleString('default', { month: 'long' });
      let dateYear = calendarMonth.value.getUTCFullYear()
      return `${dateMonth}  ${dateYear}`
    });
    

    return {
      selectedDate,
      calendarMonth,
      displayedMonth,
      calendarHeader,
    }
  }
}
</script>

<style lang="scss">
@use '@/variables.scss' as *;
.calendar-container {
  width: 25rem;
  height: 22rem;

  .calendar-header {
    display: flex;
    justify-content: space-between;
    padding: 18px 30px 9px 30px;

    .calendar-title {
      font-size: 20px;
      font-weight: 600;
    }

    .calendar-nav {
      display: flex;
      gap: 12px;
    }
  }
}
</style>