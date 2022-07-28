import Text from "components/elements/Text";
import { IPost } from "connection/types";
import { useTranslation } from "react-i18next";
import { Link, useNavigate } from "react-router-dom";
import PostFooterActions from "./PostFooterActions";
import PostHeaderProfile from "./PostHeaderProfile";
import "./Post.css";
import { TReplyLine } from "./types";
import { useAuth } from "contexts/AuthContext";
import { useMemo } from "react";

type PostItemProps = {
  post: IPost;
  linkedPost?: IPost;
  parent: IPost | undefined;
  showActions?: boolean;
  replyLine?: TReplyLine;
  isFullPost?: boolean;
};

export default function PostItem(props: PostItemProps) {
  const {
    post,
    parent,
    linkedPost,
    isFullPost = false,
    showActions = true,
    replyLine,
  } = props;

  const { t } = useTranslation();
  const navigate = useNavigate();
  const { user } = useAuth();

  const { post_type } = post;
  const source = post_type === "repost" ? parent : post;
  const sourceContent = source?.content || "";

  const contentLines = useMemo(() => {
    return sourceContent.trim().split("\n");
  }, [sourceContent]);

  const handleClick = (e: React.MouseEvent) => {
    e.stopPropagation();

    if (!isFullPost) {
      const url = `/post/${post.post_id}`;
      navigate(url);
    }
  };

  const stopPropagation = (e: React.MouseEvent) => e.stopPropagation();
  const profileUrl = "/profile/" + post.profile_username;

  if (!source) {
    return (
      <div>
        <Text>Error in post {post.post_id}</Text>
      </div>
    );
  }

  return (
    <article
      className={
        "post-item" + (isFullPost ? " post-item__full" : " post-item__link")
      }
      onClick={handleClick}
      id={`post-${post.post_id}`}
    >
      {post_type === "repost" && (
        <Link
          to={profileUrl}
          className="post-item__headline-repost"
          onClick={stopPropagation}
        >
          {post.profile_name} {t("post.reposted")}
        </Link>
      )}

      <PostHeaderProfile
        name={source.profile_name}
        linkedPost={linkedPost}
        url={"/profile/" + source.profile_username}
        username={source.profile_username}
        createdAt={source.created_at}
        isLinked={showActions}
        post={post}
        hasPopoverActions={Boolean(user && user.id === post.user_id)}
        showDate={!isFullPost}
        replyLine={replyLine}
      />

      {post_type === "reply" && parent?.user_id && (
        <Text className="post-item__headline-reply" size="sm">
          {t("post.replying-to")}{" "}
          <Link
            to={`/profile/${parent.profile_username}`}
            onClick={stopPropagation}
          >
            {parent.profile_name}
          </Link>
        </Text>
      )}

      <div className="post-content">
        {contentLines.map((line, i) => (
          <p key={i}>{line}</p>
        ))}
      </div>

      {post_type === "quote" &&
        (parent ? (
          <div className="quote-body">
            <PostItem post={parent} parent={undefined} showActions={false} />
          </div>
        ) : (
          <Text className="quote-deleted">{t("post.deleted")}</Text>
        ))}

      {showActions && (
        <PostFooterActions
          post={post}
          parent={parent}
          linkedPost={linkedPost}
          largeLayout={isFullPost}
        />
      )}
    </article>
  );
}
