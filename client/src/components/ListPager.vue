<script setup>
import { computed, ref } from "vue";

const props = defineProps({
  pages: {
    type: Number,
    required: true,
  },
  pageButtonActivatedHandler: {
    type: Function,
    required: true,
  },
  currentPage: {
    type: Number,
    required: true,
  },
});

const previewPages = computed(() => {
  let result = [];
  console.log(props.pages);
  if (props.pages > 5) {
    if (props.currentPage != 0) {
      result.push({
        page: 0,
        startButton: true,
      });
    }
    if (props.currentPage != 0) {
      result.push({ page: props.currentPage - 1 });
    }
    result.push({ page: props.currentPage });
    if (props.currentPage < props.pages - 1) {
      result.push({ page: props.currentPage + 1 });
    }
    if (props.currentPage < props.pages - 1) {
      result.push({
        page: props.pages,
        endButton: true,
      });
    }
  } else {
    for (let i = 0; i < props.pages; i++) {
      result.push({ page: i });
    }
  }
  return result;
});
</script>
<template>
  <button
    class="btn px-4 py-1 rounded bg-ui hover:bg-ui-2 active:bg-ui-3"
    v-for="p in previewPages"
    :key="p.page"
    @click="pageButtonActivatedHandler(p.page)"
    :class="currentPage === p.page ? 'text-magenta px-20' : ''"
    :title="p.page + 1"
  >
    <span v-if="p.startButton">
      <i class="fas fa-angle-left"></i>
    </span>
    <span v-else-if="p.endButton">
      <i class="fas fa-angle-right"></i>
    </span>
    <span v-else>
      {{ p.page + 1 }}
    </span>
  </button>
</template>
