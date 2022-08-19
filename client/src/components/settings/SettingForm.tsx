import Button from "components/elements/Button";
import { api } from "connection/api";
import { QUERY_KEYS } from "connection/QueryKeys";
import { useAuth } from "contexts/AuthContext";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useQueryClient } from "react-query";
import "./SettingForm.css";

function SettingForm() {
  const query = useQueryClient();
  const [loading, setLoading] = useState(false);
  const { user } = useAuth();
  const { register, handleSubmit } = useForm({
    defaultValues: {
      name: user?.name,
      bio: user?.bio,
    },
  });

  const onSubmit = async (input: any) => {
    try {
      setLoading(true);
      const response = await api.post("auth/profile", input);
      query.invalidateQueries(QUERY_KEYS.ME);
      console.log(response);
    } catch (error) {
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="settings-form">
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="form-group">
          <label htmlFor="name">نام</label>
          <div>
            <input
              dir="auto"
              style={{
                textAlign: "right",
              }}
              type="text"
              className="form-control"
              id="name"
              {...register("name", {
                required: true,
              })}
            />
          </div>
        </div>
        <div className="form-group">
          <label htmlFor="bio">بیوگرافی</label>
          <div>
            <textarea
              className="form-control"
              id="bio"
              {...register("bio", {})}
            />
          </div>
        </div>

        <Button type="submit" variant="primary" isLoading={loading}>
          ذخیره
        </Button>
      </form>
    </div>
  );
}
export default SettingForm;
