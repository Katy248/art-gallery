<script setup>
import { useAuthStore } from "../../stores/auth";
import { useRouter } from "vue-router";
const router = useRouter();
const authStore = useAuthStore();

authStore.logout();
router.push("/auth/login");
</script>
<template>
  <h1>Logging out...</h1>
</template>
