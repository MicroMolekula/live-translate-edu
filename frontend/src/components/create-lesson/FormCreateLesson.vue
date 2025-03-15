<script setup>
import {z} from 'zod'
import {AutoForm} from "@/components/ui/auto-form/index.js";

const formSchema = z.object({
  theme: z
      .string({
        required_error: "Нужно указать тему занятия"
      })
      .describe("Тема занятия"),
  group_id: z
      .string({
        required_error: "Группа"
      })
      .describe("Нужно выбрать группу"),
  languages_codes: z
      .string({
        required_error: "Нужно выбать язык"
      })
      .describe("Язык"),
  date_start: z
      .date({
        required_error: "Выберите дату занятия"
      })
      .describe("Дата и время занятия"),
  number_room: z
      .string({
        required_error: "Укажите аудиторию"
      }).describe("Аудитория занятия").optional()
})
</script>

<template>
  <AutoForm
      :schema="formSchema" class="m-5 w-2/3 space-y-6"
      :field-config = "{
        languages_codes: {
          description: 'Введите язык на который будет переводится занятие'
        }
      }"
  />
</template>

<style scoped>

</style>