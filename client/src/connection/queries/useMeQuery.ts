import { AxiosError } from "axios";
import { QUERY_KEYS } from "./../QueryKeys";
import { api } from "connection/api";
import { IUser } from "contexts/types";
import { QueryFunctionContext, useQuery } from "react-query";

const fetchUser = ({ signal }: QueryFunctionContext) => {
  return api.get("/auth/me", { signal }).then(({ data }) => data);
};

export default function useMeQuery() {
  return useQuery<IUser, AxiosError>(QUERY_KEYS.ME, fetchUser, {
    retry: 1,
  });
}
