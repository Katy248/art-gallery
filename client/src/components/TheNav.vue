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
    <header class="flex flex-row p-3 sticky top-0 bg-white w-full dark:bg-[#1C1B1A]">
        <button class="bg-white px-6 py-2 rounded-md hover:bg-gray-200 active:bg-gray-300 z-50 font-bold" @click="menuClick">ArtGallery</button>
    </header>
    <div v-show="showMenu" class="w-full h-screen fixed top-0 left-0 bg-black/60 z-40 flex flex-col items-center" @click="backdropClick">
        <nav class="left-3">
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
:root {
    --sidebar-rounded: 0.5rem;
    --sidebar-min-width: 10rem;
}

.toggle-sidebar {
    border-radius: var(--sidebar-rounded);
    min-width: var(--sidebar-min-width);
    font-size: 20px;
    /* Увеличиваем размер текста */
    font-weight: bold;
    /* Делаем текст жирным */
    background-color: #ffffff;
    /* Убираем фон (если нужно) */
    border: none;
    /* Убираем рамку */
    cursor: pointer;
    /* Указатель при наведении */
    color: #000000;
    /* Цвет текста */
    /* padding: 5px 0px; */
    /* Отступы внутри кнопки */
}

/* Выпадающее меню */
nav {
    min-width: var(--sidebar-min-width);
    /* Скрыто по умолчанию */
    position: absolute;
    top: 60px;
    /* Расстояние от кнопки */
    /* Белый фон */
    /* Тень */
    overflow: hidden;
    z-index: 500;
    animation: fadeIn 0.3s ease-in-out;
    /* Анимация */
    display: flex;
    flex-direction: column;
    gap: 0.2rem;
}

.nav {
    background-color: white;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    list-style: none;
    display: flex;
    flex-direction: column;
    /* Убираем маркеры списка */
    /* margin: 0; */
    gap: 0.5rem;
    padding: 0.5rem;
    margin: 0;
}

.nav-icon {
    width: 1rem;
    padding: 0;
    margin-right: 0.3rem;
    color: var(--accent-color);
    text-align: center;
}

.nav-item {
    padding: 0.5rem;
    padding-left: 0.6rem;
    padding-right: 0.6rem;
    border-radius: 0.3rem;
    cursor: pointer;
    text-decoration: none;
    color: #000000;
}

.nav-item:hover,
.toggle-sidebar:hover {
    /* Фон при наведении */
    background-color: #f1f1f1;
    text-decoration: none;
    filter: drop-shadow(0.1rem 0.1rem 0.1rem 20px black);
}

.nav-item:active,
.toggle-sidebar:active {
    /* Фон при наведении */
    background-color: lightgray;
}
</style>
