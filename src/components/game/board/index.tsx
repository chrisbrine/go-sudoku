import Space from "./space";
import PickType from "./pickType";
import Stats from "./stats";
import Overlay from "./overlay";
import { useAppSelector } from "../../../redux/hooks";
import { getBoard } from "../../../redux/selectors/game";
import "./board.css";

export default function Board() {
  const board = useAppSelector(getBoard);
  return (
    <>
      <PickType />
      <Overlay />
      <div className="board">
        {board.map((row, colIndex) => (
          <div key={colIndex} className="row">
            {row.map((_, rowIndex) => (
              <Space key={rowIndex} row={rowIndex} col={colIndex} />
            ))}
          </div>
        ))}
      </div>
      <Stats />
    </>
  );
}
