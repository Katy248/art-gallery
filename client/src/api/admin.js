import { runFetch } from "./shared";

const PAGE_SIZE = 20;

const getUsers = (page) => {
  return runFetch(`/api/user/get-all`, "POST", {
    page: page,
    pageSize: PAGE_SIZE,
  });
};
const deleteUser = (userId) => {
  return runFetch(`/api/user/delete`, "DELETE", { id: userId });
};

export { getUsers, deleteUser };
