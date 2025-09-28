<script setup>
import { onMounted, ref, onBeforeUnmount } from "vue";
const prompt = ref(null);
const showInstallPrompt = ref(false);

onMounted(() => {
  window.addEventListener("beforeinstallprompt", beforeInstallPromptHandler);
  window.addEventListener("appinstalled", handleAppInstalled);
});

onBeforeUnmount(() => {
  window.removeEventListener("beforeinstallprompt", beforeInstallPromptHandler);
  window.removeEventListener("appinstalled", handleAppInstalled);
});

const beforeInstallPromptHandler = (event) => {
  console.log("before prompt event", event);
  //   event.preventDefault();
  prompt.value = event;
  showInstallPrompt.value = true;
};
const installPWA = async () => {
  if (!prompt.value) return;
  prompt.value.prompt();
  const { outcome } = await prompt.value.userChoice;
  if (outcome === "accepted") {
    console.log("User accepted the install prompt");
  } else {
    console.log("User dismissed the install prompt");
  }

  showInstallPrompt.value = false;
};
const handleAppInstalled = () => {
  console.log("PWA installed");
  showInstallPrompt.value = false;
};
</script>
<template>
  <div v-if="showInstallPrompt" class="mb-6 p-4 bg-bg-2 rounded-xl">
    <p class="mb-2 text-balance">
      <span class="text-magenta font-semibold">ArtGallery</span> можно
      установить как
      <a
        class="underline active:text-magenta"
        href="https://developer.mozilla.org/ru/docs/Web/Progressive_web_apps"
        >прогрессивное веб-приложение</a
      >. Это может значительно улучшить ваш опыт использования на мобильных
      устройствах.
    </p>
    <div class="flex grow">
      <button
        class="cursor-pointer transition-all duration-300 rounded-md grow py-4 bg-magenta-2 hover:bg-magenta-2/90 active:bg-magenta-2/80 active:ring-2 ring-magenta-2"
        @click="installPWA()"
      >
        Установить
      </button>
    </div>
  </div>
  <!-- <p v-else>Прогрессивное веб-приложение уже установлено</p> -->
</template>
