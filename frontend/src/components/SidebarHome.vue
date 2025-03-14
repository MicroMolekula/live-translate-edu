<script setup>
import SidebarHome from "@/components/SidebarHome.vue";
import {SidebarProvider, SidebarTrigger, SidebarInset, Sidebar, SidebarContent} from "@/components/ui/sidebar/index.js";
import {ref} from "vue";
import AppSidebar from "@/components/AppSidebar.vue";

const open = ref(false)
</script>

<template>
  <SidebarProvider v-model:open="open">
    <Sidebar>
      <SidebarContent>

      </SidebarContent>
    </Sidebar>
    <SidebarInset>
      <SidebarTrigger/>
      <slot />
    </SidebarInset>
  </SidebarProvider>
</template>

<style scoped>

</style>