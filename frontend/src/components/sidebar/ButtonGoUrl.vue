<script setup>
import {SidebarMenuButton} from "@/components/ui/sidebar/index.js";
import {Button} from "@/components/ui/button/index.js";
import {useRouter} from "vue-router";

const props = defineProps({
  item: {
    type: Object,
    required: true
  }
})

const router = useRouter()

function goUrl() {
  router.push(props.item.url)
}
</script>

<template>
  <SidebarMenuButton asChild>
    <Button variant="link" @click="goUrl">
      <component :is="props.item.icon" />
      <span>{{props.item.title}}</span>
    </Button>
  </SidebarMenuButton>
</template>

<style scoped>

</style>