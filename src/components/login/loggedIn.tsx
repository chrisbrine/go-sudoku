import { setPage } from "../../redux/actions";
import { useAppDispatch } from "../../redux/hooks";
import { EPage } from "../../types";
import "./loggedIn.css";

export default function AlreadyLoggedIn() {
  const dispatch = useAppDispatch();
  const toDashboard = () => dispatch(setPage(EPage.DASHBOARD));
  return (
    <div className="already-logged-in">
      <h1>Already logged in</h1>
      <button onClick={toDashboard}>Go to dashboard</button>
    </div>
  );
}
