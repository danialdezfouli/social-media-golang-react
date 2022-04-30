import "./Card.css";

export default function Card({ children }: React.PropsWithChildren<{}>) {
  return <div className="card">{children}</div>;
}
