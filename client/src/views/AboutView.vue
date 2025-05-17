<script setup>
import { RouterLink } from "vue-router";
import { LINKS } from "../data";
</script>
<template>
    <section class="plain-text">
        <h2>О нас</h2>
        <p>
            Добро пожаловать в <span class="site-name">ArtGallery</span> — уникальную платформу для обмена искусством. Здесь художники и любители искусства могут делиться своими произведениями,
            находить вдохновение и взаимодействовать с другими творческими личностями.
        </p>
        <p>
            Наша цель — создать сообщество, где каждый может выразить себя, научиться чему-то новому и найти единомышленников. Мы поддерживаем разнообразие стилей и техник, от классической живописи до
            цифрового искусства.
        </p>

        <h3>Помощь</h3>
        <p>
            Если у вас возникли вопросы или вам нужна помощь, вы можете ознакомиться с нашими <RouterLink to="/help">руководством пользователя</RouterLink> или обратиться в службу поддержки. Мы всегда
            готовы помочь вам!
        </p>

        <!-- <h3>Контакты</h3>
        <p>Свяжитесь с нами по электронной почте: <a href="mailto:support@artgallery.com">support@artgallery.com</a></p>
        <p>Мы также активно ведем наши социальные сети, где вы можете следить за новостями и событиями:</p> -->
        <ul>
            <li v-for="link in LINKS">
                <a :href="link.url" target="_blank">
                    <i v-if="link.icon" :class="link.icon" class="me-1"></i>
                    {{ link.name }}</a
                >
            </li>
            <!-- <li>
                <a href="https://vk.com/artgallery" target="_blank"><i class="fab fa-vk pr-2"></i>VK</a>
            </li>
            <li>
                <a href="https://rutube.ru/artgallery" target="_blank"><i class="fab fa-youtube pr-2"></i>Rutube</a>
            </li> -->
        </ul>
    </section>
</template>
