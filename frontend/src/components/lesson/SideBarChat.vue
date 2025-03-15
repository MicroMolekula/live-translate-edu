<template>
  <div v-if="props.open" class="w-96 bg-white border-l border-gray-700 flex flex-col border-chat m-3 text-black">
    <div class="p-4 border-b-2">Чат</div>
    <ScrollArea class="p-4" style="height: calc(100% - 100px)">
      <div v-for="n in 100" :key="n">
        <div>Message {{n}}</div>
      </div>
    </ScrollArea>
    <div class="flex w-full max-w-sm items-center gap-1.5 p-4 border-t-2">
      <Input id="message" type="text" placeholder="Сообщение..." />
      <Button type="submit">
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

addIcons(BiSendFill)

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
.icon {
  scale: 1.5;
}
</style>