<script setup>
import { ref } from "vue";
import { getUserInfo } from "../../api";
import { onBeforeRouteUpdate, useRoute } from "vue-router";
import UserPostsList from "../../components/UserPostsList.vue";
import UserSavedPostsList from "../../components/UserSavedPostsList.vue";
import { TabGroup, TabList, Tab, TabPanel, TabPanels } from "@headlessui/vue";
import { useAuthStore } from "../../stores/auth";
import ErrorPresenter from "../../components/ErrorPresenter.vue";

const user = ref(null);
const selectedTabIndex = ref(0);
const authStore = useAuthStore();
const route = useRoute();

const changeTab = (tab) => {
  selectedTabIndex.value = tab;
};
const error = ref(null);

const getInfo = (id) => {
  getUserInfo(id)
    .then((r) => {
      console.log(r);

      if (r.error) {
        error.value = r;
        return;
      }
      user.value = r;
      console.log(user.value);
    })
    .catch((err) => {
      log.error(err);
    });
};

onBeforeRouteUpdate(async (to, from) => {
  getInfo(to.params.id);
  selectedTabIndex.value = 0;
  window.scrollTo(0, 0);
});
getInfo(route.params.id);
</script>
<template>
  <div v-if="user" class="flex flex-col gap-8 items-center">
    <div class="flex flex-col md:flex-row gap-4 justify-center items-center">
      <img :src="user.avatarUrl" class="h-50 rounded-full" />
      <div class="flex flex-col">
        <div class="text-2xl font-semibold">
          {{ user.name }}
        </div>
        <div class="font-semibold">
          {{ user.email }}
        </div>
        <div class="pt-2">
          {{ user.description }}
        </div>
      </div>
    </div>

    <div
      v-if="user.id"
      class="flex justify-center items-center flex-col gap-4 w-full grow"
    >
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
              <button
                class="btn hover:bg-ui-2 active:bg-ui-3 outline-none"
                :class="selected ? 'bg-ui-2' : ''"
              >
                Публикации
              </button>
            </Tab>
            <Tab as="template" v-slot="{ selected }">
              <button
                class="btn hover:bg-ui-2 active:bg-ui-3 outline-none"
                :class="selected ? 'bg-ui-2' : ''"
              >
                Сохранённые
              </button>
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
  <ErrorPresenter v-if="error?.error" :error="error" />
</template>
