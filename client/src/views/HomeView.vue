<script setup>
import { getAllPosts } from "../api";
import TheFeed from "../components/TheFeed.vue";
import { onMounted, ref } from "vue";
const dummyImg = "https://thewowstyle.com/wp-content/uploads/2015/01/nature-images..jpg";
const dummyImg2 = "https://loremflickr.com/cache/resized/defaultImage.small_1000_1000_nofilter.jpg";
const dummyAuthor = {
    id: 2,
    name: "Katy248",
    avatarUrl: "https://gravatar.com/avatar/33396cb6c169b7fa08fafb345653aee268e9e618fda5de8b2bf9889d0413ea2e?size=256",
};
const posts = ref([]);

const page = ref(0);

onMounted(() => {
    getAllPosts(page.value).then((res) => {
        posts.value = res.posts;
    });
});
</script>
<template>
    <TheFeed :posts="posts" />
</template>
