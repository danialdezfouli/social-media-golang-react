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

function mapPostToggleLike(p: IPost): IPost {
  return {
    ...p,
    liked: !p.liked,
    favorites_count: p.liked ? p.favorites_count - 1 : p.favorites_count + 1,
  };
}

export default function useLikeMutation(post: IPost) {
  const queryClient = useQueryClient();

  return useMutation(likeRequest, {
    async onMutate() {
      await queryClient.cancelQueries(QUERY_KEYS.TIMELINE);
      await queryClient.cancelQueries([QUERY_KEYS.POST, post.post_id]);
      const previousTimeline = queryClient.getQueryData(QUERY_KEYS.TIMELINE);
      const previousPost = queryClient.getQueryData([
        QUERY_KEYS.POST,
        post.post_id,
      ]);

      queryClient.setQueryData(QUERY_KEYS.TIMELINE, (timeline: any) => {
        if (!timeline) return;
        return {
          ...timeline,
          posts: timeline?.posts.map(function (p: IPost): IPost {
            if (p.post_id === post.post_id) {
              return mapPostToggleLike(p);
            } else {
              return p;
            }
          }),
        };
      });

      queryClient.setQueryData([QUERY_KEYS.POST, post.post_id], (old: any) => {
        if (!old) return;
        return {
          ...old,
          post: mapPostToggleLike(old.post),
        };
      });

      return {
        previousTimeline,
        previousPost,
      };
    },

    onError: (err, newTodo, context: any) => {
      queryClient.setQueryData(QUERY_KEYS.TIMELINE, context.previousTimeline);
      queryClient.setQueryData(
        [QUERY_KEYS.POST, post.post_id],
        context.previousTimeline
      );
    },
    onSettled: () => {
      queryClient.invalidateQueries(QUERY_KEYS.TIMELINE);
      queryClient.invalidateQueries(QUERY_KEYS.POST);
    },
  });
}
