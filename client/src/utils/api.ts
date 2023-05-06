import axios from "axios";

export const $axios = axios.create({
  baseURL: `${process.env.NEXT_PUBLIC_APP_API_URL}`,
  headers: { 'Content-Type': 'application/json' },
  responseType: 'json',
  withCredentials: true
});
