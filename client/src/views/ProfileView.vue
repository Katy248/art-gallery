<script setup>
import { ref } from "vue";
import { getProfileInfo } from "../api";
const name = ref("");
const img = ref("");
getProfileInfo().then((response) => {
    name.value = response.name;
    img.value = response["avatar.url"];
    console.log(response);
});
</script>
<template>
    <img :src="img" class="rounded-full" />
    <div>{{ name }}</div>
</template>
