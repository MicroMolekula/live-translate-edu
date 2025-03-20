<script setup>
import {
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu, SidebarMenuButton,
  SidebarMenuItem
} from "@/components/ui/sidebar/index.js";
import {useRouter} from "vue-router";


const props = defineProps({
  items: {
    type: Array,
    required: true
  }
})

const router = useRouter()
</script>

<template>
  <SidebarContent>
    <SidebarGroup>
      <SidebarGroupLabel>Занятия</SidebarGroupLabel>
      <SidebarGroupContent>
        <SidebarMenu>
          <SidebarMenuItem v-for="item in props.items" :key="item.title">
            <SidebarMenuButton asChild v-if="item.view">
              <RouterLink :to="item.url">
                <component :is="item.icon" />
                <span>{{item.title}}</span>
              </RouterLink>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarGroupContent>
    </SidebarGroup>
  </SidebarContent>
</template>

<style scoped>

</style>