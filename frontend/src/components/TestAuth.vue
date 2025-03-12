<script setup>
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
</script>

<template>
  <div class="w-full lg:grid lg:min-h-[600px] lg:grid-cols-2 xl:min-h-[800px]" style="height: 100vh">
    <div class="flex items-center justify-center py-12 border-r-4">
      <div class="mx-auto grid w-[350px] gap-6">
        <div class="grid gap-2 text-center">
          <h1 class="text-3xl font-bold">
            LiveTranslateEdu
          </h1>
          <p class="text-balance text-muted-foreground">
          </p>
        </div>
        <div class="grid gap-4">
          <div class="grid gap-2">
            <Label for="email">Email</Label>
            <Input
                id="email"
                type="email"
                placeholder="m@example.com"
                required
            />
          </div>
          <div class="grid gap-2">
            <div class="flex items-center">
              <Label for="password">Пароль</Label>
<!--              <a-->
<!--                  href="/forgot-password"-->
<!--                  class="ml-auto inline-block text-sm underline"-->
<!--              >-->
<!--                Forgot your password?-->
<!--              </a>-->
            </div>
            <Input id="password" type="password" required />
          </div>
          <div class="grid gap-2">
            <Label for="room">Комната</Label>
            <Input
                id="room"
                type="text"
                required
            />
          </div>
          <Button type="submit" class="w-full">
            Войти
          </Button>
        </div>
<!--        <div class="mt-4 text-center text-sm">-->
<!--          Don't have an account?-->
<!--          <a href="#" class="underline">-->
<!--            Sign up-->
<!--          </a>-->
<!--        </div>-->
      </div>
    </div>
    <div class="bg-muted lg:block place-items-center mt-60">
      <img
          src="/lstu.svg"
          alt="Image"
          class="w-60 object-cover dark:brightness-[0.2] dark:grayscale"
      >
    </div>
  </div>
</template>