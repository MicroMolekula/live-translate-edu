
import NotFoundPage from "@/pages/NotFoundPage.vue";
import LoginPage from "@/pages/LoginPage.vue";
import LessonPage from "@/pages/LessonPage.vue";
import HomePage from "@/pages/HomePage.vue";
import AvailableLessonsPage from "@/pages/AvailableLessonsPage.vue";
import CreateLessonPage from "@/pages/CreateLessonPage.vue";
import {roomStore} from "@/stores/stores.js";

async function getRoomData(to) {
    const roomData = roomStore()
    roomData.mapRoomToken.roomName = to.params.roomName
    await roomData.getRoomToken()
}

export const routes = [{
    path: '/auth',
    component: LoginPage,
    name: 'Login'
}, {
    path: '/lesson/:roomName',
    component: LessonPage,
    name: 'Lesson',
    beforeEnter: [getRoomData]
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