<script setup>
import { ref } from "vue";
import { Dialog, DialogPanel, TransitionRoot } from "@headlessui/vue";
import { CHANGELOGS } from "../data";
const lastChange = ref(CHANGELOGS.filter((change) => !change.preview)[0]);
const showModal = ref(false);
</script>
<template>
    <div class="bg-bg-2 p-4 rounded-lg flex flex-col gap-6">
        <h2 class="text-lg font-bold">Список изменений:</h2>
        <div class="flex flex-col gap-3">
            <h3 class="font-bold mb-2"><i class="fa-solid fa-code-commit"></i> {{ lastChange.title }}</h3>
            <ul>
                <li v-for="change in lastChange.changes" :key="change" class="ps-2">• {{ change }}</li>
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
                                <li v-for="version in CHANGELOGS" :key="version.title">
                                    <h3 class="font-bold mb-2">
                                        <i class="fa-solid fa-code-commit"></i> {{ version.title }}
                                        <span v-if="version.preview" class="text-tx-2 font-normal">preview</span>
                                    </h3>
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
