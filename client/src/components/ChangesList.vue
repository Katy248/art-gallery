<script setup>
import { ref } from "vue";
import { Dialog, DialogPanel, TransitionRoot } from "@headlessui/vue";
const changelogs = [
    { title: "Версия 0.0.3", changes: ["Добавлено описание пользователя", "Добавлено изменение текста публикации"] },
    { title: "Версия 0.0.2", changes: ["Добавлено удаление постов", "Добавлена роль администратора", "Добавлено отображение времени публикации (не только даты)"] },
];

const showModal = ref(false);
</script>
<template>
    <div class="bg-bg-2 py-2 px-4 rounded-lg flex flex-col gap-6">
        <h2 class="text-lg font-bold">Список изменений:</h2>
        <div class="flex flex-col gap-3">
            <h3 class="font-bold mb-2"><i class="fa-solid fa-code-commit"></i> {{ changelogs[0].title }}</h3>
            <ul>
                <li v-for="change in changelogs[0].changes" :key="change" class="ps-2">• {{ change }}</li>
            </ul>
        </div>

        <button class="btn btn-base" @click="showModal = true">Все изменения</button>
        <TransitionRoot as="template" :show="showModal" enter="ease-out duration-200" enter-from="opacity-20" enter-to="opacity-100">
            <Dialog @close="showModal = false" class="relative z-50">
                <div class="fixed inset-0 flex w-screen items-center justify-center backdrop-blur-sm p-8">
                    <DialogPanel>
                        <div class="bg-bg-2 p-8 rounded-lg flex flex-col gap-6 border-ui border">
                            <h2 class="text-lg font-bold">Список изменений:</h2>
                            <ul class="flex flex-col gap-3">
                                <li v-for="version in changelogs" :key="version.title">
                                    <h3 class="font-bold mb-2"><i class="fa-solid fa-code-commit"></i> {{ version.title }}</h3>
                                    <ul>
                                        <li v-for="change in version.changes" :key="change" class="ps-2">• {{ change }}</li>
                                    </ul>
                                </li>
                            </ul>
                        </div>
                    </DialogPanel>
                </div>
            </Dialog>
        </TransitionRoot>
    </div>
</template>
