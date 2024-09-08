import Header from "./components/header";
import Login from "./components/login";
import Game from "./components/game";
import { EPage } from "./types";
import { useAppSelector } from "./redux/hooks";
import { getLoading, getPage } from "./redux/selectors";
import "./App.css";

function App() {
  const loading = useAppSelector(getLoading);
  const page = useAppSelector(getPage);

  return (
    <div className="App">
      <Header />
      <div className="content">
        {loading ? (
          <div>Loading...</div>
        ) : page === EPage.LOGIN || page === EPage.REGISTER ? (
          <Login />
        ) : (
          <Game />
        )}
      </div>
    </div>
  );
}

export default App;
