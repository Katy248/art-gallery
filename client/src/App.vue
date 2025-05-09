<script setup>
import { RouterLink, RouterView } from "vue-router";
import TheNav from "./components/TheNav.vue";
import TheFooter from "./components/TheFooter.vue";
</script>

<template>
    <div class="flex flex-col min-h-screen items-center gap-5 w-full">
        <TheNav />
        <main class="grow max-w-[700px] px-5 w-full">
            <RouterView />
        </main>
        <TheFooter />
    </div>
</template>
<style>
@import "./style.css";
body {
    @apply dark:bg-[#100F0F];
}
</style>
