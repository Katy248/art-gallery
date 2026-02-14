import { useFetch } from "@vueuse/core";
import { useAuthStore } from "../stores/auth";
import { runFetch } from "./shared";

const authenticate = (email, password) => {
  return runFetch("/api/auth", "POST", {
    email: email,
    password: password,
  }).then((r) => {
    if (!r.success) {
      return { success: false };
    }
    const token = r.token;
    const store = useAuthStore();
    store.authenticate(token);
    return { success: true };
  });
};

const register = (email, name, password) => {
  return runFetch("/api/user/create", "POST", {
    email: email,
    password: password,
    name: name,
  }).catch((err) => {
    console.error("Error while registering user:", err);
    return { success: false, error: "Что-то пошло не так: " + err };
  });
};

const getProfileInfo = () => {
  return runFetch("/api/user/get", "POST", {});
};
const getUserInfo = (userId) => {
  userId = parseInt(userId);
  return runFetch("/api/user/get", "POST", { id: userId });
};

const createPost = (description, file) => {
  const data = new FormData();
  data.append("description", description);
  data.append("picture", file);
  return runFetch("/api/post/create", "POST", data, false);
};
const getPosts = (userId, page) => {
  userId = parseInt(userId);
  page = parseInt(page);
  return runFetch("/api/post/get-users", "POST", {
    userId: userId,
    page: page,
  });
};

const savePost = (postId) => {
  return runFetch("/api/post/save", "POST", { postId: postId });
};
const unsavePost = (postId) => {
  return runFetch("/api/post/unsave", "POST", { postId: postId });
};
const getSavedPosts = (userId, page) => {
  return runFetch("/api/post/get-saved", "POST", {
    userId: userId,
    page: page,
  });
};
const getAvatar = (userId) => {
  return runFetch(`/api/user/avatar/${userId}`, "GET", {});
};

const getAllPosts = (page) => {
  return runFetch("/api/post/all", "POST", { page: page });
};
const deletePost = (postId) => {
  return runFetch("/api/post/delete", "DELETE", { postId: postId });
};
const getPost = (postId) => {
  return runFetch("/api/post/get", "POST", { postId: postId });
};
const updatePost = (post) => {
  return runFetch("/api/post/update", "POST", {
    postId: post.id,
    description: post.description,
    warningMessage: post.warningMessage,
  });
};

const fetchPost = (id) => {
  const postId = parseInt(id);
  const authStore = useAuthStore();
  if (!authStore.isAuthenticated) {
    console.error("User is not authenticated to get this post info");
    return;
  }
  const token = authStore.token;

  return useFetch("/api/post/get", {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })
    .post({ postId: postId })
    .json();
};

export {
  authenticate,
  createPost,
  deletePost,
  fetchPost,
  getAllPosts,
  getAvatar,
  getPost,
  getPosts,
  getProfileInfo,
  getSavedPosts,
  getUserInfo,
  register,
  savePost,
  unsavePost,
  updatePost,
};
