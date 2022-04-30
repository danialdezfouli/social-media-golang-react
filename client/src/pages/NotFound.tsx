import { useTranslation } from "react-i18next";

export default function NotFound() {
  const { t } = useTranslation();
  return (
    <div className="p-8">
      <h1>{t("error.page-not-found")}</h1>
    </div>
  );
}
