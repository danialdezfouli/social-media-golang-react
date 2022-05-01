import { AxiosError } from "axios";
import { api } from "connection/api";
import { QueryFunctionContext, useQuery, UseQueryOptions } from "react-query";
import { QUERY_KEYS } from "../QueryKeys";
import { IPost } from "../types";

const fetchPost = ({ signal, queryKey }: QueryFunctionContext) => {
  return api.get("/post/" + queryKey[1], { signal }).then((res) => res.data);
};

interface IPostResponse {
  post: IPost;
  parents: Record<string, IPost>;
  replies: IPost[];
}

export default function usePostQuery(
  id: string,
  options?: UseQueryOptions<IPostResponse, AxiosError>
) {
  return useQuery<IPostResponse, AxiosError>(
    [QUERY_KEYS.POST, id],
    fetchPost,
    options
  );
}
