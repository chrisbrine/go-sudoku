import { useState } from "react";
import AlreadyLoggedIn from "./loggedIn";
import { EPage } from "../../types";
import { useAppDispatch, useAppSelector } from "../../redux/hooks";
import { getLoggedIn } from "../../redux/selectors";
import { doLogin, setPage } from "../../redux/actions";

export default function Login() {
  const dispatch = useAppDispatch();
  const isLoggedIn = useAppSelector(getLoggedIn);
  const onLogin = (username: string, password: string) =>
    dispatch(doLogin({ username, password }));
  const toRegisterPage = () => dispatch(setPage(EPage.REGISTER));
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const isValid = () => {
    return username.length > 0 && password.length > 0;
  };

  if (isLoggedIn) {
    return <AlreadyLoggedIn />;
  }

  return (
    <div className="login">
      <h1>Login</h1>
      <div className="fields">
        <label>Username</label>
        <input
          type="text"
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <label>Password</label>
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button
          className="action-button"
          disabled={!isValid()}
          onClick={() => {
            onLogin(username, password);
          }}
        >
          Login
        </button>
        <button className="switch-button" onClick={toRegisterPage}>
          Don't have an account? Register here.
        </button>
      </div>
    </div>
  );
}
