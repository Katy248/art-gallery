<script setup>
import { onMounted, ref } from "vue";
import { getPost, updatePost } from "../../api";
import { useRoute, useRouter } from "vue-router";
const router = useRouter();
const route = useRoute();
const postId = route.params.id;
const props = defineProps({});

const post = ref({});
onMounted(() => {
    getPost(parseInt(postId)).then((res) => {
        post.value = res.post;
    });
});
const onSubmit = () => {
    updatePost(post.value.id, post.value.description).then((res) => {
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
            <input class="text-entry" v-model="post.description" />
        </div>
        <div class="input-group">
            <input class="btn btn-primary" type="submit" value="Сохранить изменения" />
        </div>
    </form>
</template>
