import axios from "axios";
import tokenService from "utils/localStorageService";

export const HOST_URL = process.env.REACT_APP_API_URL;

export const public_api = axios.create({
  baseURL: HOST_URL,
  headers: {
    Accept: "application/json",
  },
});

export const api = axios.create({
  baseURL: HOST_URL,
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
});

api.interceptors.request.use(function (config) {
  config.headers = {
    ...config.headers,
    Authorization: "Bearer " + tokenService.getToken(),
  };

  return config;
});

let isAlreadyFetchingAccessToken = false;

api.interceptors.response.use(
  (res) => res,
  async (err) => {
    const status = err.response?.status;
    const config = err.config;

    if (status === 401 && !config._retry) {
      config._retry = true;

      try {
        if (!isAlreadyFetchingAccessToken) {
          isAlreadyFetchingAccessToken = true;
          console.log("refresh");
          const token = await refreshToken();
          isAlreadyFetchingAccessToken = false;
          tokenService.setToken(token);
          config.headers.Authorization = "Bearer " + token;
        }

        console.log("trying again");
        return axios.request(config);
      } catch (refreshErr) {
        tokenService.removeToken();
        return Promise.reject(refreshErr);
      }
    }

    return Promise.reject(err);
  }
);

function refreshToken() {
  return public_api
    .post("/auth/refresh", null, {
      withCredentials: true,
    })
    .then(({ data }) => data.access_token);
}
