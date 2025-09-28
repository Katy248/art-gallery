<script setup>
import { PLANS } from "../data";
const sortedPlans = PLANS.sort((a, b) => {
  if (a.done && !b.done) return 1;
  if (!a.done && b.done) return -1;
  return 0;
});
</script>
<template>
  <div>
    <h1 class="text-3xl">Планы</h1>
    <div class="text-xl mt-4 mb-4">
      Задачи по улучшению сайта, может именно ты выполнишь одну изх них?
    </div>

    <div class="flex flex-col gap-2">
      <div v-for="plan in sortedPlans" :key="plan.name" class="flex flex-col">
        <div class="flex gap-2 items-baseline text-lg">
          <i
            v-if="plan.done"
            title="Выполнено"
            class="fa-solid fa-circle text-green"
          ></i>
          <i
            v-else
            title="Не выполнено"
            class="fa-solid fa-circle-notch text-orange"
          ></i>
          <i
            title="Обсуждается"
            v-if="plan.underDiscussion"
            class="fa-solid fa-comment text-yellow"
          ></i>
          {{ plan.name }}
        </div>
        <div v-if="plan.justification" class="text-tx-2">
          {{ plan.justification }}
        </div>
      </div>
    </div>
  </div>
</template>
