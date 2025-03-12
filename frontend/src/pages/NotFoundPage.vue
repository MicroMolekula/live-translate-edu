<template>
  <div class="min-h-screen flex flex-col items-center justify-center bg-background text-foreground p-4">
    <div class="text-center space-y-4">
      <h1 class="text-9xl font-bold text-primary">404</h1>
      <p class="text-2xl text-muted-foreground">Oops! Страница не найдена.</p>
      <p class="text-lg text-muted-foreground">
        Кажется, вы забрели туда, куда не стоило. Давайте вернемся на главную.
      </p>
      <Button @click="goHome" class="mt-6">Вернуться на главную</Button>
    </div>
  </div>
</template>

<script setup>
import { Button } from '@/components/ui/button';
import {useRouter} from "vue-router";

const router = useRouter()

const goHome = () => {
  router.push("/test")
};
</script>

<style scoped>
/* Добавляем минимальные стили для центрирования и фона */
.min-h-screen {
  min-height: 100vh;
}
</style>