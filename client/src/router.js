import { createWebHistory, createRouter } from "vue-router";

import HomeView from "./views/HomeView.vue";
import AboutView from "./views/AboutView.vue";
import HelpView from "./views/HelpView.vue";

const routes = [
    { path: "/", component: HomeView },
    { path: "/about", component: AboutView },
    { path: "/help", component: HelpView },
    { path: "/auth/login", component: import("./views/LoginView.vue") },
];

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
});

export default router;
