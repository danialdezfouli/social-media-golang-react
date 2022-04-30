import { AxiosError } from "axios";
import { api } from "connection/api";
import { QueryFunctionContext, useQuery } from "react-query";
import { QUERY_KEYS } from "../QueryKeys";
import { IPost } from "../types";

const fetchProfileTimeline = ({ signal, queryKey }: QueryFunctionContext) => {
  return api
    .get("/profile/" + queryKey[1] + "/timeline", { signal })
    .then((res) => res.data);
};

export interface IProfileTimeline {
  posts: IPost[];
  parents: IPost[];
}

export default function useProfileTimelineQuery(id: string | undefined) {
  return useQuery<IProfileTimeline, AxiosError>(
    [QUERY_KEYS.PROFILE_TIMELINE, id],
    fetchProfileTimeline,
    {}
  );
}
