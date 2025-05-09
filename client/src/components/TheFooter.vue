<script setup>
import { ref } from "vue";
const currentYear = ref(new Date());
</script>
<template>
    <footer class="w-full">
        <div class="flex gap-10 px-20 justify-center">
            <section>
                <h5>Контакты</h5>
                <a class="item" href="mailto:support@artgallery.com">support@artgallery.com</a>
            </section>
            <section>
                <h5>Ссылки</h5>
                <ul>
                    <li><a class="item" href="https://vk.com/artgallery" target="_blank">VK</a></li>
                    <li><a class="item" href="https://rutube.ru/artgallery" target="_blank">Rutube</a></li>
                </ul>
            </section>
        </div>
        <p class="text-center py-2 text-tx-2">© 2023-{{ currentYear.getFullYear() }} ArtGallery. Все права защищены.</p>
    </footer>
</template>

<style scoped>
@reference "../style.css";
h5 {
    @apply pb-2 text-tx-2;
}
.item {
    @apply text-tx-2/50 text-sm hover:underline active:text-tx-2;
}
</style>
