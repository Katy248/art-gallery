<script setup>
import { ref } from "vue";
import { authenticate } from "../../api";
import { useAuthStore } from "../../stores/auth";
import { useRouter } from "vue-router";
const authStore = useAuthStore();
const email = ref("");
const password = ref("");
const router = useRouter();

const onSubmit = () => {
  authenticate(email.value, password.value);
};

const checkAuthAndRedirect = () => {
  if (authStore.isAuthenticated) {
    router.push("/");
  }
};

authStore.$subscribe(() => {
  checkAuthAndRedirect();
});
checkAuthAndRedirect();
</script>
<template>
  <div class="flex w-full justify-center items-center">
    <form
      class="bg-bg-2 p-4 rounded-lg flex flex-col gap-4"
      @submit.prevent="onSubmit"
    >
      <div class="input-group">
        <label>Логин:</label>
        <input
          v-model="email"
          class="text-entry"
          placeholder="example@mail.ru"
          autocomplete="username"
        />
      </div>
      <div class="input-group">
        <label>Пароль:</label>
        <input
          v-model="password"
          class="text-entry"
          type="password"
          autocomplete="current-password"
        />
      </div>
      <div class="input-group">
        <input class="btn btn-primary" type="submit" value="Войти" />
      </div>
    </form>
  </div>
</template>
