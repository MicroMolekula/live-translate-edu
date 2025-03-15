<script setup>
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import {ref} from "vue";
import login from '@/lib/request/login.js'
import {userStore} from '@/stores/userStore.js'
import {useRouter} from "vue-router";

const userData = ref({
  login: '',
  password: ''
})

const roomName = ref('')

const router = useRouter()
const userStoreData = userStore()

function handleLogin() {
  login(userData.value.login, userData.value.password)
      .then(function (response) {
        userStoreData.token += response
        router.push('/')
      })
      .catch(function (e) {
        console.log("Ошибка " + e)
      })
}

</script>

<template>
  <div class="w-full lg:grid lg:min-h-[600px] lg:grid-cols-2 xl:min-h-[800px]" style="height: 100vh">
    <div class="bg-muted lg:block place-items-center" style="background-color: rgba(24, 24, 27, 1.00)">
      <img
          src="/lstu.svg"
          alt="Image"
          class="w-60 object-cover dark:brightness-[0.2] dark:grayscale mt-60"
      >
    </div>
    <div class="flex items-center justify-center py-12">
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
                v-model="userData.login"
                required
            />
          </div>
          <div class="grid gap-2">
            <div class="flex items-center">
              <Label for="password">Пароль</Label>
            </div>
            <Input id="password" type="password" v-model="userData.password" required />
          </div>
<!--          <div class="grid gap-2">-->
<!--            <Label for="room">Комната</Label>-->
<!--            <Input-->
<!--                id="room"-->
<!--                type="text"-->
<!--                v-model="roomName"-->
<!--                required-->
<!--            />-->
<!--          </div>-->
          <Button type="submit" @click="handleLogin" class="w-full">
            Войти
          </Button>
        </div>
      </div>
    </div>
  </div>
</template>