<script setup>
import {FormControl, FormItem, FormField, FormMessage, FormDescription} from "@/components/ui/form/index.js";
import {AutoFormLabel} from "@/components/ui/auto-form/index.js";
import DateStartInput from "@/components/create-lesson/inputs/DateStartInput.vue";

const props = defineProps({
  config: {
    type: Object,
    required: true
  },
  required: {
    type: Boolean,
    required: false,
    default: true,
  },
  fieldName: {
    type: String,
    required: true
  }
})

</script>

<template>
  <FormField v-slot="slotProps" :name="fieldName">
    <FormItem v-bind="$attrs">
      <AutoFormLabel v-if="!config?.hideLabel" :required="required">
        Дата и время начала занятия
      </AutoFormLabel>
      <FormControl>
        <DateStartInput v-bind="slotProps.componentField"></DateStartInput>
      </FormControl>
      <FormDescription v-if="config?.description">
        {{ config.description }}
      </FormDescription>
      <FormMessage />
    </FormItem>
  </FormField>
</template>

<style scoped>

</style>