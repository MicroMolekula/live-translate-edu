import { createRouter, createWebHistory } from 'vue-router';
import { routes } from '@/router/routes.js'
import { userStore } from "@/stores/userStore.js";
import currentUser from '@/lib/request/currentUser'

const router = createRouter({
    history: createWebHistory(),
    routes: routes
})

export default router

// router.beforeEach(async function(to, from) {
//     const userStoreData = userStore()
//     if (to.name !== 'Login' && userStoreData.token === '') {
//         return {name: 'Login'}
//     }
//     try {
//         userStoreData.user = await currentUser(userStoreData.token)
//     } catch {
//         if (to.name !== 'Login') {
//             return { name: 'Login' }
//         }
//     }
//
// })