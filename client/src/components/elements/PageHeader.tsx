type PageHeaderProps = {};

export default function PageHeader({
  children,
}: React.PropsWithChildren<PageHeaderProps>) {
  return <div className="page-header py-4 px-4 font-bold">{children}</div>;
}
