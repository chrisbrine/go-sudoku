import { quitGame } from "../../../redux/actions";
import { useAppDispatch, useAppSelector } from "../../../redux/hooks";
import {
  getGameStatus,
  getMistakes,
  getPlaying,
} from "../../../redux/selectors";
import { EGameStatus } from "../../../types";
import "./overlay.css";

export default function OverlayGameStatus() {
  const dispatch = useAppDispatch();
  const playing = useAppSelector(getPlaying);
  const onQuitGame = () => dispatch(quitGame());
  const doNothing = (e: React.MouseEvent) => e.stopPropagation();
  const gameStatus = useAppSelector(getGameStatus);
  const mistakes = useAppSelector(getMistakes);
  const message = gameStatus === EGameStatus.WON ? "You won!" : "You lost!";
  console.log(playing);

  return (
    <div
      className={`overlay-container ${playing ? "in-game" : "done-game"}`}
      onClick={onQuitGame}
    >
      <div className="overlay" onClick={doNothing}>
        <div className="overlay-content">
          <h1>{message}</h1>
          <p>
            And with only {mistakes} mistake{mistakes === 1 ? "" : "s"}!
          </p>
        </div>
        <button onClick={onQuitGame}>Quit Game</button>
      </div>
    </div>
  );
}
