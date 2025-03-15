<script setup>
import {
  SidebarInset,
  SidebarProvider,
} from "@/components/ui/sidebar/index.js";
import AppSidebar from "@/components/sidebar/AppSidebar.vue";
import HeaderHome from "@/components/home/HeaderHome.vue";
import {onMounted, onUpdated, ref} from "vue";
import {Calendar, Home, Inbox} from "lucide-vue-next";
import {useRoute} from "vue-router";

const open = ref(true)

const items = [
  {
    title: "Доступные занятия",
    url: "/",
    icon: Home,
  },
  {
    title: "История занятий",
    url: "#",
    icon: Inbox,
  },
  {
    title: "Создать занятие",
    url: "/lesson/create",
    icon: Calendar,
  }
];

const titlePage = ref("")
onUpdated(() => {
  const route = useRoute()
  titlePage.value = route.meta.name
})

onMounted(() => {
  const route = useRoute()
  titlePage.value = route.meta.name
})

</script>

<template>
  <SidebarProvider v-model:open="open">
     <AppSidebar :pages="items"/>
      <SidebarInset>
        <HeaderHome :page-name="titlePage"/>
        <RouterView></RouterView>
      </SidebarInset>
  </SidebarProvider>
</template>

<style>
</style>