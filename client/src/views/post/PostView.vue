<script setup>
import { fetchPost } from "../../api";
import { useRoute } from "vue-router";

const route = useRoute();

const { data, error, statusCode } = fetchPost(route.params.id);
</script>
<template>
  <div v-if="error" class="bg-red-2 border-2 border-red p-4 font-semibold">
    {{ statusCode }}: {{ error }}
  </div>
  <div v-if="data" class="flex flex-col gap-2">
    <div
      v-if="data.post.warningMessage"
      class="rounded-md bg-orange-2 px-3 py-2 flex gap-4 border-orange border-2"
    >
      <div>
        <i class="fas fa-triangle-exclamation"></i>
      </div>
      <div class="font-bold">
        {{ data.post.warningMessage }}
      </div>
    </div>
    <img :src="data.post.imageUrl" class="rounded border-2 border-bg-2" />
    <div class="text-balance text-lg">
      {{ data.post.description }}
    </div>
    {{ data }}
  </div>
</template>
