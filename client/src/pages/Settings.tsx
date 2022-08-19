import PageHeader from "components/elements/PageHeader";
import SettingForm from "components/settings/SettingForm";
import { useTranslation } from "react-i18next";

export default function Settings() {
  const { t } = useTranslation();
  return (
    <div>
      <PageHeader>{t("settings.title")}</PageHeader>
      <SettingForm />
    </div>
  );
}
