<script setup>
import { useRoute, onBeforeRouteUpdate } from "vue-router";
import { ref, onMounted } from "vue";
import { getUsers } from "../../api/admin";
const route = useRoute();
const users = ref([]);

const getData = (page) => {
    getUsers(parseInt(page)).then((res) => {
        if (res.success) {
            users.value = res.users;
        }
    });
};

onMounted(() => {
    getData(route.params.page);
});
onBeforeRouteUpdate(async (to, from) => {
    getData(to.params.page);
});
</script>
<template>
    <div class="flex flex-col">
        <button
            v-for="user in users"
            :key="user.id"
            class="bg-ui first:rounded-t-xl last:rounded-b-xl px-4 py-2 not-last:border-b-2 border-ui-2 flex flex-col items-start hover:bg-ui-2 active:bg-ui-3">
            <h2 class="text-lg">{{ user.name }} <span class="text-red text-sm" v-if="user.isAdmin">админ</span></h2>
            <div class="text-tx-2">{{ user.email }}</div>
        </button>
    </div>
    <!-- <table class="w-full">
        <thead>
            <tr>
                <th>Имя</th>
                <th>Почта</th>
                <th>Администратор</th>
            </tr>
        </thead>
        <tbody class="">
            <tr v-for="user in users" :key="user.id" class="">
                <td>{{ user.name }}</td>
                <td>{{ user.email }}</td>
                <td>
                    <input type="checkbox" v-model="user.isAdmin" disabled />
                </td>
                <td>
                    <button v-if="!user.isAdmin" class="btn btn-danger">Сделать администратором</button>
                </td>
                <td>
                    <button class="btn btn-danger">Удалить</button>
                </td>
            </tr>
        </tbody>
    </table> -->
</template>

<style>
@import "../../style.css";

tr {
    @apply p-4;
}
td {
    @apply text-center py-2;
}
</style>
