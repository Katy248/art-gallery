import { createWebHistory, createRouter } from "vue-router";

import HomeView from "./views/HomeView.vue";
import AboutView from "./views/AboutView.vue";
import HelpView from "./views/HelpView.vue";

const routes = [
    { path: "/", component: HomeView },
    { path: "/about", component: AboutView },
    { path: "/help", component: HelpView },
    { path: "/auth/login", component: () => import("./views/auth/LoginView.vue") },
    { path: "/auth/register", component: () => import("./views/auth/RegisterView.vue") },
    { path: "/auth/logout", component: () => import("./views/auth/LogoutView.vue") },
    { path: "/profile", component: () => import("./views/ProfileView.vue") },
    { path: "/post/create", component: () => import("./views/CreatePostView.vue") },
    { path: "/user/:id", component: () => import("./views/user/UserView.vue"), sensitive: true },
];

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
});

export default router;
