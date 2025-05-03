import { defineStore } from "pinia";

const TOKEN_KEY = "jwt_token";

export const useAuthStore = defineStore("auth", {
    state: () => ({ token: localStorage.getItem(TOKEN_KEY) }),
    getters: {
        isAuthenticated() {
            return this.token && this.token !== "";
        },
    },
    actions: {
        authenticate(token) {
            this.token = token;
            localStorage.setItem(TOKEN_KEY, token);
        },
        logout() {
            this.token = "";
            localStorage.setItem(TOKEN_KEY, "");
        },
    },
});
