import classNames from "classnames";
import { useEffect, useRef, useState } from "react";
import { useTranslation } from "react-i18next";
import { RiSearchLine } from "react-icons/ri";
import "./SearchForm.css";

interface SearchFormProps {
  onSubmit(inputs: string): void;
  defaultValue: string | null;
}

export default function SearchForm({
  onSubmit,
  defaultValue,
}: SearchFormProps) {
  const { t } = useTranslation();
  const [focus, setFocus] = useState(false);
  const [value, setValue] = useState(defaultValue || "");
  const timer = useRef<NodeJS.Timeout>();

  useEffect(() => {
    return () => {
      if (timer.current) {
        clearTimeout(timer.current);
      }
    };
  }, []);

  function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    onSubmit(value.trim());
  }

  function handleInputChange(e: React.ChangeEvent<HTMLInputElement>) {
    setValue(e.target.value);
    if (timer.current) {
      clearTimeout(timer.current);
    }

    const DEBOUNCE_TIMER = 100;
    timer.current = setTimeout(() => {
      onSubmit(e.target.value.trim());
    }, DEBOUNCE_TIMER);
  }

  return (
    <div className="search-form">
      <form action="" onSubmit={handleSubmit}>
        <div className={classNames("wrap", { focus })}>
          <i className="icon">
            <RiSearchLine />
          </i>
          <input
            type="text"
            placeholder={t("search.input")}
            value={value}
            className="search-input"
            onChange={handleInputChange}
            onFocus={() => setFocus(true)}
            onBlur={() => setFocus(false)}
          />
        </div>
      </form>
    </div>
  );
}
