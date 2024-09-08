import { useState } from "react";
import Profile from "./profile";
import PasswordUpdate from "./password";
import { setPage } from "../../redux/actions";
import { useAppDispatch } from "../../redux/hooks";
import { EPage } from "../../types";
import "./index.css";

export default function UserDashboard() {
  const dispatch = useAppDispatch();
  const toDashboard = () => dispatch(setPage(EPage.DASHBOARD));
  const [onPassword, setOnPassword] = useState(false);

  return (
    <div className="user-dashboard">
      <div className="tabs">
        <button
          className={!onPassword ? "selected" : ""}
          onClick={() => setOnPassword(false)}
        >
          Profile
        </button>
        <button
          className={onPassword ? "selected" : ""}
          onClick={() => setOnPassword(true)}
        >
          Password
        </button>
      </div>
      {onPassword ? <PasswordUpdate /> : <Profile />}
      <button className="back-button" onClick={toDashboard}>
        Back
      </button>
    </div>
  );
}
