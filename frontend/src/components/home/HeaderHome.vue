<script setup>
import {SidebarTrigger} from "@/components/ui/sidebar/index.js";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbList,
  BreadcrumbPage
} from "@/components/ui/breadcrumb/index.js";
import {Separator} from "@/components/ui/separator/index.js";

const props = defineProps({
  pageName: {
    type: String,
    required: true
  }
})
</script>

<template>
  <header class="flex h-16 shrink-0 items-center gap-2">
    <div class="flex items-center gap-2 px-4">
      <SidebarTrigger class="-ml-1" />
      <Separator orientation="vertical" class="mr-2 h-4" />
      <Breadcrumb>
        <BreadcrumbList>
          <BreadcrumbItem>
            <BreadcrumbPage>{{props.pageName}}</BreadcrumbPage>
          </BreadcrumbItem>
        </BreadcrumbList>
      </Breadcrumb>
    </div>
  </header>
</template>

<style scoped>

</style>