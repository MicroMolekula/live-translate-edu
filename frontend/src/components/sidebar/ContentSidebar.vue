<script setup>
import { Calendar, Home, Inbox } from "lucide-vue-next"
import {
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu, SidebarMenuButton,
  SidebarMenuItem
} from "@/components/ui/sidebar/index.js";
import {useRouter} from "vue-router";
import ButtonGoUrl from "@/components/sidebar/ButtonGoUrl.vue";


const props = defineProps({
  items: {
    type: Array,
    required: true
  }
})

const router = useRouter()

function goUrl(url) {

}
</script>

<template>
  <SidebarContent>
    <SidebarGroup>
      <SidebarGroupLabel>Занятия</SidebarGroupLabel>
      <SidebarGroupContent>
        <SidebarMenu>
          <SidebarMenuItem v-for="item in props.items" :key="item.title">
            <ButtonGoUrl :item="item" />
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarGroupContent>
    </SidebarGroup>
  </SidebarContent>
</template>

<style scoped>

</style>