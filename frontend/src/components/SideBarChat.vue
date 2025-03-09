<template>
  <div v-if="props.open" class="w-96 bg-white border-l border-gray-700 flex flex-col border-chat m-3 text-black">

  </div>
</template>

<script setup>
const props = defineProps({
  open: {
    type: Boolean,
    required: true
  }
})
</script>

<style scoped>
.border-chat {
  border-radius: 8px;
}
</style>