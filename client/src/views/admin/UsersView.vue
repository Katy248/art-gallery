<script setup>
import { useRoute, onBeforeRouteUpdate, RouterLink } from "vue-router";
import { Dialog, DialogPanel, TransitionRoot } from "@headlessui/vue";
import { ref, onMounted } from "vue";
import { getUsers, deleteUser } from "../../api/admin";
import { useAuthStore } from "../../stores/auth";
const authStore = useAuthStore();
const route = useRoute();
const users = ref([]);

const getData = (page) => {
  getUsers(parseInt(page)).then((res) => {
    if (res.success) {
      users.value = res.users;
    }
  });
};

onMounted(() => {
  getData(route.params.page);
});
onBeforeRouteUpdate(async (to) => {
  getData(to.params.page);
});

const currentUser = ref({ name: "default" });
const _showModal = ref(false);

const showModal = (user) => {
  currentUser.value = user;
  _showModal.value = true;
};
const onDeleteButtonClick = () => {
  deleteUser(currentUser.value.id).then((res) => {
    if (res.success) {
      _showModal.value = false;
      getData(route.params.page);
    } else {
      console.error(res.error);
    }
  });
};
</script>
<template>
  <div class="flex gap-4 mb-4">
    <button
      class="btn btn-base"
      title="Обновить список"
      @click="getData(route.params.page)"
    >
      <i class="fas fa-rotate-right"></i>
    </button>
  </div>
  <div class="flex flex-col">
    <button
      v-for="user in users"
      :key="user.id"
      @click="showModal(user)"
      class="bg-ui first:rounded-t-xl last:rounded-b-xl px-4 py-2 not-last:border-b-2 border-ui-2 flex flex-col items-start hover:bg-ui-2 active:bg-ui-3 transition-all duration-200 cursor-pointer"
    >
      <h2 class="text-lg">
        {{ user.name }}
        <span class="text-red text-sm" v-if="user.isAdmin">админ</span>
      </h2>
      <div class="text-tx-2">{{ user.email }}</div>
    </button>
  </div>
  <TransitionRoot
    as="template"
    :show="_showModal"
    enter="ease-out duration-200"
    enter-from="opacity-20"
    enter-to="opacity-100"
  >
    <Dialog @close="_showModal = false" class="relative z-50">
      <div
        class="fixed inset-0 flex w-screen items-center justify-center backdrop-blur-sm p-8"
      >
        <DialogPanel>
          <div
            class="bg-bg-2 p-8 rounded-lg flex flex-col gap-6 border-ui border grow"
          >
            <div>
              <h2 class="text-lg">
                {{ currentUser.name }}
              </h2>
            </div>
            <div class="">
              Дата регистрации:
              {{ new Date(currentUser.createdAt).toLocaleDateString() }}
            </div>
            <div class="text-tx-2 font-mono">
              {{ currentUser.email }}
            </div>
            <div class="flex gap-2 flex-col">
              <RouterLink class="btn btn-base" :to="`/user/${currentUser.id}`">
                На страницу пользователя
              </RouterLink>
              <button class="btn btn-base">
                <i class="fas fa-envelope"></i>
                Сменить почту
              </button>
              <button class="btn btn-base">
                <i class="fas fa-key"></i>
                Сменить пароль
              </button>
              <button class="btn btn-primary" v-if="!currentUser.isAdmin">
                <i class="fas fa-user-tie"></i>
                Сделать администратором
              </button>
              <button
                v-if="currentUser.id != authStore.authData.id"
                class="btn btn-danger grow"
                @click="onDeleteButtonClick"
              >
                <i class="fas fa-trash"></i>
                Удалить пользователя
              </button>
            </div>
          </div>
        </DialogPanel>
      </div>
    </Dialog>
  </TransitionRoot>
</template>
<style>
@import "../../style.css";

tr {
  @apply p-4;
}

td {
  @apply text-center py-2;
}
</style>
