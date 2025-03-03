<script setup>
const onSubmit = (e) => {};
</script>
<template>
    <div class="bg-bg-2 p-4 rounded-lg">
        <form class="flex flex-col gap-4" @submit.prevent="onSubmit">
            <div class="input-group">
                <label>Логин:</label>
                <input class="text-entry" placeholder="example@mail.ru" />
            </div>
            <div class="input-group">
                <label>Пароль:</label>
                <input class="text-entry" type="password" />
            </div>
            <div class="input-group">
                <input class="btn btn-primary" type="submit" value="Войти" placeholder="***" />
            </div>
        </form>
    </div>
</template>
<style>
@import "../style.css";
.input-group {
    @apply flex flex-col grow gap-1;
    & > label {
        @apply text-sm;
    }
}
.text-entry {
    @apply bg-ui outline-0 px-2 py-1 rounded-md border-ui-2 border;
}
</style>
