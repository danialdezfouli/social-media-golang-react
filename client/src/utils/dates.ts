import {
  format,
  formatDistanceToNowStrict,
  parseISO,
} from "date-fns-jalali";
import { faIR } from "date-fns/locale";

export const humanTime = (datetime: string, locale = "fa_IR"): string => {
  const parsedDate = parseISO(datetime);
  let value = formatDistanceToNowStrict(parsedDate, {
    locale: locale === "fa_IR" ? faIR : undefined,
    addSuffix: true,
  });

  // value = value.replace("قبل", "پیش");

  return value;
};

export const postDateTime = (datetime: string, locale = "fa_IR"): string => {
  const parsedDate = parseISO(datetime);

  let value = format(parsedDate, "d MMMM yyyy · H:mm", {
    // locale: locale === "fa_IR" ? faIR : undefined,
  });

  return value;
};
