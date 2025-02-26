<template>
  <div class="flex">
    <!-- Sidebar -->
    <aside class="w-64 bg-green-600 text-white h-screen p-6">
      <div class="text-2xl font-semibold mb-8">LiveTranslateEdu</div>

      <!-- User info in Sidebar -->
      <div class="mb-8">
        <div class="text-lg font-medium">{{ "Иван Иванов" }}</div>
      </div>

      <!-- Navigation links in Sidebar -->
      <nav class="space-y-6">
        <ul>
          <li>
            <a
                href="/lessons"
                class="text-white hover:text-green-200 text-lg font-semibold"
            >
              Доступные занятия
            </a>
          </li>
          <li>
            <a
                href="#"
                @click="currentPage = 'history'"
                class="text-white hover:text-green-200 text-lg font-semibold"
            >
              История занятий
            </a>
          </li>
          <li>
            <a
                href="/lessons/create"
                @click="currentPage = 'create'"
                class="text-white hover:text-green-200 text-lg font-semibold"
            >
              Создать занятие
            </a>
          </li>
        </ul>
      </nav>
    </aside>

    <!-- Main Content -->
    <main class="flex-1 bg-gray-50 p-6">
      <!-- Dynamic Content for Create Lesson -->
      <div v-if="currentPage === 'create'">
        <h2 class="text-2xl font-semibold text-green-600 mb-4">Создать занятие</h2>

        <!-- Lesson Creation Form -->
        <form @submit.prevent="createLesson" class="space-y-6">
          <!-- Topic -->
          <div>
            <label for="topic" class="block text-sm font-medium text-gray-700">Тема занятия</label>
            <input
                v-model="newLesson.topic"
                type="text"
                id="topic"
                required
                class="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-green-500"
                placeholder="Введите тему занятия"
            />
          </div>

          <!-- Room -->
          <div>
            <label for="audience" class="block text-sm font-medium text-gray-700">Аудитория</label>
            <input
                v-model="newLesson.audience"
                type="text"
                id="audience"
                required
                class="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-green-500"
                placeholder="Введите аудиторию"
            />
          </div>

          <!-- Student Group Selection -->
          <div>
            <label for="group" class="block text-sm font-medium text-gray-700">Группа студентов</label>
            <select
                v-model="newLesson.group"
                id="group"
                required
                class="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-green-500"
            >
              <option value="" disabled>Выберите группу</option>
              <option value="group1">Группа 1</option>
              <option value="group2">Группа 2</option>
              <option value="group3">Группа 3</option>
            </select>
          </div>

          <!-- Start Date and Time -->
          <div>
            <label for="startTime" class="block text-sm font-medium text-gray-700">Дата и время начала</label>
            <input
                v-model="newLesson.startTime"
                type="datetime-local"
                id="startTime"
                required
                class="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-green-500"
            />
          </div>

          <!-- Submit Button -->
          <div>
            <button
                type="submit"
                class="w-full py-2 bg-green-600 text-white rounded-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500"
            >
              Создать занятие
            </button>
          </div>
        </form>
      </div>

      <!-- Other Pages -->
      <div v-if="currentPage === 'lessons'">
        <h2 class="text-2xl font-semibold text-green-600 mb-4">Доступные занятия</h2>
        <p>Здесь будет отображаться список доступных занятий.</p>
      </div>

      <div v-if="currentPage === 'history'">
        <h2 class="text-2xl font-semibold text-green-600 mb-4">История занятий</h2>
        <p>Здесь будет отображаться история ваших занятий.</p>
      </div>
    </main>
  </div>
</template>

<script>
export default {
  data() {
    return {
      currentPage: 'create', // Переключение между страницами
      newLesson: {
        topic: '',
        audience: '',
        group: '', // Группа студентов
        startTime: '',
      },
    };
  },
  methods: {
    createLesson() {
      // Логика отправки данных на сервер (например, axios)
      console.log('Новый урок создан:', this.newLesson);

      // Очистить форму после отправки
      this.newLesson = {
        topic: '',
        audience: '',
        group: '',
        startTime: '',
      };

      // Сообщение об успешном создании
      alert('Занятие успешно создано!');
    },
  },
};
</script>

<style scoped>
/* Дополнительные стили для формы и контента */
</style>