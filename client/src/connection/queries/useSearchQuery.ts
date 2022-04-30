import { AxiosError } from "axios";
import { QUERY_KEYS } from "connection/QueryKeys";
import { IProfile } from "connection/types";
import { QueryFunctionContext, useQuery } from "react-query";
import { api } from "./../api";

interface ISearchResponse {
  profiles: IProfile[];
}

function fetchSearch({ signal, queryKey }: QueryFunctionContext) {
  return api
    .get("/search", {
      signal,
      params: { q: queryKey[1] },
    })
    .then((res) => res.data);
}

export default function useSearchQuery(query: string | null) {
  return useQuery<ISearchResponse, AxiosError>(
    [QUERY_KEYS.SEARCH, query],
    fetchSearch,
    {
      enabled: Boolean(query && query.length > 0),
      // cacheTime: 0,
    }
  );
}
