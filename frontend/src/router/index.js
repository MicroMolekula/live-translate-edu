import { createRouter, createWebHistory } from 'vue-router';
import LiveKitComponent from "@/components/LiveKitComponent.vue";
import AuthComponent from "@/components/AuthComponent.vue";

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
        }
    ]
})

export default router