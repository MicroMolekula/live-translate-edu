import './assets/main.css'
import './assets/tailwind.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from "@/router/index.js";
import {createPinia} from "pinia";
import { OhVueIcon, addIcons } from "oh-vue-icons";
import * as FaIcons from "oh-vue-icons/icons/fa";

const pinia = createPinia()

const Fa = Object.values({ ...FaIcons });
addIcons(...Fa);

const app = createApp(App)
    .use(router)
    .use(pinia)

app.component("v-icon", OhVueIcon)
app.mount('#app')
