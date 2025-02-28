<template>
  <div class="flex items-center justify-center min-h-screen bg-gray-100">
    <div class="w-full max-w-md p-6 space-y-6 bg-white rounded-lg shadow-md">
      <h1 class="text-3xl font-semibold text-center text-green-600 mb-6">LiveTranslateEdu</h1>

      <!-- Login Form -->
      <form @submit.prevent="handleLogin" class="space-y-4">
        <!-- Email input -->
        <div>
          <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
          <input
              type="email"
              id="email"
              v-model="email"
              required
              class="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-green-500"
              placeholder="Введите ваш email"
          />
        </div>

        <!-- Password input -->
        <div>
          <label for="password" class="block text-sm font-medium text-gray-700">Пароль</label>
          <input
              type="password"
              id="password"
              v-model="password"
              required
              class="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-green-500"
              placeholder="Введите ваш пароль"
          />
        </div>

        <!-- Login button -->
        <div>
          <button
              type="submit"
              class="w-full px-4 py-2 bg-green-600 text-white font-semibold rounded-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500"
          >
            Войти
          </button>
        </div>
      </form>
    </div>

    <!-- Modal for error -->
    <transition name="fade">
      <div
          v-if="showErrorModal"
          class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50"
      >
        <div class="bg-white p-6 rounded-md shadow-lg max-w-sm w-full">
          <h2 class="text-lg font-semibold text-red-600">Ошибка входа</h2>
          <p class="text-sm text-gray-600">Неверный email или пароль.</p>
          <div class="mt-4 flex justify-end">
            <button @click="showErrorModal = false" class="text-green-600 hover:underline">
              Закрыть
            </button>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup>
import {ref} from "vue";
import {useRouter} from "vue-router";

const email = ref("")
const password = ref("")
const showErrorModal = ref(false)

const router = useRouter()

async function handleLogin() {
  let response = await fetch("http://localhost:8080/api/auth", {
    method: 'POST',

    body: JSON.stringify({
      "login": email.value,
      "password": password.value
    })
  })

  if (response.status !== 200) {
    showErrorModal.value = true
    return
  }

  let result = await response.json()
  localStorage.setItem("jwt", result.token)
  let responseRoom = await fetch('http://localhost:8080/api/user/room_token?room=myroom', {
    headers: {
      Authorization: "Bearer " + result.token
    },
  })
  let resultRoom = await responseRoom.json()
  localStorage.setItem("room_token", resultRoom.token)
  await router.push('/lesson')
}
</script>

<style scoped>

</style>