import { useAppSelector } from "../../redux/hooks";
import { getPage } from "../../redux/selectors";
import { EPage } from "../../types";
import Board from "./board";
import UserDashboard from "../user";
import Dashboard from "./dashboard";
import "./game.css";

export default function Game() {
  const page = useAppSelector(getPage);

  switch (page) {
    case EPage.GAME:
      return (
        <div className="game">
          <Board />
        </div>
      );
    case EPage.PROFILE:
      return <UserDashboard />;
    default:
      return <Dashboard />;
  }
}
