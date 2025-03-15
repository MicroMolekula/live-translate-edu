
import NotFoundPage from "@/pages/NotFoundPage.vue";
import LoginPage from "@/pages/LoginPage.vue";
import LessonPage from "@/pages/LessonPage.vue";
import HomePage from "@/pages/HomePage.vue";
import AvailableLessonsPage from "@/pages/AvailableLessonsPage.vue";
import CreateLessonPage from "@/pages/CreateLessonPage.vue";

export const routes = [{
    path: '/auth',
    component: LoginPage,
    name: 'Login'
}, {
    path: '/lesson/:roomName',
    component: LessonPage,
    name: 'Lesson'
}, {
    path: '/',
    component: HomePage,
    name: 'Home',
    children: [
    {
        path: '',
        component: AvailableLessonsPage,
        name: 'AvailableLesson',
        meta: {
            name: "Доступные занятия"
        }
    }, {
        path: 'lesson/create',
        component: CreateLessonPage,
        name: 'CreateLesson',
        meta: {
            name: "Создание занятия"
        }
    }
    ]
}, {
    path: '/:pathMatch(.*)*', component: NotFoundPage
},]