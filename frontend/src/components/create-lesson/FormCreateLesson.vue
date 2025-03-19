<script setup>
import {z} from 'zod'
import {AutoForm} from "@/components/ui/auto-form/index.js";
import DateStartField from "@/components/create-lesson/fields/DateStartField.vue";
import {Button} from "@/components/ui/button/index.js";
import ListField from "@/components/create-lesson/fields/ListField.vue";
import {onBeforeMount, onMounted, ref} from "vue";
import {formStore, userStore} from "@/stores/stores.js";
import {createLesson} from "@/lib/request/lesson.js";
import {useToast} from "@/components/ui/toast/index.js";
import {format} from "date-fns";
import {ru} from "date-fns/locale";

const { toast } = useToast()

const languages = ref([])
const groups = ref([])

const userData = userStore()
const formInfo = formStore()

onBeforeMount(async () => {
  await formInfo.getFormData()
  languages.value = formInfo.formData.languages
  groups.value = formInfo.formData.groups
})

const formSchema = z.object({
  theme: z
      .string({
        required_error: "Нужно указать тему занятия"
      })
      .describe("Тема занятия"),
  group_id: z.coerce
      .number({
        required_error: "Группа"
      })
      .describe("Нужно выбрать группу"),
  languages_codes: z
      .string({
        required_error: "Нужно выбать язык"
      }),
  date_start: z
      .string({
        required_error: "Выберите дату и время занятия"
      }),
  number_room: z
      .string({
        required_error: "Укажите аудиторию"
      }).describe("Аудитория занятия").optional()
})


const fieldConfig = {
  languages_codes: {
    description: 'Введите язык на который будет переводится занятие',
    component: ListField,
    label: 'Язык',
    placeholder: 'Выберите язык',
    options: languages
  },
  group_id: {
    component: ListField,
    label: 'Группа',
    placeholder: 'Выберите группу',
    options: groups
  },
  date_start: {
    component: DateStartField
  },
}

function onSubmit(values) {
  values.languages_codes = [values.languages_codes, 'ru']
  createLesson(userData.token, values)
      .then((result) => {
        if (result.status !== 200) {
          toast({
            title: 'Ошибка создания занятия',
            description: JSON.stringify(result.data),
            variant: 'destructive',
          });
        } else {
          toast({
            title: 'Занятия успешно создано',
            description: format(new Date(), 'PPP HH:mm:ss', {locale: ru}),
          });
        }
      })
      .catch((e) => {
        toast({
          title: 'Ошибка создания занятия',
          description: e.toString(),
          variant: 'destructive',
        });
      })
}
</script>

<template>
  <AutoForm
      :schema="formSchema" class="m-5 w-2/3 space-y-6"
      :field-config = "fieldConfig"
      @submit="onSubmit"
  >
    <Button type="submit">
      Создать
    </Button>
  </AutoForm>
</template>

<style scoped>

</style>