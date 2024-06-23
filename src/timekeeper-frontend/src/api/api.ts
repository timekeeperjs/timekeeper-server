import axios from "axios";

const api = axios.create({
  baseURL: import.meta.env.VITE_TIMEKEEPER_BACKEND || "http://localhost:8080",
  headers: {
    "Content-Type": "application/json",
  },
});

export default api;
