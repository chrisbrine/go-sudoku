import { useState } from "react";
import { updatePassword } from "../../redux/actions";
import { useAppDispatch } from "../../redux/hooks";

export default function PasswordUpdate() {
  const dispatch = useAppDispatch();
  const udatePassword = ({
    oldPassword,
    password,
  }: {
    oldPassword: string;
    password: string;
  }) => dispatch(updatePassword({ oldPassword, password }));

  const [oldPassword, setOldPassword] = useState("");
  const [password, setPassword] = useState("");
  const [repeatPassword, setRepeatPassword] = useState("");

  const isValidPassword = () => {
    return (
      oldPassword.length > 3 &&
      password.length > 3 &&
      oldPassword !== password &&
      password === repeatPassword
    );
  };

  return (
    <div className="password">
      <h1>Password</h1>
      <div className="fields">
        <label>Old Password</label>
        <input
          type="password"
          placeholder="Old Password"
          value={oldPassword}
          onChange={(e) => setOldPassword(e.target.value)}
        />
        <label>New Password</label>
        <input
          type="password"
          placeholder="New Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <label>Repeat Password</label>
        <input
          type="password"
          placeholder="Repeat Password"
          value={repeatPassword}
          onChange={(e) => setRepeatPassword(e.target.value)}
        />
        <button
          className="action-button"
          disabled={!isValidPassword()}
          onClick={() => {
            udatePassword({ oldPassword, password });
          }}
        >
          Update Password
        </button>
      </div>
    </div>
  );
}
