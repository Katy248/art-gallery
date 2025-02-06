<script setup>
defineProps({
    img: {},
    resetImg: Function,
});
</script>
<template>
    <div v-if="img" class="fixed z-50 top-0 left-0 w-full h-screen bg-black/50 hover:bg-black/40 flex justify-center" @click="resetImg()">
        <img :src="img" class="m-20 w-fit" />
    </div>
</template>
