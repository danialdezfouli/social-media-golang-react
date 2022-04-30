import { QUERY_KEYS } from "connection/QueryKeys";
import { useMutation, useQueryClient } from "react-query";
import { api } from "./../api";

interface IFollowDTO {
  user_id: number;
  username: string;
  type: "follow" | "unfollow";
}

interface IFollowResponse {
  data: {};
}

const followRequest = (data: IFollowDTO) => {
  if (data.type === "follow") {
    return api.post<IFollowDTO, IFollowResponse>("follow/" + data.user_id);
  } else {
    return api.delete<IFollowDTO, IFollowResponse>("unfollow/" + data.user_id);
  }
};

export function useFollowMutation(username: string | undefined) {
  const queryClient = useQueryClient();
  let key = [QUERY_KEYS.PROFILE, username];

  return useMutation(followRequest, {
    onMutate: async (data) => {
      await queryClient.cancelQueries(key);

      const previousProfile = queryClient.getQueryData(key);

      return {
        previousProfile,
      };
    },

    onError: (err, newTodo, context: any) => {
      queryClient.setQueryData(key, context.previousProfile);
    },
    onSettled: () => {
      queryClient.invalidateQueries(key);
    },
  });
}
