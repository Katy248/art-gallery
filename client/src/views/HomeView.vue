<script setup>
import { RouterLink } from "vue-router";
import { getAllPosts } from "../api";
import TheFeed from "../components/TheFeed.vue";
import { onMounted, ref } from "vue";
import ListPager from "../components/ListPager.vue";
import { useAuthStore } from "../stores/auth";
import ChangesList from "../components/ChangesList.vue";
const posts = ref([]);

const authStore = useAuthStore();

const page = ref(0);
const pagesCount = ref(0);

onMounted(() => {
  if (!authStore.isAuthenticated) {
    return;
  }
  getAllPosts(page.value).then((res) => {
    posts.value = res.posts;
    pagesCount.value = res.pages;
  });
});
const switchPage = (p) => {
  getAllPosts(p)
    .then((res) => {
      posts.value = res.posts;
      pagesCount.value = res.pages;
      page.value = p;
      window.scrollTo(0, 0);
    })
    .catch((err) => {
      console.error(err);
    });
};
</script>
<template>
  <div v-if="authStore.isAuthenticated" class="flex flex-col gap-8 w-full">
    <div class="mb-8">
      <ChangesList />
    </div>
    <div>
      <RouterLink class="btn btn-primary" to="/post/create">
        <i class="fas fa-upload"></i>
        Создать публикацию</RouterLink
      >
    </div>
    <TheFeed :posts="posts" />
    <div class="flex justify-center gap-2">
      <ListPager
        :pages="pagesCount"
        :currentPage="page"
        :pageButtonActivatedHandler="switchPage"
      />
    </div>
  </div>
  <div v-else class="flex flex-col items-center">
    <h1 class="text-center text-5xl font-bold text-magenta-2 pb-10">
      ArtGallery
    </h1>
    <div class="pb-10 flex flex-col gap-10 items-center">
      <div class="text-xl font-semibold">
        Для всякого, для всего, для ничего, для души...
      </div>
      <div class="w-full flex justify-center">
        <RouterLink
          class="text-center cursor-pointer font-semibold text-magenta p-4 hover:text-tx active:text-tx hover:bg-magenta/70 active:bg-magenta/80 border-magenta border-2 rounded-md transition-colors duration-300 w-full"
          to="/about"
          >Подробнее
          <i class="fas fa-angle-right"></i>
        </RouterLink>
      </div>
    </div>

    <div class="flex flex-col md:flex-row gap-4">
      <RouterLink class="btn btn-base" to="/auth/login">
        <i class="fas fa-arrow-right-to-bracket"></i>
        Вход</RouterLink
      >
      <RouterLink class="btn btn-base" to="/auth/register">
        <i class="fas fa-user-plus"></i>
        Регистрация</RouterLink
      >
      <RouterLink class="btn btn-base" to="/help">
        <i class="fas fa-question"></i> Помощь</RouterLink
      >
    </div>
    <div class="mt-8">
      <ChangesList />
    </div>
  </div>
</template>
