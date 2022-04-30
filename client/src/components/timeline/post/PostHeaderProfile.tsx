import { useMemo } from "react";
import { Link } from "react-router-dom";
import { humanTime } from "utils/dates";
import { TReplyLine } from "./types";

type PostProfileHeaderProps = {
  url: string;
  name: string;
  username: string;
  createdAt: string;
  isLinked: boolean;
  showDate: boolean;
  replyLine: TReplyLine | undefined;
};

export default function PostProfileHeader(props: PostProfileHeaderProps) {
  const {
    url = "",
    name = "",
    username = "",
    createdAt = "",
    isLinked = true,
    showDate = true,
    replyLine,
  } = props;

  const date = useMemo(
    () => showDate && humanTime(createdAt, "fa_IR"),
    [showDate, createdAt]
  );
  const stopPropagation = (e: React.MouseEvent) => e.stopPropagation();

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
    </header>
  );
}
