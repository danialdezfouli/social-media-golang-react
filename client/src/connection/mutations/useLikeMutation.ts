import { useLike } from "contexts/LikeContext";
import { api } from "connection/api";
import { QUERY_KEYS } from "connection/QueryKeys";
import { IPost } from "connection/types";
import { useMutation, useQueryClient } from "react-query";

interface ILikeResponse {
  liked: boolean;
}

const likeRequest = (post: IPost) => {
  return api.post<{}, ILikeResponse>("/post/" + post.post_id + "/like");
};

export default function useLikeMutation(
  post: IPost,
  source: IPost,
  username: string
) {
  const { isLiked, removeLike, addLike } = useLike();
  const queryClient = useQueryClient();

  return useMutation(likeRequest, {
    async onMutate() {
      await queryClient.cancelQueries(QUERY_KEYS.TIMELINE);
      await queryClient.cancelQueries([QUERY_KEYS.POST, post.post_id + ""]);
      await queryClient.cancelQueries([QUERY_KEYS.PROFILE_TIMELINE, username]);

      if (isLiked(source.post_id)) {
        removeLike(source.post_id);
      } else {
        addLike(source.post_id);
      }
    },

    // onError: (err, newTodo, context: any) => {},
    onSettled: () => {
      queryClient.invalidateQueries(QUERY_KEYS.TIMELINE);
      queryClient.invalidateQueries([QUERY_KEYS.PROFILE_TIMELINE, username]);
      queryClient.invalidateQueries([QUERY_KEYS.POST, post.post_id + ""]);
    },
  });
}
