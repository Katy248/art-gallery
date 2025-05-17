<script setup>
const changelogs = [{ title: "Версия 0.0.2", changes: ["Добавлено удаление постов", "Добавлена роль администратора", "Добавлено отображение времени публикации (не только даты)"] }];
</script>
<template>
    <div class="bg-bg-2 py-2 px-4 rounded-lg">
        <h2 class="text-lg font-bold mb-3">Список изменений:</h2>
        <ul>
            <li v-for="version in changelogs" :key="version.title">
                <h3 class="font-bold mb-2"><i class="fa-solid fa-code-commit"></i> {{ version.title }}</h3>
                <ul>
                    <li v-for="change in version.changes" :key="change" class="ps-2">• {{ change }}</li>
                </ul>
            </li>
        </ul>
    </div>
</template>
