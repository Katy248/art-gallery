<script setup>
import { ref } from "vue";
import { getUserInfo } from "../../api";
import { onBeforeRouteUpdate, useRoute } from "vue-router";
import UserPostsList from "../../components/UserPostsList.vue";
import UserSavedPostsList from "../../components/UserSavedPostsList.vue";
import { TabGroup, TabList, Tab, TabPanel, TabPanels } from "@headlessui/vue";
import { useAuthStore } from "../../stores/auth";

const user = ref({});
const selectedTabIndex = ref(0);
const authStore = useAuthStore();
const route = useRoute();

const changeTab = (tab) => {
    selectedTabIndex.value = tab;
};
getUserInfo(route.params.id).then((r) => {
    user.value = r;
    console.log(user.value);
});

onBeforeRouteUpdate(async (to, from) => {
    getUserInfo(to.params.id).then((r) => {
        user.value = r;
        console.log(user.value);
    });
    selectedTabIndex.value = 0;
    window.scrollTo(0, 0);
});
</script>
<template>
    <div class="flex flex-col gap-8 items-center">
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
            <!-- <div v-if="user.id == authStore.authData.id" class="flex gap-2 bg-bg-2 p-2 rounded-lg justify-center w-fit">
                <button class="btn btn-base">
                    <i class="fas fa-plus"></i>
                    Создать публикацию
                </button>
                <button class="btn btn-base">
                    <i class="fa-solid fa-image"></i>
                    Создать коллекцию
                </button>
            </div> -->

            <TabGroup :selectedIndex="selectedTabIndex" @change="changeTab">
                <TabList>
                    <div class="flex gap-2 bg-bg-2 p-2 rounded-lg w-full grow">
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
                        <UserPostsList :userId="user.id" />
                    </TabPanel>
                    <TabPanel>
                        <UserSavedPostsList :userId="user.id" />
                    </TabPanel>
                </TabPanels>
            </TabGroup>
        </div>
    </div>
</template>
