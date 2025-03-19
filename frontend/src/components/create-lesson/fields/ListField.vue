<script setup>
import {FormControl, FormItem, FormLabel, FormField, FormDescription, FormMessage} from "@/components/ui/form/index.js";
import {SelectItem, Select, SelectTrigger, SelectValue, SelectContent, SelectGroup} from "@/components/ui/select/index.js";
import {AutoFormLabel} from "@/components/ui/auto-form/index.js";

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
console.log(props.config)
</script>

<template>
  <FormField v-slot="slotProps" :name="fieldName">
    <FormItem v-bind="$attrs">
      <AutoFormLabel v-if="!config?.hideLabel" :required="required">
        {{ config?.label }}
      </AutoFormLabel>
      <FormControl>
        <slot v-bind="slotProps">
          <Select v-bind="{ ...slotProps.componentField }">

              <SelectTrigger>
                <SelectValue :placeholder="config?.placeholder" />
              </SelectTrigger>

            <SelectContent>
              <SelectGroup>
                <SelectItem v-for="item in config?.options.value" :key="item.title" :value="item.id">
                    {{ item.title }}
                </SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
        </slot>
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