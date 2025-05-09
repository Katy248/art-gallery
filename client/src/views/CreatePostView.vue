<script setup>
import { ref } from "vue";
import { createPost } from "../api";
import { useRouter } from "vue-router";

const router = useRouter();

const description = ref("");
const image = ref();

const onFileChange = (e) => {
    const file = e.target.files[0];
    console.log(file);
    image.value = file;
};
const onSubmit = () => {
    if (!image.value) {
        console.error("no image");
        return;
    }
    console.log(description.value);
    console.log(image.value);

    createPost(description.value, image.value)
        .then((r) => {
            if (r.success) {
                router.push(`/`);
            } else {
                console.error(r.error);
            }
        })
        .catch(console.error);
};
</script>
<template>
    <div>
        <form class="flex flex-col gap-4" @submit.prevent="onSubmit">
            <div class="input-group">
                <label>Описание:</label>
                <textarea v-model="description" class="text-entry" placeholder="Крутая фотка" />
            </div>
            <div class="input-group">
                <label>Изображение:</label>
                <input @change="onFileChange" accept="image/png, image/jpeg" type="file" class="text-entry" placeholder="example@mail.ru" autocomplete="username" />
            </div>
            <div class="input-group">
                <input type="submit" class="btn btn-primary" />
            </div>
        </form>
    </div>
</template>
