<script setup>
import { Calendar, Home, Inbox } from "lucide-vue-next"
import {
  Sidebar,
} from "@/components/ui/sidebar/index.js";
import HeaderSidebar from "@/components/sidebar/HeaderSidebar.vue";
import ContentSidebar from "@/components/sidebar/ContentSidebar.vue";
import FooterSidebar from "@/components/sidebar/FooterSidebar.vue";
import {userStore} from "@/stores/stores.js";

let userData = userStore()

const props = defineProps({
  pages: {
    type: Array,
    required: true
  }
})
</script>

<template>
  <Sidebar variant="inset">
    <HeaderSidebar />
    <ContentSidebar :items="props.pages"/>
    <FooterSidebar :user="userData.user"/>
  </Sidebar>
</template>
