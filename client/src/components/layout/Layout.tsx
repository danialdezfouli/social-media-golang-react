import Button from "components/elements/Button";
import { Spinner } from "components/elements/Spinner";
import Text from "components/elements/Text";
import useMeQuery from "connection/queries/useMeQuery";
import { useAuth } from "contexts/AuthContext";
import { LikeProvider } from "contexts/LikeContext";
import { useEffect } from "react";
import { useTranslation } from "react-i18next";
import { Outlet, useNavigate } from "react-router-dom";
import "./Layout.css";
import SidebarMenu from "./SidebarMenu";

export default function Layout() {
  const { t } = useTranslation();
  const { user, setUser } = useAuth();
  const { error, data, refetch } = useMeQuery();
  const navigate = useNavigate();

  useEffect(() => {
    setUser(data || null);
  }, [data, setUser]);

  useEffect(() => {
    if (
      (error as any)?.response?.status === 401 ||
      (error as any)?.response?.status === 400
    ) {
      navigate("/");
    }
  }, [error, navigate]);

  if (!user) {
    return (
      <div className="fixed left-0 top-0 w-full h-full flex items-center justify-center">
        {error && (error as any)?.message === "Network Error" ? (
          <div className="flex flex-col items-center justify-center">
            <Text size="xl" className="p-4">
              {t("error.network")}
            </Text>
            <Button onClick={() => refetch()}>{t("actions.refetch")}</Button>
          </div>
        ) : (
          <div>
            <Text size="xl" className="p-4">
              {t("actions.loading")}
            </Text>
            <Spinner />
          </div>
        )}
      </div>
    );
  }

  return (
    <div className="container">
      <main className="app-layout">
        <section className="app-layout__content">
          <LikeProvider>
            <Outlet />
          </LikeProvider>
        </section>
        <aside className="app-layout__sidebar">
          <SidebarMenu />
        </aside>
      </main>
    </div>
  );
}
