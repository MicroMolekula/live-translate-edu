<script setup>
import CardLesson from "@/components/available-lessons/CardLesson.vue";
import {onMounted, ref} from "vue";
import {userStore} from "@/stores/stores.js";
import {getLesson} from "@/lib/request/lesson.js";

// const lessons = [
//   {
//     title: "Тестовое занятия",
//     numberRoom: 255,
//     teacher: 'Teacher Teacherov',
//     datetime: '01-02-2025 10:10'
//   },
//   {
//     title: "Тестовое занятия",
//     numberRoom: 255,
//     teacher: 'Teacher Teacherov',
//     datetime: '01-02-2025 10:10'
//   },
//   {
//     title: "Тестовое занятия",
//     numberRoom: 255,
//     teacher: 'Teacher Teacherov',
//     datetime: '01-02-2025 10:10'
//   },
//   {
//     title: "Тестовое занятия",
//     numberRoom: 255,
//     teacher: 'Teacher Teacherov',
//     datetime: '01-02-2025 10:10'
//   },
//   {
//     title: "Тестовое занятия",
//     numberRoom: 255,
//     teacher: 'Teacher Teacherov',
//     datetime: '01-02-2025 10:10'
//   }
// ]

const userData = userStore()

const lessons = ref([])

function getDataLesson(data) {
  let theme = ''
  for (let item of data.contents) {
    if (item.language_code === 'ru') {
      theme = item.theme
    }
  }
  return {
    title: theme,
    numberRoom: data.number_room,
    teacher: data.teacher,
    datetime: data.date_start,
    code: data.code
  }
}

onMounted(async () => {
  try {
    let result = await getLesson(userData.token)
    if (result.status === 200) {
      for (let item of result.data) {
        lessons.value.push(getDataLesson(item))
      }
    } else {
      console.log('Ошибка получения списка занятий')
    }
  } catch {
    console.log('Ошибка получения списка занятий')
  }
})

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