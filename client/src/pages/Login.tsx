import LoginForm from "components/LoginForm";
import Text from "components/elements/Text";
import { useTranslation } from "react-i18next";
import { useEffect } from "react";
import tokenService from "utils/localStorageService";
import { useNavigate } from "react-router-dom";

export default function Login() {
  const { t } = useTranslation();
  const navigate = useNavigate();

  useEffect(() => {
    if (tokenService.hasToken()) {
      navigate("/home");
    }
  }, [navigate]);

  return (
    <div className="auth-page py-6">
      <div className="container-md">
        <Text as="h1" size="3xl" align="center" weight="bold" className="mb-2">
          {t("login.welcome")}
        </Text>
        <Text as="h2" size="xl" align="center" weight="bold" className="mb-4">
          {t("login.login")}
        </Text>
        <LoginForm />
      </div>
    </div>
  );
}
