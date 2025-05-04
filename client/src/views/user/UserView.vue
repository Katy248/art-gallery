<script setup>
import { ref } from "vue";
import { getPosts, getUserInfo, savePost, unsavePost } from "../../api";
import { useRoute } from "vue-router";
import UserPostsList from "../../components/UserPostsList.vue";
import UserSavedPostsList from "../../components/UserSavedPostsList.vue";
import { TabGroup, TabList, Tab, TabPanel, TabPanels } from "@headlessui/vue";

const route = useRoute();

const user = ref({});
const postsPage = ref(0);

const posts = ref([]);

getUserInfo(route.params.id).then((r) => {
    user.value = r;
    console.log(user.value);
});
// getPosts(route.params.id, pageXOffset.value).then((r) => {
//     posts.value = r;
//     console.log(r);
// });

// const savePostHandler = (post) => {
//     savePost(post.id).then((r) => {
//         post.saved = r.saved;
//     });
// };
// const unsavePostHandler = (post) => {
//     unsavePost(post.id).then((r) => {
//         post.saved = r.saved;
//     });
// };
</script>
<template>
    <div class="flex flex-col gap-8">
        <div class="flex flex-row gap-4 justify-center items-center">
            <img :src="user.avatarUrl" class="h-50 rounded-full" />
            <div class="text-center text-xl font-semibold">
                {{ user.name }}
            </div>
        </div>

        <div v-if="user.id" class="flex justify-center items-center flex-col gap-4">
            <TabGroup>
                <TabList>
                    <div class="flex gap-2 bg-bg-2 p-2 rounded-lg">
                        <Tab as="template" v-slot="{ selected }">
                            <button class="btn outline-none" :class="selected ? 'bg-ui-2' : ''">Публикации</button>
                        </Tab>
                        <Tab as="template" v-slot="{ selected }">
                            <button class="btn outline-none" :class="selected ? 'bg-ui-2' : ''">Сохранённые</button>
                        </Tab>
                    </div>
                </TabList>
                <TabPanels>
                    <TabPanel>
                        <div class="grid md:grid-cols-2 gap-4">
                            <UserPostsList :userId="user.id" />
                        </div>
                    </TabPanel>
                    <TabPanel>
                        <div class="grid md:grid-cols-2 gap-4">
                            <UserSavedPostsList :userId="user.id" />
                        </div>
                    </TabPanel>
                </TabPanels>
            </TabGroup>
        </div>
    </div>
</template>
