<template>
  <div class="flex">
    <!-- Sidebar -->
    <aside class="w-64 bg-green-600 text-white h-screen p-6">
      <div class="text-2xl font-semibold mb-8">LiveTranslateEdu</div>

      <!-- User info in Sidebar -->
      <div class="mb-8">
        <div class="text-lg font-medium">{{ fullName }}</div>
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
      <!-- Dynamic Content -->
      <div v-if="currentPage === 'classes'">
        <h2 class="text-2xl font-semibold text-green-600 mb-4">Доступные Занятия</h2>
        <ul class="space-y-6">
          <li v-for="(classItem, index) in classes" :key="index" class="bg-white p-6 rounded-lg shadow-lg hover:shadow-xl transition-shadow duration-300">
            <h3 class="text-xl font-semibold text-gray-800">{{ classItem.topic }}</h3>
            <p class="text-gray-600 mt-2 mb-4">{{ classItem.description }}</p>

            <!-- Display additional details -->
            <div class="text-sm text-gray-600">
              <p><strong>Аудитория:</strong> {{ classItem.audience }}</p>
              <p><strong>Преподаватель:</strong> {{ classItem.teacher }}</p>
              <p><strong>Дата и время начала:</strong> {{ formatDate(classItem.startTime) }}</p>
            </div>

            <!-- Conditional rendering for join button or message -->
            <div v-if="isClassStarted(classItem.startTime)">
              <button
                  @click="joinClass(classItem)"
                  class="w-full py-2 bg-green-600 text-white rounded-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 mt-4"
              >
                Присоединиться
              </button>
            </div>
            <div v-else>
              <p class="text-red-600 mt-4">Занятие еще не началось</p>
            </div>
          </li>
        </ul>
      </div>
    </main>
  </div>
</template>

<script>
export default {
  data() {
    return {
      currentPage: 'classes', // Переключение между страницами
      user: {
        firstName: "Иван",
        lastName: "Иванов",
      },
      classes: [
        {
          topic: "Введение в искусственный интеллект",
          description: "",
          audience: "255",
          teacher: "Александров Алексей Викторович",
          startTime: "2024-12-05T10:00:00", // ISO формат
        },
      ],
    };
  },
  computed: {
    fullName() {
      return `${this.user.firstName} ${this.user.lastName}`;
    },
  },
  methods: {
    joinClass(classItem) {
      alert(`Вы присоединились к занятию: ${classItem.topic}`);
    },
    formatDate(dateString) {
      const date = new Date(dateString);
      return date.toLocaleString(); // Форматирует дату и время в локальный формат
    },
    // Метод для проверки, началось ли занятие
    isClassStarted(startTime) {
      const currentDate = new Date();
      const classDate = new Date(startTime);
      return currentDate >= classDate; // Возвращает true, если текущее время больше или равно времени начала занятия
    },
  },
};
</script>

<style scoped>
/* Дополнительные стили для боковой панели и контента */
</style>