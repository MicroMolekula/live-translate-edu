<script setup>
import {Card, CardContent, CardFooter, CardHeader} from "@/components/ui/card/index.js";
import {Button} from "@/components/ui/button/index.js";
import {useRouter} from "vue-router";

const props = defineProps({
  lesson: {
    type: Object,
    required: true
  }
})

const router = useRouter()

function handleButton() {
  router.push('/lesson/' + props.lesson.code)
}
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
    <Button @click="handleButton">Присоедениться</Button>
  </CardFooter>
</Card>
</template>

<style scoped>

</style>