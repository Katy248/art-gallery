<script setup>
const changelogs = [{ title: "version 0.0.2", changes: ["Added post deletion", "Added administrators"] }];
</script>
<template>
    <div class="bg-bg-2 p-2 rounded-lg">
        <h1 class="text-lg font-bold mb-3">Changelogs</h1>
        <ul>
            <li v-for="version in changelogs" :key="version.title">
                <h2 class="font-bold mb-2">{{ version.title }}</h2>
                <ul>
                    <li v-for="change in version.changes" :key="change">- {{ change }}</li>
                </ul>
            </li>
        </ul>
    </div>
</template>
