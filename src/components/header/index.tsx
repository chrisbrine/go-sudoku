import { doLogout, newGame, quitGame, setPage } from "../../redux/actions";
import { useAppDispatch, useAppSelector } from "../../redux/hooks";
import { getLoggedIn, getPage } from "../../redux/selectors";
import { EPage } from "../../types";
import "./header.css";

export default function Header() {
  const dispatch = useAppDispatch();
  const onQuitGame = () => dispatch(quitGame());
  const onNewGame = () => dispatch(newGame());
  const onLogout = () => dispatch(doLogout());
  const loginPage = () => dispatch(setPage(EPage.LOGIN));

  const isLoggedIn = useAppSelector(getLoggedIn);
  const page = useAppSelector(getPage);
  return (
    <header className="header">
      <h1>Sudoku</h1>
      <div className="buttons">
        {page === EPage.GAME && (
          <button onClick={() => onQuitGame()}>Quit Game</button>
        )}
        {page === EPage.DASHBOARD && (
          <button onClick={() => onNewGame()}>New Game</button>
        )}
        {isLoggedIn ? (
          <button onClick={onLogout}>Logout</button>
        ) : (
          <button onClick={loginPage}>Login</button>
        )}
      </div>
    </header>
  );
}
