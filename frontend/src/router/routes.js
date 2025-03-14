
import NotFoundPage from "@/pages/NotFoundPage.vue";
import LoginPage from "@/pages/LoginPage.vue";
import LessonPage from "@/pages/LessonPage.vue";
import HomePage from "@/pages/HomePage.vue";

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
    name: 'Home'
}, {
    path: '/:pathMatch(.*)*', component: NotFoundPage
},]