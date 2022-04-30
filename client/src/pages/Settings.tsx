import Text from "components/elements/Text";
import { useTranslation } from "react-i18next";

export default function Settings() {
  const { t } = useTranslation();
  return (
    <div className="p-10">
      <Text size="3xl">{t("settings.title")}</Text>
    </div>
  );
}
