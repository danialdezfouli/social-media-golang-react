import classNames from "classnames";
import Button from "components/elements/Button";
import PageHeader from "components/elements/PageHeader";
import PostItem from "components/timeline/post/Post";
import { useFollowMutation } from "connection/mutations/useFollowMutation";
import useProfileQuery from "connection/queries/useProfileQuery";
import useProfileTimelineQuery, {
  IProfileTimeline,
} from "connection/queries/useProfileTimelineQuery";
import { IProfile } from "connection/types";
import { useAuth } from "contexts/AuthContext";
import { useLike } from "contexts/LikeContext";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { Link, useParams } from "react-router-dom";
import "./Profile.css";

export default function Profile() {
  const { t } = useTranslation();
  const params = useParams<"id">();
  const { data: profile } = useProfileQuery(params.id);
  const { setLikes } = useLike();
  const { data } = useProfileTimelineQuery(params.id, {
    onSuccess: (data) => {
      setLikes([...data.posts, ...Object.values(data.parents)]);
    },
    enabled: Boolean(profile),
  });

  return (
    <section className="profile-page">
      <PageHeader>
        {t("profile.title")}
        <span
          className={classNames("font-normal text-sm block", {
            "opacity-0": !data,
          })}
        >
          {data?.posts.length} {t("profile.post")}
        </span>
      </PageHeader>
      {!profile && <div className="p-8">{t("actions.loading")}</div>}
      {profile && <ProfileHeader profile={profile} />}
      {profile && <ProfilePosts data={data} />}
    </section>
  );
}

function ProfilePosts({ data }: { data?: IProfileTimeline }) {
  const { t } = useTranslation();

  return (
    <>
      <div className="profile-page__posts-header">{t("profile.posts")}</div>
      <div className="profile-page__posts">
        {data?.posts.map((post) => (
          <PostItem
            key={post.post_id}
            post={post}
            parent={post.parent_id ? data.parents[post.parent_id] : undefined}
          />
        ))}
      </div>
    </>
  );
}

function ProfileHeader({ profile }: { profile: IProfile }) {
  const { t } = useTranslation();
  const { user } = useAuth();
  const [recentlyFollowed, setRecentlyFollowed] = useState(false);
  const followMutation = useFollowMutation(profile.username);

  const handleFollow = () => {
    followMutation.mutate({
      user_id: profile.id,
      username: profile.username,
      type: "follow",
    });

    setRecentlyFollowed(true);
  };

  const handleUnfollow = () => {
    followMutation.mutate({
      user_id: profile.id,
      username: profile.username,
      type: "unfollow",
    });
  };

  return (
    <header className="profile-page__header">
      <div className="flex items-start justify-between">
        <div>
          <h1 className="profile-page__name">{profile.name}</h1>
          <div className="profile-page__username">@{profile.username}</div>
        </div>

        {user?.id === profile.id ? (
          <div>
            <Link to="/settings">
              <Button isRounded size="sm">
                {t("profile.edit-profile")}
              </Button>
            </Link>
          </div>
        ) : (
          <div className="flex flex-col">
            {profile.followed ? (
              <UnFollowButton
                onClick={handleUnfollow}
                recentlyFollowed={recentlyFollowed}
                setRecentlyFollowed={setRecentlyFollowed}
              />
            ) : (
              <FollowButton onClick={handleFollow} />
            )}
          </div>
        )}
      </div>
      <div className="profile-page__bio">{profile.bio}</div>
      <div className="profile-page__followers">
        <b className="font-bold">{profile.followers_count}</b>{" "}
        <span>{t("profile.followers")}</span>
      </div>
    </header>
  );
}

function FollowButton({ onClick }: { onClick(): void }) {
  const { t } = useTranslation();
  return (
    <Button
      isRounded
      variant="dark"
      size="sm"
      className={"transition-none w-28"}
      onClick={onClick}
    >
      {t("profile.follow")}
    </Button>
  );
}

function UnFollowButton({
  onClick,
  recentlyFollowed,
  setRecentlyFollowed,
}: {
  onClick(): void;
  setRecentlyFollowed(value: boolean): void;
  recentlyFollowed: boolean;
}) {
  const { t } = useTranslation();
  const [hover, setHover] = useState(false);
  return (
    <Button
      isRounded
      size="sm"
      variant={hover ? (recentlyFollowed ? "dark" : "danger") : "light"}
      className={"transition-none w-32 "}
      onClick={onClick}
      onMouseEnter={() => setHover(true)}
      onMouseLeave={() => {
        setHover(false);
        setRecentlyFollowed(false);
      }}
    >
      {hover && !recentlyFollowed
        ? t("profile.unfollow")
        : t("profile.following")}
    </Button>
  );
}
