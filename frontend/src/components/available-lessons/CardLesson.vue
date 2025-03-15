<script setup>
import {Card, CardContent, CardFooter, CardHeader} from "@/components/ui/card/index.js";
import {Button} from "@/components/ui/button/index.js";

let lesson = {
  title: "Тестовое занятия",
  numberRoom: 255,
  teacher: 'Teacher Teacherov',
  datetime: '01-02-2025 10:10'
}

const props = defineProps({
  lesson: {
    type: Object,
    required: true
  }
})
</script>

<template>
<Card class="m-4">
  <CardHeader class="font-bold">
    {{ props.lesson.title }}
  </CardHeader>
  <CardContent>
    <div>Аудитория: {{ props.lesson.numberRoom }}</div>
    <div>Преподаватель: {{ props.lesson.teacher }}</div>
    <div>Дата и время начала: {{ props.lesson.datetime }}</div>
  </CardContent>
  <CardFooter>
    <Button>Присоедениться</Button>
  </CardFooter>
</Card>
</template>

<style scoped>

</style>