import { api } from "connection/api";
import { QUERY_KEYS } from "connection/QueryKeys";
import { IPost } from "connection/types";
import { useMutation, useQueryClient } from "react-query";
import { PostTypesShape } from "./../types";

export interface ICreatePostDTO {
  body: string;
  type: PostTypesShape;
  reply_to?: number;
}

export interface ICreatePostResponse {
  data: {
    post: IPost;
  };
}

const createPostRequest = (data: ICreatePostDTO) => {
  return api.post<ICreatePostDTO, ICreatePostResponse>("/post", data);
};

export default function usePostMutation(
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
