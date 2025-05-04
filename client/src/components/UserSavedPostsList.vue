<script setup>
import { ref, onMounted } from "vue";
import { getAvatar, savePost, unsavePost, getSavedPosts } from "../api";
import { useRoute } from "vue-router";

const props = defineProps({
    userId: {
        type: Number,
        required: true,
    },
});

const page = ref(0);
const posts = ref([]);

const loadPosts = () => {
    getSavedPosts(props.userId, page.value).then((r) => {
        posts.value = r;
        posts.value.forEach((post) => {
            getAvatar(post.publisherId).then((r) => {
                post.publisherAvatar = r.url;
            });
        });
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
    <div v-for="post in posts" class="bg-bg-2 p-2 gap-2 flex flex-col rounded-lg">
        <div class="flex items-center gap-2">
            <img :src="post.publisherAvatar" class="h-10 rounded-full" />
            <span>{{ post.publisherName }}</span>
        </div>
        <img :src="post.imageUrl" class="rounded-md border border-ui-2" />
        <div class="grow">{{ post.description }}</div>
        <div class="flex gap-2">
            <button v-if="post.saved" class="btn btn-base" @click="() => unsavePostHandler(post)">Сохранено</button>
            <button v-else class="btn btn-primary" @click="() => savePostHandler(post)">Сохранить</button>
            <button class="btn btn-base">В коллекцию</button>
        </div>
    </div>
</template>
