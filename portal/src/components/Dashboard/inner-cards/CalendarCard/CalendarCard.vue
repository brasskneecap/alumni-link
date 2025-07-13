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
            :class="{
              // 'has-event': hasEvent(item.date),
              'is-selected': isSelected(item.date),
              'is-today': isToday(item.date)
            }"
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
    
    // const hasEvent = (dateStr) => {
    //   return eventDates.value.includes(dateStr.slice(0, 10)) // normalize to YYYY-MM-DD
    // }

    const isSelected = (date) => {
      const dateStr = new Date(date).toISOString().slice(0, 10)
      return selectedDate.value === dateStr
    }

    const isToday = (date) => {
      const dateStr = new Date(date).toISOString().slice(0, 10)
      const today = new Date().toISOString().slice(0, 10)
      return dateStr.slice(0, 10) === today
    }

    const onDayClick = (day) => {
      const year = calendarMonth.value.getFullYear()
      const month = calendarMonth.value.getMonth()

      const fullDate = new Date(year, month, day.text) 
      const dateStr = fullDate.toISOString().slice(0, 10)

      selectedDate.value = dateStr
    }

    return {
      selectedDate,
      calendarMonth,
      displayedMonth,
      calendarHeader,
      onDayClick,
      // hasEvent,
      isToday,
      isSelected,
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

.day-cell {
  border-radius: 16px;
  line-height: 3rem;
  cursor: pointer;
  transition: background 0.2s ease;

  &.has-event {
    background-color: rgba(25, 118, 210, 0.1);
    border: 1px solid #1976d2;
    color: #1976d2;
  }

  &.is-selected {
    background-color: rgba(25, 118, 210, 0.15);
    color: white;
    font-weight: bold;
  }

  &.is-today {
    background-color: #1976d2;
    color: white;
    font-weight: bold;
  }

  &:hover {
    background-color: rgba(25, 118, 210, 0.15);
  }
}
</style>