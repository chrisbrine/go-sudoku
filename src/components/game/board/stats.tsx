import { useAppSelector } from "../../../redux/hooks";
import { getMistakes, getNumbersLeft } from "../../../redux/selectors";
import "./stats.css";

export default function BoardStats() {
  const mistakes = useAppSelector(getMistakes);
  const numbersLeft = useAppSelector(getNumbersLeft);

  return (
    <div className="board-stats">
      <div className="mistakes-container">
        <h3 className="mistakes-title">Mistakes</h3>
        <div className="mistakes">{mistakes}</div>
      </div>
      <div className="numbers-left-container">
        <h3>Numbers</h3>
        <div className="numbers-left">
          {numbersLeft.map((count, numMinusOne) => (
            <div
              key={numMinusOne}
              className={`number ${count <= 0 ? "unavailable" : "available"}`}
            >
              {numMinusOne + 1}
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
