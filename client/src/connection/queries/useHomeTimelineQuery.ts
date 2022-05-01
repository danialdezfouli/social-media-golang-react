import { AxiosError } from "axios";
import { api } from "connection/api";
import { IProfile } from "connection/types";
import { QueryFunctionContext, useQuery, UseQueryOptions } from "react-query";
import { QUERY_KEYS } from "../QueryKeys";
import { IPost } from "../types";

function fetchHomeTimeline({ signal }: QueryFunctionContext) {
  return api.get("/timeline", { signal }).then((res) => res.data);
}

export interface ITimeline {
  posts: IPost[];
  parents: Record<string, IPost>;
  suggested_profiles: IProfile[];
}

export default function useHomeTimelineQuery(
  options?: UseQueryOptions<ITimeline, AxiosError>
) {
  return useQuery<ITimeline, AxiosError>(
    QUERY_KEYS.TIMELINE,
    fetchHomeTimeline,
    options
  );
}
