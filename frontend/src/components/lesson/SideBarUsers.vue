<script setup>
import {ScrollArea} from "@/components/ui/scroll-area/index.js";

const props = defineProps({
  usersChat: {
    type: Array,
  },
  open: {
    type: Boolean,
    required: true
  }
})

</script>

<template>
  <div v-if="props.open" class="w-[27rem] bg-white border-l border-gray-700 flex flex-col border-chat m-3 text-black">
    <div class="p-4 border-b-2">Участники</div>
    <ScrollArea class="p-4" style="height: calc(100% - 100px)">
      <div class="flex flex-col">
        <div v-for="(user, index) in props.usersChat" :key="index">
          {{ user.name }} {{ user.surname }}
        </div>
      </div>
    </ScrollArea>
  </div>
</template>

<style scoped>

</style>