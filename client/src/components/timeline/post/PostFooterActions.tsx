import useLikeMutation from "connection/mutations/useLikeMutation";
import { IPost } from "connection/types";
import { useLike } from "contexts/LikeContext";
import React, { useMemo } from "react";
import { useTranslation } from "react-i18next";
import { RiChat1Line, RiHeart3Fill, RiHeart3Line } from "react-icons/ri";
import { Link } from "react-router-dom";
import { postDateTime } from "utils/dates";

type PostActionsProps = {
  post: IPost;
  parent?: IPost;
  linkedPost?: IPost;
  largeLayout: boolean;
};

export default function PostActions({
  post,
  largeLayout = true,
  parent,
  linkedPost,
}: PostActionsProps) {
  const { isLiked } = useLike();
  const { t } = useTranslation();
  const source = useMemo(() => {
    if (post.post_type === "repost" && parent) {
      return parent;
    }
    return post;
  }, [parent, post]);

  const likeMutate = useLikeMutation(
    linkedPost || post,
    source,
    post.profile_username
  );

  const handleLike = (e: React.MouseEvent) => {
    e.stopPropagation();
    likeMutate.mutate(source);
  };

  const formattedDate = useMemo(
    () => postDateTime(source.created_at),
    [source.created_at]
  );

  const liked = useMemo(
    () => isLiked(source.post_id),
    [isLiked, source.post_id]
  );

  return (
    <footer className="post-footer">
      {largeLayout && (
        <Link className="date" to={"/post/" + source.post_id}>
          {formattedDate}
        </Link>
      )}

      {largeLayout && (
        <div className="charts">
          <div>
            <b>{source.favorites_count}</b>
            <span>{t("post.like")}</span>
          </div>
          <div>
            <b>{source.replies_count}</b>
            <span>{t("post.replies")}</span>
          </div>
        </div>
      )}
      <div className="actions">
        <button
          className={"like " + (liked ? "en" : "")}
          title={t("post.like")}
          onClick={handleLike}
          disabled={likeMutate.isLoading}
        >
          <i>{liked ? <RiHeart3Fill /> : <RiHeart3Line />}</i>
          {!largeLayout && <span>{source.favorites_count}</span>}
        </button>
        <button
          className="reply"
          title={t("post.reply")}
          onClick={(e) => e.stopPropagation()}
        >
          <i>
            <RiChat1Line />
          </i>
          {!largeLayout && <span>{source.replies_count}</span>}
        </button>
        {/* <button
          className={"repost " + (post.reposted ? "en" : "")}
          title={t('post.repost')}
          onClick={(e) => e.stopPropagation()}
        >
          <i><RiRepeat2Line /></i>
          <span>{source.repost_count}</span>
        </button>
        <button
          className="quote"
          title={t('post.quote')}
          onClick={(e) => e.stopPropagation()}
        >
          <i><RiChatQuoteLine /></i>
          <span>{source.quote_count}</span>
        </button> */}
      </div>
    </footer>
  );
}
