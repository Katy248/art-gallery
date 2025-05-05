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
</script>
<template>
    <div class="flex flex-col gap-8">
        <div class="flex flex-row gap-4 justify-center items-center">
            <img :src="user.avatarUrl" class="h-50 rounded-full" />
            <div class="font-semibold flex flex-col">
                <div class="text-2xl">
                    {{ user.name }}
                </div>
                <div>
                    {{ user.email }}
                </div>
            </div>
        </div>

        <div v-if="user.id" class="flex justify-center items-center flex-col gap-4 w-full grow">
            <TabGroup>
                <TabList>
                    <div class="flex gap-2 bg-bg-2 p-2 rounded-lg w-full">
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
                        <UserSavedPostsList :userId="user.id" />
                    </TabPanel>
                </TabPanels>
            </TabGroup>
        </div>
    </div>
</template>
