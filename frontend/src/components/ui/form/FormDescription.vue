<script setup>
import { cn } from '@/lib/utils';
import { useFormField } from './useFormField';

const props = defineProps({
  class: { type: null, required: false },
});

const { formDescriptionId } = useFormField();
</script>

<template>
  <p
    :id="formDescriptionId"
    :class="cn('text-sm text-zinc-500 dark:text-zinc-400', props.class)"
  >
    <slot />
  </p>
</template>
