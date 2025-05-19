<script setup>
import { RouterLink } from "vue-router";
defineProps({
  error: {
    type: Object,
    required: true,
  },
});
</script>
<template>
  <div class="flex flex-col gap-12 grow">
    <div v-if="error.code == 404">
      <h1>
        <i class="fa-solid fa-triangle-exclamation"></i>
        404 Не найдено
      </h1>
    </div>
    <div>
      <RouterLink to="/" class="btn btn-base"> На главную </RouterLink>
    </div>
  </div>
</template>

<style scoped>
@import "../style.css";

h1 {
  @apply text-3xl text-red-2 text-center font-bold;
}
</style>
