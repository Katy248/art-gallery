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

const PAGES_OFFSET = 2;

const previewPages = computed(() => {
  let result = [];
  console.log(props.pages);

  const pageOutOfRange = (p) => p < 0 || p >= props.pages;

  for (let i = props.currentPage - PAGES_OFFSET; i < props.currentPage; i++) {
    if (pageOutOfRange(i)) continue;

    result.push({ page: i });
  }

  result.push({ page: props.currentPage });

  for (
    let i = props.currentPage + 1;
    i < props.currentPage + 1 + PAGES_OFFSET;
    i++
  ) {
    if (pageOutOfRange(i)) continue;
    result.push({ page: i });
  }

  if (props.currentPage > 0) {
    result = [{ startButton: true, page: 0 }, ...result];
  }

  if (props.currentPage < props.pages - 1) {
    result.push({ endButton: true, page: props.pages - 1 });
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
    <span v-if="p.startButton" title="В начало">
      <i class="fas fa-angles-left"></i>
    </span>
    <span v-else-if="p.endButton" title="В конец">
      <i class="fas fa-angles-right"></i>
    </span>
    <span v-else>
      {{ p.page + 1 }}
    </span>
  </button>
</template>
