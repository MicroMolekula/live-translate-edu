<template>
  <div class="message" :class="{'message-own': isOwn}">
    <div class="message-user">{{ message.user }}</div>
    <div class="message-text">
      <div class="message-text-ru">{{ message.text.ru }}</div>
      <div class="message-text-en">{{ message.text.en }}</div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  message: {
    type: Object,
    required: true,
  },
  isOwn: {
    type: Boolean,
    default: false,
  },
});
</script>

<style scoped>
.message {
  display: flex;
  flex-direction: column;
  margin-bottom: 1rem;
  padding: 0.5rem;
  border-radius: 8px;
  background-color: #f3f4f6;
  max-width: 80%;
  width: 13rem;
}

.message-own {
  align-self: flex-end;
  background-color: #3b82f6;
  color: white;
}

.message-user {
  font-weight: bold;
  margin-bottom: 0.25rem;
  font-size: 0.875rem;
}

.message-text {
  font-size: 0.875rem;
}

.message-text-ru {
  margin-bottom: 0.25rem;
}

.message-text-en {
  font-style: italic;
  color: #6b7280;
}

.message-own .message-text-en {
  color: #e5e7eb;
}
</style>