import React from "react";
import { useNavigate } from "react-router-dom";
import { noOp } from "utils/common";
import tokenService from "utils/localStorageService";
import { IUser } from "./types";

const AuthContext = React.createContext<IAuthContext>({
  user: null,
  logout: noOp,
  loading: true,
  error: null,
  setError: noOp,
  setLoading: noOp,
  setUser: noOp,
});

interface IAuthContext {
  user: IUser | null;
  logout(): void;
  loading: boolean;
  error: string | null;
  setError(error: string | null): void;
  setLoading(loading: boolean): void;
  setUser(user: IUser | null): void;
}

export function useAuth() {
  const context = React.useContext(AuthContext);
  return context;
}

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = React.useState<IUser | null>(null);
  const [loading, setLoading] = React.useState(true);
  const [error, setError] = React.useState<string | null>(null);
  const navigate = useNavigate();

  const logout = () => {
    setUser(null);
    tokenService.removeToken();
    navigate("/");
  };

  React.useEffect(() => {
    setLoading(false);
  }, []);

  const context: IAuthContext = {
    user,
    logout,
    loading,
    error,
    setError,
    setLoading,
    setUser,
  };

  return (
    <AuthContext.Provider value={context}>{children}</AuthContext.Provider>
  );
}
