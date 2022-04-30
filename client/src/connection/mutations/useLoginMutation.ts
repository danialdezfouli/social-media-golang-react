import { public_api } from "./../api";
import { useMutation } from "react-query";

export interface ILoginDTO {
  username: string;
  password: string;
}

interface ILoginResponse {
  data: {
    access_token: string;
  };
}

const loginRequest = (data: ILoginDTO) => {
  return public_api.post<ILoginDTO, ILoginResponse>("/auth/login", data, {
    withCredentials: true,
  });
};

export default function useLoginMutation() {
  return useMutation(loginRequest);
}
