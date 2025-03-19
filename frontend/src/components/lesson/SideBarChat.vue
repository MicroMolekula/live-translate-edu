<template>
  <div v-if="props.open" class="w-[27rem] bg-white border-l border-gray-700 flex flex-col border-chat m-3 text-black">
    <div class="p-4 border-b-2">Чат</div>
    <ScrollArea class="p-4" style="height: calc(100% - 100px)">
      <div class="flex flex-col">
        <Message v-for="(message, index) in chatMessages" :key="index" :message="message" :is-own="message.user === 'Вы'"></Message>
      </div>
    </ScrollArea>
    <div class="flex w-full items-center gap-1.5 p-4 border-t-2">
      <Input id="message" v-model="newMessage" type="text" placeholder="Сообщение..." />
      <Button @click="sendMessage">
        <Vicon name="bi-send-fill" class="icon"></Vicon>
      </Button>
    </div>
  </div>
</template>

<script setup>
import {Button} from "@/components/ui/button/index.js";
import {Input} from "@/components/ui/input/index.js";
import { OhVueIcon as Vicon, addIcons } from "oh-vue-icons"
import { BiSendFill } from "oh-vue-icons/icons";
import {ScrollArea} from "@/components/ui/scroll-area/index.js";
import {onMounted, ref} from "vue";
import Message from "@/components/lesson/Message.vue";
import {roomStore, userStore} from "@/stores/stores.js";
addIcons(BiSendFill)

const props = defineProps({
  open: {
    type: Boolean,
    required: true
  }
})

const chatMessages = ref([]);
const userData = userStore()
const roomData = roomStore()

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

let wsClient = new WebSocket('ws://localhost:8080/api/chat/connect/' + roomData.mapRoomToken.roomName, ['auth', userData.token])

const newMessage = ref('')

wsClient.onopen = function () {
  console.log("Успешное подключение")
}
wsClient.onmessage = function(event) {
  let messageResponse = JSON.parse(event.data)
  let message = {
    user: `${messageResponse.user.name} ${messageResponse.user.surname}`,
    text: {
      ru: messageResponse.translate_content,
      en: messageResponse.content
    }
  }
  if (message.user === `${userData.user.name} ${userData.user.surname}`) {
    message.user = 'Вы'
  }
  chatMessages.value.push(message)
}

function sendMessage() {
  wsClient.send(JSON.stringify({
    content: newMessage.value
  }))
  newMessage.value = ''
}
</script>

<style scoped>
.border-chat {
  border-radius: 8px;
}
.icon {
  scale: 1.5;
}
</style>