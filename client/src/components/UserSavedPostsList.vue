<script setup>
import { ref, onMounted } from "vue";
import { getAvatar, savePost, unsavePost, getSavedPosts } from "../api";
import PostCard from "./PostCard.vue";

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
</script>
<template>
    <div class="flex flex-col gap-4">
        <div class="grid md:grid-cols-2 gap-4">
            <PostCard v-for="post in posts" :post="post" :showPublisher="true" />
        </div>
        <div class="flex gap-1 justify-center">
            <button class="btn-base px-4 py-1 rounded" v-for="p in [...Array(pages).keys()]">{{ p + 1 }}</button>
        </div>
    </div>
</template>
