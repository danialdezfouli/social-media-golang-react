import { Popover } from "@headlessui/react";
import { Spinner } from "components/elements/Spinner";
import usePostDeleteMutation from "connection/mutations/usePostDeleteMutation";
import { IPost } from "connection/types";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { RiMoreFill } from "react-icons/ri";
import { Link } from "react-router-dom";
import { humanTime } from "utils/dates";
import { TReplyLine } from "./types";

type PostProfileHeaderProps = {
  post: IPost;
  linkedPost?: IPost;
  url: string;
  name: string;
  username: string;
  createdAt: string;
  isLinked: boolean;
  showDate: boolean;
  hasPopoverActions: boolean;
  replyLine: TReplyLine | undefined;
};

export default function PostProfileHeader(props: PostProfileHeaderProps) {
  const {
    url = "",
    name = "",
    username = "",
    createdAt = "",
    isLinked = true,
    hasPopoverActions = false,
    showDate = true,
    replyLine,
    linkedPost,
    post,
  } = props;

  const [popover, setPopover] = useState(hasPopoverActions);

  const { t } = useTranslation();
  const date = showDate && humanTime(createdAt, "fa_IR");
  const postDelete = usePostDeleteMutation(
    linkedPost || post,
    post.profile_username
  );

  const stopPropagation = (e: React.MouseEvent) => e.stopPropagation();
  const handleDelete = () => {
    setPopover(false);
    postDelete.mutate({ post_id: post.post_id });
  };

  return (
    <header className="profile">
      {replyLine && (
        <div className="reply-line">
          {(replyLine === "up" || replyLine === "both") && (
            <span className="up" />
          )}

          {(replyLine === "down" || replyLine === "both") && (
            <span className="down" />
          )}
        </div>
      )}
      <div className="image"></div>
      <div className="wrap">
        {isLinked ? (
          <Link to={url} onClick={stopPropagation} className="name">
            {name}
          </Link>
        ) : (
          <div className="name">{name}</div>
        )}

        <div className="flex gap-1">
          {isLinked ? (
            <Link to={url} onClick={stopPropagation} className="username">
              @{username}
            </Link>
          ) : (
            <div className="username">@{username}</div>
          )}
          {showDate && (
            <>
              <span className="font-bold">·</span>
              <span className="date">{date}</span>
            </>
          )}
        </div>
      </div>

      {popover && (
        <Popover className="header-actions" onClick={stopPropagation}>
          <Popover.Button className="header-actions__btn">
            <RiMoreFill />
          </Popover.Button>
          <Popover.Panel className="header-actions__panel">
            <button onClick={handleDelete}>
              {t("post.delete")}
              {postDelete.isLoading && <Spinner />}
            </button>
          </Popover.Panel>
        </Popover>
      )}
    </header>
  );
}
