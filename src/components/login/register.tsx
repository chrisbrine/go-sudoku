import { useState } from "react";
import AlreadyLoggedIn from "./loggedIn";
import { EPage } from "../../types";
import { useAppDispatch, useAppSelector } from "../../redux/hooks";
import { doRegister, setPage } from "../../redux/actions";
import { getLoggedIn } from "../../redux/selectors";

export default function Register() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [repeatPassword, setRepeatPassword] = useState("");
  const [name, setName] = useState("");

  const dispatch = useAppDispatch();
  const toLoginPage = () => dispatch(setPage(EPage.LOGIN));
  const onRegister = (username: string, password: string, name: string) =>
    dispatch(doRegister({ username, password, name }));
  const isLoggedIn = useAppSelector(getLoggedIn);

  const isValid = () => {
    const validity =
      username.length > 0 &&
      password.length > 0 &&
      repeatPassword.length > 0 &&
      name.length > 0 &&
      password === repeatPassword &&
      password.length >= 4;
    return validity;
  };

  if (isLoggedIn) {
    return <AlreadyLoggedIn />;
  }

  return (
    <div className="register">
      <h1>Register</h1>
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
        <label>Repeat Password</label>
        <input
          type="password"
          placeholder="Password"
          value={repeatPassword}
          onChange={(e) => setRepeatPassword(e.target.value)}
        />
        <label>Name</label>
        <input
          type="text"
          placeholder="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <button
          className="action-button"
          disabled={!isValid()}
          onClick={() => {
            onRegister(username, password, name);
          }}
        >
          Register
        </button>
        <button className="switch-button" onClick={toLoginPage}>
          Already have an account? Login here.
        </button>
      </div>
    </div>
  );
}
