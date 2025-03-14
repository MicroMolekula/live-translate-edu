<script setup>
import ButtomPanel from "@/components/ButtomPanel.vue";
import SideBarChat from "@/components/SideBarChat.vue";
import {ref} from "vue";
import Subtitle from "@/components/Subtitle.vue";
import {userStore} from "@/stores/userStore.js";

const isOpen = ref(false)
const userData = userStore()

console.log("Токен " + userData.token)
console.log(userData.user)
function openChat(msg) {
  isOpen.value = !isOpen.value
}
</script>

<template>
  <div class="h-screen bg-main">
    <div class="flex flex-row gap-2 bg-main text-white" style="height: calc(100vh - 140px)">
      <div class="flex-1 grid gap-4 p-4 place-items-center">
        <Subtitle></Subtitle>
      </div>
      <SideBarChat :open="isOpen"/>
    </div>
    <ButtomPanel class="h-20" style="min-height: 80px" @chat="openChat"></ButtomPanel>
  </div>
</template>

<style scoped>
.bg-main {
  background-color: rgba(32, 33, 36, 1.00);
  color: white;
}
</style>