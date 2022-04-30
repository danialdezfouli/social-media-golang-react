import { IProfile } from "connection/types";
import { useTranslation } from "react-i18next";
import { Link } from "react-router-dom";
import "./ProfileCompact.css";

interface ProfileCompactProps {
  profile: IProfile;
}

export default function ProfileCompact(props: ProfileCompactProps) {
  const { t } = useTranslation();
  const { profile } = props;
  return (
    <Link to={"/profile/" + profile.username} className="profile-list-item">
      <div>
        <div className="name">{profile.name}</div>
        <div className="username">@{profile.username}</div>
      </div>
      <div className="text-xs">
        <b className="followers">{profile.followers_count}</b>{" "}
        <span>{t("profile.followers")}</span>
      </div>
    </Link>
  );
}
