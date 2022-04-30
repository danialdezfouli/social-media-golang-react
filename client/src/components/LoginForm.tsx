import useLoginMutation, {
  ILoginDTO,
} from "connection/mutations/useLoginMutation";
import { useAuth } from "contexts/AuthContext";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { useQueryClient, useQueryErrorResetBoundary } from "react-query";
import { Link, useNavigate } from "react-router-dom";
import tokenService from "utils/localStorageService";
import AlertMessage from "./elements/AlertMessage";
import Button from "./elements/Button";
import Text from "./elements/Text";
import Card from "./layout/Card";

export default function LoginForm() {
  const { reset } = useQueryErrorResetBoundary();
  const queryClient = useQueryClient();
  const { t } = useTranslation();
  const { setUser } = useAuth();
  const { isLoading, mutateAsync } = useLoginMutation();
  const { register, handleSubmit } = useForm<ILoginDTO>();

  const [errorResponse, setErrorResponse] = useState("");
  let navigate = useNavigate();

  function onSubmit(inputs: ILoginDTO) {
    setErrorResponse("");
    queryClient.clear();
    reset();

    mutateAsync(inputs)
      .then(({ data }) => {
        if (data.access_token) {
          setUser(null);
          tokenService.setToken(data.access_token);
          navigate("/home");
        }
      })
      .catch((err) => {
        if (err?.response?.status === 401) {
          setErrorResponse("اطلاعات وارد شده صحیح نیست.");
        } else {
          console.log(err);
        }
      });
  }

  return (
    <div className="flex flex-col gap-4">
      <Card>
        <form action="" onSubmit={handleSubmit(onSubmit)}>
          {errorResponse && (
            <AlertMessage feedback="fail" title="خطا" message={errorResponse} />
          )}

          <div className="form-group">
            <label htmlFor="email">{t("login.email")}</label>
            <input
              className="form-input w-full"
              id="email"
              defaultValue="danial"
              {...register("username", { required: true })}
            />
          </div>
          <div className="form-group mt-2">
            <label htmlFor="password">{t("login.password")}</label>
            <input
              type="password"
              className="form-input w-full"
              id="password"
              defaultValue="123456"
              {...register("password", { required: true })}
            />
          </div>
          <Button
            type="submit"
            variant="primary"
            className="w-full mt-4"
            isLoading={isLoading}
          >
            {t("login.submit")}
          </Button>
        </form>
      </Card>
      <Card>
        <Text align="center">
          {t("login.dont-have-an-account")}{" "}
          <Link to="/signup" className="hover:underline text-blue-500">
            {t("login.signup")}
          </Link>
        </Text>
      </Card>
    </div>
  );
}
