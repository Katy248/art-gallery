<script setup>
import { onMounted } from "vue";
import { savePost, unsavePost, getAvatar } from "../api";

const props = defineProps({
    post: {
        type: Object,
        required: true,
    },
    showPublisher: {
        type: Boolean,
        default: false,
    },
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
onMounted(() => {
    if (props.showPublisher) {
        getAvatar(props.post.publisherId).then((r) => {
            props.post.publisherAvatar = r.url;
        });
    }
});
</script>
<template>
    <div class="w-full">
        <div class="bg-bg-2 p-2 gap-2 flex flex-col rounded-lg w-full">
            <RouterLink v-if="showPublisher" class="flex items-center gap-2 hover:bg-ui-2 py-2 px-1 rounded-sm justify-start w-fit" :to="`/user/${post.publisherId}`">
                <img v-if="post.publisherAvatar" :src="post.publisherAvatar" class="h-6 rounded-full" />
                <span>{{ post.publisherName }}</span>
            </RouterLink>
            <img :src="post.imageUrl" class="rounded-md border border-ui-2" />
            <div class="grow">{{ post.description }}</div>
            <div class="grid grid-cols-2 md:flex sm:flex-row sm:flex-wrap gap-2">
                <button v-if="post.saved" class="btn bg-ui hover:bg-red-2 active:bg-red" @click="() => unsavePostHandler(post)">
                    <i class="fas fa-bookmark"></i>
                    Сохранено
                </button>
                <button v-else class="btn btn-primary" @click="() => savePostHandler(post)">
                    <i class="far fa-bookmark"></i>
                    Сохранить
                </button>
                <button class="btn btn-base">
                    <i class="fas fa-plus"></i>
                    В коллекцию
                </button>
                <button class="btn btn-base">
                    <i class="fas fa-share"></i>
                    Поделиться
                </button>
                <a class="btn btn-base">
                    <i class="fas fa-image"></i>
                    Оригинал
                </a>
            </div>
        </div>
    </div>
</template>
