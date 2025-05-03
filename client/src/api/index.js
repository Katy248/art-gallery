import { useAuthStore } from "../stores/auth";
import { useRouter } from "vue-router";

const runFetch = (url, method, body) => {
    const bodyStr = JSON.stringify(body);
    let headers = {};

    const authStore = useAuthStore();
    if (authStore.isAuthenticated) {
        console.log("Authenticated");
        console.log(authStore.token);
        headers.Authorization = `Bearer ${authStore.token}`;
        console.log(headers);
    }

    return fetch(url, {
        headers: headers,
        method: method,
        body: bodyStr,
    }).then((response) => {
        return response.json();
    });
};

const authenticate = (email, password) => {
    runFetch("/api/auth", "POST", { email: email, password: password }).then((r) => {
        const token = r.token;

        const store = useAuthStore();
        store.authenticate(token);

        // useRouter().push("/");
    });
};

const getProfileInfo = () => {
    return runFetch("/api/user/get", "POST", {});
};

export { authenticate, getProfileInfo };
