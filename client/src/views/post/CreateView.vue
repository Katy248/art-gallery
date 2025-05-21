<script setup>
import { ref } from "vue";
import { createPost } from "../../api";
import { useRouter } from "vue-router";

const router = useRouter();

const description = ref("");
const image = ref();
const previewUrl = ref("");

const onFileChange = (e) => {
  const file = e.target.files[0];
  console.log(file);
  image.value = file;
  const fr = new FileReader();
  fr.readAsDataURL(image.value);
  fr.addEventListener("load", () => {
    previewUrl.value = fr.result;
  });
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
        <textarea
          v-model="description"
          class="text-entry"
          placeholder="Крутая фотка"
        />
      </div>
      <div class="input-group">
        <div v-if="image" class="w-full flex flex-col gap-4">
          <img :src="previewUrl" class="rounded-md" />
          <button class="btn btn-base" @click="image = null">
            Удалить фото
          </button>
        </div>
        <input
          v-if="!image"
          @change="onFileChange"
          accept="image/png, image/jpeg, image/gif"
          type="file"
          class="text-entry"
          placeholder="example@mail.ru"
          autocomplete="username"
        />
      </div>
      <div class="input-group">
        <input type="submit" class="btn btn-primary" />
      </div>
    </form>
  </div>
</template>
