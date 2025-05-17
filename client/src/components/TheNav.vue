<script setup>
import { ref } from "vue";
import { RouterLink } from "vue-router";
import { useAuthStore } from "../stores/auth";
import { Menu, MenuButton, MenuItems, MenuItem } from "@headlessui/vue";

let showMenu = ref(false);

const menuClick = (event) => {
    showMenu.value = !showMenu.value;
};
const backdropClick = (event) => {
    showMenu.value = false;
};

const profileUrl = ref("");
const auth = useAuthStore();
if (auth.isAuthenticated) {
    const data = auth.authData;
    profileUrl.value = `/user/${data.id}`;
}
</script>
<template>
    <header class="flex flex-row p-3 sticky top-0 w-full bg-bg-2 justify-between">
        <button class="btn-base text-magenta px-6 py-2 rounded-md z-50 font-bold" @click="menuClick">ArtGallery</button>
        <!-- <button class="btn-base px-6 py-2 rounded-md font-bold">Profile</button> -->
    </header>
    <div v-show="showMenu" class="w-full h-screen fixed top-0 left-0 bg-black/60 backdrop-blur-sm z-10 flex flex-col items-center" @click="backdropClick">
        <nav class="absolute z-50 left-3 top-[60px] flex flex-col min-w-[10rem] rounded-md gap-2">
            <div class="nav rounded-md bg-bg">
                <RouterLink class="nav-item" to="/">
                    <i class="fas fa-home nav-icon"></i>
                    Главная
                </RouterLink>
                <!-- <a class="nav-item" href="/gallery.html" v-if="auth.isAuthenticated">
                    <i class="fas fa-image nav-icon"></i>
                    Моя галерея</a
                > -->
                <RouterLink class="nav-item" to="/post/create" v-if="auth.isAuthenticated">
                    <i class="fas fa-upload nav-icon"></i>
                    Создать публикацию
                </RouterLink>
            </div>
            <div class="nav rounded-md bg-bg">
                <RouterLink class="nav-item" to="/help">
                    <i class="fas fa-question nav-icon"></i>
                    Помощь
                </RouterLink>
                <RouterLink to="/about" class="nav-item">
                    <i class="fas fa-info nav-icon"></i>
                    О нас
                </RouterLink>
            </div>
            <div class="nav rounded-md bg-bg" v-if="auth.isAuthenticated">
                <RouterLink class="nav-item" :to="profileUrl">
                    <i class="fas fa-user-plus nav-icon"></i>
                    Профиль</RouterLink
                >
                <RouterLink to="/auth/logout" class="nav-item text-red">
                    <i class="fas fa-arrow-right-from-bracket nav-icon"></i>
                    Выход
                </RouterLink>
            </div>
            <div class="nav rounded-md bg-bg" v-else>
                <RouterLink class="nav-item" to="/auth/register">
                    <i class="fas fa-user-plus nav-icon"></i>
                    Регистрация</RouterLink
                >
                <RouterLink to="/auth/login" class="nav-item">
                    <i class="fas fa-arrow-right-to-bracket nav-icon"></i>
                    Вход
                </RouterLink>
            </div>
        </nav>
    </div>
</template>

<style scoped>
@import "../style.css";
:root {
    --sidebar-rounded: 0.5rem;
    --sidebar-min-width: 10rem;
}
</style>
