<script setup>
import { ref, onMounted } from "vue";
import { getSavedPosts } from "../api";
import PostCard from "./PostCard.vue";
import ListPager from "./ListPager.vue";

const props = defineProps({
  userId: {
    type: Number,
    required: true,
  },
});

const page = ref(0);
const pages = ref(0);
const posts = ref([]);

const loadPosts = () => {
  getSavedPosts(props.userId, page.value).then((r) => {
    posts.value = r.posts;
    pages.value = r.pages;
  });
};
onMounted(() => {
  loadPosts();
});
const switchPage = (page) => {
  page.value = page;
  loadPosts();
};
</script>
<template>
  <div class="flex flex-col gap-4">
    <div class="grid md:grid-cols-2 gap-4">
      <PostCard v-for="post in posts" :post="post" :showPublisher="true" />
    </div>
    <div class="flex gap-1 justify-center">
      <ListPager
        :pages="[...Array(pages).keys()]"
        :pageButtonActivatedHandler="switchPage"
        :currentPage="page"
      />
    </div>
  </div>
</template>
