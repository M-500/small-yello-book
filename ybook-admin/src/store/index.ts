import { defineStore } from "pinia";

export const useUserStore = defineStore("user", {
  state: () => ({
    token: null,
  }),
  actions: {
    setToken(token) {
      this.token = token;
      localStorage.setItem("token", token);
    },
    loadToken() {
      this.token = localStorage.getItem("token");
    },
    clearToken() {
      this.token = null;
      localStorage.removeItem("token");
    },
  },
});
