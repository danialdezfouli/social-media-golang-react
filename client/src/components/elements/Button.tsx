import cx from "classnames";
import "./Button.css";
import { Spinner } from "./Spinner";

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?:
    | "primary"
    | "secondary"
    | "light"
    | "dark"
    | "success"
    | "danger"
    | "warning";
  isLoading?: boolean;
  isRounded?: boolean;
  size?: "sm" | "md" | "lg";
}

export default function Button({
  children,
  variant = "light",
  className: givenClassName,
  isLoading,
  isRounded,
  size = "md",
  ...props
}: React.PropsWithChildren<ButtonProps>) {
  // btn-primary btn-success btn-danger btn-warning btn-info btn-light btn-dark
  return (
    <button
      className={cx("btn", "btn-" + variant, givenClassName, size, {
        loading: isLoading,
        "rounded-full": isRounded,
      })}
      {...props}
    >
      {isLoading ? <Spinner /> : children}
    </button>
  );
}
