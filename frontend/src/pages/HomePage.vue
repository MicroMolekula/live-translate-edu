<script setup>
import {
  Sidebar,
  SidebarContent,
  SidebarHeader,
  SidebarInset,
  SidebarProvider,
  SidebarTrigger
} from "@/components/ui/sidebar/index.js";
import {ref} from "vue";
import AppSidebar from "@/components/AppSidebar.vue";

const open = ref(false)
</script>

<template>
  <SidebarProvider style="--sidebar-width: 20rem; --sidebar-width-mobile: 20rem;" v-model:open="open" >
      <Sidebar variant="inset">
        <SidebarHeader>Hello</SidebarHeader>
        <SidebarContent>
        </SidebarContent>
      </Sidebar>
      <SidebarInset>
        <SidebarTrigger/>
        <slot />
      </SidebarInset>
  </SidebarProvider>
</template>

<style>
</style>