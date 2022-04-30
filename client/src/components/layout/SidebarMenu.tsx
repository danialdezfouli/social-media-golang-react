import { Popover } from "@headlessui/react";
import { useAuth } from "contexts/AuthContext";
import { useTranslation } from "react-i18next";
import {
  RiHome7Line,
  RiHome8Fill,
  RiMoreFill,
  RiSearch2Fill,
  RiSearch2Line,
  RiSettings2Fill,
  RiSettings2Line,
  RiUser3Fill,
  RiUser3Line,
} from "react-icons/ri";
import { Link, NavLink } from "react-router-dom";
import "./SidebarMenu.css";

export default function SidebarMenu() {
  const { t } = useTranslation();
  const { user, logout } = useAuth();

  return (
    <>
      <header className="mb-6">
        <Link
          to="/home"
          className="text-lg hover:text-blue-600 transition-colors"
        >
          {t("site.name")}
        </Link>
      </header>
      <nav>
        <NavLink to="/home">
          <span>
            <RiHome8Fill />
            <RiHome7Line />
          </span>
          <span>{t("menu.home")}</span>
        </NavLink>
        {user?.id && (
          <NavLink to={"/search"}>
            <span>
              <RiSearch2Fill />
              <RiSearch2Line />
            </span>
            <span>{t("menu.search")}</span>
          </NavLink>
        )}
        {user?.id && (
          <NavLink to={"/profile/" + user?.username}>
            <span>
              <RiUser3Fill />
              <RiUser3Line />
            </span>
            <span>{t("menu.profile")}</span>
          </NavLink>
        )}
        {user?.id && (
          <NavLink to={"/settings"}>
            <span>
              <RiSettings2Fill />
              <RiSettings2Line />
            </span>
            <span>{t("menu.settings")}</span>
          </NavLink>
        )}
      </nav>

      {user && (
        <Popover className="user-profile">
          <Popover.Button className="user-profile__btn text-right">
            <div>
              <div className="">{user.name}</div>
              <div className="text-gray-600" dir="ltr">
                @{user.username}
              </div>
            </div>
            <RiMoreFill />
          </Popover.Button>
          <Popover.Panel className="user-profile__panel">
            <div>
              <button onClick={logout}>{t("menu.logout")}</button>
            </div>
          </Popover.Panel>
        </Popover>
      )}
    </>
  );
}
