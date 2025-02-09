<script setup>
import { ref } from "vue";
import { RouterLink } from "vue-router";

let showMenu = ref(false);
let authorized = ref(false);

const menuClick = (event) => {
    showMenu.value = !showMenu.value;
};
const backdropClick = (event) => {
    showMenu.value = false;
};
</script>
<template>
    <header class="flex flex-row p-3 sticky top-0 w-full bg-(--bg-2)">
        <button class="btn-base text-magenta px-6 py-2 rounded-md z-50 font-bold" @click="menuClick">ArtGallery</button>
    </header>
    <div v-show="showMenu" class="w-full h-screen fixed top-0 left-0 bg-black/60 z-10 flex flex-col items-center" @click="backdropClick">
        <nav class="absolute z-50 left-3 top-[60px] flex flex-col min-w-[10rem] rounded-md gap-2">
            <div class="nav rounded-md">
                <RouterLink class="nav-item" to="/">
                    <i class="fas fa-home nav-icon"></i>
                    Главная
                </RouterLink>
                <a class="nav-item" href="/gallery.html">
                    <i class="fas fa-image nav-icon"></i>
                    Моя галерея</a
                >
                <a class="nav-item" href="#upload">
                    <i class="fas fa-upload nav-icon"></i>
                    Загрузить
                </a>
                <RouterLink class="nav-item" to="/help">
                    <i class="fas fa-question nav-icon"></i>
                    Помощь
                </RouterLink>
                <RouterLink to="/about" class="nav-item">
                    <i class="fas fa-info nav-icon"></i>
                    О нас
                </RouterLink>
            </div>
            <div class="nav rounded-md" v-if="authorized">
                <a class="nav-item" href="/profile">
                    <i class="fas fa-user-plus nav-icon"></i>
                    Профиль</a
                >
            </div>
            <div class="nav rounded-md" v-else>
                <a class="nav-item" href="/registration.html">
                    <i class="fas fa-user-plus nav-icon"></i>
                    Регистрация</a
                >
                <a class="nav-item" href="/login.html">
                    <i class="fas fa-arrow-right-to-bracket nav-icon"></i>
                    Вход</a
                >
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

.nav {
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    gap: 0.5rem;
    padding: 0.5rem;
    @apply bg-bg flex flex-col;
}

.nav-icon {
    @apply w-[1rem] text-center mr-[.3rem];
}

.nav-item {
    @apply py-2 px-3 rounded-sm cursor-pointer hover:bg-ui active:bg-ui-2;
}
</style>
