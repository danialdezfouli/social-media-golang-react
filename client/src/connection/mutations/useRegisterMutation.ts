import { public_api } from "./../api";
import { useMutation } from "react-query";

export interface IRegisterDTO {
  name: string;
  email: string;
  username: string;
  password: string;
}

interface IRegisterResponse {
  data: {
    access_token: string;
  };
}

const registerRequest = (data: IRegisterDTO) => {
  return public_api.post<IRegisterDTO, IRegisterResponse>(
    "/auth/register",
    data,
    {
      withCredentials: true,
    }
  );
};

export default function useRegisterMutation() {
  return useMutation(registerRequest);
}
