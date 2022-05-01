import { IPost } from "connection/types";
import React, { useCallback, useState } from "react";
import { noOp } from "utils/common";

const LikeContext = React.createContext<ILikeContext>({
  isLiked: () => false,
  setLikes: noOp,
  addLike: noOp,
  removeLike: noOp,
});

interface ILikeContext {
  isLiked(id: number): boolean;
  setLikes(posts: IPost[]): void;
  addLike(id: number): void;
  removeLike(id: number): void;
}

export function useLike() {
  const context = React.useContext(LikeContext);
  return context;
}

export function LikeProvider({ children }: { children: React.ReactNode }) {
  const [state, setState] = useState<number[]>([]);

  const addLike = useCallback((id: number) => {
    setState((prev) => [...prev, id]);
  }, []);

  const removeLike = useCallback((id: number) => {
    setState((prev) => prev.filter((i) => i !== id));
  }, []);

  const isLiked = useCallback(
    (id: number) => {
      return state.includes(id);
    },
    [state]
  );

  const setLikes = (posts: IPost[]) => {
    const ids = posts.filter((post) => post.liked).map((post) => post.post_id);

    setState((state) => {
      return Array.from(new Set([...state, ...ids]));
    });
  };

  const context: ILikeContext = {
    setLikes,
    isLiked,
    addLike,
    removeLike,
  };

  return (
    <LikeContext.Provider value={context}>{children}</LikeContext.Provider>
  );
}
