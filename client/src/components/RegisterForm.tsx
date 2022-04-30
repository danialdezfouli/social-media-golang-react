import useRegisterMutation, {
  IRegisterDTO,
} from "connection/mutations/useRegisterMutation";
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

export default function RegisterForm() {
  const { reset } = useQueryErrorResetBoundary();
  const queryClient = useQueryClient();
  const { t } = useTranslation();
  const { setUser } = useAuth();
  const { isLoading, mutateAsync } = useRegisterMutation();
  const { register, handleSubmit } = useForm<IRegisterDTO>();

  const [errorResponse, setErrorResponse] = useState("");
  let navigate = useNavigate();

  function onSubmit(inputs: IRegisterDTO) {
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
        const message = err?.response?.data.message;

        if (message === "email-exists") {
          setErrorResponse("شما قبلا با این ایمیل ثبت نام کرده اید.");
        } else if (message === "username-exists") {
          setErrorResponse("نام کاربری وارد شده قبلا انتخاب شده است.");
        } else if (message) {
          setErrorResponse(message);
        } else {
          console.error(err);
        }
        // if (err?.response?.status === 401) {
        //   setErrorResponse("اطلاعات وارد شده صحیح نیست.");
        // } else {
        // }
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
            <label htmlFor="name">{t("signup.name")}</label>
            <input
              className="form-input w-full"
              id="name"
              {...register("name", { required: true })}
            />
          </div>

          <div className="form-group mt-2">
            <label htmlFor="email">{t("signup.email")}</label>
            <input
              className="form-input w-full"
              type="email"
              id="email"
              {...register("email", { required: true })}
            />
          </div>

          <div className="form-group mt-2">
            <label htmlFor="username">{t("signup.username")}</label>
            <input
              className="form-input w-full"
              id="username"
              {...register("username", { required: true })}
            />
          </div>

          <div className="form-group mt-2">
            <label htmlFor="password">{t("signup.password")}</label>
            <input
              id="password"
              type="password"
              className="form-input w-full"
              {...register("password", { required: true })}
            />
          </div>
          <Button
            type="submit"
            variant="primary"
            className="w-full mt-4"
            isLoading={isLoading}
          >
            {t("signup.submit")}
          </Button>
        </form>
      </Card>
      <Card>
        <Text align="center">
          {t("signup.already-have-an-account")}{" "}
          <Link to="/" className="hover:underline text-blue-500">
            {t("signup.login")}
          </Link>
        </Text>
      </Card>
    </div>
  );
}
