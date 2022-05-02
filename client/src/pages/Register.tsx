import Text from "components/elements/Text";
import RegisterForm from "components/RegisterForm";
import { useTranslation } from "react-i18next";

export default function Register() {
  const { t } = useTranslation();

  return (
    <div className="auth-page py-6">
      <div className="container-md">
        <Text as="h1" size="3xl" align="center" weight="bold" className="mb-2">
          {t("login.welcome")}
        </Text>
        <Text as="h2" size="xl" align="center" weight="bold" className="mb-4">
          {t("signup.subtitle")}
        </Text>
        <RegisterForm />
      </div>
    </div>
  );
}
