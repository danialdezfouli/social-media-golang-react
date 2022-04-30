import cx from "classnames";

interface TextProps {
  size?: "xs" | "sm" | "md" | "lg" | "xl" | "2xl" | "3xl";
  as?: "p" | "h1" | "h2" | "h3" | "h4" | "h5" | "h6";
  align?: "left" | "center" | "right";
  weight?: "light" | "normal" | "medium" | "semibold" | "bold";
  className?: string;
}

export default function Text(props: React.PropsWithChildren<TextProps>) {
  const { size = "md", align, as: Component = "p", className, weight } = props;

  const classNames = cx(
    // text-xs text-sm text-md text-lg text-xl text-2xl text-3xl
    `text-${size}`,
    // text-left text-center text-right
    align && `text-${align}`,
    // font-light font-normal font-medium font-semibold font-bold
    weight && `font-${weight}`,
    className
  );

  return <Component className={classNames}>{props.children}</Component>;
}
