import { runFetch } from "./shared";

const PAGE_SIZE = 20;

const getUsers = (page) => {
    return runFetch(`/api/user/get-all`, "POST", { page: page, pageSize: PAGE_SIZE });
};

export { getUsers };
