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
const pages = ref(1);
const posts = ref([]);

const loadPosts = () => {
    getSavedPosts(props.userId, page.value).then((r) => {
        posts.value = r.posts;
        pages.value = r.pages;

        posts.value.forEach((post) => {
            getAvatar(post.publisherId).then((r) => {
                post.publisherAvatar = r.url;
            });
        });
        console.log("Response");
        console.log(r);
    });
};
const savePostHandler = (post) => {
    savePost(post.id).then((r) => {
        post.saved = r.saved;
    });
};
const unsavePostHandler = (post) => {
    unsavePost(post.id).then((r) => {
        post.saved = r.saved;
    });
};
onMounted(() => {
    loadPosts();
});
</script>
<template>
    <PostCard v-for="post in posts" :post="post" :showPublisher="true" />
    <div class="flex gap-1 justify-center">
        <button class="btn-base px-4 py-1 rounded" v-for="p in [...Array(pages).keys()]">{{ p + 1 }}</button>
    </div>
</template>
