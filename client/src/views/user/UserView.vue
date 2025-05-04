<script setup>
import { ref } from "vue";
import { getPosts, getUserInfo, savePost, unsavePost } from "../../api";
import { useRoute } from "vue-router";

const route = useRoute();

const user = ref({});
const postsPage = ref(0);

const posts = ref([]);

getUserInfo(route.params.id).then((r) => {
    user.value = r;
    console.log(user.value);
});
getPosts(route.params.id, pageXOffset.value).then((r) => {
    posts.value = r;
    console.log(r);
});

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
</script>
<template>
    <div class="flex flex-col gap-8">
        <div class="flex flex-row gap-4 justify-center items-center">
            <img :src="user.avatarUrl" class="h-50 rounded-full" />
            <div class="text-center text-xl font-semibold">
                {{ user.name }}
            </div>
        </div>
        <div class="flex justify-center items-center">
            <div class="flex flex-row gap-2 bg-bg-2 p-1 w-fit rounded-lg justify-center items-center">
                <button class="btn btn-base">Публикации</button>
                <button class="btn">Сохранённые</button>
            </div>
        </div>
        <div class="grid grid-cols-2 gap-4">
            <div v-for="post in posts" class="bg-bg-2 p-2 gap-2 flex flex-col rounded-lg">
                <img :src="post.imageUrl" class="rounded-md border border-ui-2" />
                <div class="grow">{{ post.description }}</div>
                <div class="flex gap-2">
                    <button v-if="post.saved" class="btn btn-base" @click="() => unsavePostHandler(post)">Сохранено</button>
                    <button v-else class="btn btn-primary" @click="() => savePostHandler(post)">Сохранить</button>
                    <button class="btn btn-base">В коллекцию</button>
                </div>
            </div>
        </div>
    </div>
</template>
