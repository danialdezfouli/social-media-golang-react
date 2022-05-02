import { api } from "connection/api";
import { QUERY_KEYS } from "connection/QueryKeys";
import { IPost } from "connection/types";
import { useMutation, useQueryClient } from "react-query";

export interface IDeletePostDTO {
  post_id: IPost["post_id"];
}

export interface IDeletePostResponse {}

const createPostRequest = (data: IDeletePostDTO) => {
  return api.delete<IDeletePostDTO, IDeletePostResponse>(
    "/post/" + data.post_id
  );
};

export default function usePostDeleteMutation(
  post: IPost | undefined,
  username: string
) {
  const queryClient = useQueryClient();

  return useMutation(createPostRequest, {
    async onMutate() {
      await queryClient.cancelQueries(QUERY_KEYS.TIMELINE);
      await queryClient.cancelQueries([QUERY_KEYS.PROFILE_TIMELINE, username]);
      if (post) {
        await queryClient.cancelQueries([QUERY_KEYS.POST, post.post_id + ""]);
      }
    },

    onSettled: () => {
      queryClient.invalidateQueries(QUERY_KEYS.TIMELINE);
      queryClient.invalidateQueries([QUERY_KEYS.PROFILE_TIMELINE, username]);

      if (post) {
        queryClient.invalidateQueries([QUERY_KEYS.POST, post.post_id + ""]);
      }
    },
  });
}
