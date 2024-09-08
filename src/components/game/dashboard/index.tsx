import { updateDifficulty } from "../../../redux/actions";
import { useAppDispatch, useAppSelector } from "../../../redux/hooks";
import {
  getDifficulty,
  getLosses,
  getPerfectWins,
  getPoints,
  getWins,
} from "../../../redux/selectors/game";
import Leaderboard from "./leaderboard";
import "./dashboard.css";

export default function Dashboard() {
  const dispatch = useAppDispatch();
  const onDifficultyChange = (difficulty: number) =>
    dispatch(updateDifficulty(difficulty));

  const wins = useAppSelector(getWins);
  const losses = useAppSelector(getLosses);
  const perfectWins = useAppSelector(getPerfectWins);
  const points = useAppSelector(getPoints);
  const difficulty = useAppSelector(getDifficulty);
  const difficultyOptions = ["Easy", "Medium", "Hard"];

  return (
    <div className="no-board">
      <div className="stats">
        <div className="stat">
          <div className="label">Wins</div>
          <div className="value">{wins}</div>
        </div>
        <div className="stat">
          <div className="label">Losses</div>
          <div className="value">{losses}</div>
        </div>
        <div className="stat">
          <div className="label">Perfect Wins</div>
          <div className="value">{perfectWins}</div>
        </div>
        <div className="stat">
          <div className="label">Points</div>
          <div className="value">{points}</div>
        </div>
      </div>
      <div className="difficulty">
        <h2>Difficulty</h2>
        <div className="difficulty-select">
          {difficultyOptions.map((option, index) => (
            <button
              key={index}
              className={
                `difficulty-button` +
                (difficulty === index + 1 ? " selected" : "")
              }
              disabled={difficulty === index + 1}
              onClick={() => onDifficultyChange(index + 1)}
              value={index + 1}
            >
              {option}
            </button>
          ))}
        </div>
      </div>
      <hr />
      <Leaderboard />
    </div>
  );
}
