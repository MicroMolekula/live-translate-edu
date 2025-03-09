<script setup>
import {Button} from "@/components/ui/button/index.js";

const emit = defineEmits(['chat'])

function openChat() {
  emit('chat', true)
}
</script>

<template>
<div class="fixed bottom-3 place-items-center " style="width:100%">
  <div class="flex flex-row justify-between w-11/12">
    <!-- Название занятия -->
    <div class="flex items-center">
      Тестовое занятие
    </div>
    <!-- Общие кнопки -->
    <div class="flex gap-2">
      <Button variant="secondary">Hello</Button>
      <Button variant="secondary">Hello</Button>
    </div>
    <!-- Вспомогательные кнопки -->
    <div class="flex gap-2">
      <Button @click="openChat" variant="secondary">Чат</Button>
      <Button variant="secondary">Участники</Button>
    </div>
  </div>
</div>
</template>

<style scoped>

</style>