import { createRouter, createWebHistory } from 'vue-router';
import LiveKitComponent from "@/components/LiveKitComponent.vue";
import AuthComponent from "@/components/AuthComponent.vue";
import MeComponent from "@/components/MeComponent.vue";
import Lessons from "@/components/Lessons.vue";
import CreateLesson from "@/components/CreateLesson.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/lesson',
            component: LiveKitComponent,
            name: 'lesson'
        },
        {
            path: '/lessons',
            component: Lessons,
            name: 'lessons'
        },
        {
            path: '/lessons/create',
            component: CreateLesson,
            name: 'create-lessons'
        },
        {
            path: '/auth',
            component: AuthComponent,
            name: 'auth'
        },
        {
            path: '/me',
            component: MeComponent,
            name: 'me'
        }
    ]
})

export default router