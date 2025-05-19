<script setup>
import { onMounted, ref } from "vue";
import { getPost, updatePost } from "../../api";
import { useRoute, useRouter } from "vue-router";
import { useAuthStore } from "../../stores/auth";
const router = useRouter();
const route = useRoute();
const postId = route.params.id;
const props = defineProps({});

const auth = useAuthStore();

const post = ref({});
onMounted(() => {
  getPost(parseInt(postId)).then((res) => {
    post.value = res.post;
  });
});
const onSubmit = () => {
  updatePost(post.value).then((res) => {
    if (res.success) {
      console.log("Post updated successfully");
      router.push("/");
    } else {
      console.error(res.error);
    }
  });
};
</script>
<template>
  <form class="flex flex-col gap-4" @submit.prevent="onSubmit">
    <div class="input-group">
      <img :src="post.imageUrl" class="rounded-lg" />
    </div>

    <div class="input-group">
      <label>Описание</label>
      <input class="text-entry" v-model="post.description" />
    </div>
    <div class="input-group" v-if="auth.authData.isAdmin">
      <label>Предупредительное сообщение</label>
      <input class="text-entry" v-model="post.warningMessage" />
    </div>
    <div class="input-group">
      <input
        class="btn btn-primary"
        type="submit"
        value="Сохранить изменения"
      />
    </div>
  </form>
</template>
