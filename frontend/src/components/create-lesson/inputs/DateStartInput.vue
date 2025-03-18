<script setup>
import {datetimeToBString, datetimeToString} from "@/lib/datetime/index.js";
import {Input} from "@/components/ui/input/index.js";
import {Popover, PopoverContent, PopoverTrigger} from "@/components/ui/popover/index.js";
import {Clock} from "lucide-vue-next";
import {Button} from "@/components/ui/button/index.js";
import {Calendar} from "@/components/ui/calendar/index.js";
import { Calendar as CalendarIcon } from 'lucide-vue-next';
import {computed, ref, watch} from "vue";
import {getLocalTimeZone} from "@internationalized/date";
import {useVModel} from "@vueuse/core";

const props = defineProps({
  defaultValue: { type: [String, Number], required: false },
  modelValue: { type: [String, Number], required: false },
  class: { type: null, required: false },
});

const emits = defineEmits(['update:modelValue']);

const modelValue = useVModel(props, 'modelValue', emits, {
  passive: true,
  defaultValue: props.defaultValue,
});

const selectedDate = ref(undefined)
const selectedTime = ref({
  hours: '00',
  minutes: '00'
});

const selectedDateTime = computed(() => {
  if (!selectedDate.value || !selectedTime.value) return null;
  const [hours, minutes] = [selectedTime.value.hours, selectedTime.value.minutes];
  const date = selectedDate.value.toDate(getLocalTimeZone());
  date.setHours(parseInt(hours, 10));
  date.setMinutes(parseInt(minutes, 10));
  return date;
});

function handleInput(event, value) {
  if (event.key >= "0" && event.key <= "9") {
    value = value[value.length - 1] + event.key
  } else if (event.key === "Backspace") {
    value = '000'
  }
  return value
}

function handleHours(event) {
  selectedTime.value.hours = handleInput(event, selectedTime.value.hours)
  if (parseInt(selectedTime.value.hours, 10) > 23) {
    selectedTime.value.hours = '00'
  }
}

function handleMinutes(event) {
  selectedTime.value.minutes = handleInput(event, selectedTime.value.minutes)
  if (parseInt(selectedTime.value.minutes, 10) > 60) {
    selectedTime.value.minutes = '00'
  }
}

watch(selectedDateTime, (date) => {
  modelValue.value = datetimeToString(date); // Обновляем modelValue
  if (props.componentField) {
    props.componentField.onUpdate(datetimeToString(date)); // Передаем значение в AutoForm
  }
});
</script>

<template>
  <div class="datetime-picker">
    <Popover>
      <PopoverTrigger asChild>
        <Button variant="outline" class="w-full justify-start text-left">
          <CalendarIcon class="mr-2 h-4 w-4" />
          {{
            selectedDateTime
                ? datetimeToBString(selectedDateTime)
                : "Выберите дату и время"

          }}
        </Button>
      </PopoverTrigger>
      <PopoverContent class="w-auto p-2">
        <div class="flex flex-col" >
          <Calendar
              @update:model-value="(v) => {
                        if (v) {
                          selectedDate = v
                        }
                        else {
                          selectedDate = undefined
                        }
                      }"
              :initial-focus="true"
              class="border-b-2"
              locale="ru"
          />
          <div class="flex flex-row mt-2 gap-2 place-items-center">
            <Input
                type="text"
                v-model="selectedTime.hours"
                class="w-[48px] text-center no-caret text-base"
                maxlength="2"
                minlength="2"
                @keydown="handleHours"
            />
            :
            <Input
                type="text"
                v-model="selectedTime.minutes"
                class="w-[48px] text-center no-caret text-base"
                maxlength="2"
                minlength="2"
                @keydown="handleMinutes"
            />
            <Clock />
          </div>
        </div>
      </PopoverContent>
    </Popover>
  </div>
</template>

<style scoped>
.no-caret {
  caret-color: transparent;
}
</style>