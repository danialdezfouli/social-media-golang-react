import AlertMessage from "components/elements/AlertMessage";
import Button from "components/elements/Button";
import Text from "components/elements/Text";
import usePostMutation, {
  ICreatePostDTO,
} from "connection/mutations/usePostMutation";
import { IPost, PostTypesShape } from "connection/types";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";

type PostCreateFormProps = {
  replyTo?: IPost;
  type: PostTypesShape;
};

export default function PostCreateForm(props: PostCreateFormProps) {
  const { replyTo } = props;
  const { t } = useTranslation();
  const { isLoading, mutateAsync } = usePostMutation(replyTo, "");
  const { register, handleSubmit, setValue } = useForm<ICreatePostDTO>();
  const [errorResponse, setErrorResponse] = useState("");

  function onSubmit(inputs: ICreatePostDTO) {
    setErrorResponse("");

    mutateAsync({
      body: inputs.body,
      type: props.type,
      reply_to: replyTo?.post_id,
    })
      .then(({ data }) => {
        console.log(data);
        setValue("body", "");
      })
      .catch((err) => {
        console.log(err);
      });
  }

  return (
    <div className="flex flex-col gap-4 mb-4 px-4 pb-4 border-b-2">
      {replyTo && (
        <Text className="mt-4">
          {t("post.reply-to")} {replyTo.profile_name}
        </Text>
      )}
      <form action="" onSubmit={handleSubmit(onSubmit)}>
        {errorResponse && (
          <AlertMessage feedback="fail" title="خطا" message={errorResponse} />
        )}

        <div className="form-group">
          <textarea
            className="form-input w-full"
            id="body"
            maxLength={300}
            placeholder={t("post.form.body-placeholder")}
            {...register("body", { required: true })}
          />
        </div>

        <footer>
          <Button
            type="submit"
            variant="primary"
            className="mt-2"
            isLoading={isLoading}
            isRounded
          >
            {t("post.form.send")}
          </Button>
        </footer>
      </form>
    </div>
  );
}
