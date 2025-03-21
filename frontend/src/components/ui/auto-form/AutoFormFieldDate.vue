<script setup>
import { cn } from '@/lib/utils';
import { Button } from '@/components/ui/button';
import { Calendar } from '@/components/ui/calendar';
import { ru } from 'date-fns/locale';
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormMessage,
} from '@/components/ui/form';
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover';

import { DateFormatter, getLocalTimeZone } from '@internationalized/date';
import { CalendarIcon } from 'lucide-vue-next';
import AutoFormLabel from './AutoFormLabel.vue';
import { beautifyObjectName, maybeBooleanishToBoolean } from './utils';

const ruLocale = ru;

defineProps({
  fieldName: { type: String, required: true },
  label: { type: String, required: false },
  required: { type: Boolean, required: false },
  config: { type: Object, required: false },
  disabled: { type: Boolean, required: false },
});

const df = new DateFormatter('ru-RU', {
  dateStyle: 'long',
});
</script>

<template>
  <FormField v-slot="slotProps" :name="fieldName">
    <FormItem>
      <AutoFormLabel v-if="!config?.hideLabel" :required="required">
        {{ config?.label || beautifyObjectName(label ?? fieldName) }}
      </AutoFormLabel>
      <FormControl>
        <slot v-bind="slotProps">
          <div>
            <Popover>
              <PopoverTrigger
                as-child
                :disabled="
                  maybeBooleanishToBoolean(config?.inputProps?.disabled) ??
                  disabled
                "
              >
                <Button
                  variant="outline"
                  :class="
                    cn(
                      'w-full justify-start text-left font-normal',
                      !slotProps.componentField.modelValue &&
                        'text-zinc-500 dark:text-zinc-400',
                    )
                  "
                >
                  <CalendarIcon class="mr-2 h-4 w-4" :size="16" />
                  {{
                    slotProps.componentField.modelValue
                      ? df.format(
                          slotProps.componentField.modelValue.toDate(
                            getLocalTimeZone(),
                          ),
                        )
                      : "Выберите дату"
                  }}
                </Button>
              </PopoverTrigger>
              <PopoverContent class="w-auto p-0">
                <Calendar initial-focus v-bind="slotProps.componentField" locale="ru"/>
              </PopoverContent>
            </Popover>
          </div>
        </slot>
      </FormControl>

      <FormDescription v-if="config?.description">
        {{ config.description }}
      </FormDescription>
      <FormMessage />
    </FormItem>
  </FormField>
</template>
