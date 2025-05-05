<script setup>
import { ref } from "vue";
import { register } from "../../api";
import { useRouter } from "vue-router";

const router = useRouter();

const name = ref("");
const email = ref("");
const password = ref("");
const passwordRepeated = ref("");

const onSubmit = () => {
    if (password.value !== passwordRepeated.value) {
        console.error("Пароли не совпадают");
        return;
    }

    register(email.value, name.value, password.value).then(() => {
        console.log("Регистрация прошла успешно");
        router.push("/");
    });
};
</script>
<template>
    <div class="bg-bg-2 p-4 rounded-lg">
        <form class="flex flex-col gap-4" @submit.prevent="onSubmit">
            <div class="input-group">
                <label>Имя:</label>
                <input v-model="name" class="text-entry" placeholder="CoolGuy" autocomplete="username" />
            </div>
            <div class="input-group">
                <label>Email:</label>
                <input v-model="email" class="text-entry" placeholder="example@mail.ru" autocomplete="email " />
            </div>
            <div class="input-group">
                <label>Пароль:</label>
                <input v-model="password" class="text-entry" type="password" placeholder="******" autocomplete="new-password" />
            </div>
            <div class="input-group">
                <label>Пароль (повторно):</label>
                <input v-model="passwordRepeated" class="text-entry" type="password" placeholder="******" autocomplete="new-password" />
            </div>
            <div class="input-group">
                <input class="btn btn-primary" type="submit" value="Войти" />
            </div>
        </form>
    </div>
</template>
