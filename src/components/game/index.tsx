import { useAppSelector } from "../../redux/hooks";
import { getPage } from "../../redux/selectors";
import { EPage } from "../../types";
import Board from "./board";
import Dashboard from "./dashboard";
import "./game.css";

export default function Game() {
  const page = useAppSelector(getPage);

  return page === EPage.GAME ? (
    <div className="game">
      <Board />
    </div>
  ) : (
    <Dashboard />
  );
}
