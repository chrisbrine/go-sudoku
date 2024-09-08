import { useAppSelector } from "../../../redux/hooks";
import { getLeaderboard } from "../../../redux/selectors";
import "./leaderboard.css";

export function Delim() {
  return <span className="delim">/</span>;
}

interface IItemProps {
  className: string;
  value: number;
}

export function Item({ className, value }: IItemProps) {
  return <span className={`stat-item ` + className}>{value}</span>;
}

export default function Leaderboard() {
  const leaderboard = useAppSelector(getLeaderboard);

  return leaderboard.length > 0 ? (
    <div className="leaderboard">
      <h2>Leaderboard</h2>
      <table>
        <thead>
          <tr>
            <th>Rank</th>
            <th>Username</th>
            <th>Name</th>
            <th>Wins/Losses/Perfect/Games</th>
            <th>Points</th>
          </tr>
        </thead>
        <tbody>
          {leaderboard.map((player, index) => (
            <tr key={player.Name}>
              <td>{index + 1}</td>
              <td>{player.Username}</td>
              <td>{player.Name}</td>
              <td>
                <Item className="wins" value={player.Wins} />
                <Delim />
                <Item className="losses" value={player.Losses} />
                <Delim />
                <Item className="perfect" value={player.PerfectWins} />
                <Delim />
                <Item className="games" value={player.Wins + player.Losses} />
              </td>
              <td>{player.Points}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  ) : (
    <div></div>
  );
}
