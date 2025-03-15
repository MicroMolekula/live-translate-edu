<script setup>
import CardLesson from "@/components/available-lessons/CardLesson.vue";
import {ref} from "vue";

const lessons = [
  {
    title: "Тестовое занятия",
    numberRoom: 255,
    teacher: 'Teacher Teacherov',
    datetime: '01-02-2025 10:10'
  },
  {
    title: "Тестовое занятия",
    numberRoom: 255,
    teacher: 'Teacher Teacherov',
    datetime: '01-02-2025 10:10'
  },
  {
    title: "Тестовое занятия",
    numberRoom: 255,
    teacher: 'Teacher Teacherov',
    datetime: '01-02-2025 10:10'
  },
  {
    title: "Тестовое занятия",
    numberRoom: 255,
    teacher: 'Teacher Teacherov',
    datetime: '01-02-2025 10:10'
  },
  {
    title: "Тестовое занятия",
    numberRoom: 255,
    teacher: 'Teacher Teacherov',
    datetime: '01-02-2025 10:10'
  }
]

const props = defineProps({
  openSidebar: {
    type: Boolean,
    request: false,
    default: true
  }
})

</script>

<template>
  <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-4 gap-1">
    <div v-for="lesson in lessons" :key="lesson.title">
      <CardLesson :lesson="lesson"/>
    </div>
  </div>
</template>

<style scoped>

</style>