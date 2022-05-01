import { Spinner } from "components/elements/Spinner";
import PostItem from "components/timeline/post/Post";
import usePostQuery from "connection/queries/usePostQuery";
import { IPost } from "connection/types";
import { useLike } from "contexts/LikeContext";
import { useMemo } from "react";
import { useParams } from "react-router-dom";
import "./PostPage.css";

export default function PostPage() {
  const params = useParams<{ id: string }>();
  const { setLikes } = useLike();
  const { data, isLoading } = usePostQuery(params.id!, {
    onSuccess: (data) => {
      setLikes([data.post, ...Object.values(data.parents)]);
    },
  });

  const post = data?.post;

  const isTypeReplying = post?.post_type === "reply";

  const parentsList = useMemo<IPost[]>(() => {
    if (!data || data.post.post_type !== "reply") return [];

    const post = data.post;
    const parents = data.parents;
    const result: IPost[] = [];

    if (post?.parent_id && parents[post.parent_id]) {
      result.push(parents[post.parent_id]);
    }

    while (result[0] && result[0].parent_id) {
      if (result[0].post_type === "quote") {
        break;
      }

      result.unshift(parents[result[0].parent_id]);
    }

    return result;
  }, [data]);

  return (
    <section className="post-page">
      {isLoading && (
        <div className="p-6 text-blue-600">
          <Spinner />
        </div>
      )}

      {data && isTypeReplying && (
        <div className="replied-to-parents">
          {parentsList.map((p, i) => (
            <PostItem
              key={p.post_id}
              post={p}
              linkedPost={post}
              replyLine={i === 0 ? "down" : "both"}
              parent={data.parents[p.parent_id]}
            />
          ))}
        </div>
      )}

      {post && (
        <PostItem
          key={post.post_id}
          post={post}
          replyLine={isTypeReplying && "up"}
          isFullPost
          parent={data.parents[post.parent_id]}
        />
      )}

      {data && (
        <div className="replies-list">
          {data.replies.map((p, i) => (
            <PostItem
              key={p.post_id}
              post={p}
              parent={post}
              linkedPost={post}
            />
          ))}
        </div>
      )}
    </section>
  );
}
