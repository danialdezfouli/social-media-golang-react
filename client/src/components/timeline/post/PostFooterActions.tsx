import useLikeMutation from "connection/mutations/useLikeMutation";
import { IPost } from "connection/types";
import { useMemo } from "react";
import { useTranslation } from "react-i18next";
import { RiChat1Line, RiHeart3Fill, RiHeart3Line } from "react-icons/ri";
import { Link } from "react-router-dom";
import { postDateTime } from "utils/dates";

type PostActionsProps = {
  post: IPost;
  largeLayout: boolean;
};

export default function PostActions({
  post,
  largeLayout = true,
}: PostActionsProps) {
  const { t } = useTranslation();

  const likeMutate = useLikeMutation(post);

  const handleLike = (e: React.MouseEvent) => {
    e.stopPropagation();
    likeMutate.mutate(post);
  };

  const formattedDate = useMemo(() => {
    return postDateTime(post.created_at);
  }, [post.created_at]);

  return (
    <footer className="post-footer">
      {largeLayout && (
        <Link className="date" to={"/post/" + post.post_id}>
          {formattedDate}
        </Link>
      )}

      {largeLayout && (
        <div className="charts">
          <div>
            <b>{post.favorites_count}</b>
            <span>{t("post.like")}</span>
          </div>
          <div>
            <b>{post.replies_count}</b>
            <span>{t("post.replies")}</span>
          </div>
        </div>
      )}
      <div className="actions">
        <button
          className={"like " + (post.liked ? "en" : "")}
          title={t("post.like")}
          onClick={handleLike}
          disabled={likeMutate.isLoading}
        >
          <i>{post.liked ? <RiHeart3Fill /> : <RiHeart3Line />}</i>
          {!largeLayout && <span>{post.favorites_count}</span>}
        </button>
        <button
          className="reply"
          title={t("post.reply")}
          onClick={(e) => e.stopPropagation()}
        >
          <i>
            <RiChat1Line />
          </i>
          {!largeLayout && <span>{post.replies_count}</span>}
        </button>
        {/* <button
          className={"repost " + (post.reposted ? "en" : "")}
          title={t('post.repost')}
          onClick={(e) => e.stopPropagation()}
        >
          <i><RiRepeat2Line /></i>
          <span>{post.repost_count}</span>
        </button>
        <button
          className="quote"
          title={t('post.quote')}
          onClick={(e) => e.stopPropagation()}
        >
          <i><RiChatQuoteLine /></i>
          <span>{post.quote_count}</span>
        </button> */}
      </div>
    </footer>
  );
}
