import Button from "components/elements/Button";
import Text from "components/elements/Text";
import { ReactNode } from "react";
import "./AlertMessage.css";

type FeedbackType = "info" | "warning" | "fail" | "success";

type AlertMessageProps = {
  feedback: FeedbackType;
  title: ReactNode;
  message: ReactNode;
  actionCallback?: {
    callback(): void;
    label: string;
  };
};

export default function AlertMessage({
  feedback = "info",
  title,
  message,
  actionCallback,
}: AlertMessageProps) {
  return (
    <div className={"alert " + feedback}>
      {title && <Text className="font-bold">{title}</Text>}
      <div className="flex flex-row flex-wrap gap-4">
        <Text className="col-span-4 flex-grow">{message}</Text>
        {actionCallback && (
          <Button onClick={actionCallback.callback}>
            {actionCallback.label}
          </Button>
        )}
      </div>
    </div>
  );
}
