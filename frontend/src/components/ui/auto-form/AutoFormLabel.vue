<script setup>
import { FormLabel } from '@/components/ui/form';

defineProps({
  required: { type: Boolean, required: false },
});
</script>

<template>
  <FormLabel>
    <slot />
    <span v-if="required" class="text-red-500 dark:text-red-900"> *</span>
  </FormLabel>
</template>
