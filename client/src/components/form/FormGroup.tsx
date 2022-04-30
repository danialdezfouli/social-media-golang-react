import { HTMLInputTypeAttribute } from "react";

interface FormGroupProps {
  id: string;
  label?: string;
  type?: HTMLInputTypeAttribute;
}

export function FormGroup(props: React.PropsWithChildren<FormGroupProps>) {
  const { label, id, type } = props;

  return (
    <div className="form-group">
      {label && <label htmlFor={id}>{label}</label>}

      <input type={type} className="form-control" id={id} />
      {props.children}
    </div>
  );
}
