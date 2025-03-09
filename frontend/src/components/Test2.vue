<template>
  <div class="flex h-screen bg-gray-900 text-white">
    <!-- Main Video Grid -->
    <div class="flex-1 grid grid-cols-3 gap-4 p-4">
      <div v-for="n in 9" :key="n" class="bg-gray-800 rounded-lg h-48 flex items-center justify-center">
        <span class="text-gray-400">User {{ n }}</span>
      </div>
    </div>

    <!-- Chat Sidebar (Only takes space when visible) -->
    <Transition name="slide">
      <div v-if="showChat" class="w-80 bg-gray-800 border-l border-gray-700 flex flex-col">
        <div class="p-4 border-b border-gray-700 font-semibold flex justify-between items-center">
          <span>Chat</span>
          <button @click="toggleChat" class="text-gray-400 hover:text-white">✖</button>
        </div>
        <div class="flex-1 overflow-y-auto p-4 space-y-2">
          <div v-for="msg in messages" :key="msg.id" class="bg-gray-700 p-2 rounded-lg">
            <span class="text-sm text-gray-300">{{ msg.user }}:</span>
            <p>{{ msg.text }}</p>
          </div>
        </div>
        <div class="p-4 border-t border-gray-700">
          <input type="text" placeholder="Type a message..." class="w-full p-2 bg-gray-900 rounded-lg border border-gray-700">
        </div>
      </div>
    </Transition>
  </div>

  <!-- Bottom Control Bar -->
  <div class="fixed bottom-0 left-0 right-0 bg-gray-800 p-4 flex justify-center space-x-6 border-t border-gray-700">
    <button class="bg-gray-700 p-3 rounded-full hover:bg-gray-600">
      🎤
    </button>
    <button class="bg-gray-700 p-3 rounded-full hover:bg-gray-600">
      📷
    </button>
    <button class="bg-red-600 p-3 rounded-full hover:bg-red-500">
      ⏹️
    </button>
    <button @click="toggleChat" class="bg-gray-700 p-3 rounded-full hover:bg-gray-600">
      💬
    </button>
  </div>
</template>

<script setup>
import { ref } from 'vue';
const messages = ref([
  { id: 1, user: 'Alice', text: 'Hello everyone!' },
  { id: 2, user: 'Bob', text: 'Hey Alice!' }
]);
const showChat = ref(true);
const toggleChat = () => {
  showChat.value = !showChat.value;
};
</script>

<style>
.slide-enter-active, .slide-leave-active {
  transition: transform 0.3s ease-in-out, opacity 0.3s ease-in-out;
}
.slide-enter-from {
  transform: translateX(100%);
  opacity: 0;
}
.slide-leave-to {
  transform: translateX(100%);
  opacity: 0;
}
</style>
