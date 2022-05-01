import { AxiosError } from "axios";
import { QueryFunctionContext, useQuery, UseQueryOptions } from "react-query";
import { api } from "./../api";
import { QUERY_KEYS } from "./../QueryKeys";
import { IProfile } from "./../types";

function fetchProfile({ queryKey, signal }: QueryFunctionContext) {
  return api.get("/profile/" + queryKey[1], { signal }).then((res) => res.data);
}

export default function useProfileQuery(id: string | undefined) {
  return useQuery<IProfile, AxiosError>(
    [QUERY_KEYS.PROFILE, id],
    fetchProfile,
    {
      enabled: typeof id === "string",
    }
  );
}
