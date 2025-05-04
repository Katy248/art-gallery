import { defineStore } from "pinia";

const TOKEN_KEY = "jwt_token";

export const useAuthStore = defineStore("auth", {
    state: () => ({ token: localStorage.getItem(TOKEN_KEY) }),
    getters: {
        isAuthenticated() {
            return this.token && this.token !== "";
        },
        authData() {
            if (!this.isAuthenticated) return {};
            const arrayToken = this.token.split(".");
            return JSON.parse(atob(arrayToken[1]));
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
