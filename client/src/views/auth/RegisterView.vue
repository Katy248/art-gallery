<script setup>
import { ref } from "vue";
import { register } from "../../api";
import { useRouter } from "vue-router";

const router = useRouter();
const error = ref("");

const name = ref("");
const email = ref("");
const password = ref("");
const passwordRepeated = ref("");

const onSubmit = () => {
  if (!name.value) {
    error.value = "Имя пользователя не может быть пустым";
    return;
  }
  if (!email.value) {
    error.value = "Email не может быть пустым";
    return;
  }
  if (!password.value) {
    error.value = "Пароль не может быть пустым";
    return;
  }
  if (password.value !== passwordRepeated.value) {
    console.error("Пароли не совпадают");
    error.value = "Пароли не совпадают";
    return;
  }

  register(email.value, name.value, password.value).then((r) => {
    if (r.success) {
      console.log("Регистрация прошла успешно");
      router.push("/");
    } else {
      error.value = "Регистрация не удалась" + r.error;
    }
  });
};
</script>
<template>
  <div class="w-full flex flex-col gap-8 justify-center items-center">
    <div
      v-if="error"
      class="bg-red-2 p-4 w-full rounded-lg flex flex-row items-center gap-2"
    >
      <i class="fas fa-circle-exclamation"></i>
      <div>
        {{ error }}
      </div>
    </div>
    <form
      class="bg-bg-2 p-4 rounded-lg flex flex-col gap-4 w-full sm:w-fit"
      @submit.prevent="onSubmit"
    >
      <div class="input-group">
        <label>Имя:</label>
        <input
          v-model="name"
          class="text-entry"
          placeholder="CoolGuy"
          autocomplete="name"
        />
      </div>
      <div class="input-group">
        <label>Email:</label>
        <input
          v-model="email"
          class="text-entry"
          placeholder="example@mail.ru"
          autocomplete="email "
        />
      </div>
      <div class="input-group">
        <label>Пароль:</label>
        <input
          v-model="password"
          class="text-entry"
          type="password"
          placeholder="******"
          autocomplete="new-password"
        />
      </div>
      <div class="input-group">
        <label>Пароль (повторно):</label>
        <input
          v-model="passwordRepeated"
          class="text-entry"
          type="password"
          placeholder="******"
          autocomplete="new-password"
        />
      </div>
      <div class="input-group">
        <input
          class="btn btn-primary"
          type="submit"
          value="Зарегистрироваться"
        />
      </div>
    </form>
  </div>
</template>
