
import NotFoundPage from "@/pages/NotFoundPage.vue";
import LoginPage from "@/pages/LoginPage.vue";
import LessonPage from "@/pages/LessonPage.vue";

export const routes = [{
    path: '/auth',
    component: LoginPage,
    name: 'Login'
}, {
    path: '/lesson/:roomName',
    component: LessonPage,
    name: 'Lesson'
}, {
    path: '/:pathMatch(.*)*', component: NotFoundPage
},]