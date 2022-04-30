import { IProfile } from "connection/types";
import { QUERY_KEYS } from "../QueryKeys";
import { api } from "connection/api";
import { QueryFunctionContext, useQuery } from "react-query";
import { IPost } from "../types";
import { AxiosError } from "axios";

function fetchHomeTimeline({ signal }: QueryFunctionContext) {
  return api.get("/timeline", { signal }).then((res) => res.data);
}

export interface ITimeline {
  posts: IPost[];
  parents: Record<string, IPost>;
  suggested_profiles: IProfile[];
}

export default function useHomeTimelineQuery() {
  return useQuery<ITimeline, AxiosError>(
    QUERY_KEYS.TIMELINE,
    fetchHomeTimeline,
    {}
  );
}
