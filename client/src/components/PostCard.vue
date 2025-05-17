<script setup>
import { onMounted, ref } from "vue";
import { savePost, unsavePost, getAvatar, deletePost } from "../api";
import { useAuthStore } from "../stores/auth";
import { RouterLink } from "vue-router";
import { Menu, MenuButton, MenuItems, MenuItem } from "@headlessui/vue";
const authStore = useAuthStore();
const userId = authStore.authData.id;
const deleted = ref(false);

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
const deletePostHandler = () => {
    deletePost(props.post.id).then((r) => {
        if (r.success) {
            props.post.deleted = true;
            deleted.value = true;
        }
    });
};
</script>
<template>
    <div v-if="!deleted" class="w-full">
        <div class="bg-bg-2 p-2 gap-2 flex flex-col rounded-lg w-full">
            <div class="flex flex-row justify-between">
                <div>
                    <RouterLink v-if="showPublisher" class="flex items-center gap-2 hover:bg-ui-2 py-1 px-1 rounded-sm justify-start w-fit" :to="`/user/${post.publisherId}`">
                        <img v-if="post.publisherAvatar" :src="post.publisherAvatar" class="h-6 rounded-full" />
                        <span>{{ post.publisherName }}</span>
                    </RouterLink>
                </div>
                <Menu as="div" class="relative">
                    <MenuButton class="btn-base cursor-pointer rounded-sm py-1 px-3">
                        <i class="fas fa-ellipsis"></i>
                    </MenuButton>
                    <Transition enter-active-class="transform scale-90" leave-active-class="transform scale-90">
                        <MenuItems class="absolute rounded top-10 right-0 nav shadow-xl/30 shadow-text origin-top-right">
                            <MenuItem>
                                <button class="nav-item text-left"><i class="fas fa-info-circle nav-icon"></i> Подробнее</button>
                            </MenuItem>
                            <MenuItem v-if="post.publisherId == userId">
                                <button @click="deletePostHandler" class="nav-item text-left text-red"><i class="fas fa-trash nav-icon"></i> Удалить</button>
                            </MenuItem>
                        </MenuItems>
                    </Transition>
                </Menu>
            </div>
            <img :src="post.imageUrl" class="rounded-md border border-ui-2" />
            <div class="grow">{{ post.description }}</div>
            <div v-if="post.createdAt" class="text-ui-3">
                {{ new Date(post.createdAt).toLocaleTimeString("ru-RU").slice(0, -3) }}
                {{ new Date(post.createdAt).toLocaleDateString("ru-RU") }}
            </div>
            <div class="grid grid-cols-2 md:flex sm:flex-row sm:flex-wrap gap-2 justify-evenly">
                <button v-if="post.saved" class="btn grow bg-ui hover:bg-red-2 active:bg-red" @click="() => unsavePostHandler(post)">
                    <i class="fas fa-bookmark"></i>
                    Сохранено
                </button>
                <button v-else class="btn grow btn-primary" @click="() => savePostHandler(post)">
                    <i class="far fa-bookmark"></i>
                    Сохранить
                </button>
                <!-- <button class="btn btn-base">
                    <i class="fas fa-plus"></i>
                    В коллекцию
                </button>
                <button class="btn btn-base">
                    <i class="fas fa-share"></i>
                    Поделиться
                </button> -->
                <a type="_blank" class="grow btn btn-base" :href="post.imageUrl">
                    <i class="fas fa-image"></i>
                    Оригинал
                </a>
            </div>
        </div>
    </div>
</template>
