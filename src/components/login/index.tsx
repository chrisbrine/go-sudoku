import Login from "./login";
import Register from "./register";
import AlreadyLoggedIn from "./loggedIn";
import "./login.css";
import { EPage } from "../../types";
import { useAppSelector } from "../../redux/hooks";
import { getLoggedIn, getPage } from "../../redux/selectors";

export default function LoginPage() {
  const isLoggedIn = useAppSelector(getLoggedIn);
  const page = useAppSelector(getPage);
  if (isLoggedIn) {
    return <AlreadyLoggedIn />;
  }

  return (
    <div className="login-page">
      {page === EPage.LOGIN ? <Login /> : <Register />}
    </div>
  );
}
