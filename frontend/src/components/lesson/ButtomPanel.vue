<script setup>
import {Button} from "@/components/ui/button/index.js";
import { OhVueIcon as Vicon, addIcons } from "oh-vue-icons"
import {
  FaMicrophone,
  FaMicrophoneSlash,
  BiChatLeftDots,
  BiChatLeftDotsFill,
  FaUsers,
  BiTelephoneFill
} from "oh-vue-icons/icons"
import {ref} from "vue";
import {userStore} from "@/stores/stores.js";
addIcons(FaMicrophone)
addIcons(FaUsers)
addIcons(FaMicrophoneSlash)
addIcons(BiChatLeftDots)
addIcons(BiChatLeftDotsFill)
addIcons(BiTelephoneFill)

const emit = defineEmits(['chat', 'mute', 'leave'])

const userData = userStore()

const micOn = ref(false)

function openChat() {
  emit('chat', true)
}

function handleMicro() {
  micOn.value = !micOn.value
  emit('mute', micOn.value)
}

function handleLeave() {
  emit('leave', true)
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
      <Button @click="openChat" variant="secondary">
        <Vicon name="bi-chat-left-dots-fill" class="icon"/>
        Чат
      </Button>
      <Button variant="secondary">
        <Vicon name="fa-users" class="icon"/>
        Участники
      </Button>
    </div>
    <!-- Вспомогательные кнопки -->
    <div class="flex gap-2">
      <Button v-if="userData.user.role === 'teacher'" variant="secondary" @click="handleMicro">
        <Vicon v-if="micOn" name="fa-microphone" class="icon"/>
        <Vicon v-else name="fa-microphone-slash" fill="red" class="icon"/>
      </Button>
      <Button variant="destructive" @click="handleLeave">
        <Vicon name="bi-telephone-fill" class="icon" fill="white"></Vicon>
      </Button>
    </div>
  </div>
</div>
</template>

<style scoped>
.icon {
  scale: 1.5;
}
</style>