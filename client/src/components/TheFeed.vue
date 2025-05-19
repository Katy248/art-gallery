<script setup>
import PostCard from "./PostCard.vue";
const props = defineProps({
  posts: {},
});
</script>
<template>
  <div class="flex flex-col gap-6 items-center">
    <PostCard
      v-for="post in props.posts"
      :post="post"
      :showPublisher="true"
      :key="post.id"
    />
  </div>
</template>
