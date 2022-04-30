import { Spinner } from "components/elements/Spinner";
import SearchForm from "components/layout/SearchForm";
import ProfileCompact from "components/search/ProfileCompact";
import useSearchQuery from "connection/queries/useSearchQuery";
import { useCallback } from "react";
import { useTranslation } from "react-i18next";
import { useSearchParams } from "react-router-dom";

export default function Search() {
  let [searchParams, setSearchParams] = useSearchParams();
  let query = searchParams.get("q");
  const { t } = useTranslation();
  const { data, isLoading } = useSearchQuery(query);

  const onSubmit = useCallback(
    (value: string) => {
      if (value === query) return;
      let newSearchParams = new URLSearchParams();
      newSearchParams.set("q", value);
      setSearchParams(newSearchParams);
    },
    [query, setSearchParams]
  );

  return (
    <section className="search-page">
      <SearchForm onSubmit={onSubmit} defaultValue={query} />
      <div className="result">
        {data && data.profiles.length === 0 && (
          <h1 className="p-4">{t("search.no-result")}</h1>
        )}

        {isLoading && <Spinner />}
        <div className="search-profiles">
          {data?.profiles?.map((profile) => (
            <ProfileCompact key={profile.id} profile={profile} />
          ))}
        </div>
      </div>
    </section>
  );
}
