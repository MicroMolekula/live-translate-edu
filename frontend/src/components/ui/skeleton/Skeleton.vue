<script setup>
import { cn } from '@/lib/utils';

const props = defineProps({
  class: { type: null, required: false },
});
</script>

<template>
  <div :class="cn('animate-pulse rounded-md bg-zinc-100 dark:bg-zinc-800', props.class)" />
</template>
