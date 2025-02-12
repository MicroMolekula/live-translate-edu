import { createRouter, createWebHistory } from 'vue-router';
import LiveKitComponent from "@/components/LiveKitComponent.vue";
import AuthComponent from "@/components/AuthComponent.vue";
import MeComponent from "@/components/MeComponent.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/lesson',
            component: LiveKitComponent,
            name: 'lesson'
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