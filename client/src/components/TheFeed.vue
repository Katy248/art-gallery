<script setup>
import { RouterLink } from "vue-router";
import PictureFullView from "./PictureFullView.vue";
import { ref } from "vue";
defineProps({
    posts: {},
});
const currentPic = ref(null);

const zoomIn = (picture) => {
    currentPic.value = picture;
};
const zoomOut = () => {
    currentPic.value = null;
};
</script>
<template>
    <PictureFullView :img="currentPic" :resetImg="zoomOut" />
    <div class="flex flex-col gap-5 items-center">
        <div v-for="p in posts" class="rounded-lg p-3 flex flex-col gap-3 bg-bg-2 w-full">
            <div class="flex justify-between items-stretch">
                <RouterLink class="flex gap-2 justify-baseline text-tx-2 btn-base hover:text-tx w-fit p-2 rounded-lg" :to="`/user/${p.author.id}`">
                    <div><img :src="p.author.avatarUrl" class="w-6 rounded-full" /></div>
                    <div class="">
                        {{ p.author.name }}
                    </div>
                </RouterLink>
                <div class="h-max">
                    <button class="btn btn-base">
                        <i class="fas fa-ellipsis py-2"></i>
                    </button>
                </div>
            </div>
            <!-- <div class="bg-[#282726]7 px-4 py-2 rounded-md flex flex-col gap-2"> -->
            <div class="group">
                <button class="absolute z-0 md:opacity-0 btn btn-base-overlay group-hover:opacity-100 mt-2 ml-2 transition-opacity delay-300 duration-1000" @click="zoomIn(p.imageUrl)">
                    <i class="fas fa-magnifying-glass-plus py-3"></i>
                </button>
                <img :src="p.imageUrl" class="rounded-lg w-full" />
            </div>
            <div class="">{{ p.comment }}</div>
            <!-- </div> -->
            <div class="flex gap-2 text-sm flex-wrap">
                <button class="btn btn-primary"><i class="far fa-bookmark"></i> Сохранить</button>
                <button class="btn btn-base"><i class="fas fa-share"></i> Поделиться</button>
                <button class="btn btn-base"><i class="fas fa-plus"></i> В коллекцию</button>
            </div>
        </div>
    </div>
</template>
