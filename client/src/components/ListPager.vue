<script setup>
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
</script>
<template>
  <button
    class="btn px-4 py-1 rounded bg-ui hover:bg-ui-2 active:bg-ui-3"
    v-for="p in [...Array(pages).keys()]"
    @click="pageButtonActivatedHandler(p)"
    :class="currentPage === p ? 'text-magenta px-20' : ''"
  >
    {{ p + 1 }}
  </button>
</template>
