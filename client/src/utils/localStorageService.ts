const TOKEN_KEY = "token";

const tokenService = {
  getToken() {
    return localStorage.getItem(TOKEN_KEY);
  },
  hasToken() {
    return Boolean(this.getToken());
  },
  setToken(value: string) {
    return localStorage.setItem(TOKEN_KEY, value);
  },
  removeToken() {
    return localStorage.removeItem(TOKEN_KEY);
  },
};

export default tokenService;
