import { useAuthStore } from "../stores/auth";
const runFetch = (url, method, body, inJson = true) => {
    let headers = {};

    if (inJson) {
        body = JSON.stringify(body);
        headers["Content-Type"] = "application/json";
    }

    if (method === "GET" || method === "HEAD") {
        body = null;
    }

    const authStore = useAuthStore();
    if (authStore.isAuthenticated) {
        headers.Authorization = `Bearer ${authStore.token}`;
    }

    return fetch(url, {
        headers: headers,
        method: method,
        body: body,
    })
        .then((response) => {
            return response.json();
        })
        .catch((err) => {
            console.error(err);
        });
};
export { runFetch };
