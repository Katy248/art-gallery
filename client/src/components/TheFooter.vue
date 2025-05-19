<script setup>
import { ref } from "vue";
import { CONTACTS, LINKS } from "../data";
const currentYear = ref(new Date());
</script>
<template>
  <footer class="w-full">
    <div class="flex gap-10 px-20 justify-center">
      <section>
        <h5>Контакты</h5>
        <ul>
          <li v-for="contact in CONTACTS">
            <a
              class="item"
              :href="contact.url"
              target="_blank"
              :title="contact.name"
            >
              <span v-if="contact.email">{{ contact.email }}</span>
              <span v-else>{{ contact.name }}</span>
            </a>
          </li>
        </ul>
      </section>
      <section>
        <h5>Ссылки</h5>
        <ul>
          <li v-for="link in LINKS" class="item">
            <a :href="link.url" target="_blank">
              <i v-if="link.icon" :class="link.icon" class="me-1"></i>
              {{ link.name }}</a
            >
          </li>
        </ul>
      </section>
    </div>
    <p class="text-center py-2 text-tx-2">
      © 2023-{{ currentYear.getFullYear() }} ArtGallery. Все права защищены.
    </p>
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
